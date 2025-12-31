package main

import (
	"github.com/adexcell/event-booker/internal/event"
	"github.com/wb-go/wbf/ginext"
	"github.com/wb-go/wbf/zlog"
)

func main() {
	zlog.Init()

	httprouter := ginext.New("debug")

	eventHandler := event.New()
	eventHandler.Register(httprouter)
}
