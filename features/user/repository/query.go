package repository

import (
	"be13/clean-arch/features/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// UpdateUser implements user.RepositoryInterface

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// GetAlluser implements user.RepositoryInterface
func (repo *userRepository) GetAlluser() ([]user.Core, error) {
	var user []User
	tx := repo.db.Find(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	res := toCoreList(user) // model to core
	return res, nil

}

// InsertUser implements user.RepositoryInterface
func (repo *userRepository) InsertUser(input user.Core) error {
	userGorm := fromCore(input)     //core to models
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	return nil

}

// GetById implements user.RepositoryInterface
func (repo *userRepository) GetById(id int) (user.Core, error) {
	users := User{}
	tx := repo.db.First(&users, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}
	result := users.toCore()
	return result, nil
}

// Delete implements user.RepositoryInterface
func (repo *userRepository) Delete(id int) error {
	users := User{}
	tx := repo.db.Delete(&users, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")

	}
	return nil

}

// UpdateUser implements user.RepositoryInterface
func (repo *userRepository) UpdateUser(id int, input user.Core) error {
	user := fromCore(input)
	tx := repo.db.Model(user).Where("id = ?", id).Updates(user)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
