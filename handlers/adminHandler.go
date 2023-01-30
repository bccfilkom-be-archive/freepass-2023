package handlers

import (
	"course-management/initializers"
	"course-management/models"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *gin.Context){
	var body struct{
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	var admin models.Admin
	if err:=initializers.DB.First(&admin, "username=?",body.Username).Error;err!=nil{
		message:="Failed to querying admin data"
		if admin.ID==0{
			message="Invalid email or password"
		}
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": message,
			"error":err.Error(),
		}) 
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password),[]byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid email or password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": admin.ID,
		"role":"admin",
		"exp":time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETTOKEN")))
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600*24*30,"","",false,true)
	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "Welcome, here's your token.",
		"data": gin.H{
			"username": admin.Username,
			"token": tokenString,
		},
		"error":nil,
	})
}

func AdminGetAllClasses(c *gin.Context){
	var class []models.Class
	if err:=initializers.DB.Joins("Course").Find(&class).Error;err!=nil{
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Querying class data successfully",
		"error":nil,
		"data":class,
	})
}

func CreateCourse(c *gin.Context){
	var body struct{
		Title string `json:"title" binding:"required"`
		Course_code string `json:"course_code" binding:"required"`
		Sks	int `json:"sks" binding:"required,number"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	course := models.Courses{
		Title: body.Title,
		Course_code: body.Course_code,
		Sks: body.Sks,
	}
	if err:=initializers.DB.Create(&course).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create course",
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Create course successfully",
		"error":nil,
	})
}

func CreateClass(c *gin.Context){
	var body struct{
		Name string `json:"class_name" binding:"required"`
		Course_code string `json:"course_code" binding:"required"`
		Location string `json:"location" binding:"required"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	var course models.Courses
	if err:=initializers.DB.Where("course_code = ?", body.Course_code).First(&course).Error;err != nil {
		message:="Failed to querying course data"
		if course.ID==0{
			message="Course code is not found"
		}
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": message,
			"error": err,
		})
		return
	}
	class := models.Class{
		Name: body.Name,
		Class_code: body.Course_code+"-"+body.Name,
		Location: body.Location,
		CourseID: course.ID,
	}
	if err:=initializers.DB.Create(&class).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create class",
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Create class successfully",
		"error":nil,
	})
}

func UpdateClass(c *gin.Context){
	var body struct{
		Name string `json:"class_name" binding:"required"`
		Course_code string `json:"course_code" binding:"required"`
		Location string `json:"location" binding:"required"`
	}
	if err:= c.ShouldBindJSON(&body);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	var course models.Courses
	if err:=initializers.DB.Where("course_code = ?", body.Course_code).First(&course).Error;err != nil {
		message:="Failed to querying course data"
		if course.ID==0{
			message="Course code is not found"
		}
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": message,
			"error": err,
		})
		return
	}
	
	classId:=c.Query("id")
	var classOld models.Class
	if err := initializers.DB.Joins("Course").First(&classOld,classId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}

	if body.Course_code!=classOld.Course.Course_code{
		var userClass []models.UserClass
		if err := initializers.DB.Find(&userClass,"class_id = ?", classId).Error;err!=nil{
			c.JSON(http.StatusNotFound,gin.H{
				"success": false,
				"message": "Failed to querying user's class data",
				"error":err.Error(),
			})
			return
		}
		for _,u:=range userClass{
			var user models.Users
			if err := initializers.DB.First(&user,u.UserID).Error;err!=nil{
				c.JSON(http.StatusNotFound,gin.H{
					"success": false,
					"message": "Failed to querying user data",
					"error":err.Error(),
				})
				return
			}
			user.Sks=user.Sks-classOld.Course.Sks
			if err:=initializers.DB.Save(&user).Error;err!=nil{
				c.JSON(http.StatusInternalServerError,gin.H{
					"success": false,
					"message": "Error when updating user data",
					"error":err.Error(),
				})
				return
			}
		}
		if err:=initializers.DB.Where("class_id = ?", classId).Delete(&userClass).Error;err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"success": false,
				"message": "Failed to delete user's class data",
				"error":err.Error(),
			})
			return
		}
	}
	var class models.Class
	if err := initializers.DB.First(&class,classId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	class.Name=body.Name
	class.CourseID=course.ID
	class.Location=body.Location
	class.Class_code=course.Course_code+"-"+body.Name
	if err:=initializers.DB.Save(&class).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Error when updating the database.",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "update class data successfully",
		"error":nil,
	})
}

func DeleteClass(c *gin.Context){
	var body struct{
		ID int `json:"id" binding:"required"`
	}
	if err:= c.ShouldBindJSON(&body);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	var classOld models.Class
	if err := initializers.DB.Joins("Course").First(&classOld,body.ID).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	var userClass []models.UserClass
	if err := initializers.DB.Find(&userClass,"class_id = ?", body.ID).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	for _,u:=range userClass{
		var user models.Users
		if err := initializers.DB.First(&user,u.UserID).Error;err!=nil{
			c.JSON(http.StatusNotFound,gin.H{
				"success": false,
				"message": "Failed to querying user data",
				"error":err.Error(),
			})
			return
		}
		user.Sks=user.Sks-classOld.Course.Sks
		if err:=initializers.DB.Save(&user).Error;err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"success": false,
				"message": "Error when updating user data",
				"error":err.Error(),
			})
			return
		}
	}
	var class models.Class
	if err:=initializers.DB.Where("id = ? ", body.ID).First(&class).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Can't found class data",
			"error":err.Error(),
		})
		return
	}
	if err:=initializers.DB.Where("id = ? ", body.ID).Delete(&class).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Failed to delete data",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "Delete class data successfully",
		"error":nil,
	})
}

func AddUserToClass(c *gin.Context){
	classIdString:=c.Query("id")
	classId,error:=strconv.Atoi(classIdString)
	var body struct{
		Nim string `json:"nim" binding:"required"`
	}
	if err:= c.ShouldBindJSON(&body);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	var user models.Users
	if err:=initializers.DB.Where("nim = ? ", body.Nim).First(&user).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user data",
			"error":err.Error(),
		})
		return
	}

	var userClass []models.UserClass
	if err:=initializers.DB.Joins("Class").Find(&userClass, "user_id = ?", user.ID).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	var class models.Class
	if err:=initializers.DB.Joins("Course").Where("classes.id = ?", classId).First(&class).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	for _,u:=range userClass{
		if u.Class.CourseID==class.CourseID{
			c.JSON(http.StatusBadRequest,gin.H{
				"success": false,
				"message": "Gagal menambahkan kelas, user telah terdaftar di kelas "+u.Class.Name+" mata kuliah "+class.Course.Title,
				"error": "Failed to add class, can't add classes with the same subject",
			})
			return
		}
	}
	addClass := models.UserClass{
		UserID: user.ID,
		ClassID: uint(classId),
	}
	if user.Sks+class.Course.Sks>24{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Gagal menambahkan kelas, jumlah maksimal sks telah terpenuhi",
			"error": "Failed to add class",
		})
		return
	}
	user.Sks=user.Sks+class.Course.Sks
	if err:=initializers.DB.Save(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Error when updating user data.",
			"error":err.Error(),
		})
		return
	}
	if err:=initializers.DB.Create(&addClass).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add user to class",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully added user to class",
		"error":error,
	})
}

func DeleteUserFromClass(c *gin.Context){
	classId:=c.Query("id")
	var body struct{
		ID int `json:"id" binding:"required,number"`
	}
	if err:= c.ShouldBindJSON(&body);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	var class models.Class
	if err:=initializers.DB.Joins("Course").Where("classes.id = ?", classId).First(&class,).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Error when querying class data.",
			"error":err.Error(),
		})
		return
	}
	
	var userClass models.UserClass
	if err:=initializers.DB.Where("user_id = ? AND class_id = ?", body.ID, classId).First(&userClass).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	if err:=initializers.DB.Where("user_id = ? AND class_id = ?", body.ID, classId).Delete(&userClass).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Failed to delete user from class",
			"error":err.Error(),
		})
		return
	}

	var user models.Users
	if err:=initializers.DB.First(&user,"id = ?", body.ID).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Error when querying user data.",
			"error":err.Error(),
		})
		return
	}
	user.Sks=user.Sks-class.Course.Sks
	if err:=initializers.DB.Save(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Error when updating user data.",
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "Successfully deleted user from class",
		"error":nil,
	})
}

func GetClassParticipants(c *gin.Context){
	classId:= c.Query("id")
	var userClass []models.UserClass
	if err:=initializers.DB.Select("id","nim","username","email").Joins("User").Find(&userClass,"class_id = ?", classId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	var userResponses []models.UsersResponse
	for _,c:=range userClass{
		userResponse:=func(u models.UserClass)models.UsersResponse{
			return models.UsersResponse{
				ID: int(u.User.ID),
				Nim: u.User.Nim,
				Username :u.User.Username,
				Email: u.User.Email,
			}
		}
		userResponses=append(userResponses, userResponse(c))
	}
	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "Query class's participants data successfully",
		"error":nil,
		"data":userResponses,
	})
}