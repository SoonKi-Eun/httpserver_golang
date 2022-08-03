package struct_module

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// json to go struct : https://mholt.github.io/json-to-go/

// Config Struct ===========================================================================================

type Config_ST struct {
	Project_Config    Project_Config_ST    `yaml:"project_info"`
	HttpServer_Config HttpServer_Config_ST `yaml:"httpserver_info"`
}

type Project_Config_ST struct {
	Max_Cpu_Use  int    `yaml:"max_cpu"`
	Project_name string `yaml:"project_name"`
	Project_path string `yaml:"project_path"`
}

type HttpServer_Config_ST struct {
	Port            int           `yaml:"port"`
	TimeOut         time.Duration `yaml:"timeout"`
	SSL_Active      bool          `yaml:"ssl_active"`
	PrivateKey_path string        `yaml:"private_key_path"`
	CrtCert_path    string        `yaml:"crt_certificate_path"`
}

// HTTP Sever ST =====================================================================================
const server_name = "sys_controller"

type (
	ServerStats_ST struct {
		Uptime         time.Time         `json:"uptime"`
		ServerName     string            `json:"servername"`
		Timeout        time.Duration     `json:"settimeout"`
		RequestCount   uint64            `json:"requestcount"`
		Statuses       map[string]uint64 `json:"statuses"`
		LastAccessInfo map[string]string `json:"lastaccess"`
		mutex          sync.RWMutex
	}
)

func (s *ServerStats_ST) Set_Stats(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		s.mutex.Lock()
		defer s.mutex.Unlock()

		if s.RequestCount >= math.MaxUint64 {
			s.RequestCount = 0
		}
		s.RequestCount++

		status := strconv.Itoa(c.Response().Status)
		if s.Statuses[status] >= math.MaxUint64 {
			s.Statuses[status] = 0
		}

		s.Statuses[status]++

		last_info := fmt.Sprintf("%d-%s-%s-%s", c.Response().Status, c.Request().Method, c.Request().RemoteAddr, c.Request().URL)
		s.LastAccessInfo[c.RealIP()] = last_info

		return nil
	}
}

func (s *ServerStats_ST) Set_Header(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s.ServerName = server_name
		c.Response().Header().Set(echo.HeaderServer, server_name)
		return next(c)
	}
}

func (s *ServerStats_ST) Set_TimeOut(timeout_sec time.Duration) echo.MiddlewareFunc {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Timeout = timeout_sec
	err_msg := fmt.Sprintf(`{ "err_msg" : "handler %s"}`, "timeout")
	return middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: err_msg,
		Timeout:      timeout_sec * time.Second,
	})

}

func (s *ServerStats_ST) Server_Stats(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

const (
	Error_API_Bug = 0
	Error_Unknown = 1
	Error_Parse   = 2
	Error_Timeout = 3
	Error_System  = 9
)

const (
	Error_Msg_Form_00 = `sys call api return bug : %s`
	Error_Msg_Form_01 = `unknown system type %s`
	Error_Msg_Form_02 = `parse err system type %s`
	Error_Msg_Form_03 = `handler %s`
	Error_Msg_Form_09 = `%s`
)

func Error_Server_Msg(err_msg_type int, msg string) *Err_Rep_ST {

	err_rep_st := new(Err_Rep_ST)

	switch err_msg_type {
	case 0:
		err_rep_st.Msg = fmt.Sprintf(Error_Msg_Form_00, msg)
	case 1:
		err_rep_st.Msg = fmt.Sprintf(Error_Msg_Form_01, msg)
	case 2:
		err_rep_st.Msg = fmt.Sprintf(Error_Msg_Form_02, msg)
	case 3:
		err_rep_st.Msg = fmt.Sprintf(Error_Msg_Form_03, msg)
	default:
		err_rep_st.Msg = fmt.Sprintf(Error_Msg_Form_09, msg)
	}

	return err_rep_st
}

// HTTP Sever Sys Controller ST =====================================================================================

type (
	Control_Req_ST struct {
		Id      string `json:"id" form:"id" query:"id"`
		RunType string `json:"runtype" form:"runtype" query:"runtype"`
	}
)

type (
	Control_Rep_ST struct {
		Time    time.Time `json:"time"`
		Res     bool      `json:"result"`
		SysType string    `json:"system_type"`
		RunType string    `json:"runtype"`
		Id      string    `json:"id"`
	}
)

type (
	Err_Rep_ST struct {
		Msg string `json:"err_msg"`
	}
)
