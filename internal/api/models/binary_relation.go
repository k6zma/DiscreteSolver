package models

type BinaryRelationModel struct {
	SetOfElements  []string    `json:"set_of_elements"`
	BinaryRelation [][2]string `json:"binary_relation"`
}
