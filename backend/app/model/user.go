package model

import (
	"backend/app/util"
)

type User struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Name      string    `gorm:"column:name"                 json:"name"`
	Password  string    `gorm:"column:password"             json:"password"`
	Email     string    `gorm:"column:email"                json:"email"`
	Reward    int       `gorm:"column:reward"               json:"reward"`
	Token     string    `gorm:"column:token"                json:"token"`
}

type UserDisplay struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Name      string    `gorm:"column:name"                 json:"name"`
	Email     string    `gorm:"column:email"                json:"email"`
}

type UserDetailDisplay struct {
	Id        int       `gorm:"primary_key;AUTO_INCREMENT"  json:"id"`
	Name      string    `gorm:"column:name"                 json:"name"`
	Email     string    `gorm:"column:email"                json:"email"`
	Reward    int       `gorm:"column:reward"               json:"reward"`
}

func GetUserTable() string {
	return "user"
}

func LoginUser(name string, password string) (User, error) {
	encryptPassword := util.Encrypt(password)

	user := User{}
	err := DB().Table(GetUserTable()).Where("name = ? AND password = ?", name, encryptPassword).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetAllUsers() ([]User, error) {
	user := []User{}
	err := DB().Table(GetUserTable()).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(id int) (User, error) {
	user := User{}
	err := DB().Table(GetUserTable()).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func AddUser(name string, email string, password string, reward int) (User, error) {
	encryptPassword := util.Encrypt(password)
	user := User{Name: name, Email: email, Password: encryptPassword, Reward: reward, Token: name + util.TokenGenerator()}
	err := DB().Table(GetUserTable()).Create(&user).Error

	return user, err
}

func UpdateUser(id int, name string, email string, password string, reward int) (User, error) {
	user := User{}
	user, err := GetUser(id)
	if err != nil {
		return user, err
	}

	err = DB().Table(GetUserTable()).Where("id = ?", id).Update("name", name).Error
	err = DB().Table(GetUserTable()).Where("id = ?", id).Update("email", email).Error
	err = DB().Table(GetUserTable()).Where("id = ?", id).Update("password", password).Error
	err = DB().Table(GetUserTable()).Where("id = ?", id).Update("reward", reward).Error

	user.Name = name
	user.Email = email
	user.Password = password
	user.Reward = reward

	return user, err
}

func DeleteUser(id int) (User, error) {
	user := User{}
	err := DB().Table(GetUserTable()).Where("id = ?", id).Delete(&user).Error

	return user, err
}
