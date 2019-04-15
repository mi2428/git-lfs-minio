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
	var response apiResponse
	for _, object := range r.Objects {
		oid := object.Oid
		if !m.IsExist(oid) {
			object_not_found()
		}
		url := m.DownloadURL(oid)
	}
	json.NewEncoder(w).Encode(response)
}

func upload(r apiRequest) {
	// 1. check if the requested file is exist (accept non-exist case)
	// 2. generate pre-signed url
	// 3. return response structure
}

func verify(r apiRequest) {
	// 1. check if the requested file is exist
}

func object_not_found() {
	// return 404
}

func object_already_exist() {
	// return 400?
}
