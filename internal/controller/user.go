package controller

import (
	"github.com/wb-go/wbf/ginext"
)

type userHandler struct {
}

func New() Handler {
	return &userHandler{}
}

func (h *userHandler) Register(router *ginext.Engine) {
	router.POST("/events", h.CreateEvent)
	router.POST("/events/:id/book", h.Book)
	router.POST("/events/:id/confirm", h.Payment)
	router.GET("/events/:id", h.GetInfo)
}

func (h *userHandler) CreateEvent(c *ginext.Context) {

}

func (h *userHandler) Book(c *ginext.Context) {

}

func (h *userHandler) Payment(c *ginext.Context) {

}

func (h *userHandler) GetInfo(c *ginext.Context) {

}
