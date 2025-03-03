package config

import (
	"context"

	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
)

// Config combines all available configuration parts.
type Config struct {
	Commons *shared.Commons `yaml:"-"` // don't use this directly as configuration for a service

	Service Service `yaml:"-"`

	Tracing *Tracing `yaml:"tracing"`
	Log     *Log     `yaml:"log"`
	Debug   Debug    `yaml:"debug"`

	GRPCClientTLS *shared.GRPCClientTLS `yaml:"grpc_client_tls"`

	HTTP HTTP `yaml:"http"`

	OcisPublicURL        string          `yaml:"ocis_public_url" env:"OCIS_URL;OCIS_PUBLIC_URL" desc:"URL, where oCIS is reachable for users."`
	WebdavNamespace      string          `yaml:"webdav_namespace" env:"WEBDAV_WEBDAV_NAMESPACE" desc:"CS3 path layout to use when forwarding /webdav requests"`
	RevaGateway          string          `yaml:"reva_gateway" env:"OCIS_REVA_GATEWAY;REVA_GATEWAY" desc:"CS3 gateway used to look up user metadata" deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"REVA_GATEWAY changing name for consistency" deprecationReplacement:"OCIS_REVA_GATEWAY"`
	RevaGatewayTLSMode   string          `yaml:"reva_gateway_tls_mode" env:"OCIS_REVA_GATEWAY_TLS_MODE;REVA_GATEWAY_TLS_MODE" desc:"TLS mode for grpc connection to the CS3 gateway endpoint. Possible values are 'off', 'insecure' and 'on'. 'off': disables transport security for the clients. 'insecure' allows using transport security, but disables certificate verification (to be used with the autogenerated self-signed certificates). 'on' enables transport security, including server certificate verification." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"REVA_GATEWAY_TLS_MODE changing name for consistency" deprecationReplacement:"OCIS_REVA_GATEWAY_TLS_MODE"`
	RevaGatewayTLSCACert string          `yaml:"reva_gateway_tls_cacert" env:"OCIS_REVA_GATEWAY_TLS_CACERT;REVA_GATEWAY_TLS_CACERT" desc:"The root CA certificate used to validate the gateway's TLS certificate." deprecationVersion:"3.0" removalVersion:"4.0.0" deprecationInfo:"REVA_GATEWAY_TLS_CACERT changing name for consistency" deprecationReplacement:"OCIS_REVA_GATEWAY_TLS_CACERT"`
	Context              context.Context `yaml:"-"`
}
