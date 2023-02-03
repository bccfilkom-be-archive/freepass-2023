package main

//Name: Ridha Ilham Adi Setyawan
//NIM: 225150207111070

import (
	"facebook/admin"
	"facebook/handler"
	"facebook/user"

	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:cdaaptnia404@tcp(127.0.0.1:3306)/facebook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error")
	}

	db.AutoMigrate(&admin.Admin{})

	adminRepository := admin.NewRepository(db)
	adminService := admin.NewService(adminRepository)
	adminHandler := handler.NewAdminHandler(adminService)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	//CRUD

	router := gin.Default()

	v1 := router.Group("/v1")

	//==================I AM ADMIN=========================

	//admin menambah satu user A2
	v1.POST("/admins", adminHandler.CreateAdmin)

	//admin melihat dan menampilkan semua user A4
	v1.GET("/admins", adminHandler.GetAdmins)

	//admin mendapatkan salah satu user
	v1.GET("/admins/:id", adminHandler.GetAdmin)

	//admin mengupdate atau mengedit user A3
	v1.PUT("/admins/:id", adminHandler.UpdateAdmin)

	//admin menghapus user A1
	v1.DELETE("/admins/:id", adminHandler.DeleteAdmin)
	//=======================================================

	//==================I AM USER============================

	//user melakukan registrasi akun U1
	v1.POST("/users", userHandler.CreateUser)

	//user mengedit akunnya sendiri U2
	v1.PUT("/users/id", userHandler.UpdateUser)

	//user ingin menampilkan akun sendiri
	v1.GET("/users/id", userHandler.GetUser)

	router.Run(":8888")
}

// admin := admin.Admin{
// 	Name:     "Selena Gomez",
// 	Password: "lkdiwkaj8439",
// 	Position: "Secretary",
// 	Age:      18,
// }

// err = db.Create(&admin).Error
// if err != nil {
// 	fmt.Println("===========================")
// 	fmt.Println("Error creating admin record")
// 	fmt.Println("===========================")
// }
// var admin admin.Admin

// err = db.Where("id = ?", 3).First(&admin).Error
// if err != nil {
// 	fmt.Println("===========================")
// 	fmt.Println("Error creating admin record")
// 	fmt.Println("===========================")
// }

// err = db.Delete(&admin).Error
// if err != nil {
// 	fmt.Println("===========================")
// 	fmt.Println("Error deleting admin record")
// 	fmt.Println("===========================")
// }
// admin.Name = "David Goggins"

// err = db.Save(&admin).Error

// if err != nil {
// 	fmt.Println("===========================")
// 	fmt.Println("Error updating admin record")
// 	fmt.Println("===========================")
// }

// ID       int
// Name     string
// Password string
// Position string
// Age      int

//FLOW DATA
//main
//handler
//bookinput atau request
//service
//repository
//db
//mysql
