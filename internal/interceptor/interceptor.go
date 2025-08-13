package interceptor

import (
	"context"

	"time"

	"github.com/UraharaKiska/go-auth/internal/logger"
	"github.com/UraharaKiska/go-auth/internal/metric"
	rateLimiter "github.com/UraharaKiska/go-auth/internal/rate_limiter"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validator interface {
	Validate() error
}

type RateLimiterInterceptor struct {
	rateLimiter *rateLimiter.TokenBucketLimiter
}

func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if val, ok := req.(validator); ok {
		if err := val.Validate(); err != nil {
			return nil, err
		}
	}

	return handler(ctx, req)
}

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", req))
	}

	logger.Info("request", zap.String("method", info.FullMethod), zap.Any("req", req), zap.Any("res", res), zap.Duration("duration", time.Since(now)))

	return res, err
}

func MetricsInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	metric.IncRequestCounter()
	res, err := handler(ctx, req)
	if err != nil {
		metric.IncResponseCounter("error", info.FullMethod)
	} else {
		metric.IncResponseCounter("success", info.FullMethod)
	}

	return res, err
}

func TraceIDInjectorInterceptor(
    ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler,
) (interface{}, error) {
    resp, err := handler(ctx, req)

    span := trace.SpanFromContext(ctx)
    if span.SpanContext().IsValid() {
        traceID := span.SpanContext().TraceID().String()
        md := metadata.Pairs("x-trace-id", traceID)
        grpc.SendHeader(ctx, md)
    }

    return resp, err
}


func NewRateLimiterInterceptor(rateLimiter *rateLimiter.TokenBucketLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{rateLimiter: rateLimiter}
}

func (r *RateLimiterInterceptor) Unary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if !r.rateLimiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "too many requests")
	}
	return handler(ctx, req)
}