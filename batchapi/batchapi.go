package batchapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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
		var resobj apiResObject
		oid := object.Oid
		size := object.Size

		resobj.Oid = oid
		resobj.Size = size

		if !m.IsExist(oid) {
			resobj.Error = &apiResObjError{
				Code:    "400",
				Message: "something wrong",
			}
			response.Objects = append(response.Objects, resobj)
			continue
		}
		resobj.Error = nil

		url := m.DownloadURL(oid)
		expires_at := time.Now().Add(m.URLExpires)

		resobj.Actions.Download = nil
		resobj.Actions.Verify = nil
		resobj.Actions.Upload = &apiResObjActUpload{
			ExpiresAt: expires_at.Format(time.RFC3339),
			Header:    nil,
			Href:      url.String(),
		}
		response.Objects = append(response.Objects, resobj)
	}

	json.NewEncoder(w).Encode(response)
}

func upload(r apiRequest) {
	var response apiResponse

	for _, object := range r.Objects {
		var resobj apiResObject
		oid := object.Oid
		size := object.Size

		resobj.Oid = oid
		resobj.Size = size

		if !m.IsExist(oid) {
			resobj.Error = &apiResObjError{
				Code:    "400",
				Message: "something wrong",
			}
			response.Objects = append(response.Objects, resobj)
			continue
		}
		resobj.Error = nil

		url := m.DownloadURL(oid)
		expires_at := time.Now().Add(m.URLExpires)

		resobj.Actions.Download = nil
		resobj.Actions.Verify = nil
		resobj.Actions.Upload = &apiResObjActUpload{
			ExpiresAt: expires_at.Format(time.RFC3339),
			Header:    nil,
			Href:      url.String(),
		}
		response.Objects = append(response.Objects, resobj)
	}

	json.NewEncoder(w).Encode(response)
}

func verify(r apiRequest) {
	// 1. check if the requested file is exist
}
