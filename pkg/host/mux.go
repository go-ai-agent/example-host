package host

import (
	"github.com/go-ai-agent/example-host/pkg/handler"
	"net/http"
	"net/http/pprof"
)

const (
	healthLivenessPattern = "/health/liveness"

	indexPattern   = "/debug/pprof/"
	cmdLinePattern = "/debug/pprof/cmdline"
	profilePattern = "/debug/pprof/profile" // ?seconds=30
	symbolPattern  = "/debug/pprof/symbol"
	tracePattern   = "/debug/pprof/trace"
)

func initMux(r *http.ServeMux) {
	addRoutes(r)
	//r.Handle(runPattern, http.HandlerFunc(handler.RunHandler))
	//r.Handle(actuatorPattern, http.HandlerFunc(middleware.ActuatorHandler))
	r.Handle(healthLivenessPattern, http.HandlerFunc(handler.HealthLivenessHandler))
}

func addRoutes(r *http.ServeMux) {
	r.Handle(indexPattern, http.HandlerFunc(pprof.Index))
	r.Handle(cmdLinePattern, http.HandlerFunc(pprof.Cmdline))
	r.Handle(profilePattern, http.HandlerFunc(pprof.Profile))
	r.Handle(symbolPattern, http.HandlerFunc(pprof.Symbol))
	r.Handle(tracePattern, http.HandlerFunc(pprof.Trace))

}
