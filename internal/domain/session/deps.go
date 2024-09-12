package session

import (
	"context"
	"notification-bot/internal/domain/session/entity"
)

type Repository interface {
	Save(ctx context.Context, session entity.Session) error
	Get(ctx context.Context, chatId int64) (*entity.Session, error)
}
