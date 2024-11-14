package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

// Получение профиля студента по ID
func (s *Server) getStudentProfile(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	profile, err := s.hh.StudentProfileByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, profile)
}

// Получение всех профилей студентов
func (s *Server) getAllStudentProfiles(ctx echo.Context) error {
	total, profiles, err := s.hh.StudentProfiles(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "profiles": profiles})
}

// Обновление профиля студента
func (s *Server) updateStudentProfile(ctx echo.Context) error {
	var profile models.StudentProfile
	if err := ctx.Bind(&profile); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.StudentProfileUpdate(ctx.Request().Context(), &profile); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, profile)
}
