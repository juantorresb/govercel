package handler

import (
	"fmt"
	"net/http"

	"myapp/cloudinary/controllers"

	vercel "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := vercel.New()

	server.Use(vercel.Recovery(func(err interface{}, c *vercel.Context) {
		if httpError, ok := err.(vercel.HttpError); ok {
			c.JSON(httpError.Status, vercel.H{
				"message": httpError.Error(),
			})
		} else {
			message := fmt.Sprintf("%s", err)
			c.JSON(500, vercel.H{
				"message": message,
			})
		}
	}))

	server.GET("/", controllers.Status)
	server.POST("/upload", controllers.FileUpload)

	server.Handle(w, r)
}

func helloMethod(context *vercel.Context) {
	name := context.Query("name")
	if name == "" {
		context.JSON(400, vercel.H{
			"message": "name not found",
		})
	} else {
		context.JSON(200, vercel.H{
			"data": fmt.Sprintf("Hello %s!", name),
		})
	}
}

/*

	server.GET("/hello", helloMethod)

	server.GET("/user/:id", func(context *vercel.Context) {
		context.JSON(400, vercel.H{
			"data": vercel.H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *vercel.Context) {
		context.JSON(200, vercel.H{
			"data": vercel.H{
				"url": context.Path,
			},
		})
	})

	server.POST("/upload", func(context *vercel.Context) {
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			context.JSON(400, vercel.H{
				"message": "param file not found",
			})
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		//		fmt.Fprintf(w, "Successfully Uploaded File\n")
		context.JSON(400, vercel.H{
			"message": "Successfully Uploaded File",
		})
	})
*/
