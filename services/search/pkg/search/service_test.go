package search_test

import (
	"context"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	sprovider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	typesv1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	revactx "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/rgrpc/status"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	cs3mocks "github.com/cs3org/reva/v2/tests/cs3mocks/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/owncloud/ocis/v2/ocis-pkg/log"
	searchmsg "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/search/v0"
	searchsvc "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/search/v0"
	"github.com/owncloud/ocis/v2/services/search/pkg/config"
	"github.com/owncloud/ocis/v2/services/search/pkg/content"
	contentMocks "github.com/owncloud/ocis/v2/services/search/pkg/content/mocks"
	engineMocks "github.com/owncloud/ocis/v2/services/search/pkg/engine/mocks"
	"github.com/owncloud/ocis/v2/services/search/pkg/search"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

var _ = Describe("Searchprovider", func() {
	var (
		s               search.Searcher
		extractor       *contentMocks.Extractor
		gatewayClient   *cs3mocks.GatewayAPIClient
		gatewaySelector pool.Selectable[gateway.GatewayAPIClient]
		indexClient     *engineMocks.Engine
		ctx             context.Context
		logger          = log.NewLogger()
		user            = &userv1beta1.User{
			Id: &userv1beta1.UserId{
				OpaqueId: "user",
			},
		}
		otherUser = &userv1beta1.User{
			Id: &userv1beta1.UserId{
				OpaqueId: "otheruser",
			},
		}
		personalSpace = &sprovider.StorageSpace{
			Opaque: &typesv1beta1.Opaque{
				Map: map[string]*typesv1beta1.OpaqueEntry{
					"path": {
						Decoder: "plain",
						Value:   []byte("/foo"),
					},
				},
			},
			Id:   &sprovider.StorageSpaceId{OpaqueId: "storageid$personalspace!personalspace"},
			Root: &sprovider.ResourceId{StorageId: "storageid", SpaceId: "personalspace", OpaqueId: "personalspace"},
			Name: "personalspace",
		}

		ri = &sprovider.ResourceInfo{
			Id: &sprovider.ResourceId{
				StorageId: "storageid",
				OpaqueId:  "opaqueid",
			},
			ParentId: &sprovider.ResourceId{
				StorageId: "storageid",
				OpaqueId:  "parentopaqueid",
			},
			Path:  "foo.pdf",
			Size:  12345,
			Mtime: &typesv1beta1.Timestamp{Seconds: 4000},
		}
	)

	BeforeEach(func() {
		pool.RemoveSelector("GatewaySelector" + "com.owncloud.api.gateway")
		gatewayClient = &cs3mocks.GatewayAPIClient{}
		gatewaySelector = pool.GetSelector[gateway.GatewayAPIClient](
			"GatewaySelector",
			"com.owncloud.api.gateway",
			func(cc *grpc.ClientConn) gateway.GatewayAPIClient {
				return gatewayClient
			},
		)

		ctx = revactx.ContextSetUser(context.Background(), user)
		indexClient = &engineMocks.Engine{}
		extractor = &contentMocks.Extractor{}

		s = search.NewService(gatewaySelector, indexClient, extractor, logger, &config.Config{})

		gatewayClient.On("Authenticate", mock.Anything, mock.Anything).Return(&gateway.AuthenticateResponse{
			Status: status.NewOK(ctx),
			Token:  "authtoken",
		}, nil)
		gatewayClient.On("Stat", mock.Anything, mock.Anything).Return(&sprovider.StatResponse{
			Status: status.NewOK(context.Background()),
			Info:   ri,
		}, nil)
		gatewayClient.On("GetPath", mock.Anything, mock.MatchedBy(func(req *sprovider.GetPathRequest) bool {
			return req.ResourceId.OpaqueId == ri.Id.OpaqueId
		})).Return(&sprovider.GetPathResponse{
			Status: status.NewOK(context.Background()),
			Path:   ri.Path,
		}, nil)
		indexClient.On("DocCount").Return(uint64(1), nil)
	})

	Describe("New", func() {
		It("returns a new instance", func() {
			s := search.NewService(gatewaySelector, indexClient, extractor, logger, &config.Config{})
			Expect(s).ToNot(BeNil())
		})
	})

	Describe("IndexSpace", func() {
		It("walks the space and indexes all files", func() {
			gatewayClient.On("GetUserByClaim", mock.Anything, mock.Anything).Return(&userv1beta1.GetUserByClaimResponse{
				Status: status.NewOK(context.Background()),
				User:   user,
			}, nil)
			extractor.On("Extract", mock.Anything, mock.Anything, mock.Anything).Return(content.Document{}, nil)
			indexClient.On("Upsert", mock.Anything, mock.Anything).Return(nil)
			indexClient.On("Search", mock.Anything, mock.Anything).Return(&searchsvc.SearchIndexResponse{}, nil)

			err := s.IndexSpace(&sprovider.StorageSpaceId{OpaqueId: "storageid$spaceid!spaceid"}, user.Id)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("Search", func() {
		It("fails when an empty query is given", func() {
			res, err := s.Search(ctx, &searchsvc.SearchRequest{
				Query: "",
			})
			Expect(err).To(HaveOccurred())
			Expect(res).To(BeNil())
		})

		Context("with a personal space", func() {
			BeforeEach(func() {
				gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&sprovider.ListStorageSpacesResponse{
					Status:        status.NewOK(ctx),
					StorageSpaces: []*sprovider.StorageSpace{personalSpace},
				}, nil)
				indexClient.On("Search", mock.Anything, mock.Anything).Return(&searchsvc.SearchIndexResponse{
					TotalMatches: 1,
					Matches: []*searchmsg.Match{
						{
							Score: 1,
							Entity: &searchmsg.Entity{
								Ref: &searchmsg.Reference{
									ResourceId: &searchmsg.ResourceID{
										StorageId: personalSpace.Root.StorageId,
										SpaceId:   personalSpace.Root.SpaceId,
										OpaqueId:  personalSpace.Root.OpaqueId,
									},
									Path: "./path/to/Foo.pdf",
								},
								Id: &searchmsg.ResourceID{
									StorageId: personalSpace.Root.StorageId,
									OpaqueId:  "foo-id",
								},
								Name: "Foo.pdf",
							},
						},
					},
				}, nil)
			})

			It("does not mess with field-based searches", func() {
				_, err := s.Search(ctx, &searchsvc.SearchRequest{
					Query: "Size:<10",
				})
				Expect(err).ToNot(HaveOccurred())
				indexClient.AssertCalled(GinkgoT(), "Search", mock.Anything, mock.MatchedBy(func(req *searchsvc.SearchIndexRequest) bool {
					return req.Query == "Size:<10"
				}))
			})

			It("searches the personal user space", func() {
				res, err := s.Search(ctx, &searchsvc.SearchRequest{
					Query: "foo",
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(res).ToNot(BeNil())
				Expect(res.TotalMatches).To(Equal(int32(1)))
				Expect(len(res.Matches)).To(Equal(1))
				match := res.Matches[0]
				Expect(match.Entity.Id.OpaqueId).To(Equal("foo-id"))
				Expect(match.Entity.Name).To(Equal("Foo.pdf"))
				Expect(match.Entity.Ref.ResourceId.OpaqueId).To(Equal(personalSpace.Root.OpaqueId))
				Expect(match.Entity.Ref.Path).To(Equal("./path/to/Foo.pdf"))
			})
		})

		Context("with received shares", func() {
			var (
				grantSpace      *sprovider.StorageSpace
				mountpointSpace *sprovider.StorageSpace
			)

			BeforeEach(func() {
				grantSpace = &sprovider.StorageSpace{
					SpaceType: "grant",
					Owner:     otherUser,
					Id:        &sprovider.StorageSpaceId{OpaqueId: "storageproviderid$spaceid!otherspacegrant"},
					Root:      &sprovider.ResourceId{StorageId: "storageproviderid", SpaceId: "spaceid", OpaqueId: "otherspacegrant"},
					Name:      "grantspace",
				}
				mountpointSpace = &sprovider.StorageSpace{
					SpaceType: "mountpoint",
					Owner:     otherUser,
					Id:        &sprovider.StorageSpaceId{OpaqueId: "storageproviderid$spaceid!otherspacemountpoint"},
					Root:      &sprovider.ResourceId{StorageId: "storageproviderid", SpaceId: "spaceid", OpaqueId: "otherspacemountpoint"},
					Name:      "mountpointspace",
					Opaque: &typesv1beta1.Opaque{
						Map: map[string]*typesv1beta1.OpaqueEntry{
							"grantStorageID": {Decoder: "plain", Value: []byte("storageproviderid")},
							"grantSpaceID":   {Decoder: "plain", Value: []byte("spaceid")},
							"grantOpaqueID":  {Decoder: "plain", Value: []byte("otherspacegrant")},
						},
					},
				}
				gatewayClient.On("GetPath", mock.Anything, mock.Anything).Return(&sprovider.GetPathResponse{
					Status: status.NewOK(ctx),
					Path:   "/grant/path",
				}, nil)
			})

			It("searches the received spaces", func() {
				gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&sprovider.ListStorageSpacesResponse{
					Status:        status.NewOK(ctx),
					StorageSpaces: []*sprovider.StorageSpace{grantSpace, mountpointSpace},
				}, nil)
				indexClient.On("Search", mock.Anything, mock.Anything).Return(&searchsvc.SearchIndexResponse{
					TotalMatches: 1,
					Matches: []*searchmsg.Match{
						{
							Entity: &searchmsg.Entity{
								Ref: &searchmsg.Reference{
									ResourceId: &searchmsg.ResourceID{
										StorageId: grantSpace.Root.StorageId,
										SpaceId:   grantSpace.Root.SpaceId,
										OpaqueId:  grantSpace.Root.OpaqueId,
									},
									Path: "./grant/path/to/Shared.pdf",
								},
								Id: &searchmsg.ResourceID{
									StorageId: grantSpace.Root.StorageId,
									OpaqueId:  "grant-shared-id",
								},
								Name: "Shared.pdf",
							},
						},
					},
				}, nil)

				res, err := s.Search(ctx, &searchsvc.SearchRequest{
					Query: "Foo",
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(res).ToNot(BeNil())
				Expect(len(res.Matches)).To(Equal(1))
				match := res.Matches[0]
				Expect(match.Entity.Id.OpaqueId).To(Equal("grant-shared-id"))
				Expect(match.Entity.Name).To(Equal("Shared.pdf"))
				Expect(match.Entity.Ref.ResourceId.OpaqueId).To(Equal(mountpointSpace.Root.OpaqueId))
				Expect(match.Entity.Ref.Path).To(Equal("./to/Shared.pdf"))
			})

			Context("when searching both spaces", func() {
				BeforeEach(func() {
					gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything).Return(&sprovider.ListStorageSpacesResponse{
						Status:        status.NewOK(ctx),
						StorageSpaces: []*sprovider.StorageSpace{personalSpace, grantSpace, mountpointSpace},
					}, nil)
					indexClient.On("Search", mock.Anything, mock.MatchedBy(func(req *searchsvc.SearchIndexRequest) bool {
						return req.Ref.ResourceId.OpaqueId == grantSpace.Root.SpaceId &&
							req.Ref.ResourceId.SpaceId == grantSpace.Root.SpaceId
					})).Return(&searchsvc.SearchIndexResponse{
						TotalMatches: 2,
						Matches: []*searchmsg.Match{
							{
								Score: 2,
								Entity: &searchmsg.Entity{
									Ref: &searchmsg.Reference{
										ResourceId: &searchmsg.ResourceID{
											StorageId: grantSpace.Root.StorageId,
											SpaceId:   grantSpace.Root.SpaceId,
											OpaqueId:  grantSpace.Root.OpaqueId,
										},
										Path: "./grant/path/to/Shared.pdf",
									},
									Id: &searchmsg.ResourceID{
										StorageId: grantSpace.Root.StorageId,
										OpaqueId:  "grant-shared-id",
									},
									Name: "Shared.pdf",
								},
							},
							{
								Score: 0.01,
								Entity: &searchmsg.Entity{
									Ref: &searchmsg.Reference{
										ResourceId: &searchmsg.ResourceID{
											StorageId: grantSpace.Root.StorageId,
											SpaceId:   grantSpace.Root.SpaceId,
											OpaqueId:  grantSpace.Root.OpaqueId,
										},
										Path: "./grant/path/to/Irrelevant.pdf",
									},
									Id: &searchmsg.ResourceID{
										StorageId: grantSpace.Root.StorageId,
										OpaqueId:  "grant-irrelevant-id",
									},
									Name: "Irrelevant.pdf",
								},
							},
						},
					}, nil)
					indexClient.On("Search", mock.Anything, mock.MatchedBy(func(req *searchsvc.SearchIndexRequest) bool {
						return req.Ref.ResourceId.OpaqueId == personalSpace.Root.OpaqueId &&
							req.Ref.ResourceId.SpaceId == personalSpace.Root.SpaceId
					})).Return(&searchsvc.SearchIndexResponse{
						TotalMatches: 1,
						Matches: []*searchmsg.Match{
							{
								Score: 1,
								Entity: &searchmsg.Entity{
									Ref: &searchmsg.Reference{
										ResourceId: &searchmsg.ResourceID{
											StorageId: personalSpace.Root.StorageId,
											SpaceId:   personalSpace.Root.SpaceId,
											OpaqueId:  personalSpace.Root.OpaqueId,
										},
										Path: "./path/to/Foo.pdf",
									},
									Id: &searchmsg.ResourceID{
										StorageId: personalSpace.Root.StorageId,
										OpaqueId:  "foo-id",
									},
									Name: "Foo.pdf",
								},
							},
						},
					}, nil)
				})

				It("considers the search Ref parameter", func() {
					res, err := s.Search(ctx, &searchsvc.SearchRequest{
						Query: "foo",
						Ref: &searchmsg.Reference{
							ResourceId: &searchmsg.ResourceID{
								StorageId: "storageid",
								SpaceId:   "personalspace",
								OpaqueId:  "personalspace",
							},
						},
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(len(res.Matches)).To(Equal(1))
					Expect(res.Matches[0].Entity.Id.OpaqueId).To(Equal("foo-id"))
				})

				It("finds matches in both the personal space AND the grant", func() {
					res, err := s.Search(ctx, &searchsvc.SearchRequest{
						Query: "foo",
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(len(res.Matches)).To(Equal(3))
					ids := []string{res.Matches[0].Entity.Id.OpaqueId, res.Matches[1].Entity.Id.OpaqueId, res.Matches[2].Entity.Id.OpaqueId}
					Expect(ids).To(ConsistOf("foo-id", "grant-shared-id", "grant-irrelevant-id"))
				})

				It("sorts and limits the combined results from all spaces", func() {
					res, err := s.Search(ctx, &searchsvc.SearchRequest{
						Query:    "foo",
						PageSize: 2,
					})
					Expect(err).ToNot(HaveOccurred())
					Expect(res).ToNot(BeNil())
					Expect(len(res.Matches)).To(Equal(2))
					ids := []string{res.Matches[0].Entity.Id.OpaqueId, res.Matches[1].Entity.Id.OpaqueId}
					Expect(ids).To(Equal([]string{"grant-shared-id", "foo-id"}))
				})
			})
		})
	})
})
