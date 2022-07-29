package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Tender = tenderApi{}

type tenderApi struct{}

// 登录接口
type submitReq struct {
	Time string `json:"time"`
}

func (*tenderApi) Submit(r *ghttp.Request) {
	//var a submitReq
	//reqData, _ := io.ReadAll(r.Body)
	//_ = json.Unmarshal(reqData, &a)
	//timeArr := strings.Split(a.Time, ":")
	//if len(timeArr) != 2 {
	//	r.Response.Status = 500
	//	_ = r.Response.WriteJsonExit(&Response{
	//		Code: -7,
	//		Data: "time format error",
	//	})
	//	return
	//}
	//h, err1 := strconv.Atoi(timeArr[0])
	//m, err2 := strconv.Atoi(timeArr[1])
	//if err1 != nil || err2 != nil {
	//	r.Response.Status = 500
	//	_ = r.Response.WriteJsonExit(&Response{
	//		Code: -8,
	//		Data: "time format error",
	//	})
	//	return
	//}

}
