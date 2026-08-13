package domain

import "time"

type TaskStatisticsQueryResult struct {
	TaskCreated              int
	TaskCompleted            int
	TaskAvgCompletionSeconds *float64
}

type Statistics struct {
	TaskCreated           int
	TaskCompleted         int
	TaskCompletionPercent *float64
	TaskAvgCompletionTime *time.Duration
}

func NewTaskStatistics(
	taskCreated int,
	taskCompleted int,
	taskCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TaskCreated:           taskCreated,
		TaskCompleted:         taskCompleted,
		TaskCompletionPercent: taskCompletedRate,
		TaskAvgCompletionTime: tasksAverageCompletionTime,
	}
}
