package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
)

type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service
	Service Service         `yaml:"-"`
	Tracing *Tracing        `yaml:"tracing"`
	Log     *Log            `yaml:"log"`
	Debug   Debug           `yaml:"debug"`

	GRPC GRPCConfig `yaml:"grpc"`

	TokenManager *TokenManager `yaml:"token_manager"`
	Reva         *shared.Reva  `yaml:"reva"`

	SkipUserGroupsInToken bool `yaml:"skip_user_groups_in_token" env:"USERS_SKIP_USER_GROUPS_IN_TOKEN" desc:"Disables the loading of user's group memberships from the reva access token."`

	Driver  string  `yaml:"driver" env:"USERS_DRIVER" desc:"The driver which should be used by the users service. Supported values are 'ldap' and 'owncloudsql'."`
	Drivers Drivers `yaml:"drivers"`

	Supervised bool            `yaml:"-"`
	Context    context.Context `yaml:"-"`
}
type Tracing struct {
	Enabled   bool   `yaml:"enabled" env:"OCIS_TRACING_ENABLED;USERS_TRACING_ENABLED" desc:"Activates tracing."`
	Type      string `yaml:"type" env:"OCIS_TRACING_TYPE;USERS_TRACING_TYPE" desc:"The type of tracing. Defaults to \"\", which is the same as \"jaeger\". Allowed tracing types are \"jaeger\" and \"\" as of now."`
	Endpoint  string `yaml:"endpoint" env:"OCIS_TRACING_ENDPOINT;USERS_TRACING_ENDPOINT" desc:"The endpoint of the tracing agent."`
	Collector string `yaml:"collector" env:"OCIS_TRACING_COLLECTOR;USERS_TRACING_COLLECTOR" desc:"The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces. Only used if the tracing endpoint is unset."`
}

type Log struct {
	Level  string `yaml:"level" env:"OCIS_LOG_LEVEL;USERS_LOG_LEVEL" desc:"The log level. Valid values are: \"panic\", \"fatal\", \"error\", \"warn\", \"info\", \"debug\", \"trace\"."`
	Pretty bool   `yaml:"pretty" env:"OCIS_LOG_PRETTY;USERS_LOG_PRETTY" desc:"Activates pretty log output."`
	Color  bool   `yaml:"color" env:"OCIS_LOG_COLOR;USERS_LOG_COLOR" desc:"Activates colorized log output."`
	File   string `yaml:"file" env:"OCIS_LOG_FILE;USERS_LOG_FILE" desc:"The path to the log file. Activates logging to this file if set."`
}

type Service struct {
	Name string `yaml:"-"`
}

type Debug struct {
	Addr   string `yaml:"addr" env:"USERS_DEBUG_ADDR" desc:"Bind address of the debug server, where metrics, health, config and debug endpoints will be exposed."`
	Token  string `yaml:"token" env:"USERS_DEBUG_TOKEN" desc:"Token to secure the metrics endpoint."`
	Pprof  bool   `yaml:"pprof" env:"USERS_DEBUG_PPROF" desc:"Enables pprof, which can be used for profiling."`
	Zpages bool   `yaml:"zpages" env:"USERS_DEBUG_ZPAGES" desc:"Enables zpages, which can be used for collecting and viewing in-memory traces."`
}

type GRPCConfig struct {
	Addr      string                 `yaml:"addr" env:"USERS_GRPC_ADDR" desc:"The bind address of the GRPC service."`
	TLS       *shared.GRPCServiceTLS `yaml:"tls"`
	Namespace string                 `yaml:"-"`
	Protocol  string                 `yaml:"protocol" env:"USERS_GRPC_PROTOCOL" desc:"The transport protocol of the GPRC service."`
}

type Drivers struct {
	LDAP        LDAPDriver        `yaml:"ldap"`
	OwnCloudSQL OwnCloudSQLDriver `yaml:"owncloudsql"`

	JSON JSONDriver   `yaml:"json,omitempty"` // not supported by the oCIS product, therefore not part of docs
	REST RESTProvider `yaml:"rest,omitempty"` // not supported by the oCIS product, therefore not part of docs
}

type JSONDriver struct {
	File string `yaml:"file"`
}
type LDAPDriver struct {
	URI                      string          `yaml:"uri" env:"OCIS_LDAP_URI;LDAP_URI;USERS_LDAP_URI" desc:"URI of the LDAP Server to connect to. Supported URI schemes are 'ldaps://' and 'ldap://'" deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_URI changing name for consistency" deprecationReplacement:"OCIS_LDAP_URI"`
	CACert                   string          `yaml:"ca_cert" env:"OCIS_LDAP_CACERT;LDAP_CACERT;USERS_LDAP_CACERT" desc:"Path/File name for the root CA certificate (in PEM format) used to validate TLS server certificates of the LDAP service. If not defined, the root directory derives from $OCIS_BASE_DATA_PATH:/idm." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_CACERT changing name for consistency" deprecationReplacement:"OCIS_LDAP_CACERT"`
	Insecure                 bool            `yaml:"insecure" env:"OCIS_LDAP_INSECURE;LDAP_INSECURE;USERS_LDAP_INSECURE" desc:"Disable TLS certificate validation for the LDAP connections. Do not set this in production environments." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_INSECURE changing name for consistency" deprecationReplacement:"OCIS_LDAP_INSECURE"`
	BindDN                   string          `yaml:"bind_dn" env:"OCIS_LDAP_BIND_DN;LDAP_BIND_DN;USERS_LDAP_BIND_DN" desc:"LDAP DN to use for simple bind authentication with the target LDAP server." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_BIND_DN changing name for consistency" deprecationReplacement:"OCIS_LDAP_BIND_DN"`
	BindPassword             string          `yaml:"bind_password" env:"LDAP_BIND_PASSWORD;USERS_LDAP_BIND_PASSWORD" desc:"Password to use for authenticating the 'bind_dn'."`
	UserBaseDN               string          `yaml:"user_base_dn" env:"OCIS_LDAP_USER_BASE_DN;LDAP_USER_BASE_DN;USERS_LDAP_USER_BASE_DN" desc:"Search base DN for looking up LDAP users." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_BASE_DN changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_BASE_DN"`
	GroupBaseDN              string          `yaml:"group_base_dn" env:"OCIS_LDAP_GROUP_BASE_DN;LDAP_GROUP_BASE_DN;USERS_LDAP_GROUP_BASE_DN" desc:"Search base DN for looking up LDAP groups." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_BASE_DN changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_BASE_DN"`
	UserScope                string          `yaml:"user_scope" env:"OCIS_LDAP_USER_SCOPE;LDAP_USER_SCOPE;USERS_LDAP_USER_SCOPE" desc:"LDAP search scope to use when looking up users. Supported values are 'base', 'one' and 'sub'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCOPE changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCOPE"`
	GroupScope               string          `yaml:"group_scope" env:"OCIS_LDAP_GROUP_SCOPE;LDAP_GROUP_SCOPE;USERS_LDAP_GROUP_SCOPE" desc:"LDAP search scope to use when looking up groups. Supported values are 'base', 'one' and 'sub'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCOPE changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCOPE"`
	UserSubstringFilterType  string          `yaml:"user_substring_filter_type" env:"LDAP_USER_SUBSTRING_FILTER_TYPE;USERS_LDAP_USER_SUBSTRING_FILTER_TYPE" desc:"Type of substring search filter to use for substring searches for users. Possible values: 'initial' for doing prefix only searches, 'final' for doing suffix only searches or 'any' for doing full substring searches"`
	UserFilter               string          `yaml:"user_filter" env:"OCIS_LDAP_USER_FILTER;LDAP_USER_FILTER;USERS_LDAP_USER_FILTER" desc:"LDAP filter to add to the default filters for user search like '(objectclass=ownCloud)'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_FILTER changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_FILTER"`
	GroupFilter              string          `yaml:"group_filter" env:"OCIS_LDAP_GROUP_FILTER;LDAP_GROUP_FILTER;USERS_LDAP_GROUP_FILTER" desc:"LDAP filter to add to the default filters for group searches." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_FILTER changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_FILTER"`
	UserObjectClass          string          `yaml:"user_object_class" env:"OCIS_LDAP_USER_OBJECTCLASS;LDAP_USER_OBJECTCLASS;USERS_LDAP_USER_OBJECTCLASS" desc:"The object class to use for users in the default user search filter like 'inetOrgPerson'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_OBJECTCLASS changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_OBJECTCLASS"`
	GroupObjectClass         string          `yaml:"group_object_class" env:"OCIS_LDAP_GROUP_OBJECTCLASS;LDAP_GROUP_OBJECTCLASS;USERS_LDAP_GROUP_OBJECTCLASS" desc:"The object class to use for groups in the default group search filter like 'groupOfNames'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_OBJECTCLASS changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_OBJECTCLASS"`
	IDP                      string          `yaml:"idp" env:"OCIS_URL;OCIS_OIDC_ISSUER;USERS_IDP_URL" desc:"The identity provider value to set in the userids of the CS3 user objects for users returned by this user provider."`
	DisableUserMechanism     string          `yaml:"disable_user_mechanism" env:"OCIS_LDAP_DISABLE_USER_MECHANISM;LDAP_DISABLE_USER_MECHANISM;USERS_LDAP_DISABLE_USER_MECHANISM" desc:"An option to control the behavior for disabling users. Valid options are 'none', 'attribute' and 'group'. If set to 'group', disabling a user via API will add the user to the configured group for disabled users, if set to 'attribute' this will be done in the ldap user entry, if set to 'none' the disable request is not processed." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_DISABLE_USER_MECHANISM changing name for consistency" deprecationReplacement:"OCIS_LDAP_DISABLE_USER_MECHANISM"`
	UserTypeAttribute        string          `yaml:"user_type_attribute" env:"OCIS_LDAP_USER_SCHEMA_USER_TYPE;LDAP_USER_SCHEMA_USER_TYPE;USERS_LDAP_USER_TYPE_ATTRIBUTE" desc:"LDAP Attribute to distinguish between 'Member' and 'Guest' users. Default is 'ownCloudUserType'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCHEMA_USER_TYPE changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCHEMA_USER_TYPE"`
	LdapDisabledUsersGroupDN string          `yaml:"ldap_disabled_users_group_dn" env:"OCIS_LDAP_DISABLED_USERS_GROUP_DN;LDAP_DISABLED_USERS_GROUP_DN;USERS_LDAP_DISABLED_USERS_GROUP_DN" desc:"The distinguished name of the group to which added users will be classified as disabled when 'disable_user_mechanism' is set to 'group'." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_DISABLED_USERS_GROUP_DN changing name for consistency" deprecationReplacement:"OCIS_LDAP_DISABLED_USERS_GROUP_DN"`
	UserSchema               LDAPUserSchema  `yaml:"user_schema"`
	GroupSchema              LDAPGroupSchema `yaml:"group_schema"`
}

type LDAPUserSchema struct {
	ID              string `yaml:"id" env:"OCIS_LDAP_USER_SCHEMA_ID;LDAP_USER_SCHEMA_ID;USERS_LDAP_USER_SCHEMA_ID" desc:"LDAP Attribute to use as the unique ID for users. This should be a stable globally unique ID like a UUID." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCHEMA_ID changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCHEMA_ID"`
	IDIsOctetString bool   `yaml:"id_is_octet_string" env:"OCIS_LDAP_USER_SCHEMA_ID_IS_OCTETSTRING;LDAP_USER_SCHEMA_ID_IS_OCTETSTRING;USERS_LDAP_USER_SCHEMA_ID_IS_OCTETSTRING" desc:"Set this to true if the defined 'ID' attribute for users is of the 'OCTETSTRING' syntax. This is e.g. required when using the 'objectGUID' attribute of Active Directory for the user ID's." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCHEMA_ID_IS_OCTETSTRING changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCHEMA_ID_IS_OCTETSTRING"`
	Mail            string `yaml:"mail" env:"OCIS_LDAP_USER_SCHEMA_MAIL;LDAP_USER_SCHEMA_MAIL;USERS_LDAP_USER_SCHEMA_MAIL" desc:"LDAP Attribute to use for the email address of users." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCHEMA_MAIL changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCHEMA_MAIL"`
	DisplayName     string `yaml:"display_name" env:"OCIS_LDAP_USER_SCHEMA_DISPLAYNAME;LDAP_USER_SCHEMA_DISPLAYNAME;USERS_LDAP_USER_SCHEMA_DISPLAYNAME" desc:"LDAP Attribute to use for the displayname of users." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCHEMA_DISPLAYNAME changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCHEMA_DISPLAYNAME"`
	Username        string `yaml:"user_name" env:"OCIS_LDAP_USER_SCHEMA_USERNAME;LDAP_USER_SCHEMA_USERNAME;USERS_LDAP_USER_SCHEMA_USERNAME" desc:"LDAP Attribute to use for username of users." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_SCHEMA_USERNAME changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_SCHEMA_USERNAME"`
	Enabled         string `yaml:"user_enabled" env:"OCIS_LDAP_USER_ENABLED_ATTRIBUTE;LDAP_USER_ENABLED_ATTRIBUTE;USERS_LDAP_USER_ENABLED_ATTRIBUTE" desc:"LDAP attribute to use as a flag telling if the user is enabled or disabled." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_USER_ENABLED_ATTRIBUTE changing name for consistency" deprecationReplacement:"OCIS_LDAP_USER_ENABLED_ATTRIBUTE"`
}

type LDAPGroupSchema struct {
	ID              string `yaml:"id" env:"OCIS_LDAP_GROUP_SCHEMA_ID;LDAP_GROUP_SCHEMA_ID;USERS_LDAP_GROUP_SCHEMA_ID" desc:"LDAP Attribute to use as the unique ID for groups. This should be a stable globally unique ID like a UUID." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCHEMA_ID changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCHEMA_ID"`
	IDIsOctetString bool   `yaml:"id_is_octet_string" env:"OCIS_LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING;LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING;USERS_LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING" desc:"Set this to true if the defined 'id' attribute for groups is of the 'OCTETSTRING' syntax. This is e.g. required when using the 'objectGUID' attribute of Active Directory for the group ID's." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCHEMA_ID_IS_OCTETSTRING"`
	Mail            string `yaml:"mail" env:"OCIS_LDAP_GROUP_SCHEMA_MAIL;LDAP_GROUP_SCHEMA_MAIL;USERS_LDAP_GROUP_SCHEMA_MAIL" desc:"LDAP Attribute to use for the email address of groups (can be empty)." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCHEMA_MAIL changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCHEMA_MAIL"`
	DisplayName     string `yaml:"display_name" env:"OCIS_LDAP_GROUP_SCHEMA_DISPLAYNAME;LDAP_GROUP_SCHEMA_DISPLAYNAME;USERS_LDAP_GROUP_SCHEMA_DISPLAYNAME" desc:"LDAP Attribute to use for the displayname of groups (often the same as groupname attribute)." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCHEMA_DISPLAYNAME changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCHEMA_DISPLAYNAME"`
	Groupname       string `yaml:"group_name" env:"OCIS_LDAP_GROUP_SCHEMA_GROUPNAME;LDAP_GROUP_SCHEMA_GROUPNAME;USERS_LDAP_GROUP_SCHEMA_GROUPNAME" desc:"LDAP Attribute to use for the name of groups." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCHEMA_GROUPNAME changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCHEMA_GROUPNAME"`
	Member          string `yaml:"member" env:"OCIS_LDAP_GROUP_SCHEMA_MEMBER;LDAP_GROUP_SCHEMA_MEMBER;USERS_LDAP_GROUP_SCHEMA_MEMBER" desc:"LDAP Attribute that is used for group members." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"LDAP_GROUP_SCHEMA_MEMBER changing name for consistency" deprecationReplacement:"OCIS_LDAP_GROUP_SCHEMA_MEMBER"`
}

type OwnCloudSQLDriver struct {
	DBUsername         string `yaml:"db_username" env:"USERS_OWNCLOUDSQL_DB_USERNAME" desc:"Database user to use for authenticating with the owncloud database."`
	DBPassword         string `yaml:"db_password" env:"USERS_OWNCLOUDSQL_DB_PASSWORD" desc:"Password for the database user."`
	DBHost             string `yaml:"db_host" env:"USERS_OWNCLOUDSQL_DB_HOST" desc:"Hostname of the database server."`
	DBPort             int    `yaml:"db_port" env:"USERS_OWNCLOUDSQL_DB_PORT" desc:"Network port to use for the database connection."`
	DBName             string `yaml:"db_name" env:"USERS_OWNCLOUDSQL_DB_NAME" desc:"Name of the owncloud database."`
	IDP                string `yaml:"idp" env:"USERS_OWNCLOUDSQL_IDP" desc:"The identity provider value to set in the userids of the CS3 user objects for users returned by this user provider."`
	Nobody             int64  `yaml:"nobody" env:"USERS_OWNCLOUDSQL_NOBODY" desc:"Fallback number if no numeric UID and GID properties are provided."`
	JoinUsername       bool   `yaml:"join_username" env:"USERS_OWNCLOUDSQL_JOIN_USERNAME" desc:"Join the user properties table to read usernames"`
	JoinOwnCloudUUID   bool   `yaml:"join_owncloud_uuid" env:"USERS_OWNCLOUDSQL_JOIN_OWNCLOUD_UUID" desc:"Join the user properties table to read user IDs."`
	EnableMedialSearch bool   `yaml:"enable_medial_search" env:"USERS_OWNCLOUDSQL_ENABLE_MEDIAL_SEARCH" desc:"Allow 'medial search' when searching for users instead of just doing a prefix search. This allows finding 'Alice' when searching for 'lic'."`
}
type RESTProvider struct {
	ClientID          string
	ClientSecret      string
	RedisAddr         string
	RedisUsername     string
	RedisPassword     string
	IDProvider        string
	APIBaseURL        string
	OIDCTokenEndpoint string
	TargetAPI         string
}
