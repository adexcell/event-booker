package controller

import (
	"github.com/wb-go/wbf/ginext"
)

type eventHandler struct {
}

func NewevEntHandler() Handler {
	return &eventHandler{}
}

func (h *eventHandler) Register(router *ginext.Engine) {
	router.POST("/events", h.CreateEvent)
	router.POST("/events/:id/book", h.Book)
	router.POST("/events/:id/confirm", h.Payment)
	router.GET("/events/:id", h.GetInfo)
}

func (h *eventHandler) CreateEvent(c *ginext.Context) {

}

func (h *eventHandler) Book(c *ginext.Context) {

}

func (h *eventHandler) Payment(c *ginext.Context) {

}

func (h *eventHandler) GetInfo(c *ginext.Context) {

}
