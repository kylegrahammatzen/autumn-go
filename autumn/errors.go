package autumn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("autumn: %s (code: %s, status: %d)", e.Message, e.Code, e.StatusCode)
}

func handleErrorResponse(resp *http.Response) error {
	var apiError Error
	apiError.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&apiError); err != nil {
		return &Error{
			Code:       "UNKNOWN_ERROR",
			Message:    fmt.Sprintf("HTTP %d: %s", resp.StatusCode, resp.Status),
			StatusCode: resp.StatusCode,
		}
	}

	return &apiError
}