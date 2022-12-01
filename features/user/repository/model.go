package repository

import (
	_user "be13/clean-arch/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string `gorm:"type:varchar(15)"`
	Address  string
	Role     string
	// Books    []Book
}

func fromCore(dataCore _user.Core) User {
	userGorm := User{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Phone:    dataCore.Phone,
		Address:  dataCore.Address,
		Role:     dataCore.Role,
	}
	return userGorm
}
func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Phone:     dataModel.Phone,
		Address:   dataModel.Address,
		Role:      dataModel.Role,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}

}

func toCoreList(models []User) []_user.Core {
	var userCore []_user.Core
	for _, v := range models {
		userCore = append(userCore, v.toCore())

	}
	return userCore
}

func toModelList(core []_user.Core) []User {
	var model []User
	for _, v := range core {
		model = append(model, fromCore(v))

	}
	return model

}
