package config

import (
	"ImageManagement/m-v0/Utils"
	"context"

	"github.com/cloudinary/cloudinary-go/v2"
)

func Credentials() (*cloudinary.Cloudinary, context.Context) {
	// Add your Cloudinary credentials, set configuration parameter
	// Secure=true to return "https" URLs, and create a context
	//===================

	var cld, _ = cloudinary.New()
	cld.Config.Cloud.APIKey = utils.GetEnvVar("API_KEY")
	cld.Config.Cloud.APISecret = utils.GetEnvVar("API_SECRET")
	cld.Config.Cloud.CloudName = utils.GetEnvVar("CLOUD_NAME")
	cld.Config.URL.Secure = true
	ctx := context.Background()
	return cld, ctx
}
