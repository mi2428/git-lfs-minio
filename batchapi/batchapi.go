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
	case "upload":
		upload(w, reqbody, m)
		break
	case "download":
		download(w, reqbody, m)
		break
	case "verify":
		verify(w, reqbody, m)
		break
	}
}

func upload(w http.ResponseWriter, r apiRequest, m *miniolfs.MinioLFS) {
	var response apiResponse

	for _, object := range r.Objects {
		var resobj apiResObject
		oid := object.Oid
		size := object.Size

		resobj.Oid = oid
		resobj.Size = size

		if m.IsExist(oid) {
			resobj.Error = &apiResObjError{
				Code:    "422",
				Message: "Object already exist",
			}
			response.Objects = append(response.Objects, resobj)
			continue
		}
		resobj.Error = nil

		url := m.UploadURL(oid)
		expires_at := time.Now().Add(m.URLExpires)

		resobj.Actions = &apiResObjActions{
			Upload: &apiResObjActUpload{
				ExpiresAt: expires_at.Format(time.RFC3339),
				Header:    nil,
				Href:      url.String(),
			},
			Download: nil,
			Verify:   nil,
		}
		response.Objects = append(response.Objects, resobj)
	}

	json.NewEncoder(w).Encode(response)
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
				Code:    "404",
				Message: "Object not found",
			}
			response.Objects = append(response.Objects, resobj)
			continue
		}
		resobj.Error = nil

		url := m.DownloadURL(oid)
		expires_at := time.Now().Add(m.URLExpires)

		resobj.Actions = &apiResObjActions{
			Upload: nil,
			Download: &apiResObjActDownload{
				ExpiresAt: expires_at.Format(time.RFC3339),
				Header:    nil,
				Href:      url.String(),
			},
			Verify: nil,
		}
		response.Objects = append(response.Objects, resobj)
	}

	json.NewEncoder(w).Encode(response)
}

func verify(w http.ResponseWriter, r apiRequest, m *miniolfs.MinioLFS) {
	// TBD
}
