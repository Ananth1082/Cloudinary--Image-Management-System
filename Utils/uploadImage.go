package utils

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(cld *cloudinary.Cloudinary, ctx context.Context, filePath string, imageName string) error {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := cld.Upload.Upload(ctx, filePath, uploader.UploadParams{
		PublicID:       imageName,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),})

	if err != nil {
		fmt.Println("error uploading image")
		return err

	}
	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
	return nil
}
