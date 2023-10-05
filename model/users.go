package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blogs    []Blog `json:"blog" form:"blog" gorm:"ForeignKey:UserID;"`
}

type LoginModel struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserModel struct {
	DB *gorm.DB
}

func (usermodel *UserModel) Init(DB *gorm.DB) {
	usermodel.DB = DB
}

func (userModel *UserModel) GetAllUser() []User {
	var listUser []User
	if err := userModel.DB.Find(&listUser).Error; err != nil {
		logrus.Error("Model : Get all user data error, ", err.Error())
		return nil
	}

	return listUser
}

func (userModel *UserModel) GetUserById(id int) User {
	var user User
	user.ID = uint(id)
	if err := userModel.DB.First(&user).Error; err != nil {
		logrus.Error("Model : Get user data error, ", err.Error())
		return User{}
	}

	return user
}

func (userModel *UserModel) CreateUser(newUser User) *User {
	if err := userModel.DB.Create(&newUser).Error; err != nil {
		logrus.Error("Model : Create user data error, ", err.Error())
		return nil
	}

	return &newUser
}

func (userModel *UserModel) UpdateUser(updateUser User) *User {
	var query = userModel.DB.Updates(updateUser)
	if err := query.Error; err != nil {
		logrus.Error("Model : Update user data error, ", err.Error())
		return nil
	}

	if dataCount := query.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
		return &User{}
	}

	return &updateUser
}

func (userModel *UserModel) DeleteUser(id int) {
	var deleteUser User
	deleteUser.ID = uint(id)
	if err := userModel.DB.Delete(&deleteUser).Error; err != nil {
		logrus.Error("Model : Delete user data error, ", err.Error())
	}
}

func (userModel *UserModel) GetUserBlogs() []User {
	var user []User
	if err := userModel.DB.Preload("Blogs").Find(&user).Error; err != nil {
		logrus.Error("Model : Get user blogs data error, ", err.Error())
		return []User{}
	}

	return user
}

func (usermodel *UserModel) Login(email string, password string) *User {
	var user User
	if err := usermodel.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		logrus.Error("Model : Login data error,", err.Error())
		return nil
	}
	if user.ID < 1 {
		logrus.Error("Model : not found")
		return &user
	}

	return &user
}
