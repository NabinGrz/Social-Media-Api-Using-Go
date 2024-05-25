package cloudinaryService

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func UploadCloudinary(file *multipart.FileHeader, fileNamePrefix string) (string, error) {
	// Remove from local
	defer func() {
		os.Remove("assets/uploads/" + file.Filename)
	}()

	cloudinary_url := "cloudinary://645717742131472:aDjTSlRXXOjY6lHv7cDlW7h8QoM@dkmgqyugj"
	cld, err := cloudinary.NewFromURL(cloudinary_url)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Upload the image on the cloud
	var ctx = context.Background()
	filename := fileNamePrefix + "-" + file.Filename + "-" + uuid.New().String()
	resp, err := cld.Upload.Upload(ctx, "assets/uploads/"+file.Filename,
		uploader.UploadParams{PublicID: filename,
			ResourceType: "auto"})

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Return the image url
	return resp.SecureURL, nil
}
