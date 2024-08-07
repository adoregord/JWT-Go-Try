package usecase

import (
	"errors"
	"jwt-try/internal/domain"
	middleware "jwt-try/internal/middleware/jwt"
	"jwt-try/internal/repository"
	utils "jwt-try/internal/utils/hash"
)

type UserUsecaseInterface interface {
	RegisterUser(user domain.User) error
	CheckCredential(user domain.User) (string, error)
	VerifyJWT(token string) (*middleware.Claimsasdasda, error)
}

type UserUsecase struct {
	UserRepo repository.UserRepoInterface
}

func NewUserUsecase(userRepo repository.UserRepoInterface) UserUsecaseInterface {
	return UserUsecase{
		UserRepo: userRepo,
	}
}

func (uc UserUsecase) RegisterUser(User domain.User) error {
	// encrypt the password
	pass, err := utils.HashPassword(User.Password)
	if err != nil {
		return err
	}
	User.Password = pass

	return uc.UserRepo.RegisterUser(&User)
}

func (uc UserUsecase) CheckCredential(user domain.User) (string, error) {
	ok := uc.UserRepo.CheckCredential(&user)
	if !ok {
		return "", errors.New("invalid credentials")
	}

	token, err := middleware.GenerateJwt(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc UserUsecase) VerifyJWT(token string) (*middleware.Claimsasdasda, error){

	return middleware.VerifyJWT(token)

}
