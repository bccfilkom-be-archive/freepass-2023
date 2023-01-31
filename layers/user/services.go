package layers

import "freepass-2023/models"

type Service interface {
	Create() (models.User, error)
	Read() (models.User, error)
	Update() (models.User, error)
	Delete() (models.User, error)
}
