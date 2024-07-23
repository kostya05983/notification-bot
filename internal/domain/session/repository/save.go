package repository

import (
	"context"

	"notification-bot/internal/domain/session/entity"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

const saveQuery = `UPSERT INTO sessions(chat_id, state, google_token) VALUES($chat_id, $state, $google_token)`

func (r *Repository) save(ctx context.Context, session entity.Session) error {
	writeControl := table.TxControl(
		table.BeginTx(
			table.WithSerializableReadWrite()))

	err := r.driver.Table().Do(ctx, func(ctx context.Context, s table.Session) error {
		params := table.NewQueryParameters(
			table.ValueParam("$chat_id", types.Int64Value(session.ChatId)),
			table.ValueParam("$state", types.Int64Value(int64(session.State))),
			table.ValueParam("$google_token", types.NullableStringValueFromString(session.GoogleToken)),
		)

		_, res, err := s.Execute(ctx, writeControl, saveQuery, params)
		if err != nil {
			return err
		}
		defer res.Close()

		return res.Err()
	})

	if err != nil {
		return err
	}

	return nil
}
