package controllers

import (
	cloudinaryWrapper "ImageManagement/m-v0/Cloudinary"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/gin-gonic/gin"
)

func ReqGetImageHandler(ctx *gin.Context, cld *cloudinary.Cloudinary) {
	publicID := ctx.Query("publicID")
	toDownload := ctx.Query("toDownload")
	image, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: publicID})
	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Failed to get image URL: %v", err))
		return
	}

	if toDownload == "true" {
		ctx.Header("Content-Disposition", "attachment; filename="+publicID+"."+image.Format)
	}
	ctx.String(http.StatusOK, image.SecureURL)

}

func ReqUploadImageHandler(c *gin.Context, cld *cloudinary.Cloudinary, ctx context.Context) {
	// Parse the multipart form, limit the maximum memory to 10 MB
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.String(http.StatusBadRequest, "File is too large")
		return
	}

	// Get the file from the form input name 'file'
	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Failed to get file: %v", err))
		return
	}
	defer file.Close()

	// Create the uploads folder if it doesn't exist
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create uploads directory: %v", err))
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(filepath.Join("./uploads", handler.Filename))
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to create file: %v", err))
		return
	}

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to save file: %v", err))
		return
	}

	if code := cloudinaryWrapper.UploadImage(cld, ctx, filepath.Join("./uploads", handler.Filename), handler.Filename, "Home"); code != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to upload image: %v", err))
		return
	}
	dst.Close()
	//delete the file after uploading
	if err := os.Remove(filepath.Join("./uploads", handler.Filename)); err != nil {
		fmt.Println("error deleting file", err)
	}
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully.", handler.Filename))
}

// func getImageFromAFolder(requestCtx *gin.Context, cld *cloudinary.Cloudinary, cloudinaryCtx context.Context, folder string) {
// 	image,err := cld.Admin.AssetsByAssetFolder(cloudinaryCtx,admin.AssetsByAssetFolderParams{
// 		AssetFolder: ,
// 	})
// }

func UploadImageFromAFolder(cld *cloudinary.Cloudinary, cloudinaryCtx context.Context, localFolder string, cloudPath string) {
	//uploads all images in a folder of the local device to cloudinary
	//follows the uploads the same file structure

	//get all files in the folder
	children, err := os.ReadDir(localFolder)

	if err != nil {
		fmt.Println("error reading folder", err)
		return
	}

	if _, err := cld.Admin.CreateFolder(cloudinaryCtx, admin.CreateFolderParams{
		Folder: filepath.Join(cloudPath, localFolder),
	}); err != nil {
		println("Error creating cloudinary folder")
	}
	for _, file := range children {
		path := filepath.Join(".", localFolder, file.Name())
		err := cloudinaryWrapper.UploadImage(cld, cloudinaryCtx, path, file.Name(), cloudPath)
		if err != nil {
			fmt.Printf("Error uploading the file %s\n %v", file.Name(), err)
		} else {
			fmt.Printf("Successfully uploaded the file %s", file.Name())
		}
		os.Remove(path)
	}
}
