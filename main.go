package main

import (
	cloudinary_wrapper "ImageManagement/m-v0/Cloudinary"
	config "ImageManagement/m-v0/Config"
	"log"
)

func main() {
	// routes.Routes()
	cld, ctx := config.Credentials()
	// res, err := cloudinary_wrapper.CreateCollection(cld, ctx, "Vendor")
	// if err != nil {
	// 	log.Fatal("Error occured")
	// }
	// log.Println("res: ", res)
	// img, err := os.Open("./Upload/image.png")
	// if err != nil {
	// 	log.Fatal("Error opening file")
	// }
	// doc := cloudinary_wrapper.NewDocument("Fisheries App/Vendor", "Vendor", "", img)
	// res, err := doc.NewDoc(cld, ctx)
	// log.Println(res, err)
	res, err := cloudinary_wrapper.GetDoc(cld, ctx, "Vendor")
	log.Println(res, err)
}
