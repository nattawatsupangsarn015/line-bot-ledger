package public

import (
	"errors"
	"example/line-bot-ledger/controller"
	"example/line-bot-ledger/model"
	"example/line-bot-ledger/request"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user request.Register) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newUser := &model.User{
		ID:        primitive.NewObjectID(),
		Email:     user.Email,
		Password:  string(hashedPassword),
		Name:      user.Name,
		Phone:     user.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = controller.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return "OK", nil
}

func LoginUser(user request.Login, lineId string) (interface{}, error) {
	findUser, err := controller.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if findUser == nil {
		return nil, nil
	}

	structureUser, ok := findUser.(model.User)
	if !ok {
		return nil, errors.New("Cannot convert structure")
	}

	plaintextPassword := []byte(user.Password)
	hashedPassword := []byte(structureUser.Password)

	err = bcrypt.CompareHashAndPassword(hashedPassword, plaintextPassword)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Forbidden")
	}

	err = controller.UpdateLineIdUser(user.Email, lineId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return findUser, nil
}

func LogoutUser(lineId string) error {
	findUser, err := controller.GetUserByLineId(lineId)
	if err != nil {
		return err
	}

	if findUser == nil {
		return nil
	}

	err = controller.UpdateLogoutUser(lineId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
