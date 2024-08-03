package service

import (
	"log"
	"net"
	"net/http"
)

type Service struct {
	state State
}

// New creates a new instance of the user management service.
func New(state State) *Service {
	return &Service{
		state: state,
	}
}

// Listen starts the user management service.
func (s *Service) Listen(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /on-sign-up", s.onSignUp)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("User Management service is listening on %s\n", addr)

	return http.Serve(listener, mux)
}
