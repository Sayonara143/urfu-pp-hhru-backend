package api

import (
	"net/http"

	"github.com/Sayonara143/urfu-pp-hhru-backend/models"
	"github.com/labstack/echo/v4"
)

// Получение логов поиска
func (s *Server) getSearchLogs(ctx echo.Context) error {
	total, logs, err := s.hh.SearchLogs(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "logs": logs})
}

// Получение сохранённых поисковых запросов
func (s *Server) getSavedSearches(ctx echo.Context) error {
	total, searches, err := s.hh.SavedSearches(ctx.Request().Context(), 0, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"total": total, "searches": searches})
}

// Сохранение поискового запроса
func (s *Server) saveSearchQuery(ctx echo.Context) error {
	var search models.SavedSearch
	if err := ctx.Bind(&search); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.SavedSearchInsert(ctx.Request().Context(), &search); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, search)
}

// Удаление сохранённого поискового запроса
func (s *Server) deleteSavedSearchQuery(ctx echo.Context) error {
	id, err := ParamID(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := s.hh.SavedSearchDelete(ctx.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusOK)
}
