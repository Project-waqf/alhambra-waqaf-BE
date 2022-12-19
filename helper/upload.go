package helper

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"os"

	"github.com/codedius/imagekit-go"
	"github.com/labstack/echo/v4"
)



func Upload(c echo.Context, file multipart.File, fileheader *multipart.FileHeader) error{
	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGEIO_PUBLIC"),
		PrivateKey: os.Getenv("IMAGEIO_PRIVATE"),
	}
	
	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}
	

	ur := imagekit.UploadRequest{
		File:              buf.Bytes(), // []byte OR *url.URL OR url.URL OR base64 string
		FileName:          fileheader.Filename,
		UseUniqueFileName: false,
		Tags:              []string{},
		Folder:            "/",
		IsPrivateFile:     false,
		CustomCoordinates: "",
		ResponseFields:    nil,
	}
	
	ctx := context.Background()
	
	_, err = ik.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		return err
	}
	return nil
}