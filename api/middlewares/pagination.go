package middlewares

import (
	"strconv"

	"github.com/Sayonara143/urfu-pp-hhru-backend/pagination"
	"github.com/labstack/echo/v4"
)

func Pagination() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			q := ctx.Request().URL.Query()
			p := pagination.New()

			page, err := strconv.Atoi(q.Get("page"))
			if err == nil {
				p.SetPage(page)
			}

			perPage, err := strconv.Atoi(q.Get("per-page"))
			if err == nil {
				p.SetPerPage(perPage)
			}

			ctx.SetRequest(ctx.Request().WithContext(pagination.Set(ctx.Request().Context(), p)))

			return next(ctx)
		}
	}
}
