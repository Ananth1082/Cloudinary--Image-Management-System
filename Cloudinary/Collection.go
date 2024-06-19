package cloudinary_wrapper

import (
	config "ImageManagement/m-v0/Config"
	"context"
	"errors"
	"path"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func CreateCollection(cld *cloudinary.Cloudinary, ctx context.Context, collectionName string) (*admin.CreateFolderResult, error) {
	res, err := cld.Admin.CreateFolder(ctx, admin.CreateFolderParams{Folder: config.Root + "/" + collectionName})
	if err != nil {
		return nil, errors.New("an error occured while creating the folder")
	}
	return res, nil
}

func DeleteCollection(cld *cloudinary.Cloudinary, ctx context.Context, collectionName string) error {
	_, err := cld.Admin.DeleteFolder(ctx, admin.DeleteFolderParams{Folder: path.Join(config.Root, collectionName)})
	return err
}

func (doc *Document) NewDoc(cld *cloudinary.Cloudinary, ctx context.Context) (*uploader.UploadResult, error) {
	res, err := cld.Upload.Upload(ctx, doc.image, uploader.UploadParams{
		PublicID:                       doc.docRef,
		UniqueFilename:                 api.Bool(true),
		Overwrite:                      api.Bool(false),
		Folder:                         path.Dir(doc.docRef),
		UseAssetFolderAsPublicIDPrefix: api.Bool(true),
	})
	doc.assetURL = res.URL //basically binds the doc object to the object in cloudinary
	return res, err
}

func GetDoc(cld *cloudinary.Cloudinary, ctx context.Context, docID string) (*admin.AssetResult, error) {
	return cld.Admin.Asset(ctx, admin.AssetParams{
		PublicID: docID,
	})
}
