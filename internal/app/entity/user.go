package entity

type UserLogin struct {
	//UserId 对外的业务主键
	UserId string `gorm:"uniqueIndex"`

	//下面三种信息在库中都是不可重复的
	//UserName 用户名
	UserName string  `gorm:"uniqueIndex"`
	Email    *string `gorm:"uniqueIndex"`
	Phone    *string `gorm:"uniqueIndex"`

	PassWordMd5 string
}
