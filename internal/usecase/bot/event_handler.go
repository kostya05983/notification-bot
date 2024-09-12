package bot

import (
	"context"

	"notification-bot/internal/domain/session"
	"notification-bot/internal/domain/session/entity"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

type Handler struct {
	session session.Repository
}

func (h *Handler) Handle(ctx context.Context, update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
	if !update.Message.IsCommand() {
		return nil, nil
	}

	chatId := update.Message.Chat.ID

	session, err := h.session.Get(ctx, chatId)

	if err != nil {
		return nil, errors.Wrap(err, "sessions.Get:")
	}

	if session == nil {
		session = &entity.Session{
			ChatId:      chatId,
			State:       entity.StateCreated,
			GoogleToken: nil,
		}
	}

	switch session.State {
	case entity.StateCreated:
		//look into command
		switch update.Message.Command() {
		case "start":
			if !session.IsAuthenticated() {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы уже авторизованы для работы с ботом, доступные команды: \n editCalendar - изменить календарь")
				msg.ReplyToMessageID = update.Message.MessageID

				return &msg, nil
			}

		case "editCalendar":
			//todo validation
		}
	case entity.StateWaitToken:

	}

}
