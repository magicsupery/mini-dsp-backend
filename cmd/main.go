package main

import (
	"log"
	"net/http"
	"time"

	"mini-dsp-backend/routers"
)

func main() {
	// 初始化 Gin 路由
	r := routers.SetupRouter()

	// 可根据需要自定义 http.Server
	srv := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("DSP Backend is running at :8080 ...")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("ListenAndServe error: %v\n", err)
	}
}
