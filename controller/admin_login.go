package controller

import (
	"encoding/json"
	"time"

	"github.com/CoderTH/go_gateway/dao"
	"github.com/CoderTH/go_gateway/dto"
	"github.com/CoderTH/go_gateway/middleware"
	"github.com/CoderTH/go_gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct{}

func AdminLoginRegister(router *gin.RouterGroup) {
	adminLogin := AdminLoginController{}
	router.POST("/login", adminLogin.AdminLogin)
	router.GET("/logout", adminLogin.AdminLoginOut)
}

// AdminLogin godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	admin := &dao.Admin{}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	//设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	sess.Save()

	out := dto.AdminLoginOutput{
		Token: admin.UserName,
	}
	middleware.ResponseSuccess(c, out)
}

// AdminLogin godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [get]
func (adminlogin *AdminLoginController) AdminLoginOut(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "")
}
