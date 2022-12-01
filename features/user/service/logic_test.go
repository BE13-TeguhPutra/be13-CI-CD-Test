package service

import (
	"be13/clean-arch/features/user"
	"be13/clean-arch/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//syarat unit test
// membuat file akhiran _test.go
//membuat fucn diawalai nama test
//

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepo)
	returnData := []user.Core{{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}}
	t.Run("Succes Get All", func(t *testing.T) {
		repo.On("GetAlluser").Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		response, err := srv.GetAlluser()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t) //memastikan kalau semua sudah sesuai, assert nya sudah dicek dan terpenui.
	})

	t.Run("Failed Get All", func(t *testing.T) {
		repo.On("GetAlluser").Return(nil, errors.New("failed to get data")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		response, err := srv.GetAlluser()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t) //memastikan kalau semua sudah sesuai, assert nya sudah dicek dan terpenui.
	})

}

func TestCreate(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Succes Create User", func(t *testing.T) {
		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		repo.On("InsertUser", inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.InsertUser(inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Create, Duplicate", func(t *testing.T) {
		inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		repo.On("InsertUser", inputRepo).Return(errors.New("Failed Insert")).Once()
		srv := New(repo)
		err := srv.InsertUser(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "Failed Insert Data", err.Error()) // samakan dengan di logic
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create, Empty Name", func(t *testing.T) {
		// inputRepo := user.Core{Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"} tidak diperulkan kareana tidak dijalankan kareana sudah di return "Lihat Logic"
		inputData := user.Core{Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta"}
		// repo.On("InsertUser", inputRepo).Return(errors.New("Failed Insert,error querry")).Once()
		srv := New(repo)
		err := srv.InsertUser(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetbyId(t *testing.T) {
	repo := new(mocks.UserRepo)
	returnData := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
	t.Run("Succes Get user by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(returnData, nil).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		response, err := srv.GetById(1)
		assert.Nil(t, err)
		assert.Equal(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed to Get user by id", func(t *testing.T) {
		repo.On("GetById", 1).Return(user.Core{}, errors.New("Id not Found")).Once() //menentukan fungsi yang akan dijalankan//

		srv := New(repo) //membuat dependency injection
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, returnData.Name, response.Name)
		repo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success Update", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		repo.On("UpdateUser", int(inputRepo.ID), inputRepo).Return(nil).Once()
		srv := New(repo)
		err := srv.UpdateUser(int(inputData.ID), inputData)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		repo.On("UpdateUser", int(inputRepo.ID), inputRepo).Return(errors.New("Id Error")).Once()
		srv := New(repo)
		err := srv.UpdateUser(int(inputData.ID), inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "Id not Found", err.Error())

		repo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success Delete", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		repo.On("Delete", int(inputRepo.ID)).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(int(inputData.ID))
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Update data", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		inputData := user.Core{ID: 1, Name: "alta", Email: "alta@mail.id", Password: "qwerty", Phone: "081234", Address: "Jakarta", Role: "user"}
		repo.On("Delete", int(inputRepo.ID)).Return(errors.New("Id not Found")).Once()
		srv := New(repo)
		err := srv.Delete(int(inputData.ID))
		assert.NotNil(t, err)

		repo.AssertExpectations(t)
	})

}
