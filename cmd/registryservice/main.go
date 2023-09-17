package main

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

//你个registry服务怎么干的（注册handler，启动registry服务）

func main() {
	registry.SetupRegistryService()
	//绑定路径和处理函数
	http.Handle("/services", registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("Registry service started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shoutting down registry service")

}
