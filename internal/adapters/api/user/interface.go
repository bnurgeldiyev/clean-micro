package user

import (
	"clean-micro/internal/adapters/api"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/peer"
)

type userHandler struct {
	apiService Service
}

func NewUserHandler(service Service) api.AuthServer {
	return &userHandler{
		apiService: service,
	}
}

func handleGrpc(handleName string, ctx context.Context) (clog *log.Entry) {

	p, _ := peer.FromContext(ctx)
	clog = log.WithFields(log.Fields{
		"remote-addr": p.Addr.String(),
		"handle":      handleName,
	}).WithContext(ctx)

	return
}
