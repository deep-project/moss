package entity

import (
	"errors"
	"github.com/brianvoe/sjwt"
	"github.com/duke-git/lancet/v2/random"
	"golang.org/x/crypto/bcrypt"
	"moss/infrastructure/utils/timex"
	"time"
)

type Admin struct {
	Username    string         `json:"username"`     // 后台管理用户名
	Password    string         `json:"password"`     // 后台管理用户密码
	JwtKey      string         `json:"jwt_key"`      // jwt key
	LoginExpire timex.Duration `json:"login_expire"` // 登录过期时间
}

func (*Admin) ConfigID() string {
	return "admin"
}

func NewAdmin() *Admin {
	return &Admin{
		JwtKey:      random.RandString(20),
		LoginExpire: timex.Duration{Number: 24, Unit: "hour"},
	}
}

// ResetJwtKey 重置JWT key
func (a *Admin) ResetJwtKey() {
	a.JwtKey = random.RandString(20)
}

// Update 更新管理员配置，密码为空则不更新密码
func (a *Admin) Update(username, password string, loginExpire timex.Duration) error {
	a.LoginExpire = loginExpire
	if password != "" {
		if err := a.UpdatePassword(password); err != nil {
			return err
		}
	}
	return a.UpdateUsername(username)
}

func (a *Admin) UpdateUsername(username string) error {
	if username == "" {
		return errors.New("username is required")
	}
	a.Username = username
	return nil
}

// UpdatePassword 更新密码
func (a *Admin) UpdatePassword(password string) error {
	p, err := a.EncryptPassword(password)
	if err != nil {
		return err
	}
	a.Password = p
	return nil
}

// EncryptPassword 加密密码
func (*Admin) EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password is required")
	}
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// VerifyPassword 验证密码
func (a *Admin) VerifyPassword(val string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(val)) == nil
}

// InitAdministrator 初始化管理员
func (a *Admin) InitAdministrator(username, password string) error {
	if err := a.UpdateUsername(username); err != nil {
		return err
	}
	return a.UpdatePassword(password)
}

// Login 管理员登录
func (a *Admin) Login(username, password string) (token string, err error) {
	if err = a.LoginVerify(username, password); err != nil {
		return
	}
	return a.GenerateJwtToken(), nil
}

// LoginVerify 登陆验证
func (a *Admin) LoginVerify(username, password string) error {
	if username == "" {
		return errors.New("username is required")
	}
	if password == "" {
		return errors.New("password is required")
	}
	if username != a.Username || !a.VerifyPassword(password) {
		return errors.New("username or password error")
	}
	return nil
}

// GenerateJwtToken 生成JWT token
func (a *Admin) GenerateJwtToken() string {
	claims := sjwt.New()
	claims.SetIssuedAt(time.Now())
	if d := a.LoginExpire.Duration(); d > 0 {
		claims.SetExpiresAt(time.Now().Add(d)) // 过期时间
	}
	return claims.Generate([]byte(a.JwtKey))
}

// VerifyJwtToken 验证JWT token
func (a *Admin) VerifyJwtToken(token string) bool {
	if !sjwt.Verify(token, []byte(a.JwtKey)) {
		return false
	}
	claims, err := sjwt.Parse(token)
	if err != nil {
		return false
	}
	if err := claims.Validate(); err != nil { // 验证是否过期
		return false
	}
	return true
}
