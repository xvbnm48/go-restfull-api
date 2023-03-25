package user

type Service interface {
	CreateUser(input CreateUserInput) (User, error)
}
