package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"service/internal/logger"
	servicemetrics "service/internal/metrics/service"

	"service/internal/transport"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Version string

func main() {
	// uncommit for profiling service with pprof
	// mux := http.NewServeMux()
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	// go func() {
	// 	http.ListenAndServe(":8081", mux)
	// }()

	lg := logger.New("service", true, "debug")
	lg.LogInfo(fmt.Sprintf("Logger init for %v with %v level", "service", "debug"))

	ServiceWithMetrics := servicemetrics.NewServiceMetrics("ns")

	lg.LogInfo(fmt.Sprintf("Server started on %s with: version - %s", ":8080", Version))

	transportMiddlewares := []transport.Option{}
	servicesOpts := []transport.Option{
		transport.Service(transport.NewMyService(ServiceWithMetrics)),
	}

	srv := transport.New(lg.Logger, append(transportMiddlewares, servicesOpts...)...).WithLog()

	srvMetrics := fiber.New(fiber.Config{DisableStartupMessage: true})
	srvMetrics.All("/", adaptor.HTTPHandler(promhttp.Handler()))
	go func() {
		err := srvMetrics.Listen(":9090")
		if err != nil {
			log.Fatal("error in metrics: " + err.Error())
		}
	}()
	srv.Fiber().Get("/readiness", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON("OK")
	})
	srv.Fiber().Get("/liveness", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusOK)
		return ctx.JSON("OK")
	})

	go func() {
		if err := srv.Fiber().Listen(":8080"); err != nil {
			lg.LogInfo(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	srv.Shutdown()
	lg.LogInfo("shutting down")
}
