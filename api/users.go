package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) getUser(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := s.hh.UserByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (s *Server) getAllUsers(ctx echo.Context) error {
	total, users, err := s.hh.Users(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "users": users})
}

func (s *Server) updateUser(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.UserUpdate(ctx.Request().Context(), &user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, user)
}

func (s *Server) deleteUser(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.UserDelete(ctx.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
