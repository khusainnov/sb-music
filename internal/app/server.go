package app

import (
	"context"
	"net"

	"github.com/khusainnov/sb-music/internal/config"
	"github.com/khusainnov/sb-music/internal/http"
	"go.uber.org/zap"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

var (
	log, _ = zap.NewProduction()
)

func Run(ctx context.Context, cfg *config.Config) error {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	reflection.Register(s)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		cfg.Logger.Fatal("failed to listen tcp", zap.Error(err))
	}

	httpServer := http.New(cfg)
	httpServer.RunHTTP()

	cfg.Logger.Info("starting listen grpc", zap.String("PORT", cfg.GRPCAddr))
	if err = s.Serve(l); err != nil {
		log.Fatal("error due listen grpc", zap.Error(err))
	}

	return nil
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info("unary interceptor", zap.Any("method", info.FullMethod))

	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Info("stream interceptor", zap.Any("method", info.FullMethod))

	return handler(srv, stream)
}
