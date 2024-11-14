package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

// Получение профиля работодателя по ID
func (s *Server) getEmployerProfile(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	profile, err := s.hh.EmployerProfileByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, profile)
}

// Получение всех профилей работодателей
func (s *Server) getAllEmployerProfiles(ctx echo.Context) error {
	total, profiles, err := s.hh.EmployerProfiles(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "profiles": profiles})
}

// Создание профиля работодателя
func (s *Server) createEmployerProfile(ctx echo.Context) error {
	var profile models.EmployerProfile
	if err := ctx.Bind(&profile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.EmployerProfileInsert(ctx.Request().Context(), &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, profile)
}

// Обновление профиля работодателя
func (s *Server) updateEmployerProfile(ctx echo.Context) error {
	var profile models.EmployerProfile
	if err := ctx.Bind(&profile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.EmployerProfileUpdate(ctx.Request().Context(), &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, profile)
}
