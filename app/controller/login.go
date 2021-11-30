package controller

import (
	"go-api/app/request"
	"go-api/app/services"
	"go-api/app/services/core"
	"go-api/core/session"
	"go-api/tool"

	"github.com/gin-gonic/gin"
)

// @Summary 登录页
// @Description  登录页
// @Router /admin/login [get]
func Login(c *gin.Context) {
	data := tool.M{
		"titles":     "管理后台",
		"captchaUrl": "/api/captcha",
	}
	tool.HTML(c, "common/login.html", data)
}

// @Summary 验证码
// @Description  验证码
// @Router /api/captcha [get]
func Captcha(c *gin.Context) {
	if ret := core.Captcha(); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		tool.JSONP(c, 0, ret.GetMsg(), ret["data"])
	}
}

// @Summary 登录
// @Description  登录
// @Router /admin/Loginin [post]
func Loginin(c *gin.Context) {
	var login request.Login
	c.ShouldBind(&login)
	if err := request.Verify(login, request.LoginVerify); err != nil {
		tool.JSONP(c, 400, err.Error(), nil)
		return
	}

	//图片验证码验证
	if !core.Verify(login.CaptchaId, login.Captcha) {
		tool.JSONP(c, tool.CAPTCHA_ERROR, "验证码验证失败", nil)
		return
	}

	if ret := services.Login(login); !ret.GetStatus() {
		tool.JSONP(c, 40001, ret.GetMsg(), ret["data"])
		return
	} else {
		//c.SetCookie("uid", tool.StructToJson(ret["data"]), 86400, "/", "", false, true)
		//tool.NewSecureCookie(c).SetCookie("uid", tool.StructToJson(ret["data"]), 86400, "/", "", false, true)
		s := session.Default(c)
		s.Set("waitUse", tool.StructToJson(ret["data"]))
		s.Save()
		tool.JSONP(c, 0, ret.GetMsg(), tool.M{"url": "/admin/index"})
	}
}

// @Summary 退出登录
// @Description  退出登录
// @Router /admin/Loginout [post]
func Loginout(c *gin.Context) {
	s := session.Default(c)
	s.Delete("waitUse")
	s.Save()
	//tool.NewSecureCookie(c).SetCookie("uid", "", -1, "/", "", false, true)
	tool.JSONP(c, 0, "退出成功", tool.M{"url": "/admin/login"})
}
