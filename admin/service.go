package admin

type Service interface {
	FindAll() ([]Admin, error)
	FindByID(ID int) (Admin, error)
	Create(adminRequest AdminRequest) (Admin, error)
	Update(ID int, adminRequest AdminRequest) (Admin, error)
	Delete(ID int) (Admin, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Admin, error) {

	admins, err := s.repository.FindAll()

	return admins, err
}

func (s *service) FindByID(ID int) (Admin, error) {
	admin, err := s.repository.FindByID(ID)

	return admin, err
}

func (s *service) Create(adminRequest AdminRequest) (Admin, error) {

	admin := Admin{
		Name:     adminRequest.Name,
		Company:  adminRequest.Company,
		Password: adminRequest.Password,
		Position: adminRequest.Position,
	}

	newAdmin, err := s.repository.Create(admin)
	return newAdmin, err
}

func (s *service) Update(ID int, adminRequest AdminRequest) (Admin, error) {
	admin, err := s.repository.FindByID(ID)

	admin.Name = adminRequest.Name
	admin.Company = adminRequest.Company
	admin.Password = adminRequest.Password
	admin.Position = adminRequest.Position

	newAdmin, err := s.repository.Update(admin)
	return newAdmin, err
}

func (s *service) Delete(ID int) (Admin, error) {
	admin, err := s.repository.FindByID(ID)

	newAdmin, err := s.repository.Delete(admin)
	return newAdmin, err
}
