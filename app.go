package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"./batchapi"
	"./miniolfs"
	"github.com/gorilla/mux"
)

const (
	serverVersionString = "version: 0.2.0"
	configFile          = "./config.json"
)

type config struct {
	ServerListenAddr string `json:"serverListenAddr"`
	MinioHost        string `json:"minioHost"`
	MinioAccessKey   string `json:"minioAccessKey"`
	MinioSecretKey   string `json:"minioSecretKey"`
	MinioBucket      string `json:"minioBucket"`
	MinioURLExpires  uint64 `json:"minioURLExpires"`
}

var m *miniolfs.MinioLFS
var runningConf config

func main() {
	raw, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	json.Unmarshal(raw, &runningConf)

	m = miniolfs.NewMinioLFS(miniolfs.MinioLFSInitParams{
		Host:       runningConf.MinioHost,
		AccessKey:  runningConf.MinioAccessKey,
		SecretKey:  runningConf.MinioSecretKey,
		Bucket:     runningConf.MinioBucket,
		URLExpires: runningConf.MinioURLExpires,
	})

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/version", versionHandler)
	router.HandleFunc("/objects/batch", batchHandler)
	router.Use(setHTTPHeader)

	server := &http.Server{
		Addr:           runningConf.ServerListenAddr,
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
