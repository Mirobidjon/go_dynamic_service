package helper

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	mincre "github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mirobidjon/go_dynamic_service/config"
)

func MinioUploader(cfg config.Config, object multipart.File, fileName, contentType string, fileSize int64) error {

	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  mincre.NewStaticV4(cfg.MinioAccessKeyID, cfg.MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return fmt.Errorf("error while getting %s minio client err: %v", fileName, err)
	}

	exists, err := minioClient.BucketExists(context.Background(), cfg.MinioBucketName)
	if err != nil {
		return fmt.Errorf("error while getting %s bucket err: %v", fileName, err)
	}

	if !exists {
		err = minioClient.MakeBucket(context.Background(), cfg.MinioBucketName, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			return fmt.Errorf("error while create %s bucket, err: %v", fileName, err)
		}
	}

	_, err = minioClient.PutObject(context.Background(), cfg.MinioBucketName, fileName, object, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return fmt.Errorf("error while uploading file to minIO, err: %v", err)
	}

	return nil
}

func FileUpload(cfg config.Config, object multipart.File, fileName, contentType string, fileSize int64) (string, error) {
	link := fmt.Sprintf("%s/client-api/download-file?link=/%s", cfg.BaseURL, fileName)

	return link, MinioUploader(cfg, object, fileName, contentType, fileSize)
}
