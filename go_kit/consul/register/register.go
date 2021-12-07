package register

import (
	consulApi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

const (
	ip          = "192.168.175.143"
	port        = "8500"
	address     = ip + ":" + port
	health_http = "http://" + address + "/health"
)

func RegService() {
	config := consulApi.DefaultConfig()
	config.Address = address
	reg := consulApi.AgentServiceRegistration{}
	reg.Name = "userservice" //注册service的名字
	reg.Address = ip         //注册service的ip
	reg.Port = 8080          //注册service的端口
	reg.Tags = []string{"primary"}

	check := consulApi.AgentServiceCheck{} //创建consul的检查器
	check.Interval = "5s"                  //设置consul心跳检查时间间隔
	check.HTTP = health_http               //设置检查使用的url

	reg.Check = &check

	client, err := consulApi.NewClient(config) //创建客户端
	if err != nil {
		log.Fatal(err)
	}
	err = client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}
