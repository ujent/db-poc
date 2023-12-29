package main

import (
	"log"
	"skeleton/settings"
	"skeleton/storage"
)

type server struct {
	logger   *log.Logger
	store    storage.Storage
	settings settings.Settings
}
