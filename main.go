package main

import (
	"context"
	"log/slog"
	"os"
	"skeleton/internal/config"
	"skeleton/internal/settings"
	"skeleton/internal/storage/db/contract"
	"skeleton/internal/storage/db/storage"
)

// точка сборки приложения
func main() {
	// читаем конфиг
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	//в зависимости от используемой БД создаем storage
	st, err := storage.NewStorage(&contract.Settings{ConnStr: conf.DbConnStr, Dialect: conf.Dialect})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	defer st.Close()

	serv, err := NewServer(&settings.ServerSettings{Port: conf.ServerPort, EnvType: conf.EnvType}, logger, st)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	ctx := context.Background()

	err = serv.Start(ctx)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	defer serv.Stop(ctx)
}
