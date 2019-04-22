package miniolfs

import (
	"log"
	"net/url"
	"time"

	minio "github.com/minio/minio-go"
)

type MinioLFSInitParams struct {
	Host       string
	AccessKey  string
	SecretKey  string
	Bucket     string
	URLExpires time.Duration
}

type MinioLFS struct {
	api        *minio.Client
	bucket     string
	urlExpires time.time
}

func NewMinioLFS(p MinioLFSInitParams) *MinioLFS {
	m := new(MinioLFS)
	api, err := minio.New(p.Host, p.AccessKey, p.SecretKey, false)
	if err != nil {
		log.Fatal(err)
	} else {
		m.api = api
	}
	m.bucket = p.Bucket
	m.urlExpires = p.URLExpires
	return m
}

func (m *MinioLFS) IsExist(oid string) bool {
	if _, err := m.api.StatObject(m.bucket, oid, minio.StatObjectOptions{}); err != nil {
		res := minio.ToErrorResponse(err)
		switch res.Code {
		case "NoSuchBucket":
		case "NoSuchKey":
			return false
		default:
			log.Fatal(err)
		}
	}
	return true
}

func (m *MinioLFS) DownloadURL(oid string) *url.URL {
	reqParams := make(url.Values)
	presignedURL, err := m.api.PresignedGetObject(m.bucket, oid, m.urlExpires*time.Second, reqParams)
	if err != nil {
		log.Fatal(err)
	}
	return presignedURL
}

func (m *MinioLFS) UploadURL(oid string) *url.URL {
	presignedURL, err := m.api.PresignedPutObject(m.bucket, oid, m.urlExpires*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	return presignedURL
}
