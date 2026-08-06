package task_service

import (
	"context"
	"fmt"
)

func (s *TaskService) DeleteTask(
	ctx context.Context,
	id int,
) error {

	err := s.taskRepository.DeleteTask(ctx, id)

	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}
