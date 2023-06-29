package helper

import (
	"context"
	config "myapp/configs"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	config.ValidateEnvFile()
	cld, err := cloudinary.NewFromParams(config.EnvCloudName(), config.EnvCloudAPIKey(), config.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.EnvCloudUploadFolder()})
	if err != nil {
		return "", err
	}

	return uploadParam.SecureURL, nil
}

func ImagesBulkUploadHelper(folderPath string) (string, error) {
	var response string
	// Read in all files in this folder.
	//	folder := "/Users/sam/test/"

	// For waiting on threads.
	//var wg sync.WaitGroup

	// Get files in stage.
	dirRead, _ := os.Open(folderPath)
	dirFiles, _ := dirRead.Readdir(0)
	for dirIndex := range dirFiles {
		//	fileHere := dirFiles[dirIndex]
		//		fileNameHere := fileHere.Name()
		response, _ = ImageUploadHelper(dirIndex)
		// Increment the WaitGroup counter.
		//wg.Add(1)

		// Thread.
		//go ReadFileSafely(fileNameHere, folder, &wg)
	}
	return response, nil
}
