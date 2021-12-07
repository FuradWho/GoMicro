package main

import (
	"fmt"
	"github.com/FuradWho/GoMicro/go_kit/consul/register"
	"github.com/FuradWho/GoMicro/go_kit/endpoint"
	"github.com/FuradWho/GoMicro/go_kit/service"
	"github.com/FuradWho/GoMicro/go_kit/transport"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("!!!!")
	user := service.UserService{}
	point := endpoint.GenUserEndPoint(user)
	serverHandler := httpTransport.NewServer(point, transport.DecodeUserRequest, transport.EncodeUserResponse)

	r := mux.NewRouter() //使用mux来使服务支持路由
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status":"ok"}`))
	})

	errChan := make(chan error)
	go func() {
		register.RegService()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Errorln(err)
			errChan <- err
		}
	}()

	go func() {
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-signalChan)
	}()

	getErr := <-errChan //只要报错 或者service关闭阻塞在这里的会进行下去
	register.UnRegService()
	log.Infoln(getErr)
}
