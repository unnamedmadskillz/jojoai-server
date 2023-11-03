package model

type ReplicRow struct {
	Name  string `bson:"model_name"`
	Text  string `bson:"tts_text"`
	Path  string `bson:"output_file_path"`
	Order int    `bson:"order"`
}

type ReplicDB struct {
	Data []ReplicRow `bson:"data"`
}
