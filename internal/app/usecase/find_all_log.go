package usecase

import (
	"context"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
)

type FindAllLogUseCase interface {
	Execute(ctx context.Context, input FindAllLogInput) ([]FindAllLogOutput, error)
}

type FindAllLogInput struct {
	Limit int64
}

type FindAllLogOutput struct {
	ID        string
	CreatedAt time.Time
	Level     string
	Message   any
}

type FindAllLogUseCaseImpl struct {
	logRepo protocols.LogRepository
}

func NewFindAllLogUseCaseImpl(logRepo protocols.LogRepository) FindAllLogUseCase {
	return FindAllLogUseCaseImpl{
		logRepo: logRepo,
	}
}

func (uc FindAllLogUseCaseImpl) Execute(ctx context.Context, input FindAllLogInput) ([]FindAllLogOutput, error) {
	logs, err := uc.logRepo.FindAll(ctx, input.Limit)
	if err != nil {
		return nil, err
	}

	var result []FindAllLogOutput
	for _, log := range logs {
		result = append(result, FindAllLogOutput{
			ID:        log.ID.Hex(),
			CreatedAt: log.CreatedAt,
			Level:     log.Level,
			Message:   log.Message,
		})
	}

	return result, nil
}
