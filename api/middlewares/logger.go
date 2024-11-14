package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
)

func AddonInfo(l Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			request := c.Request()

			err := next(c)

			fields := map[string]interface{}{
				"latency":     int(time.Since(start) / time.Millisecond),
				"path":        request.URL.Path,
				"method":      request.Method,
				"ctx-path":    c.Path(),
				"unique-path": request.Method + "-" + c.Path(),
			}

			switch err := err.(type) {
			case *echo.HTTPError:
				fields["status"] = err.Code
				fields["error"] = err.Error()
			default:
				fields["status"] = c.Response().Status
			}

			l.SendWithFields(fields)
			return err
		}
	}
}
