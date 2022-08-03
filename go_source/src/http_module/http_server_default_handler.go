package http_module

import (
	"kbell/struct_module"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Http_Default_Handler(e *echo.Echo, s *struct_module.ServerStats_ST) {

	e.GET("/", route_main)

	e.GET("/timeout", route_timeout_test)

	e.GET("/stats", s.Server_Stats)

}

func route_main(c echo.Context) error {
	return c.String(http.StatusOK, `sys_controller active`)
}

func route_timeout_test(c echo.Context) error {
	time.Sleep(time.Second * (Config.HttpServer_Config.TimeOut + 1))
	return c.String(http.StatusRequestTimeout, `{ "err_msg" : "timeout" }\n`) // this line not work!!
}
