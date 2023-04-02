package main

import (
	"github.com/gin-gonic/gin"
	appContext "go_service_food_organic/component/app_context"
	"go_service_food_organic/middleware"
	foodTransport "go_service_food_organic/module/food/transport"
	userTransport "go_service_food_organic/module/user/transport"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := "cool_organic:@Klov3x124n@tcp(127.0.0.1:3307)/cool_organic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err)
	}

	db.Debug()

	SecretKey := os.Getenv("SYSTEM_SECRET")

	appCtx := appContext.NewAppContext(db, SecretKey)

	rt := gin.Default()
	rt.Use(middleware.Recover(appCtx))

	{
		food := rt.Group("food")
		food.GET("/listfood", foodTransport.GinListFood(appCtx))
		food.POST("/createfood", foodTransport.GinCreateFood(appCtx))
	}
	{
		user := rt.Group("user")
		user.POST("/register", userTransport.GinRegister(appCtx))
		user.POST("/authenticate", userTransport.GinLogin(appCtx))
	}

	rt.Run()
}
