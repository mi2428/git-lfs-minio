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
	URLExpires uint64
}

type MinioLFS struct {
	api        *minio.Client
	Bucket     string
	URLExpires time.Duration
}

func NewMinioLFS(p MinioLFSInitParams) *MinioLFS {
	m := new(MinioLFS)
	api, err := minio.New(p.Host, p.AccessKey, p.SecretKey, false)
	if err != nil {
		log.Fatal(err)
	} else {
		m.api = api
	}
	m.Bucket = p.Bucket
	m.URLExpires = time.Duration(p.URLExpires) * time.Second
	return m
}

func (m *MinioLFS) IsExist(oid string) bool {
	if _, err := m.api.StatObject(m.Bucket, oid, minio.StatObjectOptions{}); err != nil {
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
	presignedURL, err := m.api.PresignedGetObject(m.Bucket, oid, m.URLExpires, reqParams)
	if err != nil {
		log.Fatal(err)
	}
	return presignedURL
}

func (m *MinioLFS) UploadURL(oid string) *url.URL {
	presignedURL, err := m.api.PresignedPutObject(m.Bucket, oid, m.URLExpires)
	if err != nil {
		log.Fatal(err)
	}
	return presignedURL
}
