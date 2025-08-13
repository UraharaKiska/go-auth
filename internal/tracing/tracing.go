package tracing

import (
    "context"
    "log"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/sdk/resource"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
    "go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

func Init(serviceName string) func(context.Context) error {
    ctx := context.Background()

    client, err := otlptracegrpc.New(ctx,
        otlptracegrpc.WithEndpoint("localhost:4317"),
        otlptracegrpc.WithInsecure(),
    )
    if err != nil {
        log.Fatal("Failed to create OTLP Trace client: ", err)
    }

    res, err := resource.New(ctx,
        resource.WithAttributes(
            semconv.ServiceNameKey.String(serviceName),
        ),
    )
    if err != nil {
        log.Fatal(err)
    }

    tp := sdktrace.NewTracerProvider(
        sdktrace.WithBatcher(client),
        sdktrace.WithResource(res),
    )
    otel.SetTracerProvider(tp)

    // Создаем и сохраняем глобальный трейсер один раз
    Tracer = otel.Tracer(serviceName)

    return tp.Shutdown
}

// Функция для доступа к глобальному трейсеру из других пакетов
// func Tracer() trace.Tracer {
//     return tracer
// }
