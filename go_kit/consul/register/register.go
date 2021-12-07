package register

import (
	consulApi "github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

const (
	ip             = "192.168.175.143"
	consulPort     = 8500
	consulAddress  = ip + ":8500"
	serviceName    = "userService"
	servicePort    = 8080
	serviceAddress = ip + ":8080"
	healthHttp     = "http://" + serviceAddress + "/health"
)

var ConsulClient *consulApi.Client

func init() {
	config := consulApi.DefaultConfig()
	config.Address = consulAddress
	//创建客户端
	client, err := consulApi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	ConsulClient = client
}

func RegService() {

	reg := consulApi.AgentServiceRegistration{}
	reg.Name = serviceName //注册service的名字
	reg.Address = ip       //注册service的ip
	reg.Port = servicePort //注册service的端口
	reg.Tags = []string{"primary"}

	check := consulApi.AgentServiceCheck{} //创建consul的检查器
	check.Interval = "5s"                  //设置consul心跳检查时间间隔
	check.HTTP = healthHttp                //设置检查使用的url

	reg.Check = &check

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func UnRegService() {
	err := ConsulClient.Agent().ServiceDeregister(serviceName)
	if err != nil {
		log.Errorln(err)
		return
	}
}
