package api

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	customValidator "github.com/Sayonara143/urfu-pp-hhru-backend/api/custom-validator"
	"github.com/Sayonara143/urfu-pp-hhru-backend/api/middlewares"
)

type Server struct {
	e *echo.Echo

	hh   HHService
	l    Logger
	auth AuthService
}

func New(hh HHService, auth AuthService, l Logger) *Server {
	s := &Server{
		e:  echo.New(),
		hh: hh,
		l:  l,
	}
	cv := customValidator.New(validator.New())
	s.e.Validator = cv
	s.e.HidePort = true
	s.e.HideBanner = true
	s.e.Debug = true

	s.e.Use(
		middlewares.AddonInfo(l),
		middlewares.Pagination(),
		// middlewares.CheckAuth(auth, s.l),
	)

	s.setupRoutes()
	return s
}

func (s *Server) Start(addr string) error {
	return s.e.Start(addr)
}

func (s *Server) Close(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
