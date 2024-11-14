package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

// Получение резюме по ID
func (s *Server) getResume(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resume, err := s.hh.ResumeByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resume)
}

// Получение всех резюме
func (s *Server) getAllResumes(ctx echo.Context) error {
	total, resumes, err := s.hh.Resumes(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "resumes": resumes})
}

// Создание резюме
func (s *Server) createResume(ctx echo.Context) error {
	var resume models.Resume
	if err := ctx.Bind(&resume); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.ResumeInsert(ctx.Request().Context(), &resume); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, resume)
}

// Обновление резюме
func (s *Server) updateResume(ctx echo.Context) error {
	var resume models.Resume
	if err := ctx.Bind(&resume); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.ResumeUpdate(ctx.Request().Context(), &resume); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, resume)
}

// Удаление резюме
func (s *Server) deleteResume(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.ResumeDelete(ctx.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
