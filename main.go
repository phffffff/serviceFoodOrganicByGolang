package main

import (
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	uploadProvider "go_service_food_organic/component/upload_provider"
	"go_service_food_organic/middleware"
	foodTransport "go_service_food_organic/module/food/transport"
	profileTransport "go_service_food_organic/module/profile/transport"
	imageTransport "go_service_food_organic/module/upload/image/transport"
	userTransport "go_service_food_organic/module/user/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("err:", err)
	}

	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err)
	}

	db.Debug()

	SecretKey := os.Getenv("SYSTEM_SECRET")
	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")

	s3Provider := uploadProvider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appContext.NewAppContext(db, SecretKey, s3Provider)

	rt := gin.Default()
	rt.Use(middleware.Recover(appCtx))

	{
		admin := rt.Group(
			"/admin",
			middleware.RequiredAuth(appCtx),
			middleware.RoleRequired(appCtx, common.Admin),
		)

		{
			upload := admin.Group("image")
			upload.POST("upload", imageTransport.GinUploadImage(appCtx))
			upload.GET("list", imageTransport.GinListImage(appCtx))
			upload.DELETE("delete/:id", imageTransport.GinDeleteImage(appCtx))
		}

		{
			food := admin.Group("food")
			food.GET("/listfood", foodTransport.GinListFood(appCtx))
			food.POST("/createfood", foodTransport.GinCreateFood(appCtx))
		}

		{
			user := admin.Group("user")
			user.GET("/list", userTransport.GinListUser(appCtx))
			user.DELETE("/delete/:id", userTransport.GinDeleteUser(appCtx))
			user.PATCH("/update-pass/:id", userTransport.GinUpdateUser(appCtx))
		}

		{
			profile := admin.Group("profile")
			profile.GET("/list", profileTransport.GinListProfile(appCtx))
			profile.PUT("update/:id", profileTransport.GinUpdateProfile(appCtx))
		}
	}
	//user
	{
		user := rt.Group("user")
		user.POST("register", userTransport.GinRegister(appCtx))
		user.POST("authenticate", userTransport.GinLogin(appCtx))
		user.DELETE("delete/:id",
			middleware.RequiredAuth(appCtx),
			userTransport.GinDeleteUser(appCtx),
		)
		user.PATCH("update-pass/:id",
			middleware.RequiredAuth(appCtx),
			userTransport.GinUpdateUser(appCtx),
		)
	}

	{
		profile := rt.Group("profile", middleware.RequiredAuth(appCtx))
		profile.PUT("update/:id", profileTransport.GinUpdateProfile(appCtx))
	}

	rt.Run()
}
