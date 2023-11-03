package model

type Dialogue struct {
	Replic []struct {
		Speaker   string `json:"speaker"`
		Utterance string `json:"text"`
	} `json:"dialogue"`
}
