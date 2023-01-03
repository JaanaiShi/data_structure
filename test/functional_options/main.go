package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      *tls.Config
}


type Option func(*Server)

func WithProtocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConns(maxcounts int) Option {
	return func(s *Server) {
		s.MaxConns = maxcounts
	}
}

func WithTLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func NewServer(addr string, port int, options ...Option) *Server {
	server := Server{
		Addr: addr,
		Port: port,
	}

	for _, option := range options {
		option(&server)
	}

	return &server
}

func main() {
	server := NewServer("https://github.com", 1234, WithTimeout(100*time.Second))
	fmt.Println(server)
}