package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) getVacancy(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	vacancy, err := s.hh.JobVacancyByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, vacancy)
}

func (s *Server) getAllVacancies(ctx echo.Context) error {
	total, vacancies, err := s.hh.JobVacancies(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "vacancies": vacancies})
}

func (s *Server) createVacancy(ctx echo.Context) error {
	var vacancy models.JobVacancy
	if err := ctx.Bind(&vacancy); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.JobVacancyInsert(ctx.Request().Context(), &vacancy); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, vacancy)
}

func (s *Server) updateVacancy(ctx echo.Context) error {
	var vacancy models.JobVacancy
	if err := ctx.Bind(&vacancy); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.JobVacancyUpdate(ctx.Request().Context(), &vacancy); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, vacancy)
}

func (s *Server) deleteVacancy(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.JobVacancyDelete(ctx.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
