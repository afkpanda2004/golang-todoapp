package statistics_service

import (
	"context"
	"fmt"
	"time"

	"github.com/afkpanda2004/golang-todoapp/internal/core/domain"
	core_errors "github.com/afkpanda2004/golang-todoapp/internal/core/errors"
)

func (s *StatisticsService) GetTaskStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil && from.After(*to) {
		return domain.Statistics{}, fmt.Errorf(
			"'from' must be before or equal to 'to': %w",
			core_errors.ErrInvalidArgument,
		)
	}

	statsModel, err := s.statisticsRepository.GetTaskStatistics(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("get task statistics from repository: %w", err)
	}

	var avgCompletionTime *time.Duration
	if statsModel.TaskAvgCompletionSeconds != nil {
		duration := time.Duration(*statsModel.TaskAvgCompletionSeconds * float64(time.Second))
		avgCompletionTime = &duration
	}

	return domain.NewTaskStatistics(
		statsModel.TaskCreated,
		statsModel.TaskCompleted,
		statsModel.TaskAvgCompletionSeconds,
		avgCompletionTime,
	), nil
}
