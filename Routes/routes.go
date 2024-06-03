package routespkg

import (
	config "ImageManagement/m-v0/Config"
	controllers "ImageManagement/m-v0/Controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	cld, ctx := config.Credentials()
	router := gin.Default()

	// Define the upload endpoint
	router.POST("/upload", func(c *gin.Context) {
		controllers.UploadImageHandler(c, cld, ctx)
	})
	
	router.GET("/getImageByPublicID", func(ctx *gin.Context) {
		controllers.GetImageHandler(ctx, cld)
	})
	router.Run(":8080") // Start the server on port 8080
}