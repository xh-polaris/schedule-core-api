// Code generated by hertz generator. DO NOT EDIT.

package core_api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	core_api "github.com/xh-polaris/schedule-core-api/biz/adaptor/controller/core_api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_schedule := root.Group("/schedule", _scheduleMw()...)
		_schedule.POST("/create", append(_createscheduleMw(), core_api.CreateSchedule)...)
		_schedule.POST("/create_list", append(_createschedulesMw(), core_api.CreateSchedules)...)
		_schedule.POST("/create_ori", append(_createschedulefromoriginMw(), core_api.CreateScheduleFromOrigin)...)
		_schedule.GET("/delete", append(_deletescheduleMw(), core_api.DeleteSchedule)...)
		_schedule.POST("/list", append(_getschedulesMw(), core_api.GetSchedules)...)
		_schedule.POST("/update", append(_updatescheduleMw(), core_api.UpdateSchedule)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/get_info", append(_getuserinfoMw(), core_api.GetUserInfo)...)
		_user.POST("/send_verify_code", append(_sendverifycodeMw(), core_api.SendVerifyCode)...)
		_user.POST("/sign_in", append(_signinMw(), core_api.SignIn)...)
		_user.POST("/sign_up", append(_signupMw(), core_api.SignUp)...)
		_user.POST("/update_info", append(_updateuserinfoMw(), core_api.UpdateUserInfo)...)
	}
}
