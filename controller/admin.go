package controller

import (
	"encoding/json"

	"github.com/CoderTH/go_gateway/dao"
	"github.com/CoderTH/go_gateway/dto"
	"github.com/CoderTH/go_gateway/middleware"
	"github.com/CoderTH/go_gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func AdminRegister(router *gin.RouterGroup) {
	admin := AdminController{}
	router.GET("/admin_info", admin.AdminInfo)
	router.POST("/change_pwd", admin.ChangePwd)
}

// Admin godoc
// @Summary 管理员信息获取
// @Description 管理员信息获取
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminlogin *AdminController) AdminInfo(c *gin.Context) {
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	sessInfoStr := sessInfo.(string)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfoStr), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	//1. duqussionKey对应的sjon转换为结构体
	//2. 取出数据
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://www.google.com/imgres?imgurl=https%3A%2F%2Fupload.wikimedia.org%2Fwikipedia%2Fcommons%2Fthumb%2Fe%2Fe8%2FCandymyloveYasu.png%2F220px-CandymyloveYasu.png&imgrefurl=https%3A%2F%2Fzh.wikipedia.org%2Fwiki%2F%25E5%25A4%25B4%25E5%2583%258F&tbnid=lLmpObEzLEyh_M&vet=12ahUKEwjd3qyA8-_0AhWF0IsBHUjWD60QMygAegUIARCaAQ..i&docid=ryEBP9pMcfMEoM&w=220&h=220&itg=1&q=touxiang&ved=2ahUKEwjd3qyA8-_0AhWF0IsBHUjWD60QMygAegUIARCaAQ",
		Introduction: "I am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (adminlogin *AdminController) ChangePwd(c *gin.Context) {
	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	//1. session读取用户信息sessInof
	//2. sessInfo.ID 读取数据库信息adminInfo
	//3. params.password+adminInfo.salt sha256 slatPassword
	//4.数据库保存
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	sessInfoStr := sessInfo.(string)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfoStr), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	//从数据库中读取adminInfo
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, (&dao.Admin{UserName: adminSessionInfo.UserName}))
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	//生成新的加盐密码
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = saltPassword
	//数据库保存
	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
