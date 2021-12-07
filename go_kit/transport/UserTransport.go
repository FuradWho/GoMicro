package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/FuradWho/GoMicro/go_kit/endpoint"
	"net/http"
	"strconv"
)

func DecodeUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	if r.URL.Query().Get("uid") != "" {
		uid, err := strconv.Atoi(r.URL.Query().Get("uid"))
		if err != nil {
			return nil, err
		}
		return endpoint.UserRequest{Uid: uid}, nil
	}

	return nil, errors.New("param error")
}

func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type", "application/json") //设置响应格式为json，这样客户端接收到的值就是json，就是把我们设置的UserResponse给json化了

	return json.NewEncoder(w).Encode(response) //判断响应格式是否正确
}
