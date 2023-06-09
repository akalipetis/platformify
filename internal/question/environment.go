package question

import (
	"context"

	"github.com/platformsh/platformify/internal/question/models"
)

type Environment struct{}

func (q *Environment) Ask(ctx context.Context) error {
	answers, ok := models.FromContext(ctx)
	if !ok {
		return nil
	}

	answers.Environment = make(map[string]string)
	for _, dm := range answers.DependencyManagers {
		switch dm {
		case models.Poetry:
			answers.Environment["POETRY_VERSION"] = "1.4.0"
			answers.Environment["POETRY_VIRTUALENVS_IN_PROJECT"] = "true"
		case models.Pipenv:
			answers.Environment["PIPENV_VERSION"] = "2023.2.18"
			answers.Environment["PIPENV_VENV_IN_PROJECT"] = "1"
		}
	}

	return nil
}
