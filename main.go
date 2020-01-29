package main

import (

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	hospitalHandler "github.com/ribhavsharma/my-go-app/hospital/delivery"
	hospitalRepository "github.com/ribhavsharma/my-go-app/hospital/repository"
	hospitalUsecase "github.com/ribhavsharma/my-go-app/hospital/usecase"
	locationHandler "github.com/ribhavsharma/my-go-app/location/delivery"
	locationRepository "github.com/ribhavsharma/my-go-app/location/repository"
	locationUsecase "github.com/ribhavsharma/my-go-app/location/usecase"
	patientHandler "github.com/ribhavsharma/my-go-app/patient/delivery"
	patientRepository "github.com/ribhavsharma/my-go-app/patient/repository"
	patientUsecase "github.com/ribhavsharma/my-go-app/patient/usecase"
)


func main() {
	var db *sqlx.DB
	router := gin.Default()
	var err error
	db = sqlx.MustConnect("postgres", "postgres://tueisrkv:4caHkg9pKgWatez67bwcydlM1C_U3o4B@rosie.db.elephantsql.com:5432/tueisrkv")
	if err != nil {
		panic(err)
	}
	hospitalRepo := hospitalRepository.NewhospitalRepository(db)
	locationRepo := locationRepository.NewLocationRepository(db)
	patientRepo := patientRepository.NewPatientRepository(db)
	hospitalUsecase := hospitalUsecase.NewHospitalUsecase(hospitalRepo)
	locationUsecase := locationUsecase.NewLocationUsecase(locationRepo)
	patientUsecase := patientUsecase.NewPatientUsecase(patientRepo)
	hospitalGroup := router.Group("/hospital")
	locationGroup := router.Group("/location")
	patientGroup := router.Group("/patient")
	hospitalHandler.NewHospitalHandler(hospitalGroup, hospitalUsecase)
	locationHandler.NewLocationHandler(locationGroup, locationUsecase)
	patientHandler.NewPatientHandler(patientGroup, patientUsecase)
	err = router.Run()
	if err != nil{
		panic("Could not run the app")
	}
}
