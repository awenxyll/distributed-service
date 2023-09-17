package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

//服务们怎么干（启动服务（包括log web服务启动，log服务注册））

func Start(ctx context.Context, host, port string, reg registry.Registration, registerHandlerFunc func()) (context.Context, error) {
	registerHandlerFunc()
	//启动web服务
	ctx = startService(ctx, reg.ServiceName, host, port)
	//服务注册
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = host + ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v started. Press any key to stop. \n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
