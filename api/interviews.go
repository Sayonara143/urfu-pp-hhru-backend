package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

func (s *Server) getInterview(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	interview, err := s.hh.InterviewByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, interview)
}

func (s *Server) getAllInterviews(ctx echo.Context) error {
	total, interviews, err := s.hh.Interviews(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "interviews": interviews})
}

func (s *Server) createInterview(ctx echo.Context) error {
	var interview models.Interview
	if err := ctx.Bind(&interview); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.InterviewInsert(ctx.Request().Context(), &interview); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, interview)
}

func (s *Server) updateInterview(ctx echo.Context) error {
	var interview models.Interview
	if err := ctx.Bind(&interview); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.InterviewUpdate(ctx.Request().Context(), &interview); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, interview)
}

func (s *Server) deleteInterview(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.InterviewDelete(ctx.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
