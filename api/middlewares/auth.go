package middlewares

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/Sayonara143/urfu-pp-hhru-backend/services/auth"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func CheckAuth(auther Authorizer, log Logger) echo.MiddlewareFunc {
// 	return middleware.BasicAuth(func(u, p string, ctx echo.Context) (bool, error) {
// 		user, err := auther.BasicAuth(ctx.Request().Context(), u, p)
// 		if err != nil || user == nil {
// 			log.Error(fmt.Errorf("basic auth failed by username: %s", u), "wrong login or password")
// 			return false, echo.NewHTTPError(http.StatusUnauthorized, "auth failed: wrong username or password")
// 		}
// 		ctx.Set(auth.CtxUserID, (user.ID).String())
// 		return true, nil

// 	})
// }
