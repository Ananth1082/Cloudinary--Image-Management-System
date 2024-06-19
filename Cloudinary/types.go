package cloudinary_wrapper

type Document struct {
	docRef   string
	assetURL string
	image    interface{}
}

func NewDocument(docRef, assetURL string, image interface{}) *Document {
	return &Document{docRef: docRef, assetURL: assetURL, image: image}
}
