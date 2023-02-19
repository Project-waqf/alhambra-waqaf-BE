package helper

import (
	"bytes"
	"context"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"strconv"
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


func Upload(c echo.Context, file multipart.File, fileheader *multipart.FileHeader, tipe string) (string, string, error){
	var ur  imagekit.UploadRequest
	
	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGEIO_PUBLIC"),
		PrivateKey: os.Getenv("IMAGEIO_PRIVATE"),
	}
	
	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		panic(err)
	}

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	fileName := autoGenerate(10, fileheader.Filename) + timestamp
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", "", err
	}
	
	if tipe == "news"{
		ur = imagekit.UploadRequest{
			File:              buf.Bytes(), // []byte OR *url.URL OR url.URL OR base64 string
			FileName:          fileName,
			UseUniqueFileName: false,
			Tags:              []string{},
			Folder:            "/news",
			IsPrivateFile:     false,
			CustomCoordinates: "",
			ResponseFields:    nil,
		}
	} else if tipe == "wakaf" {
		ur = imagekit.UploadRequest{
			File:              buf.Bytes(), // []byte OR *url.URL OR url.URL OR base64 string
			FileName:          fileName,
			UseUniqueFileName: false,
			Tags:              []string{},
			Folder:            "/wakaf",
			IsPrivateFile:     false,
			CustomCoordinates: "",
			ResponseFields:    nil,
		}
	}
	
	ctx := context.Background()
	
	r, err := ik.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		return "", "", err
	}
	return r.FileID, r.URL, nil
}

// func Delete(filename string) (string, error) 