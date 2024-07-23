package repository

import (
	"context"
	"path"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func (r *Repository) createTable(ctx context.Context) error {
	return r.driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		return s.CreateTable(ctx, path.Join(r.driver.Name(), tableName),
			options.WithColumn(chatIdField, types.TypeInt64),
			options.WithColumn(stateField, types.TypeInt64),
			options.WithColumn(googleTokenField, types.TypeString),
		)
	})
}
