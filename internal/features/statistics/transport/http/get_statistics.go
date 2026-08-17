package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/afkpanda2004/golang-todoapp/internal/core/domain"
	core_logger "github.com/afkpanda2004/golang-todoapp/internal/core/logger"
	core_http_request "github.com/afkpanda2004/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/afkpanda2004/golang-todoapp/internal/core/transport/http/response"
)

type GetTaskStatisticsResponse struct {
	TaskCreated           int      `json:"task_created"              example:"10"`
	TaskCompleted         int      `json:"task_completed"            example:"7"`
	TaskCompletionPercent *float64 `json:"task_completion_percent"   example:"70.0"`
	TaskAvgCompletionTime *string  `json:"task_avg_completion_time_sec" example:"3600s"`
}

// GetStatistics godoc
// @Summary      Получить статистику по задачам
// @Description  Возвращает агрегированную статистику по задачам с опциональной фильтрацией
// @Tags         statistics
// @Produce      json
// @Param        user_id  query int    false "Фильтр по ID пользователя"
// @Param        from     query string false "Начало периода (RFC3339)" format(date-time)
// @Param        to       query string false "Конец периода (RFC3339)" format(date-time)
// @Success      200      {object}  GetTaskStatisticsResponse
// @Failure      400      {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      500      {object}  core_http_response.ErrorResponse "Internal server error"
// @Router       /statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIdFromToQueryParams(r)

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task statistics query params",
		)
		return
	}

	stats, err := h.statisticsService.GetTaskStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get task statistics",
		)
		return
	}

	response := taskStatisticsResponseFromDomain(stats)
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getUserIdFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetTimeQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetTimeQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}

func taskStatisticsResponseFromDomain(stats domain.Statistics) GetTaskStatisticsResponse {

	var avgTime *string
	if stats.TaskAvgCompletionTime != nil {
		duration := stats.TaskAvgCompletionTime.String()
		avgTime = &duration
	}

	return GetTaskStatisticsResponse{
		TaskCreated:           stats.TaskCreated,
		TaskCompleted:         stats.TaskCompleted,
		TaskCompletionPercent: stats.TaskCompletionPercent,
		TaskAvgCompletionTime: avgTime,
	}
}
