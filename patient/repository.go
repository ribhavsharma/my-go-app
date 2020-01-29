package patient

import (
"github.com/ribhavsharma/my-go-app/models"
)

type Repository interface {
Fetch() ([]models.Patient, error)
GetById(id int64) (models.Patient, error)
New(patient models.Patient) error
}
