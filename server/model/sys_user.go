package model

type SysUser struct {
	Uid int `json:"uid" gorm:"comment:用户uid"`
	Username string `json:"username" gorm:"comment:用户名"`
	Password string `json:"password" gorm:"comment:密码"`
	RoleId string `json:"role_id" gorm:"comment:角色id"`

}

func Login() {

}