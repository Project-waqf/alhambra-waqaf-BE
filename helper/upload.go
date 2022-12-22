package helper

import (
	"bytes"
	"context"
	"math/rand"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/codedius/imagekit-go"
	"github.com/labstack/echo/v4"
)

// CREATE RANDOM STRING

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func autoGenerate(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return autoGenerate(length, charset)
}


func Upload(c echo.Context, file multipart.File, fileheader *multipart.FileHeader) (string, error){
	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGEIO_PUBLIC"),
		PrivateKey: os.Getenv("IMAGEIO_PRIVATE"),
	}
	
	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		panic(err)
	}

	fileName := autoGenerate(10, fileheader.Filename)

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}
	

	ur := imagekit.UploadRequest{
		File:              buf.Bytes(), // []byte OR *url.URL OR url.URL OR base64 string
		FileName:          fileName,
		UseUniqueFileName: false,
		Tags:              []string{},
		Folder:            "/",
		IsPrivateFile:     false,
		CustomCoordinates: "",
		ResponseFields:    nil,
	}
	
	ctx := context.Background()
	
	r, err := ik.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		return "", err
	}
	return r.URL, nil
}