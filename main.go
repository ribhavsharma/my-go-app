package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var router *gin.Engine

type Hospital struct {
	HospitalId       int    `db:"id" json:"id"`
	HospitalName     string `db:"name" json:"name"`
	MaxPatientAmount int    `db:"max_patient_amount" json:"max_patient_amount"`
}

type Location struct {
	LocationId   int    `db:"id" json:"id"`
	LocationName string `db:"name" json:"name"`
	HospitalId   int    `db:"hospital_id" json:"hospital_id"`
}

type Patient struct {
	PatientId   int    `db:"id" json:"id"`
	PatientName string `db:"name" json:"name"`
	Illness     string `db:"illness" json:"illness"`
	BirthDate   string `db:"birth_date" json:"birth_date"`
	LocationId  int    `db:"location_id" json:"location_id"`
}

/*
connect with the database instance
*/
func initDatabase() {
	var err error
	db, err = sqlx.Connect("postgres", "postgres://tueisrkv:4caHkg9pKgWatez67bwcydlM1C_U3o4B@rosie.db.elephantsql.com:5432/tueisrkv")
	if err != nil {
		panic(err)
	}
}

/*
fetch location data from the database
*/
func GetAllLocations(c *gin.Context) {
	var location []Location
	error := db.Select(&location, "SELECT * from location")
	if error != nil {
		panic(error)
	}
	c.JSON(http.StatusOK, location)
}

/*
fetch hospital data from the database
*/
func GetAllHospitals(c *gin.Context) {
	var hospital []Hospital
	error := db.Select(&hospital, "SELECT * from hospital")
	if error != nil {
		panic(error)
	}
	c.JSON(http.StatusOK, hospital)
}

/*
fetch patient data from the database
*/
func GetAllPatients(c *gin.Context) {
	var patient []Patient
	error := db.Select(&patient, "SELECT * from patient")
	if error != nil {
		panic(error)
	}
	c.JSON(http.StatusOK, patient)
}

/*
initialize all routes
*/
func initAPI() {
	router = gin.Default()
	router.GET("/patient/all", GetAllPatients)
	router.GET("/location/all", GetAllLocations)
	router.GET("/hospital/all", GetAllHospitals)
}

func main() {
	initDatabase()
	initAPI()
	err := router.Run(":8080")
	if err != nil {
		panic("Could not run the app")
	}
}
