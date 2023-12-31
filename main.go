package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	fmt.Println("go http2 start")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This request is served over %s protocol!", r.Proto)
	}))

	h2srv := &http2.Server{}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: h2c.NewHandler(http.DefaultServeMux, h2srv),
	}

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Printf("start server http://localhost:%s \n", port)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
