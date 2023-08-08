package config

import "path/filepath"

const (
	AppDataFolderName = "sendplan-data"
	AppDataFileName   = "data.json"
)

var AppDataFolderPath = filepath.Join(".", AppDataFolderName)
var AppDataFilePath = filepath.Join(AppDataFolderPath, AppDataFileName)
