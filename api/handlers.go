package api

import (
	"net/http"

	apiv1 "github.com/lab1/errors-logs/api/v1"
	"github.com/lab1/errors-logs/pkg/lib/errors"
)

func (a *API) checkInternetHandler(w http.ResponseWriter, r *http.Request) {
	req := &apiv1.ConnectionCheckRequest{}
	err := decodeRequest(r, req)
	if err != nil {
		a.handleError(w, errors.BadRequest.Wrap(err))

		return
	}

	resp := apiv1.ConnectionCheckResponse{
		Status: "up",
	}
	err = a.connService.CheckSet(req.Addresses, req.Timeout)
	if err != nil {
		if errors.NoConnection.Is(err) {
			resp.Status = "down"
			a.logger.Debugf("down error: %v", err)
		} else {
			a.handleError(w, err)

			return
		}
	}

	_ = encodeResponse(w, resp)
}
