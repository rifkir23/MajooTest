package service

import (
	"github.com/rifkir23/MjTest/entity"
	"github.com/rifkir23/MjTest/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var (
	repoUser = repository.NewUserRepo()
)

type AuthService interface {
	VerifyCredential(userName string, password string) interface{}
}

func NewAuthService() AuthService {
	return &authService{}
}

type authService struct {
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := repoUser.VerifyCredential(username)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.UserName == username && comparedPassword {
			return res
		}
		return false
	}
	return false
}

//func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
//	userToCreate := entity.User{}
//	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
//	if err != nil {
//		log.Fatalf("Failed map %v", err)
//	}
//	res := service.userRepository.InsertUser(userToCreate)
//	return res
//}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
