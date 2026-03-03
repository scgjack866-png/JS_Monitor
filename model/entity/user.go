package entity

import (
	"time"
)

type User struct {
	// 主键
	ID uint64 `gorm:"Column:id;type:bigint;PRIMARY_KEY;AUTO_INCREMENT" mson:"UserId"`
	// 用户名
	UserName string `gorm:"Column:username"`
	// 昵称
	NickName string `gorm:"Column:nickname"  mson:"NickName"`
	// 性别(1:男;2:女)
	Gender int `gorm:"Column:gender"`
	// 用户头像
	Avatar string `gorm:"Column:avatar" mson:"Avatar"`
	// 联系方式
	Mobile string `gorm:"Column:mobile"`
	// 用户状态(1:正常;0:禁用)
	Status *int `gorm:"Column:status;default:1" mson:"status"`
	// 邮箱
	Email string `gorm:"Column:email"`
	// 密码
	Password string `gorm:"Column:password;NOT NULL"`
	// 部门ID
	DeptId uint64 `gorm:"Column:dept_id;NOT NULL"`
	// 创建时间
	CreateTime time.Time `gorm:"Column:create_time;NOT NULL" mson:"createTime"`
	// 更新时间
	UpdateTime time.Time `gorm:"Column:update_time;NOT NULL"`
	// 逻辑删除标识(0:未删除;1:已删除)
	Deleted int `gorm:"Column:deleted"`

	Roles []Role `gorm:"many2many:sys_user_role"`
}

// 指定表名
func (User) TableName() string {
	return "sys_user"
}
