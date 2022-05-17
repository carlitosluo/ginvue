package controller

import (
	"ginvue/common"
	"ginvue/model"
	"ginvue/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
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
		name = util.Randomstring(10)
	}
	log.Println(name, telephone, password)

	//查询手机号
	if isTelephoneExits(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422, "msg": "用户已存在"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if (user.ID) != 0 {
		return true
	}
	return false
}
