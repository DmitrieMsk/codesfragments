package user

type UserHandler struct {
}

func NewUserHadnler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetAll() *[]User {
	//должен просто принимать get запрос на /users
	//to do...
	return nil
}
