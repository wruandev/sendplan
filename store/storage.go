package store

import "wruandev/sendplan/entity"

type IStorage interface {
	ReadData() (entity.AppData, error)
	InitData() error
	WriteData(appData entity.AppData) error
}

var Storage IStorage

func SetStorage(str IStorage) {
	Storage = str
}
