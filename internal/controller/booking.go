package controller

import (
	"github.com/wb-go/wbf/ginext"
)

type bookingHandler struct {
}

func NewBookingHandler() Handler {
	return &bookingHandler{}
}

func (h *bookingHandler) Register(router *ginext.Engine) {
	router.POST("/events", h.CreateEvent)
	router.POST("/events/:id/book", h.Book)
	router.POST("/events/:id/confirm", h.Payment)
	router.GET("/events/:id", h.GetInfo)
}

func (h *bookingHandler) CreateEvent(c *ginext.Context) {

}

func (h *bookingHandler) Book(c *ginext.Context) {

}

func (h *bookingHandler) Payment(c *ginext.Context) {

}

func (h *bookingHandler) GetInfo(c *ginext.Context) {

}
