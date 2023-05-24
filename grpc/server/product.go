package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	protos "webservices/grpc/protos/proto/product"
)

type ProductServer struct {
	log hclog.Logger
}

func NewProductServer(l hclog.Logger) *ProductServer {
	return &ProductServer{l}
}

func (p *ProductServer) GetProduct(ctx context.Context, req *protos.GetProductRequest) (*protos.GetProductResponse, error) {
	p.log.Info("Handle GetProduct", "id", req.GetId())

	return &protos.GetProductResponse{
		Id:          req.GetId(),
		Name:        "Test",
		Description: "Test",
		Price:       "199.9",
	}, nil
}
