package main

import (
	"context"
	"fmt"
	"net/http"
)

type HttpPortNumber int

type ApiServer struct {
	Addr string
	Port HttpPortNumber

	HttpHandlers map[string]http.HandlerFunc

	server *http.Server
}

func (s *ApiServer) Start() error {
	if s.HttpHandlers == nil {
		return fmt.Errorf("channels HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	for key, fun := range s.HttpHandlers {
		handler.HandleFunc(key, fun)
	}

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Addr, s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *ApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
