package hospital

import (
	"github.com/ribhavsharma/my-go-app/models"
)

type Usecase interface {
	Fetch() ([]models.Hospital, error)
	GetById(id int64) (models.Hospital, error)
	New(hospital models.Hospital) error
}
