package delivery

import "be13/clean-arch/features/user"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Role    string `json:"role"`
}

func coreToResponse(core user.Core) UserResponse {
	response := UserResponse{
		ID:      core.ID,
		Name:    core.Name,
		Email:   core.Email,
		Phone:   core.Phone,
		Address: core.Address,
		Role:    core.Role,
	}
	return response

}



func responseList(listRes []user.Core) []UserResponse {
	var resList []UserResponse
	for _, v := range listRes {
		resList = append(resList, coreToResponse(v))

	}
	return resList

}
