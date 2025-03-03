package defaults

import (
	"path/filepath"
	"time"

	"github.com/owncloud/ocis/v2/ocis-pkg/config/defaults"
	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	"github.com/owncloud/ocis/v2/ocis-pkg/structs"
	"github.com/owncloud/ocis/v2/services/storage-users/pkg/config"
)

// FullDefaultConfig returns a fully initialized default configuration
func FullDefaultConfig() *config.Config {
	cfg := DefaultConfig()
	EnsureDefaults(cfg)
	Sanitize(cfg)
	return cfg
}

// DefaultConfig returns a basic default configuration
func DefaultConfig() *config.Config {
	return &config.Config{
		Debug: config.Debug{
			Addr:   "127.0.0.1:9159",
			Token:  "",
			Pprof:  false,
			Zpages: false,
		},
		GRPC: config.GRPCConfig{
			Addr:      "127.0.0.1:9157",
			Namespace: "com.owncloud.api",
			Protocol:  "tcp",
		},
		HTTP: config.HTTPConfig{
			Addr:      "127.0.0.1:9158",
			Namespace: "com.owncloud.web",
			Protocol:  "tcp",
			Prefix:    "data",
		},
		Service: config.Service{
			Name: "storage-users",
		},
		Reva:             shared.DefaultRevaConfig(),
		DataServerURL:    "http://localhost:9158/data",
		DataGatewayURL:   "https://localhost:9200/data",
		TransferExpires:  86400,
		UploadExpiration: 24 * 60 * 60,
		Driver:           "ocis",
		Drivers: config.Drivers{
			OwnCloudSQL: config.OwnCloudSQLDriver{
				Root:                  filepath.Join(defaults.BaseDataPath(), "storage", "owncloud"),
				ShareFolder:           "/Shares",
				UserLayout:            "{{.Username}}",
				UploadInfoDir:         filepath.Join(defaults.BaseDataPath(), "storage", "uploadinfo"),
				DBUsername:            "owncloud",
				DBPassword:            "owncloud",
				DBHost:                "",
				DBPort:                3306,
				DBName:                "owncloud",
				UsersProviderEndpoint: "com.owncloud.api.users",
			},
			S3NG: config.S3NGDriver{
				MetadataBackend:            "messagepack",
				Root:                       filepath.Join(defaults.BaseDataPath(), "storage", "users"),
				ShareFolder:                "/Shares",
				UserLayout:                 "{{.Id.OpaqueId}}",
				Region:                     "default",
				PersonalSpaceAliasTemplate: "{{.SpaceType}}/{{.User.Username | lower}}",
				GeneralSpaceAliasTemplate:  "{{.SpaceType}}/{{.SpaceName | replace \" \" \"-\" | lower}}",
				PermissionsEndpoint:        "com.owncloud.api.settings",
				MaxAcquireLockCycles:       20,
				LockCycleDurationFactor:    30,
			},
			OCIS: config.OCISDriver{
				MetadataBackend:            "messagepack",
				Root:                       filepath.Join(defaults.BaseDataPath(), "storage", "users"),
				ShareFolder:                "/Shares",
				UserLayout:                 "{{.Id.OpaqueId}}",
				PersonalSpaceAliasTemplate: "{{.SpaceType}}/{{.User.Username | lower}}",
				GeneralSpaceAliasTemplate:  "{{.SpaceType}}/{{.SpaceName | replace \" \" \"-\" | lower}}",
				PermissionsEndpoint:        "com.owncloud.api.settings",
				MaxAcquireLockCycles:       20,
				LockCycleDurationFactor:    30,
			},
		},
		Events: config.Events{
			Addr:      "127.0.0.1:9233",
			ClusterID: "ocis-cluster",
			EnableTLS: false,
		},
		StatCache: config.StatCache{
			Store:    "noop",
			Database: "ocis",
			TTL:      300 * time.Second,
		},
		FilemetadataCache: config.FilemetadataCache{
			Store:    "memory",
			Database: "ocis",
			TTL:      24 * 60 * time.Second,
		},
		Tasks: config.Tasks{
			PurgeTrashBin: config.PurgeTrashBin{
				ProjectDeleteBefore:  30 * 24 * time.Hour,
				PersonalDeleteBefore: 30 * 24 * time.Hour,
			},
		},
	}
}

// EnsureDefaults adds default values to the configuration if they are not set yet
func EnsureDefaults(cfg *config.Config) {
	// provide with defaults for shared logging, since we need a valid destination address for "envdecode".
	if cfg.Log == nil && cfg.Commons != nil && cfg.Commons.Log != nil {
		cfg.Log = &config.Log{
			Level:  cfg.Commons.Log.Level,
			Pretty: cfg.Commons.Log.Pretty,
			Color:  cfg.Commons.Log.Color,
			File:   cfg.Commons.Log.File,
		}
	} else if cfg.Log == nil {
		cfg.Log = &config.Log{}
	}
	// provide with defaults for shared tracing, since we need a valid destination address for "envdecode".
	if cfg.Tracing == nil && cfg.Commons != nil && cfg.Commons.Tracing != nil {
		cfg.Tracing = &config.Tracing{
			Enabled:   cfg.Commons.Tracing.Enabled,
			Type:      cfg.Commons.Tracing.Type,
			Endpoint:  cfg.Commons.Tracing.Endpoint,
			Collector: cfg.Commons.Tracing.Collector,
		}
	} else if cfg.Tracing == nil {
		cfg.Tracing = &config.Tracing{}
	}

	if cfg.Reva == nil && cfg.Commons != nil {
		cfg.Reva = structs.CopyOrZeroValue(cfg.Commons.Reva)
	}

	if cfg.TokenManager == nil && cfg.Commons != nil && cfg.Commons.TokenManager != nil {
		cfg.TokenManager = &config.TokenManager{
			JWTSecret: cfg.Commons.TokenManager.JWTSecret,
		}
	} else if cfg.TokenManager == nil {
		cfg.TokenManager = &config.TokenManager{}
	}

	if cfg.GRPC.TLS == nil && cfg.Commons != nil {
		cfg.GRPC.TLS = structs.CopyOrZeroValue(cfg.Commons.GRPCServiceTLS)
	}

	if cfg.Tasks.PurgeTrashBin.UserID == "" && cfg.Commons != nil {
		cfg.Tasks.PurgeTrashBin.UserID = cfg.Commons.AdminUserID
	}
}

// Sanitize sanitized the configuration
func Sanitize(cfg *config.Config) {
	// nothing to sanitize here atm
}
