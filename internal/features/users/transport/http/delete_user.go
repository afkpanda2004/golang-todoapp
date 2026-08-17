package users_transport_http

import (
	"net/http"

	core_logger "github.com/afkpanda2004/golang-todoapp/internal/core/logger"
	core_http_request "github.com/afkpanda2004/golang-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/afkpanda2004/golang-todoapp/internal/core/transport/http/response"
)

type DeleteUserResponse UserDTOResponse

// DeleteUser godoc
// @Summary      Удалить пользователя
// @Description  Удаляет пользователя по идентификатору
// @Tags         users
// @Param        id   path  int  true  "ID удаляемого пользователя"
// @Success      204  "Пользователь успешно удалён"
// @Failure      400  {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse "Not found"
// @Failure      500  {object}  core_http_response.ErrorResponse "Internal server error"
// @Router       /users/{id} [delete]

func (h *UserHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")

	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failde to get userID path value",
		)
		return
	}

	if err := h.usersService.DeleteUser(ctx, userID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete user",
		)
		return
	}

	responseHandler.NoContentResponse()

}
