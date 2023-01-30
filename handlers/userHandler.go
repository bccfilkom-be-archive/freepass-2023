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

func UserSignup(c *gin.Context) {
	var body struct{
		Email string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Nim	string `json:"nim" binding:"required"`
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
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if _, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to hash password",
			"error":   err.Error(),
		})
		return
	}
	user := models.Users{
		Email	 : body.Email,
		Username : body.Username,
		Nim		 : body.Nim,
		Password : string(hash),
	}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create user",
			"error":   result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user registered successfully",
		"error":nil,
	})
}

func UsersLogin(c *gin.Context){
	var body struct{
		Nim string `json:"nim" binding:"required"`
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
	var user models.Users
	if err:=initializers.DB.First(&user, "nim=?",body.Nim).Error;err!=nil{
		message:="Failed to querying user data"
		if user.ID==0{
			message="Invalid NIM or password"
		}
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": message,
			"error": err,
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid NIM or password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"role":"user",
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
			"username": user.Username,
			"token": tokenString,
		},
		"error":nil,
	})
}

func GetUserProfile(c *gin.Context){
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	var user []models.Users
	if err := initializers.DB.Select("id","nim","username", "email", "sks").First(&user, "id = ?", userId).Error; err != nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user data",
			"error":err.Error(),
		}) 
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "success querying user data",
		"data":user,
	})
}

func UpdateProfile(c *gin.Context){
	var body models.Users
	if err:= c.ShouldBindJSON(&body);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	var user models.Users
	if err := initializers.DB.First(&user,userId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user data",
			"error":err.Error(),
		})
		return
	}
	hash, hashError := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if hashError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to hash password",
			"error":   hashError.Error(),
		})
		return
	}
	user.Nim=body.Nim
	user.Email=body.Email 
	user.Username=body.Username
	user.Password=string(hash)
	if err:=initializers.DB.Save(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Error when updating the user data",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "update user data successfully",
		"error":nil,
	})
}
func GetUserIdClass(userId any)([]int,error){
	var userClass []models.UserClass
	err:=initializers.DB.Where("user_id = ?", userId).Find(&userClass).Error
	var idClass []int
	for _,c:=range userClass{
		idClass=append(idClass,int(c.ClassID))
	}
	return idClass,err
}
func GetAllClasses(c *gin.Context){
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	idClass,err := GetUserIdClass(userId)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	var classes []models.Class
	if err:=initializers.DB.Not(idClass).Joins("Course").Find(&classes).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Query data successfully",
		"error":nil,
		"data":classes,
	})
}

func SearchClass(c *gin.Context){
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	courseCode:= c.Query("course_code")
	idClass,err := GetUserIdClass(userId)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	var classes []models.Class
	if err:=initializers.DB.Not(idClass).Joins("Course").Where("course_code = ?", courseCode).Find(&classes).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Query data successfully",
		"error":nil,
		"data":classes,
	})
}

func AddClass(c *gin.Context){
	var body struct{
		ClassID int `json:"id" binding:"required"`
	}
	if err:= c.ShouldBindJSON(&body);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to read body",
			"error":   err.Error(),
		})
		return
	}
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	var userClass []models.UserClass
	if err:=initializers.DB.Joins("Class").Find(&userClass, "user_id = ?", userId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	var class models.Class
	if err:=initializers.DB.Joins("Course").Where("classes.id = ?", body.ClassID).First(&class).Error;err!=nil{
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
				"message": "Gagal menambahkan kelas, anda telah terdaftar di kelas "+u.Class.Name+" mata kuliah "+class.Course.Title,
				"error": "Failed to add class, can't add classes with the same subject",
			})
			return
		}
	}
	var user models.Users
	if err:= initializers.DB.First(&user,userId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user data",
			"error":err.Error(),
		})
		return
	}
	addClass := models.UserClass{
		UserID: user.ID,
		ClassID: uint(body.ClassID),
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
			"message": "Error when updating user data",
			"error":err.Error(),
		})
		return
	}
	if err:=initializers.DB.Create(&addClass).Error;err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to add class",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully added class",
		"error":nil,
	})
}

func GetUserClasses(c *gin.Context){
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	idClass,error:=GetUserIdClass(userId)
	if error!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":error.Error(),
		})
		return
	}
	var class []models.Class
	if err:=initializers.DB.Joins("Course").Where(map[string]interface{}{"classes.id": idClass}).Find(&class).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "Query user's class data successfully",
		"error":nil,
		"data":class,
	})
}

func DeleteUserClass(c *gin.Context){
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	idString:= c.Param("id")
	classId,_:= strconv.Atoi(idString)
	var userClass models.UserClass
	if err:=initializers.DB.Where("user_id = ? AND class_id = ?", userId, classId).First(&userClass).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user's class data",
			"error":err.Error(),
		})
		return
	}
	if err:=initializers.DB.Where("user_id = ? AND class_id = ?", userId, classId).Delete(&userClass).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Failed to delete user's class data",
			"error":err.Error(),
		})
		return
	}
	var user models.Users
	if err:= initializers.DB.First(&user,userId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying user data",
			"error":err.Error(),
		})
		return
	}
	var class models.Class
	if err:= initializers.DB.Joins("Course").First(&class,classId).Error;err!=nil{
		c.JSON(http.StatusNotFound,gin.H{
			"success": false,
			"message": "Failed to querying class data",
			"error":err.Error(),
		})
		return
	}
	user.Sks=user.Sks-class.Course.Sks
	if err:=initializers.DB.Save(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"success": false,
			"message": "Error when updating user data",
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success": true,
		"message": "Delete user's class data successfully",
		"error":nil,
	})
}

func GetUserClassParticipants(c *gin.Context){
	classId:= c.Query("id")
	userId,idError:=c.Get("user")
	if !idError {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	var userClass []models.UserClass
	if err:=initializers.DB.Select("id","nim","username","email").Joins("User").Find(&userClass,"user_id = ? AND class_id = ?",userId, classId).Error;err!=nil{
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