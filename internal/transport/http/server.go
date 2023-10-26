package transport

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type HttpServer struct {
	httpServer *http.Server
}

func (s *HttpServer) Run(address string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        handler,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// graceful shutdown
	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig
		shutdownCtx, _ := context.WithTimeout(serverCtx, 5*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Info().Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := s.httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Error().Msg(err.Error())
		}
		serverStopCtx()
	}()

	err := s.httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Error().Msg(err.Error())
	}
	<-serverCtx.Done()
	return err
}
