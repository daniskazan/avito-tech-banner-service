package handlers

import (
	"context"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Server struct {
	Client     *ent.Client
	Router     *echo.Echo
	emailAdmin string
}

func NewServer() *Server {
	srv := &Server{Router: echo.New()}
	return srv
}

func (s *Server) Start() {
	s.setupDatabase()
	s.ConfigureRouting()

	s.Router.Logger.Fatal(s.Router.Start(":8000"))
}

func (s *Server) setupDatabase() {
	client, err := ent.Open("postgres", "postgresql://banner:banner@localhost:5432/banner?sslmode=disable")
	if err != nil {
		log.Fatalf("error creating client db: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	s.Client = client
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)

}
