package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Server struct {
	DB         *gorm.DB
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
	db, err := gorm.Open(postgres.Open("postgresql://banner:banner@localhost:5432/banner"))
	if err != nil {
		os.Exit(1)
	}
	s.DB = db
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)

}
