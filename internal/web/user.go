package web

import (
	"errors"

	"syntax/global/code"
	"syntax/internal/model"
	"syntax/internal/service"
	"syntax/utils/pkg"

	"github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) RegisterRouter(serve *gin.Engine) {
	ug := serve.Group("/user")
	ug.POST("/signup", u.Signup)
	ug.POST("/login", u.Login)
	ug.POST("/profile", u.Profile)
	ug.POST("/edit", u.Edit)
}

const (
	EmailRegexpPattern    = "^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
	PasswordRegexpPattern = "^(?=.*[a-zA-Z])(?=.*\\d)(?=.*[!@#$%^&*()`~_+])[a-zA-Z\\d!@#$%^&`~*()_+]{8,}$"
)

type UserHandler struct {
	emailRexExp    *regexp2.Regexp
	passwordRexExp *regexp2.Regexp
	svc            *service.UserService
}

func NewUserHandle(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRexExp:    regexp2.MustCompile(EmailRegexpPattern, regexp2.None),
		passwordRexExp: regexp2.MustCompile(PasswordRegexpPattern, regexp2.None),
		svc:            svc,
	}
}

// Signup 注册逻辑管控 (register logic controller area)
func (u *UserHandler) Signup(c *gin.Context) {
	type SignUpReq struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		Repassword string `json:"repassword"`
	}
	var signUserData SignUpReq

	err := c.BindJSON(&signUserData)
	// 绑定失败执行返回处理.
	if err != nil {
		pkg.Failure(c, code.BindError, "内部错误")
		return
	}
	isEmail, err := u.emailRexExp.MatchString(signUserData.Email)
	// 正则检查邮箱是否合法时候正则是否出错
	if err != nil {
		pkg.Failure(c, code.SYSInsideError, "内部错误")
		return
	}
	// 验证邮箱格式
	if !isEmail {
		pkg.Failure(c, code.EmailPatternError, "非法邮箱格式")
		return
	}
	// 检查用户名和密码是否合法. 双密码是否相同.
	if signUserData.Repassword != signUserData.Password {
		pkg.Failure(c, code.BindError, "两次输入的密码不一致")
		return
	}
	// 正则检查密码是否符合规则 (大于6位,最少一个大写字母和一个小写字母和若干数字.)
	ISPasswordExist, err := u.passwordRexExp.MatchString(signUserData.Password)
	if err != nil {
		pkg.Failure(c, code.SYSInsideError, "内部错误")
		return
	}
	if !ISPasswordExist {
		pkg.Failure(c, code.PasswordPatternError, "密码必须包含字母,数字,特殊符号,并且不少于八位.")
		return
	}
	err = u.svc.Signup(c, &model.User{
		Email:    signUserData.Email,
		Password: signUserData.Password,
	})
	switch {
	case err == nil:
		pkg.Success(c, "注册成功", nil)
	case errors.Is(err, service.EmailReduplicateError):
		pkg.Failure(c, code.EmailExistError, "该邮箱已被注册")
	default:
		pkg.Failure(c, code.SYSInsideError, "内部错误")
	}
}

func (this *UserHandler) Login(c *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	err := c.BindJSON(&req)
	if err != nil {
		pkg.Failure(c, code.BindError, "内部错误")
		return
	}
	u, err := this.svc.Login(c, req.Email, req.Password)
	switch {
	case err == nil:
		session := sessions.Default(c)
		session.Set("session_id", u.Id)
		session.Options(sessions.Options{MaxAge: 15 * 60})
		err := session.Save()
		if err != nil {
			pkg.Failure(c, code.SYSInsideError, "内部错误")
			return
		}
		pkg.Success(c, "登录成功", nil)
		return
	case errors.Is(err, service.ErrorInvalidUserOrPassword):
		pkg.Failure(c, code.InvalidUserOrPassword, "用户名或密码错误")
		return
	default:
		pkg.Failure(c, code.SYSInsideError, "内部错误")
		return
	}
}

func (this *UserHandler) Edit(c *gin.Context) {
	type Req struct {
		Email     string `json:"email"`
		NickName  string `json:"nickname"`
		Birthday  int64  `json:"birthday"`
		Introduce string `json:"introduce"`
	}
	var req Req
	err := c.BindJSON(&req)
	if err != nil {
		pkg.Failure(c, code.BindError, "内部错误bind")
		return
	}
	_, err = this.svc.Edit(c, req.Email, req.NickName, req.Birthday, req.Introduce)
	switch {
	case err == nil:
		pkg.Success(c, "获取成功", nil)
	default:
		pkg.Failure(c, code.SYSInsideError, "内部错误edit")
		return
	}
}

func (this *UserHandler) Profile(c *gin.Context) {
	type Req struct {
		Email string `json:"email"`
	}
	var req Req
	err := c.BindJSON(&req)
	if err != nil {
		pkg.Failure(c, code.BindError, "内部错误")
		return
	}
	u, err := this.svc.Profile(c, req.Email)
	switch {
	case err == nil:
		pkg.Success(c, "获取成功", u)
	default:
		pkg.Failure(c, code.SYSInsideError, "内部错误")
		return
	}
}
