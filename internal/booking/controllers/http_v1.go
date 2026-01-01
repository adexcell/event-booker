package event

import (
	"github.com/adexcell/event-booker/internal/controller"
	"github.com/wb-go/wbf/ginext"
)

type handler struct{

}

func New() controller.Handler {
	return &handler{}
}

func (h *handler) Register(router *ginext.Engine) {
	router.POST("/events", h.CreateEvent)
	router.POST("/events/:id/book", h.Book)
	router.POST("/events/:id/confirm", h.Payment)
	router.GET("/events/:id", h.GetInfo)
}

func (h *handler) CreateEvent(c *ginext.Context) {

}

func (h *handler) Book(c *ginext.Context) {

}

func (h *handler) Payment(c *ginext.Context) {

}

func (h *handler) GetInfo(c *ginext.Context) {

}

