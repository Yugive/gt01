package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	srv    []*http.Server
	ctx    context.Context
	cancel func()
}

func NewApp(srv []*http.Server) *App {
	c, cancel := context.WithCancel(context.Background())
	return &App{
		srv:    srv,
		ctx:    c,
		cancel: cancel,
	}
}

func (a *App) Run() error {
	g, ctx := errgroup.WithContext(a.ctx)

	for _, srv := range a.srv {
		s := srv
		g.Go(func() error {
			<-ctx.Done()
			return s.Shutdown(ctx)
		})
		g.Go(func() error {
			return s.ListenAndServe()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				a.Stop()
			}
		}
	})
	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

type myHandler struct{}

func (my *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	debugServer := &http.Server{
		Addr:    ":6060",
		Handler: nil,
	}
	app := NewApp([]*http.Server{debugServer, httpServer})
	time.AfterFunc(time.Second*5, func() {
		app.Stop()
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
