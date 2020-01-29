package location

import (
"github.com/ribhavsharma/my-go-app/models"
)

type Repository interface {
Fetch() ([]models.Location, error)
GetById(id int64) (models.Location, error)
New(hospital models.Location) error
}
