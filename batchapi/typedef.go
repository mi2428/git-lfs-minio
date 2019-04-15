package batchapi

import "time"

type apiRequest struct {
	Operation string         `json:"operation"`
	Objects   []apiReqObject `json:"objects"`
}

type apiReqObject struct {
	Oid  string `json:"oid"`
	Size uint   `json:"size"`
}

type apiResponse struct {
	Operation string         `json:"operation"`
	Objects   []apiResObject `json:"objects"`
}

type apiResObject struct {
	apiReqObject
	Actions apiResObjActions `json:"actions"`
}

type apiResObjActions struct {
	Upload   apiResObjActUpload   `json:"upload"`
	Download apiResObjActDownload `json:"download"`
	Verify   apiResObjActVerify   `json:"verify"`
}

type apiResObjActUpload struct {
	Href      string             `json:"href"`
	Header    apiResObjActHeader `json:"header"`
	ExpiresAt time.Time          `json:"expires_at"`
}

type apiResObjActDownload struct {
	Href      string             `json:"href"`
	Header    apiResObjActHeader `json:"header"`
	ExpiresAt time.Time          `json:"expires_at"`
}

type apiResObjActVerify struct {
	Href      string             `json:"href"`
	Header    apiResObjActHeader `json:"header"`
	ExpiresAt time.Time          `json:"expires_at"`
}

type apiResObjActHeader struct {
	Key           string `json:"Key"`
	Authorization string `json:"Authorization"`
}
