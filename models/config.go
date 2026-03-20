package models

type Gateway struct {
	Name string `json:"name"`
}

type Logging struct {
	Level     string `json:"level"`
	AccessLog struct {
		Enabled  bool   `json:"enabled"`
		Format   string `json:"format"`
		FileName string `json:"file_name"`
	}
}

type Config struct {
	Gateway  Gateway   `json:"gateway"`
	Services []Service `json:"services"`
	Logging  Logging   `json:"logging"`
}
