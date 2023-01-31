package user

type Service interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(userRequest UserRequest) (User, error)
	Update(ID int, userRequest UserRequest) (User, error)
	Delete(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {

	users, err := s.repository.FindAll()

	return users, err
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	return user, err
}

func (s *service) Create(userRequest UserRequest) (User, error) {

	user := User{
		Name:     userRequest.Name,
		Company:  userRequest.Company,
		Password: userRequest.Password,
		Position: userRequest.Position,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *service) Update(ID int, userRequest UserRequest) (User, error) {
	user, err := s.repository.FindByID(ID)

	user.Name = userRequest.Name
	user.Company = userRequest.Company
	user.Password = userRequest.Password
	user.Position = userRequest.Position

	newUser, err := s.repository.Update(user)
	return newUser, err
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	newUser, err := s.repository.Delete(user)
	return newUser, err
}
