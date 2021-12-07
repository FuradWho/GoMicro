package service

type IUserService interface {
	GetName(userid int) string
}

type UserService struct{}

func (u UserService) GetName(userid int) string {
	if userid == 01 {
		return "jerry"
	}
	return "guest"
}
