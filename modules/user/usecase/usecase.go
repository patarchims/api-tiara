package usecase

import (
	"vincentcoreapi/modules/user"
	"vincentcoreapi/modules/user/entity"
)

type userUseCase struct {
	userRepository entity.UserRepository
}

func NewUserUseCase(ur entity.UserRepository) entity.UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu *userUseCase) GetByUserRepository(userName string) (user user.ApiUser, exist bool) {
	user, exist = uu.userRepository.GetByUserRepository(userName)
	return
}
