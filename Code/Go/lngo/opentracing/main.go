package main

import (
	"log"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := cfg.New(
		"your_service_name",
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		log.Println(err)
	}
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	someFunction()
}

func someFunction() {
	parent := opentracing.GlobalTracer().StartSpan("hello")
	defer parent.Finish()
	child := opentracing.GlobalTracer().StartSpan(
		"world", opentracing.ChildOf(parent.Context()))
	defer child.Finish()
}
