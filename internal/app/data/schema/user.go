package schema

import (
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// User 用户基础信息
type User struct {
	Model
	//UserId 对外的业务主键
	UserId int64 `gorm:"uniqueIndex"`

	//下面三种信息在库中都是不可重复的
	//UserName 用户名
	UserName string  `gorm:"uniqueIndex"`
	Email    *string `gorm:"uniqueIndex"`
	Phone    *string `gorm:"uniqueIndex"`

	PassWordMd5 string
}

var (
	mId uint16 = 0
)

func GetMachineId() (uint16, error) {
	return mId, nil
}

// BeforeCreate 雪花算法初始化业务主键
func (m *User) BeforeCreate(tx *gorm.DB) (err error) {

	st := sonyflake.Settings{
		StartTime: time.Time{},
		MachineID: GetMachineId,
		CheckMachineID: func(u uint16) bool {
			return true
		},
	}
	sf := sonyflake.NewSonyflake(st)
	id, err := sf.NextID()
	if err != nil {
		return err
	}
	m.UserId = int64(id)
	m.UserName = "pero" + strconv.FormatInt(m.UserId, 10)
	return
}
