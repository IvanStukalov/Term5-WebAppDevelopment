package handler

import (
	"log"
	"net/http"
	"github.com/IvanStukalov/Term5-WebAppDevelopment/internal/api"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo api.Repo
}

func NewHandler(repo api.Repo) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping", h.Ping)

	starRouter := r.Group("star") 
	{
		starRouter.GET("/", h.GetStarList)
		starRouter.GET("/:id", h.GetStar)
		starRouter.PUT("/delete/:id", h.DeleteStar)
		starRouter.PUT("/update/:id", h.UpdateStar)
		starRouter.POST("/", h.CreateStar)
	}

	eventRouter := r.Group("event")
	{
		eventRouter.GET("/", h.GetEventList)
		eventRouter.GET("/:id", h.GetEvent)
	}

	// listen and serve on 127.0.0.1:8080
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

// ping
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "pong",
		})
}
