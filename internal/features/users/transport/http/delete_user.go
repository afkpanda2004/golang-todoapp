package users_transport_http

import (
	"net/http"

	core_logger "github.com/afkpanda2004/golang-todoapp/internal/core/logger"
	core_http_response "github.com/afkpanda2004/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/afkpanda2004/golang-todoapp/internal/core/transport/http/utils"
)

type DeleteUserResponse UserDTOResponse

func (h *UserHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")

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
