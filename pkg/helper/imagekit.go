package helper

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func fileCheck(fileheader *multipart.FileHeader, tipe string) bool {

	// file size max 5MB
	size := fileheader.Size
	if size > 5*1024*1024 {
		return false
	}

	fileExt := filepath.Ext(fileheader.Filename)
	fileExtension := strings.ToLower(fileExt)

	if tipe == "image" {
		if fileExtension == ".jpeg" || fileExtension == ".jpg" || fileExtension == ".png" {
			return true
		}
	}
	return false
}

func Upload(c echo.Context, file multipart.File, fileheader *multipart.FileHeader, tipe string) (string, string, error) {
	var ur imagekit.UploadRequest

	// check file type
	if isImage := fileCheck(fileheader, "image"); isImage == false {
		return "", "", errors.New("file not an image")
	}

	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGEIO_PUBLIC"),
		PrivateKey: os.Getenv("IMAGEIO_PRIVATE"),
	}

	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		panic(err)
	}

	timestamp := strconv.Itoa(int(time.Now().Unix()))
	fileName := autoGenerate(10, fileheader.Filename) + timestamp + filepath.Ext(fileheader.Filename)
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", "", err
	}

	fmt.Println(fileName)

	if tipe == "news" {
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
	} else if tipe == "asset" {
		ur = imagekit.UploadRequest{
			File:              buf.Bytes(), // []byte OR *url.URL OR url.URL OR base64 string
			FileName:          fileName,
			UseUniqueFileName: false,
			Tags:              []string{},
			Folder:            "/asset",
			IsPrivateFile:     false,
			CustomCoordinates: "",
			ResponseFields:    nil,
		}
	}

	ctx := context.Background()

	r, err := ik.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		return "", "", errors.New("error when upload file: " + err.Error())
	}
	return r.FileID, r.URL, nil
}

func Delete(fileId string) error {

	var fileIdArr []string
	fileIdArr = append(fileIdArr, fileId)

	ctx := context.Background()

	opts := imagekit.Options{
		PublicKey:  os.Getenv("IMAGEIO_PUBLIC"),
		PrivateKey: os.Getenv("IMAGEIO_PRIVATE"),
	}

	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		panic(err)
	}

	resp, err := ik.Media.DeleteFiles(ctx, &imagekit.DeleteFilesRequest{fileIdArr})
	if err != nil {
		return err
	}
	
	if len(resp.SuccessfullyDeletedFileIDs) == 0 {
		return errors.New("nothing file has deleted")
	}

	return nil
}
