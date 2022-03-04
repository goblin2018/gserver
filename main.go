package main

import (
	"fmt"
	"gserver/controllers"
	"gserver/pkg/conf"
	"gserver/pkg/env"
	"net/http"
)

func main() {
	env.Init()
	fmt.Print("Hello Gserver")

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Server.Port),
		Handler:        controllers.InitRoute(),
		ReadTimeout:    conf.Server.ReadTimeout,
		WriteTimeout:   conf.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(fmt.Errorf("listen and server error: %v", err))
		return
	}

	// go func() {
	// 	if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
	// 		panic(err)
	// 	}
	// }()

	// quit := make(chan os.Signal)

	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// <-quit
	// log.Println("Shutting down server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatalf("Server forced to shutdown: %v", err)
	// }
	// log.Println("Server exiting")

}
