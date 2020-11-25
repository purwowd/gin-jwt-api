package user

type RegisterUserInput struct {
	Name     string
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type CheckEmailInput struct {
	Email string
}