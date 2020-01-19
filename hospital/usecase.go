package hospital

type Usecase interface {
	Fetch([]models.hospital, error)
	GetById(id int64) (models.hospital, error)
	New(hospital models.hospital) error
}
