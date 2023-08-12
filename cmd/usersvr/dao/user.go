package dao

import (
	"errors"
	"github.com/shixiaocaia/tiktok/cmd/usersvr/constant"
	"github.com/shixiaocaia/tiktok/cmd/usersvr/dao/mysql"
	"github.com/shixiaocaia/tiktok/cmd/usersvr/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserNameIsExist(username string) (bool, error) {
	db := mysql.GetDB()
	user := User{}
	err := db.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		if err.Error() != gorm.ErrRecordNotFound.Error() {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func InsertUser(username, password string) (*User, error) {
	db := mysql.GetDB()
	// 加密密文，明文存储密码不安全
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// mysql创建用户
	user := User{
		Name:            username,
		Password:        string(hashPassword),
		Follow:          0,
		Follower:        0,
		TotalFav:        0,
		FavCount:        0,
		WorkCount:       0,
		Avatar:          "https://tse1-mm.cn.bing.net/th/id/R-C.d83ded12079fa9e407e9928b8f300802?rik=Gzu6EnSylX9f1Q&riu=http%3a%2f%2fwww.webcarpenter.com%2fpictures%2fGo-gopher-programming-language.jpg&ehk=giVQvdvQiENrabreHFM8x%2fyOU70l%2fy6FOa6RS3viJ24%3d&risl=&pid=ImgRaw&r=0",
		BackgroundImage: "https://tse2-mm.cn.bing.net/th/id/OIP-C.sDoybxmH4DIpvO33-wQEPgHaEq?pid=ImgDet&rs=1",
		Signature:       "test sign",
	}
	result := db.Model(&User{}).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Infof("create user", user)

	// todo redis缓存

	return &user, nil
}

func GetUserInfo(u interface{}) (User, error) {
	db := mysql.GetDB()
	user := User{}
	var err error

	switch u := u.(type) {
	case int64:
		err = db.Where("id = ?", u).First(&user).Error
	case string:
		err = db.Where("user_name = ?", u).First(&user).Error
	default:
		err = errors.New(constant.ErrorUserInfo)
	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserListInfo(userIdList []int64) ([]*User, error) {
	db := mysql.GetDB()
	var users []*User
	log.Debugf("userIdList: %v", userIdList)
	err := db.Where("id in ?", userIdList).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}
