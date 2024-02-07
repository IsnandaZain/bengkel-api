package controllers

import (
	"bengkel-api/database"
	"bengkel-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCar godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int tru "ID for the car"
// @Success 200 {object} models.Car
// @Router /cars/{carID} [get]
func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	db := database.GetDB()

	car := models.Car{}
	err := db.First(&car, "id=?", carID).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

// GetAllCar godoc
// @Summary Get list
// @Description Get list of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCar(ctx *gin.Context) {
	db := database.GetDB()

	car := models.Car{}
	err := db.Find(&car).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Failed get data",
			"error_message": "Something wrong when try to get data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

// CreateCar godoc
// @Summary Post details for a given Id
// @Description Post details of car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	db := database.GetDB()
	newCar := models.Car{}

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(&newCar)
	err := db.Create(&newCar).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Error insert data",
			"error_message": "Something wrong when try to insert data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	db := database.GetDB()

	carUpdate := models.Car{}

	if err := ctx.ShouldBindJSON(&carUpdate); err != nil { // get body request
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Model(&)
	ctx.JSON(http.StatusOK, gin.H{
		"car": "success",
	})

}
