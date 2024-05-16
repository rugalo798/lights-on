package service

type Request struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

type Header struct {
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	PayloadVersion int64  `json:"payloadVersion"`
}

type Payload struct {
	AccessToken string `json:"accessToken"`
	DevID       string `json:"devId"`
	Value       int64  `json:"value"`
}


