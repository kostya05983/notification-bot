package di

import (
	baseErrors "errors"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

func InitDi(di *dig.Container) error {
	err := baseErrors.Join(
		initInfrastructure(di),
	)
	if err != nil {
		return errors.Wrap(err, "initDi:")
	}

	return nil
}
