package internal

import (
	"flag"
	"time"

	"github.com/go-logr/logr"
	"github.com/kyverno/kyverno/pkg/leaderelection"
	"github.com/kyverno/kyverno/pkg/logging"
	"github.com/kyverno/kyverno/pkg/toggle"
	"github.com/sigstore/sigstore/pkg/tuf"
)

var (
	// logging
	loggingFormat   string
	loggingTsFormat string
	disableLogColor bool
	// profiling
	profilingEnabled bool
	profilingAddress string
	profilingPort    string
	// tracing
	tracingEnabled bool
	tracingAddress string
	tracingPort    string
	tracingCreds   string
	// metrics
	otel                 string
	otelCollector        string
	metricsPort          string
	transportCreds       string
	disableMetricsExport bool
	// kubeconfig
	kubeconfig           string
	clientRateLimitQPS   float64
	clientRateLimitBurst int
	eventsRateLimitQPS   float64
	eventsRateLimitBurst int
	// engine
	enablePolicyException  bool
	exceptionNamespace     string
	enableConfigMapCaching bool
	openreportsEnabled     bool
	// cosign
	enableTUF  bool
	tufMirror  string
	tufRoot    string
	tufRootRaw string
	// registry client
	imagePullSecrets          string
	allowInsecureRegistry     bool
	registryCredentialHelpers string
	// leader election
	leaderElectionRetryPeriod time.Duration
	// cleanupServerPort is the kyverno cleanup server port
	cleanupServerPort string
	// image verify cache
	imageVerifyCacheEnabled     bool
	imageVerifyCacheTTLDuration time.Duration
	imageVerifyCacheMaxSize     int64
	// global context
	enableGlobalContext bool
	// reporting
	enableReporting string
	// resync
	resyncPeriod time.Duration
	// custom resource watch
	crdWatcher bool
)

func initLoggingFlags() {
	logging.InitFlags(nil)
	flag.StringVar(&loggingFormat, "loggingFormat", logging.TextFormat, "This determines the output format of the logger.")
	flag.BoolVar(&disableLogColor, "disableLogColor", false, "Disable colored output in logs.")
	flag.StringVar(&loggingTsFormat, "loggingtsFormat", logging.DefaultTime, "This determines the timestamp format of the logger.")
	checkErr(flag.Set("v", "2"), "failed to init flags")
}

func initProfilingFlags() {
	flag.BoolVar(&profilingEnabled, "profile", false, "Set this flag to 'true', to enable profiling.")
	flag.StringVar(&profilingPort, "profilePort", "6060", "Profiling server port, defaults to '6060'.")
	flag.StringVar(&profilingAddress, "profileAddress", "", "Profiling server address, defaults to ''.")
}

func initTracingFlags() {
	flag.BoolVar(&tracingEnabled, "enableTracing", false, "Set this flag to 'true', to enable tracing.")
	flag.StringVar(&tracingPort, "tracingPort", "4317", "Tracing receiver port, defaults to '4317'.")
	flag.StringVar(&tracingAddress, "tracingAddress", "", "Tracing receiver address, defaults to ''.")
	flag.StringVar(&tracingCreds, "tracingCreds", "", "Set this flag to the CA secret containing the certificate which is used by our Opentelemetry Tracing Client. If empty string is set, means an insecure connection will be used")
}

func initMetricsFlags() {
	flag.StringVar(&otel, "otelConfig", "prometheus", "Set this flag to 'grpc', to enable exporting metrics to an Opentelemetry Collector. The default collector is set to \"prometheus\"")
	flag.StringVar(&otelCollector, "otelCollector", "opentelemetrycollector.kyverno.svc.cluster.local", "Set this flag to the OpenTelemetry Collector Service Address. Kyverno will try to connect to this on the metrics port.")
	flag.StringVar(&transportCreds, "transportCreds", "", "Set this flag to the CA secret containing the certificate which is used by our Opentelemetry Metrics Client. If empty string is set, means an insecure connection will be used")
	flag.StringVar(&metricsPort, "metricsPort", "8000", "Expose prometheus metrics at the given port, default to 8000.")
	flag.BoolVar(&disableMetricsExport, "disableMetrics", false, "Set this flag to 'true' to disable metrics.")
}

func initKubeconfigFlags(qps float64, burst int, eventsQPS float64, eventsBurst int) {
	if f := flag.Lookup("kubeconfig"); f == nil {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	}
	flag.Float64Var(&clientRateLimitQPS, "clientRateLimitQPS", qps, "Configure the maximum QPS to the Kubernetes API server from Kyverno. Uses the client default if zero.")
	flag.IntVar(&clientRateLimitBurst, "clientRateLimitBurst", burst, "Configure the maximum burst for throttle. Uses the client default if zero.")
	flag.Float64Var(&eventsRateLimitQPS, "eventsRateLimitQPS", eventsQPS, "Configure the maximum QPS to the Kubernetes API server from Kyverno for events. Uses the client default if zero.")
	flag.IntVar(&eventsRateLimitBurst, "eventsRateLimitBurst", eventsBurst, "Configure the maximum burst for throttle for events. Uses the client default if zero.")
	flag.DurationVar(&resyncPeriod, "resyncPeriod", 15*time.Minute, "Configure the resync period for informer factory")
	flag.BoolVar(&crdWatcher, "crdWatcher", false, "Enable watching of custom resources to invalidate discovery cache on changes.")
}

func initPolicyExceptionsFlags() {
	flag.StringVar(&exceptionNamespace, "exceptionNamespace", "", "Configure the namespace to accept PolicyExceptions. If it is set to '*', exceptions are allowed in all namespaces.")
	flag.BoolVar(&enablePolicyException, "enablePolicyException", false, "Enable PolicyException feature.")
}

func initConfigMapCachingFlags() {
	flag.BoolVar(&enableConfigMapCaching, "enableConfigMapCaching", true, "Enable config maps caching.")
}

func initDeferredLoadingFlags() {
	flag.Func(toggle.EnableDeferredLoadingFlagName, toggle.EnableDeferredLoadingDescription, toggle.EnableDeferredLoading.Parse)
}

func initCosignFlags() {
	flag.BoolVar(&enableTUF, "enableTuf", false, "enable tuf for private sigstore deployments")
	flag.StringVar(&tufMirror, "tufMirror", tuf.DefaultRemoteRoot, "Alternate TUF mirror for sigstore. If left blank, public sigstore one is used for cosign verification.")
	flag.StringVar(&tufRoot, "tufRoot", "", "Path to alternate TUF root.json for sigstore (url or env). If left blank, public sigstore one is used for cosign verification.")
	flag.StringVar(&tufRootRaw, "tufRootRaw", "", "The raw body of alternate TUF root.json for sigstore. If left blank, public sigstore one is used for cosign verification.")
}

func initRegistryClientFlags() {
	flag.BoolVar(&allowInsecureRegistry, "allowInsecureRegistry", false, "Whether to allow insecure connections to registries. Don't use this for anything but testing.")
	flag.StringVar(&imagePullSecrets, "imagePullSecrets", "", "Secret resource names for image registry access credentials.")
	flag.StringVar(&registryCredentialHelpers, "registryCredentialHelpers", "", "Credential helpers to enable (default,google,amazon,azure,github). No helpers are added when this flag is empty.")
}

func initImageVerifyCacheFlags() {
	flag.BoolVar(&imageVerifyCacheEnabled, "imageVerifyCacheEnabled", true, "Enable a TTL cache for verified images.")
	flag.Int64Var(&imageVerifyCacheMaxSize, "imageVerifyCacheMaxSize", 1000, "Maximum number of keys that can be stored in the TTL cache. Keys are a combination of policy elements along with the image reference. Default is 1000. 0 sets the value to default.")
	flag.DurationVar(&imageVerifyCacheTTLDuration, "imageVerifyCacheTTLDuration", 60*time.Minute, "Maximum TTL value for a cache expressed as duration. Default is 60m. 0 sets the value to default.")
}

func initLeaderElectionFlags() {
	flag.DurationVar(&leaderElectionRetryPeriod, "leaderElectionRetryPeriod", leaderelection.DefaultRetryPeriod, "Configure leader election retry period.")
}

func initCleanupFlags() {
	flag.StringVar(&cleanupServerPort, "cleanupServerPort", "9443", "kyverno cleanup server port, defaults to '9443'.")
}

func initReportingFlags() {
	flag.StringVar(&enableReporting, "enableReporting", "validate,mutate,mutateExisting,generate,imageVerify", "Comma separated list to enables reporting for different rule types. (validate,mutate,mutateExisting,generate,imageVerify)")
}

func initOpenreportsFlagSet() *flag.FlagSet {
	flagset := flag.NewFlagSet("openreports", flag.ExitOnError)
	flagset.BoolVar(&openreportsEnabled, "openreportsEnabled", false, "Use openreports.io/v1alpha1 for the reporting group")
	return flagset
}

func lookupKubeconfigFlag() {
	if f := flag.Lookup("kubeconfig"); f != nil {
		kubeconfig = f.Value.String()
	}
}

type options struct {
	clientRateLimitQPS   float64
	clientRateLimitBurst int
	eventsRateLimitQPS   float64
	eventsRateLimitBurst int
}

func newOptions() options {
	return options{
		clientRateLimitQPS:   100,
		clientRateLimitBurst: 200,
		eventsRateLimitQPS:   1000,
		eventsRateLimitBurst: 2000,
	}
}

type Option = func(*options)

func WithDefaultQps(qps float64) Option {
	return func(o *options) {
		o.clientRateLimitQPS = qps
	}
}

func WithDefaultBurst(burst int) Option {
	return func(o *options) {
		o.clientRateLimitBurst = burst
	}
}

func initFlags(config Configuration, opts ...Option) {
	options := newOptions()
	for _, o := range opts {
		if o != nil {
			o(&options)
		}
	}
	// logging
	initLoggingFlags()
	// profiling
	if config.UsesProfiling() {
		initProfilingFlags()
	}
	// tracing
	if config.UsesTracing() {
		initTracingFlags()
	}
	// metrics
	if config.UsesMetrics() {
		initMetricsFlags()
	}
	// kubeconfig
	if config.UsesKubeconfig() {
		initKubeconfigFlags(options.clientRateLimitQPS, options.clientRateLimitBurst, options.eventsRateLimitQPS, options.eventsRateLimitBurst)
	}
	// policy exceptions
	if config.UsesPolicyExceptions() {
		initPolicyExceptionsFlags()
	}
	// config map caching
	if config.UsesConfigMapCaching() {
		initConfigMapCachingFlags()
	}
	// deferred loading
	if config.UsesDeferredLoading() {
		initDeferredLoadingFlags()
	}
	// cosign
	if config.UsesCosign() {
		initCosignFlags()
	}
	// registry client
	if config.UsesRegistryClient() {
		initRegistryClientFlags()
	}
	// image verify cache
	if config.UsesImageVerifyCache() {
		initImageVerifyCacheFlags()
	}
	// leader election
	if config.UsesLeaderElection() {
		initLeaderElectionFlags()
	}
	// reporting
	if config.UsesReporting() {
		initReportingFlags()
	}

	if config.UsesOpenreports() {
		config.AddFlagSet(initOpenreportsFlagSet())
	}

	initCleanupFlags()
	for _, flagset := range config.FlagSets() {
		flagset.VisitAll(func(f *flag.Flag) {
			flag.CommandLine.Var(f.Value, f.Name, f.Usage)
		})
	}
}

func showWarnings(config Configuration, logger logr.Logger) {
}

func ParseFlags(config Configuration, opts ...Option) {
	initFlags(config, opts...)
	flag.Parse()
	lookupKubeconfigFlag()
}

func ExceptionNamespace() string {
	return exceptionNamespace
}

func PolicyExceptionEnabled() bool {
	return enablePolicyException
}

func LeaderElectionRetryPeriod() time.Duration {
	return leaderElectionRetryPeriod
}

func CleanupServerPort() string {
	return cleanupServerPort
}

func GlobalContextEnabled() bool {
	return enableGlobalContext
}

func printFlagSettings(logger logr.Logger) {
	logger = logger.WithName("flag")
	flag.VisitAll(func(f *flag.Flag) {
		logger.V(2).Info("", f.Name, f.Value)
	})
}
