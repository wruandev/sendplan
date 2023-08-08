package store

import (
	"encoding/json"
	"io"
	"os"
	"wruandev/sendplan/entity"
)

type JsonFileStorage struct {
	FolderPath string
	FilePath   string
}

func (storage JsonFileStorage) ReadData() (entity.AppData, error) {
	rfile, err := os.OpenFile(storage.FilePath, os.O_RDONLY, 0664)

	if err != nil {
		return entity.AppData{}, err
	}

	defer rfile.Close()

	var appData entity.AppData

	b, err := io.ReadAll(rfile)

	if err != nil {
		return entity.AppData{}, err
	}

	if err := json.Unmarshal(b, &appData); err != nil {
		return entity.AppData{}, err
	}

	return appData, nil
}

func (storage JsonFileStorage) InitData() error {
	err := os.MkdirAll(storage.FolderPath, os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(storage.FilePath, os.O_RDONLY|os.O_CREATE, 0664)

	if err != nil {
		return err
	}

	fi, err := file.Stat()

	if err != nil {
		return err
	}

	if fi.Size() == 0 {

		appData := entity.AppData{
			Version: 1,
			LastId:  0,
			Data:    []entity.ApplyPlan{},
		}

		if err := storage.WriteData(appData); err != nil {
			return err
		}

	}

	return nil
}

func (storage JsonFileStorage) WriteData(appData entity.AppData) error {
	file, err := os.OpenFile(storage.FilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(appData); err != nil {
		return err
	}

	return nil
}
