package grpc

import (
	"database/sql"
	"fmt"

	"github.com/minhvuongrbs/financial-service-example/config"
	"github.com/minhvuongrbs/financial-service-example/pkg/database"
	pkgredis "github.com/minhvuongrbs/financial-service-example/pkg/redis"
	"github.com/redis/go-redis/v9"
)

type infrastructureDependencies struct {
	db          *sql.DB
	redisClient *redis.Client
}

func initInfrastructure(cfg config.Config) (infrastructureDependencies, error) {
	db, err := database.NewMysqlDatabaseConn(&cfg.Database)
	if err != nil {
		return infrastructureDependencies{}, fmt.Errorf("failed to init database: %w", err)
	}
	redisClient, err := pkgredis.NewRedisClient(cfg.RedisConnection)
	if err != nil {
		return infrastructureDependencies{}, fmt.Errorf("failed to init redis client: %w", err)
	}

	return infrastructureDependencies{
		db:          db,
		redisClient: redisClient,
	}, nil
}
