package usercenter

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chhz0/usercenter-go/internal/pkg/options"
	"github.com/chhz0/usercenter-go/internal/usercenter/router"
	"github.com/gin-gonic/gin"
)

type httpServer struct {
	srv *http.Server
}

func newGinHTTPServer(opts *options.HTTPOptions) *httpServer {
	g := gin.New()

	router.Register(g)

	return &httpServer{
		srv: &http.Server{
			Addr:    opts.Addr,
			Handler: g.Handler(),
		},
	}
}

func (hs *httpServer) ListenAndServe() error {
	go func() {
		if err := hs.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := hs.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	return nil
}
