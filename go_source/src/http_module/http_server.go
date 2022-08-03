package http_module

import (
	"fmt"
	"kbell/struct_module"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func HttpServerStart() {

	echo, stats_st := server_init()
	server_reg_handlers(echo, stats_st)
	http_server_start(echo)
}

func create_server_status() *struct_module.ServerStats_ST {
	return &struct_module.ServerStats_ST{
		Uptime:         time.Now(),
		Statuses:       map[string]uint64{},
		LastAccessInfo: map[string]string{},
	}
}

func server_err_handler(err error, c echo.Context) {

	code := http.StatusInternalServerError

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	err_msg := struct_module.Error_Server_Msg(struct_module.Error_System, err.Error())

	c.JSON(code, err_msg)
}

func server_init() (*echo.Echo, *struct_module.ServerStats_ST) {

	echo := echo.New()
	stats_st := create_server_status()

	echo.Debug = false
	echo.HideBanner = true
	echo.HidePort = true

	echo.Use(stats_st.Set_Header)
	echo.Use(stats_st.Set_TimeOut(Config.HttpServer_Config.TimeOut))
	echo.Use(stats_st.Set_Stats)
	echo.HTTPErrorHandler = server_err_handler

	return echo, stats_st
}

func server_reg_handlers(echo *echo.Echo, stats_st *struct_module.ServerStats_ST) {
	Http_Default_Handler(echo, stats_st)
	Http_Sys_Controller_Handler(echo)
}

func http_server_start(e *echo.Echo) {

	ssh_active_flag := Config.HttpServer_Config.SSL_Active
	host_addr := fmt.Sprintf(":%d", Config.HttpServer_Config.Port)

	if ssh_active_flag {

		crt_path := Config.HttpServer_Config.CrtCert_path
		key_path := Config.HttpServer_Config.PrivateKey_path

		Logger.Info("Active Https >>>>> ")
		Logger.Info("crt_path : ", crt_path)
		Logger.Info("key_path : ", key_path)

		// Https Server
		if err := e.StartTLS(host_addr, crt_path, key_path); err != http.ErrServerClosed {
			Logger.Fatal(err)
		}

	} else {

		// Http Server
		Logger.Fatal(e.Start(host_addr))

	}

}
