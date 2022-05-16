package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422, "msg": "手机号码必须为11位"})
			return
		}
		if len(password) < 6 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422, "msg": "密码不能少于6位"})
			return
		}
		if len(name) == 0 {
			name = Randomstring(10)
		}
		log.Println(name, telephone, password)

		//查询手机号
		if isTelephoneExits(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422, "msg": "用户已存在"})
			return
		}
		newUser := User{
			Name:      name,
			Telephone: telephone,
			Password:  password,
		}
		db.Create(&newUser)

		c.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic((r.Run(":8000"))) // 监听并在 0.0.0.0:8080 上启动服务
}
func Randomstring(n int) string {
	var letters = []byte("qwertyuioplkjhgfdsamnbvcxzMNBVCXZLKJHGFDSAPOIUYTREWQ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)
	if (user.ID) != 0 {
		return true
	}
	return false
}

func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "localhost"
	port := "23306"
	database := "ginvues"
	username := "root"
	password := "sup157"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database ,err: " + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}