package api

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ParamID(ctx echo.Context) (uuid.UUID, error) {
	paramID := ctx.Param("id")
	if paramID == "" {
		return uuid.Nil, nil
	}

	id, err := uuid.Parse(paramID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// func RequesterID(ctx echo.Context) (uuid.UUID, error) {
// 	requesterID := ctx.Get(auth.CtxUserID)
// 	if requesterID == nil {
// 		return uuid.Nil, nil
// 	}

// 	userID, ok := requesterID.(string)
// 	if !ok {
// 		return uuid.Nil, errors.New("invalid requesterID type")
// 	}

// 	id, err := uuid.Parse(userID)
// 	if err != nil {
// 		return uuid.Nil, err
// 	}

// 	return id, nil

// }
