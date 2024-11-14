package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/Sayonara143/urfu-pp-hhru-backend/pagination"
	"github.com/labstack/echo/v4"
)

// Получение заявки по ID (GET /applications/:id)
func (s *Server) getApplication(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	application, err := s.hh.JobApplicationByID(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, application)
}

// Получение всех заявок с поддержкой пагинации (GET /applications)
func (s *Server) getAllApplications(ctx echo.Context) error {
	p := pagination.Get(ctx.Request().Context())

	total, applications, err := s.hh.JobApplications(ctx.Request().Context(), p.Limit(), p.Offset())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	p.SetItemsTotal(total)

	response := responseWrapper{
		Response:   applications,
		Pagination: p.PaginationResponse(),
	}

	return ctx.JSON(http.StatusOK, response)
}

// Создание новой заявки (POST /applications)
func (s *Server) createApplication(ctx echo.Context) error {
	var application models.JobApplication
	if err := ctx.Bind(&application); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.JobApplicationInsert(ctx.Request().Context(), &application); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, application)
}

// Обновление заявки (PUT /applications/:id)
func (s *Server) updateApplication(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var application models.JobApplication
	if err := ctx.Bind(&application); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	application.ID = id

	if err := s.hh.JobApplicationUpdate(ctx.Request().Context(), &application); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, application)
}

// Удаление заявки (DELETE /applications/:id)
func (s *Server) deleteApplication(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.JobApplicationDelete(ctx.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
