package api

import (
	"encoding/json"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"io"
	"toubiaogo/app/dao"
	"toubiaogo/app/model"
)

var Login = LoginApi{}

type LoginApi struct{}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

// 登录接口
type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*LoginApi) Index(r *ghttp.Request) {
	var a loginReq
	reqData, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(reqData, &a)
	user, err := dao.User.FindOne("username = ?", a.Username)
	if err != nil || user == nil {
		r.Response.Status = 401
		_ = r.Response.WriteJsonExit(&Response{
			Code: -1,
			Data: "can't find user",
		})
		return
	}
	if user.Password != a.Password {
		r.Response.Status = 401
		_ = r.Response.WriteJsonExit(&Response{
			Code: -2,
			Data: "password wrong",
		})
		return
	}
	_ = r.Session.Set("uid", user.Id)
	_ = r.Session.Set("username", user.Username)
	_ = r.Session.Set("time", gtime.Timestamp())
	_ = r.Response.WriteJson(&Response{
		Code: 0,
		Data: "login ok",
	})
}

// 注册接口
type registerReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (*LoginApi) Register(r *ghttp.Request) {
	var a registerReq
	reqData, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(reqData, &a)
	if len(a.Username) == 0 || len(a.Password) == 0 {
		r.Response.Status = 500
		_ = r.Response.WriteJsonExit(&Response{
			Code: -3,
			Data: "Username or Password error",
		})
	}
	user, err := dao.User.FindOne("username = ?", a.Username)
	if err != nil {
		r.Response.Status = 500
		_ = r.Response.WriteJsonExit(&Response{
			Code: -4,
			Data: "register fail",
		})
	}
	if user != nil {
		r.Response.Status = 402
		_ = r.Response.WriteJsonExit(&Response{
			Code: -5,
			Data: "user exist",
		})
	}

	_, err = dao.User.Insert(model.User{
		Username: a.Username,
		Password: a.Password,
	})
	if err != nil {
		r.Response.Status = 500
		_ = r.Response.WriteJsonExit(&Response{
			Code: -6,
			Data: "insert mysql error",
		})
	}
	_ = r.Response.WriteJson(&Response{
		Code: 0,
		Data: "register ok",
	})
}
