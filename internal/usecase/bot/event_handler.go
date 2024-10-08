package bot

import (
	"context"
	"encoding/json"
	"fmt"

	calendarService "notification-bot/internal/domain/calendar/service"
	"notification-bot/internal/domain/session"
	"notification-bot/internal/domain/session/entity"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"golang.org/x/oauth2"
)

type Handler struct {
	session session.Repository
	config  *oauth2.Config
}

func (h *Handler) Handle(ctx context.Context, update tgbotapi.Update) (*tgbotapi.MessageConfig, error) {
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
				msg := tgbotapi.NewMessage(chatId, "Вы уже авторизованы для работы с ботом, доступные команды: \n editCalendar - изменить календарь")
				msg.ReplyToMessageID = update.Message.MessageID

				return &msg, nil
			}
			authLink := calendarService.GetAuthLink(h.config)

			msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Перейдите по ссылке и отправьте код боту %s", authLink))
			msg.ReplyToMessageID = update.Message.MessageID

			session.State = entity.StateWaitToken

		case "editCalendar":
			//todo validation
		}
	case entity.StateWaitToken:
		code := update.Message.Text

		token, err := calendarService.GetToken(h.config, code)
		if err != nil {
			return nil, errors.Wrap(err, "session.GetToken:")
		}

		json, err := json.Marshal(token)
		if err != nil {
			return nil, errors.Wrap(err, "json.Marshal:")
		}

		session.AddToken(string(json))
	case entity.StateChooseCalendar:
		service, err := h.initService(ctx, *session)
		if err != nil {
			return nil, errors.Wrap(err, "h.initService")
		}

		calendars, err := service.GetCalendars()
		if err != nil {
			return nil, errors.Wrap(err, "service.GetCalendars:")
		}

		row := tgbotapi.NewInlineKeyboardRow()

		for _, item := range calendars.Items {
			button := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("Название календаря:%s", item.Description), item.Id)
			row = append(row, button)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите календарь:")

		msg.ReplyMarkup = row

		return &msg, nil
	case entity.StateWaitCalendar:
		calendarId := update.CallbackData()
		if calendarId == "" { //user don't push button
			return nil, nil
		}

	}

	err = h.session.Save(ctx, *session)
	if err != nil {
		return nil, errors.Wrap(err, "session.Save")
	}

	return nil, nil
}

func (h *Handler) initService(ctx context.Context, session entity.Session) (*calendarService.CalendarService, error) {
	var oauth *oauth2.Token

	err := json.Unmarshal([]byte(*session.GoogleToken), oauth)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal:")
	}

	httpClient := h.config.Client(ctx, oauth)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, errors.Wrap(err, "calendar.NewService:")
	}

	entityService := calendarService.New(*srv)

	return &entityService, nil
}
