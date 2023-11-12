package routes

import (
	"raihpeduli/features/fundraise"
	"raihpeduli/helpers"
	"raihpeduli/middlewares"

	"github.com/labstack/echo/v4"
)

func Fundraises(e *echo.Echo, handler fundraise.Handler, jwt helpers.JWTInterface) {
	fundraises := e.Group("/fundraises")

	fundraises.GET("", handler.GetFundraises())
	fundraises.POST("", handler.CreateFundraise(), middlewares.AuthorizeJWT(jwt))

	fundraises.GET("/:id", handler.FundraiseDetails())
	fundraises.PUT("/:id", handler.UpdateFundraise())
	fundraises.DELETE("/:id", handler.DeleteFundraise())
}
