package app

import "users/internal/config"

type provider struct {
	grpcConfig *config.Grpc
}

func newProvider(grpcConfig *config.Grpc) *provider {
	return &provider{
		grpcConfig: grpcConfig,
	}
}
