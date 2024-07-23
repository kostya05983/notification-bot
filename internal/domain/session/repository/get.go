package repository

import (
	"context"

	"notification-bot/internal/domain/session/entity"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func (r *Repository) get(ctx context.Context, chatId int64) (*entity.Session, error) {
	readTx := table.TxControl(
		table.BeginTx(
			table.WithOnlineReadOnly(),
		),
		table.CommitTx(),
	)

	var res result.Result
	var state int64
	var googleToken string

	err := r.driver.Table().Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
			_, res, err = s.Execute(ctx, readTx,
				`
			DECLARE $chat_id as int64
			SELECT state, google_token FROM sessions where chat_id=$chat_id;
		`,
				table.NewQueryParameters(
					table.ValueParam("$chat_id", types.Int64Value(chatId)),
				),
			)
			if err != nil {
				return err
			}
			defer res.Close()

			for res.NextResultSet(ctx) {
				for res.NextRow() {
					err = res.ScanNamed(
						named.Required(stateField, &state),
						named.Optional(googleTokenField, &googleToken),
					)
					if err != nil {
						return err
					}
				}
			}

			return res.Err()
		},
	)

	if err != nil {
		return nil, err
	}

	mappedState, err := r.mapState(state)
	if err != nil {
		return nil, err
	}

	return &entity.Session{
		ChatId:      chatId,
		State:       *mappedState,
		GoogleToken: &googleToken,
	}, nil
}
