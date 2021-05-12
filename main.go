package main

import (
	"fmt"
	"github.com/WelfareSystem/log"
	"github.com/WelfareSystem/router"
	"net/http"

)

func init() {
	err := log.LogInit()
	if err != nil {
		fmt.Println("panic: log init error")
	}
}

func main ()  {
	router := router.NewRouter()
	log.Info("server is running...")
	log.ERROR.Fatal(http.ListenAndServe(":8001", router))
}
