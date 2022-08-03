package http_module

import (
	"kbell/struct_module"
	"kbell/task_module"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Http_Sys_Controller_Handler(e *echo.Echo) {

	//REST JSON Sample(system_type = openstack)
	//=> curl -XPOST http://ip:port/system/control/openstack -H 'Content-Type: application/json' -d '{"id":"id-001" , "runtype" : "start"}'
	//REST FORM Sample(system_type = openstack)
	//=> curl -XPOST http://ip:port/system/control/openstack -d 'id=id001' -d 'runtype=start'
	e.POST("/system/control/:system_type", recv_sys_control)

	e.GET("/system/state/:system_type/:id", recv_sys_status)

}

func recv_sys_control(c echo.Context) (err error) {

	system_type := c.Param("system_type")

	con_req_st := new(struct_module.Control_Req_ST)
	err = c.Bind(con_req_st)
	if (err != nil) || (con_req_st.Id == "" && con_req_st.RunType == "") {
		return c.JSON(http.StatusBadRequest,
			struct_module.Error_Server_Msg(struct_module.Error_Parse, system_type))
	}

	res := task_module.Sys_Control_Task(system_type, con_req_st.RunType, con_req_st.Id)
	con_rep_st := set_rep_sys_control(system_type, con_req_st, res)
	show_sys_control_task_res(c, system_type, con_req_st, res)

	if res {
		return c.JSON(http.StatusOK, con_rep_st)
	} else {
		return c.JSON(http.StatusBadRequest, con_rep_st)
	}
}

func show_sys_control_task_res(c echo.Context, system_type string, con_req_st *struct_module.Control_Req_ST, res_flag bool) {
	Logger.Info("--------------------------------------------------------------")
	Logger.Info("Http Reponse Res (", c.Request().RemoteAddr, ") ::: system_type : ", system_type, ", id : ", con_req_st.Id, ", runtype : ", con_req_st.RunType, ", res : ", res_flag)
	Logger.Info("--------------------------------------------------------------")
}

func set_rep_sys_control(system_type string, con_req_st *struct_module.Control_Req_ST, res bool) *struct_module.Control_Rep_ST {

	con_rep_st := new(struct_module.Control_Rep_ST)
	con_rep_st.Time = time.Now()
	con_rep_st.Id = con_req_st.Id
	con_rep_st.RunType = con_req_st.RunType
	con_rep_st.SysType = system_type
	con_rep_st.Res = res

	return con_rep_st
}

func recv_sys_status(c echo.Context) (err error) {

	//system_type := c.Param("system_type")
	//target_id := c.Param("id")
	//status_res := task_module.Sys_Status_Task(system_type, target_id)
	return c.String(http.StatusServiceUnavailable, `{ "err_msg" : "not support func" }}\n`)
}
