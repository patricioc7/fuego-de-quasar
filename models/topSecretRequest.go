package models

type Satellite struct {
	Name string `json:"name"`
	Distance float32 `json:"distance"`
	Message []string `json:"message"`
}

type TopSecret struct {
	Satellites  [3]Satellite `json:"satellites"`
}