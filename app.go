package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"./batchapi"
	"./miniolfs"
	"github.com/gorilla/mux"
)

const (
	minioHost           = "localhost:9000"
	minioAccessKey      = "minio-access-key"
	minioSecretKey      = "minio-secret-key"
	minioBucket         = "git-lfs"
	minioURLExpires     = 3600
	serverVersionString = "version: 0.1.0"
)

var m *miniolfs.MinioLFS

func main() {
	m = miniolfs.NewMinioLFS(miniolfs.MinioLFSInitParams{
		Host:       minioHost,
		AccessKey:  minioAccessKey,
		SecretKey:  minioSecretKey,
		Bucket:     minioBucket,
		URLExpires: minioURLExpires
	})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/version", versionHandler)
	router.HandleFunc("/objects/batch", batchHandler)
	router.HandleFunc("/verify", verifyHandler)
	router.Use(setHTTPHeader)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}

func setHTTPHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept", "application/vnd.git-lfs+json")
		w.Header().Set("Content-Type", "application/vnd.git-lfs+json")
		next.ServeHTTP(w, r)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It works!")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, serverVersionString)
}

func batchHandler(w http.ResponseWriter, r *http.Request) {
	batchapi.RequestHandler(w, r, m)
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It works!")
}
