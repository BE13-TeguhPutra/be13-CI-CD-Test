package service

import (
	"be13/clean-arch/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Delete implements user.ServiceInterface
func (service *userService) Delete(id int) error {
	errDelete := service.userRepository.Delete(id)
	if errDelete != nil {
		return errDelete
	}
	return nil
}

// GetById implements user.ServiceInterface
func (service *userService) GetById(id int) (user.Core, error) {
	data, errData := service.userRepository.GetById(id)
	if errData != nil {
		return user.Core{}, errors.New("error get id")
	}
	return data, nil

}

// GetAlluser implements user.ServiceInterface
func (service *userService) GetAlluser() (data []user.Core, err error) {
	//memanggil fungsi yang ada di repositori
	data, err = service.userRepository.GetAlluser()
	return data, err
}

// InsertUser implements user.ServiceInterface
func (service *userService) InsertUser(input user.Core) error {
	input.Role = "user"
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	err := service.userRepository.InsertUser(input)
	if err != nil {
		return errors.New("Failed Insert Data")
	}
	return nil
}

// UpdateUser implements user.ServiceInterface
func (service *userService) UpdateUser(id int, input user.Core) error {
	err := service.userRepository.UpdateUser(id, input)
	if err != nil {
		return errors.New("Id not Found")
	}
	return nil
}
