package transport

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/FuradWho/GoMicro/go_kit_consumer/endpoint"
	"net/http"
	"strconv"
)

func GetUserInfoRequest(_ context.Context, request *http.Request, r interface{}) error {
	userRequest := r.(endpoint.UserRequest)
	request.URL.Path += "/user/" + strconv.Itoa(userRequest.Uid)
	return nil
}

func GetUserInfoResponse(_ context.Context, res *http.Response) (response interface{}, err error) {
	if res.StatusCode > 400 {
		return nil, errors.New("no data")
	}
	var userResponse endpoint.UserResponse
	err = json.NewDecoder(res.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}
	return userResponse, err
}
