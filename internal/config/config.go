package config

import "go.uber.org/zap"

type Config struct {
	Logger *zap.Logger

	GRPCAddr string `env:"GRPC_ADDR" default:":9090"`
	HTTPAddr string `env:"HTTP_ADDR" default:":8080"`
}
