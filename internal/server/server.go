package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/odpf/compass/internal/server/health"
	handlersv1beta1 "github.com/odpf/compass/internal/server/v1beta1"
	"github.com/odpf/compass/internal/store/postgres"
	"github.com/odpf/compass/pkg/grpc_interceptor"
	"github.com/odpf/compass/pkg/statsd"
	compassv1beta1 "github.com/odpf/compass/proto/odpf/compass/v1beta1"
	"github.com/odpf/salt/log"
	"github.com/odpf/salt/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

type Config struct {
	Host     string `mapstructure:"host" default:"0.0.0.0"`
	Port     int    `mapstructure:"port" default:"8080"`
	GRPCPort int    `mapstructure:"grpc_port" default:"8081"`
	BaseUrl  string `mapstructure:"baseurl" default:"localhost:8080"`

	// User Identity
	Identity IdentityConfig `mapstructure:"identity"`
}

func (cfg Config) addr() string     { return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port) }
func (cfg Config) grpcAddr() string { return fmt.Sprintf("%s:%d", cfg.Host, cfg.GRPCPort) }

type IdentityConfig struct {
	// User Identity
	HeaderKeyUUID       string `mapstructure:"headerkey_uuid" default:"Compass-User-UUID"`
	HeaderValueUUID     string `mapstructure:"headervalue_uuid" default:"odpf@email.com"`
	HeaderKeyEmail      string `mapstructure:"headerkey_email" default:"Compass-User-Email"`
	ProviderDefaultName string `mapstructure:"provider_default_name" default:""`
}

func Serve(
	ctx context.Context,
	config Config,
	logger *log.Logrus,
	pgClient *postgres.Client,
	nrApp *newrelic.Application,
	statsdReporter *statsd.Reporter,
	assetService handlersv1beta1.AssetService,
	starService handlersv1beta1.StarService,
	discussionService handlersv1beta1.DiscussionService,
	tagService handlersv1beta1.TagService,
	tagTemplateService handlersv1beta1.TagTemplateService,
	userService handlersv1beta1.UserService,
) error {

	v1beta1Handler := handlersv1beta1.NewAPIServer(
		logger,
		assetService,
		starService,
		discussionService,
		tagService,
		tagTemplateService,
		userService,
	)

	healthHandler := health.NewHandler()

	// init grpc
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logger.Entry()),
			nrgrpc.UnaryServerInterceptor(nrApp),
			grpc_interceptor.StatsD(statsdReporter),
			grpc_interceptor.UserHeaderCtx(config.Identity.HeaderKeyUUID, config.Identity.HeaderKeyEmail),
		)),
	)
	reflection.Register(grpcServer)

	compassv1beta1.RegisterCompassServiceServer(grpcServer, v1beta1Handler)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthHandler)

	// init http proxy
	grpcDialCtx, grpcDialCancel := context.WithTimeout(ctx, time.Second*5)
	defer grpcDialCancel()

	headerMatcher := makeHeaderMatcher(config)

	grpcConn, err := grpc.DialContext(grpcDialCtx, config.grpcAddr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	runtimeCtx, runtimeCancel := context.WithCancel(ctx)
	defer runtimeCancel()

	gwmux := runtime.NewServeMux(
		runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
		runtime.WithIncomingHeaderMatcher(headerMatcher),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithHealthEndpointAt(grpc_health_v1.NewHealthClient(grpcConn), "/ping"),
	)

	if err := compassv1beta1.RegisterCompassServiceHandler(runtimeCtx, gwmux, grpcConn); err != nil {
		return err
	}

	defer func() {
		if pgClient != nil {
			logger.Warn("closing db...")
			if err := pgClient.Close(); err != nil {
				logger.Error("error when closing db", "err", err)
			}
			logger.Warn("db closed...")
		}
	}()

	logger.Info("Starting server", "http_port", config.addr(), "grpc_port", config.grpcAddr())
	if err := mux.Serve(
		ctx,
		mux.WithHTTPTarget(config.addr(), &http.Server{
			Handler:      gwmux,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		}),
		mux.WithGRPCTarget(config.grpcAddr(), grpcServer),
		mux.WithGracePeriod(5*time.Second),
	); !errors.Is(err, context.Canceled) {
		logger.Error("mux serve error", "err", err)
	}

	logger.Info("server stopped")

	return nil
}

func makeHeaderMatcher(c Config) func(key string) (string, bool) {
	return func(key string) (string, bool) {
		switch strings.ToLower(key) {
		case strings.ToLower(c.Identity.HeaderKeyUUID):
			return key, true
		case strings.ToLower(c.Identity.HeaderKeyEmail):
			return key, true
		default:
			return runtime.DefaultHeaderMatcher(key)
		}
	}
}
