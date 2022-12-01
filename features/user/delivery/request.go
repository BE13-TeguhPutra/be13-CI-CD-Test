package delivery

import "be13/clean-arch/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

func (req *UserRequest) reqToCore() user.Core {
	return user.Core{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Address:  req.Address,
		Role:     req.Role,
	}

}
