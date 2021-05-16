package main

import (
	"context"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"week4/internal/repository"
	"week4/internal/repository/ent"
	"week4/internal/service"
	"week4/internal/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func InitConfig() (*viper.Viper, error){
	viper.AddConfigPath("../../config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return viper.GetViper(), nil
}

func NewDB(v *viper.Viper) (*ent.Client, error) {
	client, err := ent.Open(
			v.Sub("db").GetString("type"),
			v.Sub("db").GetString("dsn"),
		)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

var UserSet = wire.NewSet(
	service.NewUserService,
	repository.NewRepository,
	usecase.NewUserUsecase,
)