package user

type UserService struct{}

func (us *UserService) GetUser(userId string) (User, error) {
    return User{ID: userId, Name: "John Doe"}, nil
}
