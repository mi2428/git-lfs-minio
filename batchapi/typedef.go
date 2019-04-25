package batchapi

type apiRequest struct {
	Operation string         `json:"operation"`
	Objects   []apiReqObject `json:"objects"`
}

type apiReqObject struct {
	Oid  string `json:"oid"`
	Size uint   `json:"size"`
}

type apiResponse struct {
	Objects []apiResObject `json:"objects"`
}

type apiResObject struct {
	apiReqObject
	Actions *apiResObjActions `json:"actions,omitempty"`
	Error   *apiResObjError   `json:"error,omitempty"`
}

type apiResObjActions struct {
	Upload   *apiResObjActUpload   `json:"upload,omitempty"`
	Download *apiResObjActDownload `json:"download,omitempty"`
	Verify   *apiResObjActVerify   `json:"verify,omitempty"`
}

type apiResObjActUpload struct {
	Href      string              `json:"href"`
	Header    *apiResObjActHeader `json:"header,omitempty"`
	ExpiresAt string              `json:"expires_at,omitempty"`
}

type apiResObjActDownload struct {
	Href      string              `json:"href"`
	Header    *apiResObjActHeader `json:"header,omitempty"`
	ExpiresAt string              `json:"expires_at,omitempty"`
}

type apiResObjActVerify struct {
	Href      string              `json:"href"`
	Header    *apiResObjActHeader `json:"header,omitempty"`
	ExpiresAt string              `json:"expires_at,omitempty"`
}

type apiResObjActHeader struct {
	Key           string `json:"Key"`
	Authorization string `json:"Authorization"`
}

type apiResObjError struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}
