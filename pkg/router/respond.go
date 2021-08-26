package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Respond responds with a JSON object from the object passed in body
func (r *Router) Respond(ctx context.Context, w http.ResponseWriter, code int, body interface{}) {
	reqID := r.GetRequestID(ctx)

	b, err := json.Marshal(body)
	if err != nil {
		r.Error(
			ctx,
			w,
			http.StatusInternalServerError,
			"Internal Server Error",
			err,
		)
		return
	}

	r.setDefaultHeaders(ctx, w)
	w.Header().Set("Content-Type", "application/json")

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	r.logger.Debugw("response",
		"request_id", reqID,
		"status_code", code,
		"response_body", string(b),
	)

	_, err = w.Write(b)
	if err != nil {
		r.logger.Errorw("response.fail",
			"request_id", reqID,
			"error", err,
		)
		return
	}
}

// RespondRaw responds to a request with a raw byte slice
func (r *Router) RespondRaw(ctx context.Context, w http.ResponseWriter, code int, body []byte) {
	reqID := r.GetRequestID(ctx)

	r.setDefaultHeaders(ctx, w)

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	r.logger.Debugw("response",
		"request_id", reqID,
		"status_code", code,
		"response_body", string(body),
	)

	_, err := w.Write(body)
	if err != nil {
		r.logger.Errorw("response.fail",
			"request_id", reqID,
			"error", err,
		)
		return
	}
}

// Error responds with our default ErrorResponse object as a response. Pass the
// user-facing response to msg and the error for our logs to err
// nolint:lll
func (r *Router) Error(ctx context.Context, w http.ResponseWriter, code int, msg, err interface{}) {
	reqID := r.GetRequestID(ctx)

	if err != nil {
		r.logger.Errorw("response.error",
			"request_id", reqID,
			"error", err,
		)
	}

	resp := ErrorResponse{
		Error: fmt.Sprintf("%v", msg),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		// Will be caught by recoverer middleware
		panic(fmt.Sprintf("[%s] %s", reqID, err))
	}

	r.setDefaultHeaders(ctx, w)
	w.Header().Set("Content-Type", "application/json")

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	r.logger.Debugw("response",
		"request_id", reqID,
		"status_code", code,
		"response_body", string(b),
	)

	_, err = w.Write(b)
	if err != nil {
		r.logger.Errorw("response.fail",
			"request_id", reqID,
			"error", err,
		)
		return
	}
}

// Success responds with the given code and no body
// nolint:lll
func (r *Router) Success(ctx context.Context, w http.ResponseWriter, code int) {
	reqID := r.GetRequestID(ctx)

	r.setDefaultHeaders(ctx, w)

	if code != http.StatusOK {
		w.WriteHeader(code)
	}

	r.logger.Debugw("response",
		"request_id", reqID,
		"status_code", code,
	)
}

// ServeFile responds with the image contents over the http connection
// nolint:lll
func (r *Router) ServeFile(ctx context.Context, w http.ResponseWriter, req *http.Request, f string) {
	reqID := r.GetRequestID(ctx)

	r.logger.Debugw("response.image",
		"request_id", reqID,
		"file_name", f,
	)

	http.ServeFile(w, req, f)
}

// WriteHeader responds with a raw status code
// nolint:lll
func (r *Router) WriteHeader(ctx context.Context, w http.ResponseWriter, code int) {
	reqID := r.GetRequestID(ctx)

	r.logger.Debugw("response",
		"request_id", reqID,
		"status_code", code,
	)

	w.WriteHeader(code)
}
