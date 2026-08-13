package statistics_postgres_repository

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/afkpanda2004/golang-todoapp/internal/core/domain"
)

func (r *StatisticsRepository) GetTaskStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.TaskStatisticsQueryResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var (
		conditions []string
		args       []any
	)

	if userID != nil {
		args = append(args, *userID)
		conditions = append(conditions, fmt.Sprintf("author_user_id=$%d", len(args)))
	}

	if from != nil {
		args = append(args, *from)
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", len(args)))
	}

	if to != nil {
		args = append(args, *to)
		conditions = append(conditions, fmt.Sprintf("created_at <= $%d", len(args)))
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(`
	SELECT
		COUNT(*) AS total_created,
		COUNT(*) FILTER (WHERE completed) AS total_completed,
		AVG(EXTRACT(EPOCH FROM (completed_at - created_at)))
			FILTER (WHERE completed) AS avg_completion_seconds
	FROM todoapp.tasks
	%s
	`, whereClause)

	row := r.pool.QueryRow(ctx, query, args...)

	var result domain.TaskStatisticsQueryResult

	err := row.Scan(
		&result.TaskCreated,
		&result.TaskCompleted,
		&result.TaskAvgCompletionSeconds,
	)
	if err != nil {
		return domain.TaskStatisticsQueryResult{}, fmt.Errorf("select task statistics: %w", err)
	}

	return result, nil
}
