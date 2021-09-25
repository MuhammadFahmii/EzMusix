package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Medicines struct {
	ID          int       `gorm:"primaryKey" json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Detail      string    `json:"detail" form:"detail"`
	Dose        string    `json:"dose" form:"dose"`
	Side_Effect string    `json:"side_effect" form:"side_effect"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "find_medicines",
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["DB_Username"], config["DB_Password"], config["DB_Host"], config["DB_Port"], config["DB_Name"])
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&Medicines{})
}

func main() {
	InitDB()
	e := echo.New()
	g := e.Group("/medicines")
	g.POST("", AddMedicines)
	g.GET("", GetMedicines)
	g.PUT("/:id", UpdateMedicines)
	g.DELETE("/:id", DeleteMedicines)
	e.Start(":8000")
}

func AddMedicines(c echo.Context) error {
	medicines := Medicines{}
	if e := c.Bind(&medicines); e != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg": "failed",
		})
	}
	if err := DB.Create(&medicines).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"msg":  "Success",
		"data": medicines,
	})
}

func GetMedicines(c echo.Context) error {
	medicines := Medicines{}
	if err := DB.Find(&medicines).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	if medicines.ID == 0 {
		return c.JSON(http.StatusFound, echo.Map{
			"msg": "Empty",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg":  "Success",
		"data": medicines,
	})
}

func UpdateMedicines(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	medicines := Medicines{}
	if err := DB.Where("id=?", id).First(&medicines).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"msg": "ID not found",
		})
	}
	c.Bind(&medicines)
	if err := DB.Where("id=?", id).Updates(&medicines).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "Success",
	})
}

func DeleteMedicines(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	medicines := Medicines{}
	if err := DB.Where("id = ?", id).Delete(&medicines).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"msg": "failed",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "Success",
	})
}
