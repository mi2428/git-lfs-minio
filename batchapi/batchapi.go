package batchapi

import (
	"encoding/json"
	"log"
	"net/http"

	"../miniolfs"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, m *miniolfs.MinioLFS) {
	var reqbody apiRequest
	if err := json.NewDecoder(r.Body).Decode(&reqbody); err != nil {
		log.Fatal(err)
	}
	switch reqbody.Operation {
	case "download":
		download(w, reqbody, m)
		break
	case "upload":
		upload(reqbody)
		break
	case "verify":
		verify(reqbody)
		break
	}
}

func download(w http.ResponseWriter, r apiRequest, m *miniolfs.MinioLFS) {
	// 1. check if the requested file is exist (accept exist case)
	// 2. generate pre-signed url
	// 3. return response structure
}

func upload(r apiRequest) {
	// 1. check if the requested file is exist (accept non-exist case)
	// 2. generate pre-signed url
	// 3. return response structure
}

func verify(r apiRequest) {
	// 1. check if the requested file is exist
}
