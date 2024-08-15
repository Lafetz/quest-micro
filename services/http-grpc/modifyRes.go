package httpgrpc

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StandardResp struct {
	Details interface{} `json:"details"`
	Error   string      `json:"error"`
}

const (
	proxyFlag = "__succ__"
)

func mapGRPCCodeToHTTPStatus(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.InvalidArgument:
		return http.StatusUnprocessableEntity
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func HttpErrorHandler(ctx context.Context, mux *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	// success proxy
	raw := err.Error()
	if strings.HasPrefix(raw, proxyFlag) {
		raw = raw[len(proxyFlag):]
		w.Write([]byte(raw))
		return
	}

	// error handler
	s, ok := status.FromError(err)
	httpStatus := mapGRPCCodeToHTTPStatus(s.Code())

	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}
	var errMsg string
	if httpStatus >= 500 { // prevent returning port when there is connection error
		errMsg = "internal server error"
	} else {
		errMsg = s.Message()
	}
	resp := StandardResp{
		Details: s.Details(),
		Error:   errMsg,
	}
	bs, _ := json.Marshal(&resp)

	w.WriteHeader(httpStatus)
	w.Write(bs)

}
