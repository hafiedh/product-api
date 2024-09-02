package server

import (
	"api-product/internal/infrastructure/container"
	"api-product/internal/server/handler"
	"fmt"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/color"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func StartService(container *container.Container) {
	e := echo.New()

	handler.SetupRouter(e, container)

	e.Server.Addr = fmt.Sprintf("%s:%s", container.Config.Apps.Address, container.Config.Apps.HttpPort)

	color.Println(color.Green(fmt.Sprintf("⇨ h2c server started on port: %s\n", container.Config.Apps.HttpPort)))

	// * HTTP/2 Cleartext Server (HTTP2 over HTTP)
	gracehttp.Serve(&http.Server{Addr: e.Server.Addr, Handler: h2c.NewHandler(e, &http2.Server{MaxConcurrentStreams: 500, MaxReadFrameSize: 1048576})})
}
