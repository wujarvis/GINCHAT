package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Salt          string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     string
	HeartBeatTime string
	LogoutTime    string
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(&UserBasic{}).Create(user)
}

func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(&UserBasic{}).Delete(user)
}

func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(&UserBasic{}).Where("id = ?", user.ID).Updates(UserBasic{
		Name:     user.Name,
		Password: user.Password,
	})
}

func FindUserByName(name string) *UserBasic {
	user := UserBasic{}
	first := utils.DB.Where("name = ?", name).First(&user)
	if first.Error == nil {
		// token
		stamp := fmt.Sprintf("%d", time.Now().Unix())
		tmp := utils.Md5Encode(stamp)
		utils.DB.Model(&user).Where("name = ?", name).Update("identity", tmp)
		return &user
	} else {
		return nil
	}
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}
