package usecase

import (
	"users-api/model"
	"users-api/respository"
)

type UserUsecase struct {
	userRepository respository.UserRepository
}

func NewUserUsecase(respository respository.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository: respository,
	}
}

func (uu *UserUsecase) GetUsers() ([]model.User, error) {
	return uu.userRepository.GetUsers()
}

func (uu *UserUsecase) CreateUser(user model.User) (model.User, error) {
	userID, err := uu.userRepository.CreateUser(user)
	if err != nil {
		return model.User{}, err
	}

	user.ID = userID

	return user, err
}

func (uu *UserUsecase) GetUserByID(id int) (*model.User, error) {
	user, err := uu.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (uu *UserUsecase) UpdateUser(user model.User) (*model.User, error) {
	updatedUser, err := uu.userRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, err
}

func (uu *UserUsecase) DeleteUserByID(id int) (*model.User, error) {
	deletedUser, err := uu.userRepository.DeleteUserByID(id)
	if err != nil {
		return nil, err
	}

	return deletedUser, err
}
