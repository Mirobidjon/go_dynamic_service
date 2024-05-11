package handlers

import (
	"context"
	"fmt"
	"os"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	mincre "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mirobidjon/go_dynamic_service/api/http"
	"github.com/mirobidjon/go_dynamic_service/api/models"
	"github.com/mirobidjon/go_dynamic_service/config"
	"github.com/mirobidjon/go_dynamic_service/pkg/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Download godoc
// @Router /client-api/download-file [get]
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param link query string true "link"
// @Success 200 {object} http.Response{data=string} "Success"
// @Response 400 {object} http.Response{error=string} "Bad Request"
// @Failure 500 {object} http.Response{error=string} "Server Error"
func (h *Handler) DownloadFile(c *fiber.Ctx) error {
	link := c.Query("link")

	var (
		fileName string
		arr      = strings.Split(link, "/")
	)

	fileName = arr[len(arr)-1]

	file, err := os.Create(fileName)

	if err != nil {
		return h.handleResponse(c, http.BadEnvironment, "os create"+err.Error(), "", "")
	}

	defer func() {
		err = file.Close()
		if err != nil {
			h.log.Error("error file close msg:" + err.Error())
		}
		err = os.Remove(fileName)
		if err != nil {
			h.log.Error("error file remove msg:" + err.Error())
		}
	}()

	// Initialize a Minio client.
	minioClient, err := minio.New(config.Load().MinioEndpoint, &minio.Options{
		Creds:  mincre.NewStaticV4(config.Load().MinioAccessKeyID, config.Load().MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return err
	}

	// Open the file for writing.
	file1, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file1.Close()

	// Download the object from Minio and write it to the file.
	err = minioClient.FGetObject(context.Background(), h.cfg.MinioBucketName, fileName, fmt.Sprintf("./%s", fileName), minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	err = c.Download(fileName)

	return nil
}

// UploadFile godoc
// @Security ApiKeyAuth
// @Param platform-id header string true "uuid" default(7d4a4c38-dd84-4902-b744-0488b80a4c01)
// @Router /client-api/upload [POST]
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 201 {object} http.Response{data=models.PhotoUrl} "Success"
// @Response 422 {object} http.Response{error=string} "Validation Error"
// @Response 400 {object} http.Response{error=string} "Bad Request"
// @Failure 500 {object} http.Response{error=string} "Server Error"
func (h *Handler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return h.handleResponse(c, http.BadEnvironment, err.Error(), "", "")
	}

	contentType := "application/octet-stream"

	object, err := file.Open()
	if err != nil {
		return h.handleResponse(c, http.BadEnvironment, err.Error(), "", "")
	}
	defer object.Close()

	id := primitive.NewObjectID().Hex()
	arr := strings.Split(file.Filename, ".")
	fileName := arr[0]
	fileType := "." + arr[len(arr)-1]
	link, err := helper.FileUpload(h.cfg, object, fileName+id+fileType, contentType, file.Size)
	if err != nil {
		return h.handleResponse(c, http.BadEnvironment, err.Error(), "", "")
	}

	return h.handleResponse(c, http.OK, models.Url{Link: link,
		FileName: fileName + id + fileType}, "", "")
}
