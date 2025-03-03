package svc_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	userv1beta1 "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	typesv1beta1 "github.com/cs3org/go-cs3apis/cs3/types/v1beta1"
	revactx "github.com/cs3org/reva/v2/pkg/ctx"
	"github.com/cs3org/reva/v2/pkg/rgrpc/status"
	"github.com/cs3org/reva/v2/pkg/rgrpc/todo/pool"
	cs3mocks "github.com/cs3org/reva/v2/tests/cs3mocks/mocks"
	"github.com/go-chi/chi/v5"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	libregraph "github.com/owncloud/libre-graph-api-go"
	ogrpc "github.com/owncloud/ocis/v2/ocis-pkg/service/grpc"
	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	settingsmsg "github.com/owncloud/ocis/v2/protogen/gen/ocis/messages/settings/v0"
	settings "github.com/owncloud/ocis/v2/protogen/gen/ocis/services/settings/v0"
	"github.com/owncloud/ocis/v2/services/graph/mocks"
	"github.com/owncloud/ocis/v2/services/graph/pkg/config"
	"github.com/owncloud/ocis/v2/services/graph/pkg/config/defaults"
	identitymocks "github.com/owncloud/ocis/v2/services/graph/pkg/identity/mocks"
	service "github.com/owncloud/ocis/v2/services/graph/pkg/service/v0"
	"github.com/stretchr/testify/mock"
	"go-micro.dev/v4/client"
	"google.golang.org/grpc"
)

type userList struct {
	Value []*libregraph.User
}

var _ = Describe("Users", func() {
	var (
		svc             service.Service
		ctx             context.Context
		cfg             *config.Config
		gatewayClient   *cs3mocks.GatewayAPIClient
		gatewaySelector pool.Selectable[gateway.GatewayAPIClient]
		eventsPublisher mocks.Publisher
		roleService     *mocks.RoleService
		identityBackend *identitymocks.Backend

		rr *httptest.ResponseRecorder

		currentUser = &userv1beta1.User{
			Id: &userv1beta1.UserId{
				OpaqueId: "user",
			},
		}
	)

	BeforeEach(func() {
		eventsPublisher.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		pool.RemoveSelector("GatewaySelector" + "com.owncloud.api.gateway")
		gatewayClient = &cs3mocks.GatewayAPIClient{}
		gatewaySelector = pool.GetSelector[gateway.GatewayAPIClient](
			"GatewaySelector",
			"com.owncloud.api.gateway",
			func(cc *grpc.ClientConn) gateway.GatewayAPIClient {
				return gatewayClient
			},
		)

		identityBackend = &identitymocks.Backend{}
		roleService = &mocks.RoleService{}

		rr = httptest.NewRecorder()
		ctx = context.Background()

		cfg = defaults.FullDefaultConfig()
		cfg.Identity.LDAP.CACert = "" // skip the startup checks, we don't use LDAP at all in this tests
		cfg.TokenManager.JWTSecret = "loremipsum"
		cfg.Commons = &shared.Commons{}
		cfg.GRPCClientTLS = &shared.GRPCClientTLS{}
		cfg.Application.ID = "some-application-ID"

		_ = ogrpc.Configure(ogrpc.GetClientOptions(cfg.GRPCClientTLS)...)
		svc, _ = service.NewService(
			service.Config(cfg),
			service.WithGatewaySelector(gatewaySelector),
			service.EventsPublisher(&eventsPublisher),
			service.WithIdentityBackend(identityBackend),
			service.WithRoleService(roleService),
		)
	})

	Describe("GetMe", func() {
		It("handles missing user", func() {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me", nil)
			svc.GetMe(rr, r)

			Expect(rr.Code).To(Equal(http.StatusInternalServerError))
		})

		It("gets the information", func() {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me", nil)
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.GetMe(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
		})

		It("expands the memberOf", func() {
			user := &libregraph.User{
				Id: libregraph.PtrString("user1"),
				MemberOf: []libregraph.Group{
					{DisplayName: libregraph.PtrString("somegroup")},
				},
			}
			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me?$expand=memberOf", nil)
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.GetMe(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			responseUser := &libregraph.User{}
			err = json.Unmarshal(data, &responseUser)
			Expect(err).ToNot(HaveOccurred())

			Expect(responseUser.GetId()).To(Equal("user1"))
			Expect(responseUser.GetMemberOf()).To(HaveLen(1))
			Expect(responseUser.GetMemberOf()[0].GetDisplayName()).To(Equal("somegroup"))

		})

		It("expands the appRoleAssignments", func() {
			assignments := []*settingsmsg.UserRoleAssignment{
				{
					Id:          "some-appRoleAssignment-ID",
					AccountUuid: "user",
					RoleId:      "some-appRole-ID",
				},
			}
			roleService.On("ListRoleAssignments", mock.Anything, mock.Anything, mock.Anything).Return(&settings.ListRoleAssignmentsResponse{Assignments: assignments}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/me?$expand=appRoleAssignments", nil)
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.GetMe(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))

			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			responseUser := &libregraph.User{}
			err = json.Unmarshal(data, &responseUser)
			Expect(err).ToNot(HaveOccurred())

			Expect(responseUser.GetId()).To(Equal("user"))
			Expect(responseUser.GetAppRoleAssignments()).To(HaveLen(1))
			Expect(responseUser.GetAppRoleAssignments()[0].GetId()).To(Equal("some-appRoleAssignment-ID"))
			Expect(responseUser.GetAppRoleAssignments()[0].GetAppRoleId()).To(Equal("some-appRole-ID"))
			Expect(responseUser.GetAppRoleAssignments()[0].GetPrincipalId()).To(Equal("user"))
			Expect(responseUser.GetAppRoleAssignments()[0].GetResourceId()).To(Equal("some-application-ID"))
		})
	})

	Describe("GetUsers", func() {
		It("handles invalid requests", func() {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$invalid=true", nil)
			svc.GetUsers(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("lists the users", func() {
			user := &libregraph.User{}
			user.SetId("user1")
			users := []*libregraph.User{user}

			identityBackend.On("GetUsers", mock.Anything, mock.Anything, mock.Anything).Return(users, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users", nil)
			svc.GetUsers(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			res := userList{}
			err = json.Unmarshal(data, &res)
			Expect(err).ToNot(HaveOccurred())

			Expect(len(res.Value)).To(Equal(1))
			Expect(res.Value[0].GetId()).To(Equal("user1"))
		})

		It("sorts", func() {
			user := &libregraph.User{}
			user.SetId("user1")
			user.SetMail("z@example.com")
			user.SetDisplayName("9")
			user.SetOnPremisesSamAccountName("9")
			user2 := &libregraph.User{}
			user2.SetId("user2")
			user2.SetMail("a@example.com")
			user2.SetDisplayName("1")
			user2.SetOnPremisesSamAccountName("1")
			users := []*libregraph.User{user, user2}

			identityBackend.On("GetUsers", mock.Anything, mock.Anything, mock.Anything).Return(users, nil)

			getUsers := func(path string) []*libregraph.User {
				r := httptest.NewRequest(http.MethodGet, path, nil)
				rec := httptest.NewRecorder()
				svc.GetUsers(rec, r)

				Expect(rec.Code).To(Equal(http.StatusOK))
				data, err := io.ReadAll(rec.Body)
				Expect(err).ToNot(HaveOccurred())

				res := userList{}
				err = json.Unmarshal(data, &res)
				Expect(err).ToNot(HaveOccurred())
				return res.Value
			}

			unsorted := getUsers("/graph/v1.0/users")
			Expect(len(unsorted)).To(Equal(2))
			Expect(unsorted[0].GetId()).To(Equal("user1"))
			Expect(unsorted[1].GetId()).To(Equal("user2"))

			byMail := getUsers("/graph/v1.0/users?$orderby=mail")
			Expect(len(byMail)).To(Equal(2))
			Expect(byMail[0].GetId()).To(Equal("user2"))
			Expect(byMail[1].GetId()).To(Equal("user1"))
			byMail = getUsers("/graph/v1.0/users?$orderby=mail%20asc")
			Expect(len(byMail)).To(Equal(2))
			Expect(byMail[0].GetId()).To(Equal("user2"))
			Expect(byMail[1].GetId()).To(Equal("user1"))
			byMail = getUsers("/graph/v1.0/users?$orderby=mail%20desc")
			Expect(len(byMail)).To(Equal(2))
			Expect(byMail[0].GetId()).To(Equal("user1"))
			Expect(byMail[1].GetId()).To(Equal("user2"))

			byDisplayName := getUsers("/graph/v1.0/users?$orderby=displayName")
			Expect(len(byDisplayName)).To(Equal(2))
			Expect(byDisplayName[0].GetId()).To(Equal("user2"))
			Expect(byDisplayName[1].GetId()).To(Equal("user1"))
			byDisplayName = getUsers("/graph/v1.0/users?$orderby=displayName%20asc")
			Expect(len(byDisplayName)).To(Equal(2))
			Expect(byDisplayName[0].GetId()).To(Equal("user2"))
			Expect(byDisplayName[1].GetId()).To(Equal("user1"))
			byDisplayName = getUsers("/graph/v1.0/users?$orderby=displayName%20desc")
			Expect(len(byDisplayName)).To(Equal(2))
			Expect(byDisplayName[0].GetId()).To(Equal("user1"))
			Expect(byDisplayName[1].GetId()).To(Equal("user2"))

			byOnPremisesSamAccountName := getUsers("/graph/v1.0/users?$orderby=onPremisesSamAccountName")
			Expect(len(byOnPremisesSamAccountName)).To(Equal(2))
			Expect(byOnPremisesSamAccountName[0].GetId()).To(Equal("user2"))
			Expect(byOnPremisesSamAccountName[1].GetId()).To(Equal("user1"))
			byOnPremisesSamAccountName = getUsers("/graph/v1.0/users?$orderby=onPremisesSamAccountName%20asc")
			Expect(len(byOnPremisesSamAccountName)).To(Equal(2))
			Expect(byOnPremisesSamAccountName[0].GetId()).To(Equal("user2"))
			Expect(byOnPremisesSamAccountName[1].GetId()).To(Equal("user1"))
			byOnPremisesSamAccountName = getUsers("/graph/v1.0/users?$orderby=onPremisesSamAccountName%20desc")
			Expect(len(byOnPremisesSamAccountName)).To(Equal(2))
			Expect(byOnPremisesSamAccountName[0].GetId()).To(Equal("user1"))
			Expect(byOnPremisesSamAccountName[1].GetId()).To(Equal("user2"))

			// Handles invalid sort field
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$orderby=invalid", nil)
			svc.GetUsers(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("expands the appRoleAssignments", func() {

			user := &libregraph.User{}
			user.SetId("user1")
			user.SetMail("z@example.com")
			user.SetDisplayName("9")
			user.SetOnPremisesSamAccountName("9")
			user2 := &libregraph.User{}
			user2.SetId("user2")
			user2.SetMail("a@example.com")
			user2.SetDisplayName("1")
			user2.SetOnPremisesSamAccountName("1")
			users := []*libregraph.User{user, user2}
			identityBackend.On("GetUsers", mock.Anything, mock.Anything, mock.Anything).Return(users, nil)

			roleService.On("ListRoleAssignments", mock.Anything, mock.Anything, mock.Anything).Return(func(ctx context.Context, in *settings.ListRoleAssignmentsRequest, opts ...client.CallOption) *settings.ListRoleAssignmentsResponse {
				return &settings.ListRoleAssignmentsResponse{Assignments: []*settingsmsg.UserRoleAssignment{
					{
						Id:          "some-appRoleAssignment-ID",
						AccountUuid: in.GetAccountUuid(),
						RoleId:      "some-appRole-ID",
					},
				}}
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$expand=appRoleAssignments", nil)
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.GetUsers(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))

			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			res := userList{}
			err = json.Unmarshal(data, &res)
			Expect(err).ToNot(HaveOccurred())

			responseUsers := res.Value
			Expect(len(responseUsers)).To(Equal(2))
			Expect(responseUsers[0].GetId()).To(Equal("user1"))
			Expect(responseUsers[0].GetAppRoleAssignments()).To(HaveLen(1))
			Expect(responseUsers[0].GetAppRoleAssignments()[0].GetId()).To(Equal("some-appRoleAssignment-ID"))
			Expect(responseUsers[0].GetAppRoleAssignments()[0].GetAppRoleId()).To(Equal("some-appRole-ID"))
			Expect(responseUsers[0].GetAppRoleAssignments()[0].GetPrincipalId()).To(Equal("user1"))
			Expect(responseUsers[0].GetAppRoleAssignments()[0].GetResourceId()).To(Equal("some-application-ID"))

			Expect(responseUsers[1].GetId()).To(Equal("user2"))
			Expect(responseUsers[1].GetAppRoleAssignments()).To(HaveLen(1))
			Expect(responseUsers[1].GetAppRoleAssignments()[0].GetId()).To(Equal("some-appRoleAssignment-ID"))
			Expect(responseUsers[1].GetAppRoleAssignments()[0].GetAppRoleId()).To(Equal("some-appRole-ID"))
			Expect(responseUsers[1].GetAppRoleAssignments()[0].GetPrincipalId()).To(Equal("user2"))
			Expect(responseUsers[1].GetAppRoleAssignments()[0].GetResourceId()).To(Equal("some-application-ID"))

		})
	})

	DescribeTable("GetUsers handles unsupported or invalid filters",
		func(filter string, status int) {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$filter="+url.QueryEscape(filter), nil)
			svc.GetUsers(rr, r)

			Expect(rr.Code).To(Equal(status))
		},
		Entry("with invalid filter", "invalid", http.StatusBadRequest),
		Entry("with unsupported filter for user property", "mail eq 'unsupported'", http.StatusNotImplemented),
		Entry("with unsupported filter operation", "mail add 10", http.StatusNotImplemented),
		Entry("with unsupported logical operation", "memberOf/any(n:n/id eq 1) or memberOf/any(n:n/id eq 2)", http.StatusNotImplemented),
		Entry("with unsupported lambda query ", `drives/any(n:n/id eq '1')`, http.StatusNotImplemented),
		Entry("with unsupported lambda token ", "memberOf/all(n:n/id eq 1)", http.StatusNotImplemented),
		Entry("with unsupported filter operation ", "memberOf/any(n:n/id ne 1)", http.StatusNotImplemented),
		Entry("with unsupported filter operand type", "memberOf/any(n:n/id eq 1)", http.StatusNotImplemented),
		Entry("with unsupported memberOf lambda filter property", "memberOf/any(n:n/name eq 'name')", http.StatusNotImplemented),
		Entry("with unsupported appRoleAssignments lambda filter property", "appRoleAssignments/any(n:n/id eq 'id')", http.StatusNotImplemented),
		Entry("with unsupported appRoleAssignments lambda filter property",
			"appRoleAssignments/any(n:n/id eq 'id') and appRoleAssignments/any(n:n/id eq 'id')", http.StatusNotImplemented),
		Entry("with unsupported appRoleAssignments lambda filter operation",
			"appRoleAssignments/all(n:n/appRoleId eq 'id') and appRoleAssignments/any(n:n/appRoleId eq 'id')", http.StatusNotImplemented),
		Entry("with unsupported appRoleAssignments lambda filter operation",
			"appRoleAssignments/any(n:n/appRoleId ne 'id') and appRoleAssignments/any(n:n/appRoleId eq 'id')", http.StatusNotImplemented),
		Entry("with unsupported appRoleAssignments lambda filter operation",
			"appRoleAssignments/any(n:n/appRoleId eq 1) and appRoleAssignments/any(n:n/appRoleId eq 'id')", http.StatusNotImplemented),
	)

	DescribeTable("With a valid filter",
		func(filter string, status int) {
			user := &libregraph.User{}
			user.SetId("25cb7bc0-3168-4a0c-adbe-396f478ad494")
			users := []*libregraph.User{user}
			identityBackend.On("GetGroupMembers", mock.Anything, "25cb7bc0-3168-4a0c-adbe-396f478ad494", mock.Anything).Return(users, nil)
			identityBackend.On("GetGroupMembers", mock.Anything, "2713f1d5-6822-42bd-ad56-9f6c55a3a8fa", mock.Anything).Return([]*libregraph.User{}, nil)
			identityBackend.On("GetUsers", mock.Anything, mock.Anything).Return([]*libregraph.User{user}, nil)
			roleService.On("ListRoleAssignments", mock.Anything, mock.Anything, mock.Anything).
				Return(func(ctx context.Context, in *settings.ListRoleAssignmentsRequest, opts ...client.CallOption) *settings.ListRoleAssignmentsResponse {
					return &settings.ListRoleAssignmentsResponse{Assignments: []*settingsmsg.UserRoleAssignment{
						{
							Id:          "some-appRoleAssignment-ID",
							AccountUuid: user.GetId(),
							RoleId:      "some-appRole-ID",
						},
					}}
				}, nil)
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$filter="+url.QueryEscape(filter), nil)
			svc.GetUsers(rr, r)

			Expect(rr.Code).To(Equal(status))
		},
		Entry("with memberOf lambda filter with UUID", "memberOf/any(n:n/id eq 25cb7bc0-3168-4a0c-adbe-396f478ad494)", http.StatusOK),
		Entry("with memberOf lambda filter with UUID string", "memberOf/any(n:n/id eq '25cb7bc0-3168-4a0c-adbe-396f478ad494')", http.StatusOK),
		Entry("with appRoleAssignments lambda filter with appRoleId", "appRoleAssignments/any(n:n/appRoleId eq 'some-appRole-ID')", http.StatusOK),
		Entry("with two memberOf lambda filters combined with and",
			"memberOf/any(n:n/id eq 25cb7bc0-3168-4a0c-adbe-396f478ad494) and memberOf/any(n:n/id eq 2713f1d5-6822-42bd-ad56-9f6c55a3a8fa)",
			http.StatusOK),
		Entry("with two memberOf lambda filters combined with or",
			"memberOf/any(n:n/id eq 25cb7bc0-3168-4a0c-adbe-396f478ad494) or memberOf/any(n:n/id eq 2713f1d5-6822-42bd-ad56-9f6c55a3a8fa)",
			http.StatusOK),
		Entry("with supported appRoleAssignments lambda filter property",
			"appRoleAssignments/any(n:n/appRoleId eq 'some-appRoleAssignment-ID') and memberOf/any(n:n/id eq 2713f1d5-6822-42bd-ad56-9f6c55a3a8fa)",
			http.StatusOK),
	)

	Describe("GetUser", func() {
		It("handles missing userids", func() {
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users", nil)
			svc.GetUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("gets the user", func() {
			user := &libregraph.User{}
			user.SetId("user1")

			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", *user.Id)
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.GetUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)

			Expect(err).ToNot(HaveOccurred())
			responseUser := &libregraph.User{}
			err = json.Unmarshal(data, &responseUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(responseUser.GetId()).To(Equal("user1"))
			Expect(len(responseUser.GetDrives())).To(Equal(0))
		})

		It("includes the personal space if requested", func() {
			user := &libregraph.User{}
			user.SetId("user1")

			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status:     status.NewOK(ctx),
				TotalBytes: 10,
			}, nil)
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "drive1"},
						Root:      &provider.ResourceId{SpaceId: "space", OpaqueId: "space"},
						SpaceType: "project",
					},
					{
						Id:        &provider.StorageSpaceId{OpaqueId: "personal"},
						Owner:     &userv1beta1.User{Id: &userv1beta1.UserId{OpaqueId: "user1"}},
						Root:      &provider.ResourceId{SpaceId: "personal", OpaqueId: "personal"},
						SpaceType: "personal",
					},
				},
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$expand=drive", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", *user.Id)
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.GetUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())
			responseUser := &libregraph.User{}
			err = json.Unmarshal(data, &responseUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(responseUser.GetId()).To(Equal("user1"))
			Expect(*responseUser.GetDrive().Id).To(Equal("personal"))
		})

		It("includes the drives if requested", func() {
			user := &libregraph.User{}
			user.SetId("user1")

			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
			gatewayClient.On("GetQuota", mock.Anything, mock.Anything, mock.Anything).Return(&provider.GetQuotaResponse{
				Status:     status.NewOK(ctx),
				TotalBytes: 10,
			}, nil)
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Id:   &provider.StorageSpaceId{OpaqueId: "drive1"},
						Root: &provider.ResourceId{SpaceId: "space", OpaqueId: "space"},
					},
				},
			}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users?$expand=drives", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", *user.Id)
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.GetUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())
			responseUser := &libregraph.User{}
			err = json.Unmarshal(data, &responseUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(responseUser.GetId()).To(Equal("user1"))
			Expect(len(responseUser.GetDrives())).To(Equal(1))
		})

		It("expands the appRoleAssignments", func() {
			user := &libregraph.User{}
			user.SetId("user1")

			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)

			assignments := []*settingsmsg.UserRoleAssignment{
				{
					Id:          "some-appRoleAssignment-ID",
					AccountUuid: "user1",
					RoleId:      "some-appRole-ID",
				},
			}
			roleService.On("ListRoleAssignments", mock.Anything, mock.Anything, mock.Anything).Return(&settings.ListRoleAssignmentsResponse{Assignments: assignments}, nil)

			r := httptest.NewRequest(http.MethodGet, "/graph/v1.0/users/user1?$expand=appRoleAssignments", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", user.GetId())
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.GetUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))

			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			responseUser := &libregraph.User{}
			err = json.Unmarshal(data, &responseUser)
			Expect(err).ToNot(HaveOccurred())

			Expect(responseUser.GetId()).To(Equal("user1"))
			Expect(responseUser.GetAppRoleAssignments()).To(HaveLen(1))
			Expect(responseUser.GetAppRoleAssignments()[0].GetId()).To(Equal("some-appRoleAssignment-ID"))
			Expect(responseUser.GetAppRoleAssignments()[0].GetAppRoleId()).To(Equal("some-appRole-ID"))
			Expect(responseUser.GetAppRoleAssignments()[0].GetPrincipalId()).To(Equal("user1"))
			Expect(responseUser.GetAppRoleAssignments()[0].GetResourceId()).To(Equal("some-application-ID"))
		})
	})

	Describe("PostUser", func() {
		var (
			user *libregraph.User

			assertHandleBadAttributes = func(user *libregraph.User) {
				userJson, err := json.Marshal(user)
				Expect(err).ToNot(HaveOccurred())

				r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users", bytes.NewBuffer(userJson))
				svc.PostUser(rr, r)

				Expect(rr.Code).To(Equal(http.StatusBadRequest))
			}
		)

		BeforeEach(func() {
			user = &libregraph.User{}
			user.SetDisplayName("Display Name")
			user.SetOnPremisesSamAccountName("user")
			user.SetMail("user@example.com")
		})

		It("handles invalid bodies", func() {
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users?$invalid=true", nil)
			svc.PostUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("handles missing display names", func() {
			user.DisplayName = nil
			assertHandleBadAttributes(user)

		})

		It("handles missing OnPremisesSamAccountName", func() {
			user.OnPremisesSamAccountName = nil
			assertHandleBadAttributes(user)

			user.SetOnPremisesSamAccountName("")
			assertHandleBadAttributes(user)
		})

		It("handles bad Mails", func() {
			user.SetMail("not-a-mail-address")
			assertHandleBadAttributes(user)
		})

		It("handles set Ids - they are read-only", func() {
			user.SetId("/users/user")
			assertHandleBadAttributes(user)
		})

		It("handles invalid userType", func() {
			user.SetUserType("Clown")
			assertHandleBadAttributes(user)
		})

		It("creates a user", func() {
			roleService.On("AssignRoleToUser", mock.Anything, mock.Anything).Return(&settings.AssignRoleToUserResponse{}, nil)
			identityBackend.On("CreateUser", mock.Anything, mock.Anything).Return(func(ctx context.Context, user libregraph.User) *libregraph.User {
				user.SetId("/users/user")
				return &user
			}, nil)
			userJson, err := json.Marshal(user)
			Expect(err).ToNot(HaveOccurred())

			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users", bytes.NewBuffer(userJson))
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.PostUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			createdUser := libregraph.User{}
			err = json.Unmarshal(data, &createdUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(createdUser.GetUserType()).To(Equal("Member"))
		})

		It("creates a guest user", func() {
			roleService.On("AssignRoleToUser", mock.Anything, mock.Anything).Return(&settings.AssignRoleToUserResponse{}, nil)
			identityBackend.On("CreateUser", mock.Anything, mock.Anything).Return(func(ctx context.Context, user libregraph.User) *libregraph.User {
				user.SetId("/users/user")
				return &user
			}, nil)

			user.SetUserType("Guest")
			userJson, err := json.Marshal(user)
			Expect(err).ToNot(HaveOccurred())

			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users", bytes.NewBuffer(userJson))
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.PostUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			createdUser := libregraph.User{}
			err = json.Unmarshal(data, &createdUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(createdUser.GetUserType()).To(Equal("Guest"))
		})

		It("creates a member user", func() {
			roleService.On("AssignRoleToUser", mock.Anything, mock.Anything).Return(&settings.AssignRoleToUserResponse{}, nil)
			identityBackend.On("CreateUser", mock.Anything, mock.Anything).Return(func(ctx context.Context, user libregraph.User) *libregraph.User {
				user.SetId("/users/user")
				return &user
			}, nil)

			user.SetUserType("Member")
			userJson, err := json.Marshal(user)
			Expect(err).ToNot(HaveOccurred())

			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users", bytes.NewBuffer(userJson))
			r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
			svc.PostUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err := io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			createdUser := libregraph.User{}
			err = json.Unmarshal(data, &createdUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(createdUser.GetUserType()).To(Equal("Member"))
		})

		Describe("Handling usernames with spaces", func() {
			var (
				newSvc = func(usernameMatch string) service.Service {
					localCfg := defaults.FullDefaultConfig()
					localCfg.Identity.LDAP.CACert = "" // skip the startup checks, we don't use LDAP at all in this tests
					localCfg.TokenManager.JWTSecret = "loremipsum"
					localCfg.Commons = &shared.Commons{}
					localCfg.GRPCClientTLS = &shared.GRPCClientTLS{}

					localCfg.API.UsernameMatch = usernameMatch

					_ = ogrpc.Configure(ogrpc.GetClientOptions(cfg.GRPCClientTLS)...)
					localSvc, _ := service.NewService(
						service.Config(localCfg),
						service.WithGatewaySelector(gatewaySelector),
						service.EventsPublisher(&eventsPublisher),
						service.WithIdentityBackend(identityBackend),
						service.WithRoleService(roleService),
					)

					return localSvc
				}
			)

			BeforeEach(func() {
				user.SetOnPremisesSamAccountName("username with spaces")
			})

			It("rejects a username with spaces if match regex is default", func() {
				userJson, err := json.Marshal(user)
				Expect(err).ToNot(HaveOccurred())

				r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/me/users", bytes.NewBuffer(userJson))
				r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))

				newSvc("default").PostUser(rr, r)
				Expect(rr.Code).To(Equal(http.StatusBadRequest))
			})

			It("creates a user with spaces in username if regex is none", func() {
				roleService.On("AssignRoleToUser", mock.Anything, mock.Anything).Return(&settings.AssignRoleToUserResponse{}, nil)
				identityBackend.On("CreateUser", mock.Anything, mock.Anything).Return(func(ctx context.Context, user libregraph.User) *libregraph.User {
					user.SetId("/users/user")
					return &user
				}, nil)
				userJson, err := json.Marshal(user)
				Expect(err).ToNot(HaveOccurred())

				r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/me/users", bytes.NewBuffer(userJson))
				r = r.WithContext(revactx.ContextSetUser(ctx, currentUser))
				newSvc("none").PostUser(rr, r)

				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})

	})

	Describe("DeleteUser", func() {
		It("handles missing userids", func() {
			r := httptest.NewRequest(http.MethodDelete, "/graph/v1.0/users/{userid}", nil)
			svc.DeleteUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("prevents a user from deleting themselves", func() {
			lu := libregraph.User{}
			lu.SetId(currentUser.Id.OpaqueId)
			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(&lu, nil)

			r := httptest.NewRequest(http.MethodDelete, "/graph/v1.0/users/{userid}", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", currentUser.Id.OpaqueId)
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.DeleteUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusForbidden))
		})

		It("deletes a user from deleting themselves", func() {
			otheruser := &userv1beta1.User{
				Id: &userv1beta1.UserId{
					OpaqueId: "otheruser",
				},
			}

			lu := libregraph.User{}
			lu.SetId(otheruser.Id.OpaqueId)
			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(&lu, nil)
			identityBackend.On("DeleteUser", mock.Anything, mock.Anything).Return(nil)
			gatewayClient.On("DeleteStorageSpace", mock.Anything, mock.Anything).Return(&provider.DeleteStorageSpaceResponse{
				Status: status.NewOK(ctx),
			}, nil)
			gatewayClient.On("ListStorageSpaces", mock.Anything, mock.Anything, mock.Anything).Return(&provider.ListStorageSpacesResponse{
				Status: status.NewOK(ctx),
				StorageSpaces: []*provider.StorageSpace{
					{
						Opaque:    &typesv1beta1.Opaque{},
						Id:        &provider.StorageSpaceId{OpaqueId: "drive1"},
						Root:      &provider.ResourceId{SpaceId: "space", OpaqueId: "space"},
						SpaceType: "personal",
						Owner:     otheruser,
					},
				},
			}, nil)

			r := httptest.NewRequest(http.MethodDelete, "/graph/v1.0/users/{userid}", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", lu.GetId())
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.DeleteUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusNoContent))
			gatewayClient.AssertNumberOfCalls(GinkgoT(), "DeleteStorageSpace", 2) // 2 calls for the home space. first trash, then purge
		})
	})

	Describe("PatchUser", func() {
		var (
			user *libregraph.User
		)

		BeforeEach(func() {
			user = &libregraph.User{}
			user.SetDisplayName("Display Name")
			user.SetOnPremisesSamAccountName("user")
			user.SetMail("user@example.com")
			user.SetId("/users/user")

			identityBackend.On("GetUser", mock.Anything, mock.Anything, mock.Anything).Return(user, nil)
		})

		It("handles missing userids", func() {
			r := httptest.NewRequest(http.MethodPatch, "/graph/v1.0/users/{userid}", nil)
			svc.PatchUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("handles invalid bodies", func() {
			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users?$invalid=true", nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", user.GetId())
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.PatchUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("handles invalid email", func() {
			user.SetMail("invalid")
			data, err := json.Marshal(user)
			Expect(err).ToNot(HaveOccurred())

			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users?$invalid=true", bytes.NewBuffer(data))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", user.GetId())
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.PatchUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("handles invalid userType", func() {
			user.SetUserType("Clown")
			data, err := json.Marshal(user)
			Expect(err).ToNot(HaveOccurred())

			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users?$invalid=true", bytes.NewBuffer(data))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", user.GetId())
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.PatchUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusBadRequest))
		})

		It("updates attributes", func() {
			user.SetUserType("Member")
			identityBackend.On("UpdateUser", mock.Anything, user.GetId(), mock.Anything).Return(user, nil)

			user.SetUserType(("Member"))
			user.SetDisplayName("New Display Name")
			data, err := json.Marshal(user)
			Expect(err).ToNot(HaveOccurred())

			r := httptest.NewRequest(http.MethodPost, "/graph/v1.0/users", bytes.NewBuffer(data))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("userID", user.GetId())
			r = r.WithContext(context.WithValue(revactx.ContextSetUser(ctx, currentUser), chi.RouteCtxKey, rctx))
			svc.PatchUser(rr, r)

			Expect(rr.Code).To(Equal(http.StatusOK))
			data, err = io.ReadAll(rr.Body)
			Expect(err).ToNot(HaveOccurred())

			updatedUser := libregraph.User{}
			err = json.Unmarshal(data, &updatedUser)
			Expect(err).ToNot(HaveOccurred())
			Expect(updatedUser.GetUserType()).To(Equal("Member"))
			Expect(updatedUser.GetDisplayName()).To(Equal("New Display Name"))
		})
	})
})
