package server

type createStationRequest struct {
	Name string `json:"name"`
}

type createStationResponse struct {
	ID  string  `json:"id"`
	Err *ErrObj `json:"err,omitempty"`
}

type ErrObj struct {
	ErrCode    string `json:"err_code"`
	ErrMessage string `json:"err_message"`
}
