package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
	"sync"

	mgrpcc "github.com/go-micro/plugins/v4/client/grpc"
	mbreaker "github.com/go-micro/plugins/v4/wrapper/breaker/gobreaker"
	mtracer "github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	"github.com/owncloud/ocis/v2/ocis-pkg/registry"
	"github.com/owncloud/ocis/v2/ocis-pkg/shared"
	"go-micro.dev/v4/client"
	"go.opentelemetry.io/otel/trace"
)

var (
	defaultClient client.Client
	once          sync.Once
)

// ClientOptions represent options (e.g. tls settings) for the grpc clients
type ClientOptions struct {
	tlsMode string
	caCert  string
	tp      trace.TracerProvider
}

// Option is used to pass client options
type ClientOption func(opts *ClientOptions)

// WithTLSMode allows setting the TLSMode option for grpc clients
func WithTLSMode(v string) ClientOption {
	return func(o *ClientOptions) {
		o.tlsMode = v
	}
}

// WithTLSCACert allows setting the CA Certificate for grpc clients
func WithTLSCACert(v string) ClientOption {
	return func(o *ClientOptions) {
		o.caCert = v
	}
}

// WithTraceProvider allows to set the trace Provider for grpc clients
func WithTraceProvider(tp trace.TracerProvider) ClientOption {
	return func(o *ClientOptions) {
		o.tp = tp
	}
}

// Configure configures the default oOCIS grpc client (e.g. TLS settings)
func Configure(opts ...ClientOption) error {
	var options ClientOptions
	for _, opt := range opts {
		opt(&options)
	}
	var outerr error
	once.Do(func() {
		reg := registry.GetRegistry()
		var tlsConfig *tls.Config
		cOpts := []client.Option{
			client.Registry(reg),
			client.Wrap(mbreaker.NewClientWrapper()),
			client.Wrap(mtracer.NewClientWrapper(
				mtracer.WithTraceProvider(options.tp),
			)),
		}
		switch options.tlsMode {
		case "insecure":
			tlsConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
			cOpts = append(cOpts, mgrpcc.AuthTLS(tlsConfig))
		case "on":
			tlsConfig = &tls.Config{}
			// Note: If caCert is empty we use the system's default set of trusted CAs
			if options.caCert != "" {
				certs := x509.NewCertPool()
				pemData, err := os.ReadFile(options.caCert)
				if err != nil {
					outerr = err
					return
				}
				if !certs.AppendCertsFromPEM(pemData) {
					outerr = errors.New("could not initialize default client, adding CA cert failed")
					return
				}
				tlsConfig.RootCAs = certs
			}
			cOpts = append(cOpts, mgrpcc.AuthTLS(tlsConfig))
			//case "off":
			//default:
		}

		defaultClient = mgrpcc.NewClient(cOpts...)
	})
	return outerr
}

// DefaultClient returns a custom oCIS grpc configured client.
func DefaultClient() client.Client {
	return defaultClient
}

func GetClientOptions(t *shared.GRPCClientTLS) []ClientOption {
	opts := []ClientOption{
		WithTLSMode(t.Mode),
		WithTLSCACert(t.CACert),
	}
	return opts
}

func NewClient(opts ...ClientOption) (client.Client, error) {
	var options ClientOptions
	for _, opt := range opts {
		opt(&options)
	}

	reg := registry.GetRegistry()
	var tlsConfig *tls.Config
	cOpts := []client.Option{
		client.Registry(reg),
		client.Wrap(mbreaker.NewClientWrapper()),
		client.Wrap(mtracer.NewClientWrapper(
			mtracer.WithTraceProvider(options.tp),
		)),
	}
	switch options.tlsMode {
	case "insecure":
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
		cOpts = append(cOpts, mgrpcc.AuthTLS(tlsConfig))
	case "on":
		tlsConfig = &tls.Config{}
		// Note: If caCert is empty we use the system's default set of trusted CAs
		if options.caCert != "" {
			certs := x509.NewCertPool()
			pemData, err := os.ReadFile(options.caCert)
			if err != nil {
				return nil, err
			}
			if !certs.AppendCertsFromPEM(pemData) {
				return nil, errors.New("could not initialize client, adding CA cert failed")
			}
			tlsConfig.RootCAs = certs
		}
		cOpts = append(cOpts, mgrpcc.AuthTLS(tlsConfig))
		//case "off":
		//default:
	}

	return mgrpcc.NewClient(cOpts...), nil
}
