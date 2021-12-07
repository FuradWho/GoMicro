package service

import "errors"

type IUserService interface {
	GetName(userid int) string
	DelUser(uid int) error
}

type UserService struct{}

func (u UserService) GetName(userid int) string {
	if userid == 01 {
		return "jerry"
	}
	return "guest"
}

func (u UserService) DelUser(uid int) error {

	if uid == 01 {
		return nil
	}
	return errors.New("the user not exist")

}
