package di

import (
	"context"
	"errors"
	"os"

	wrapErr "github.com/pkg/errors"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func initInfrastructure(di *dig.Container) error {
	return errors.Join(di.Provide(func() (*ydb.Driver, error) {
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
			return nil, wrapErr.Wrap(err, "ydb.Open")
		}

		return db, nil
	}),
		di.Provide(func() (*oauth2.Config, error) {
			googleJson := os.Getenv("GOOGLE_JSON")

			return google.ConfigFromJSON([]byte(googleJson), calendar.CalendarReadonlyScope)
		}),
	)
}
