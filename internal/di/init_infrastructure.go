package di

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"go.uber.org/dig"
)

func initInfrastructure(di dig.Container) error {
	return di.Provide(func() (*ydb.Driver, error) {
		dsn := os.Getenv("YDB-DSN")
		if len(dsn) == 0 {
			return nil, errors.New("YDB-DSN is empty, please provide it")
		}

		token := os.Getenv("YDB_IAM_TOKEN")
		if len(token) == 0 {
			return nil, errors.New("YDB_IAM_TOKEN is empty, please provide it")
		}

		db, err := ydb.Open(context.Background(), dsn, ydb.WithAccessTokenCredentials(token))

		if err != nil {
			return nil, errors.Wrap(err, "")
		}

		return db, nil
	})
}
