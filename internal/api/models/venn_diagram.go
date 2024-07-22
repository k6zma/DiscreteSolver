package models

type VennDiagramRequest struct {
	Expression string `json:"expression"`
}

type VennDiagramResponse struct {
	ImageURL string `json:"image_url"`
}
