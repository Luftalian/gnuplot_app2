package handler

import (
	"github.com/ras0q/go-backend-template/internal/repository"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) SetupRoutes(api *echo.Group) {
	// ping API
	pingAPI := api.Group("/ping")
	{
		pingAPI.GET("", h.Ping)
	}

	// user API
	userAPI := api.Group("/users")
	{
		userAPI.GET("", h.GetUsers)
		userAPI.POST("", h.CreateUser)
		userAPI.GET("/:userID", h.GetUser)
	}

	// files API
	filesAPI := api.Group("/files")
	{
		filesAPI.POST("/edit", h.EditFigure)
	}

	// upload API
	uploadAPI := api.Group("/upload")
	{
		uploadAPI.POST("", h.UploadFile)
	}
}
