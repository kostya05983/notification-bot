package repository

import (
	"fmt"
	"notification-bot/internal/domain/session/entity"
)

func (r *Repository) mapState(state int64) (*entity.State, error) {
	switch state {
	case 1:
		return &entity.Created, nil
	case 2:
		return &entity.WaitToken, nil
	case 3:
		return &entity.ChooseCalendar, nil
	}

	return nil, fmt.Errorf("Can't define state from %d", state)
}
