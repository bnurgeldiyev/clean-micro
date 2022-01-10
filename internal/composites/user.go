package composites

import (
	"clean-micro/internal/adapters/api"
	user2 "clean-micro/internal/adapters/api/user"
	user3 "clean-micro/internal/adapters/pgdb/user"
	"clean-micro/internal/adapters/service/cache"
	"clean-micro/internal/domain/user"
)

type UserComposite struct {
	Storage user.Storage
	Service user2.Service
	Handler api.AuthServer
}

func NewUserComposite(postgresComposite *PostgresComposite, cache *cache.RedisService) (*UserComposite, error) {
	storage := user3.NewStorage(postgresComposite.DB)
	service := user.NewService(storage, cache)
	handler := user2.NewUserHandler(service)

	return &UserComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
