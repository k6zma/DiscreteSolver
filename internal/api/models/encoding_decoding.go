package models

type EncodeDecodeRequest struct {
	String string `json:"string"`
}

type DecodeRequest struct {
	EncodedString string            `json:"encoded_string"`
	Alphabet      map[string]string `json:"alphabet"`
}

type EncodeResponse struct {
	EncodedString     string            `json:"encoded_string"`
	Alphabet          map[string]string `json:"alphabet"`
	AverageCodeLength float64           `json:"average_code_length"`
}

type DecodeResponse struct {
	DecodedString string `json:"decoded_string"`
}
