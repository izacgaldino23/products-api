package config

import (
	"os"
	"reflect"
)

type Config struct {
	Database *Database `name:"DATABASE"`
}

type Database struct {
	Host     string `name:"HOST"`
	Name     string `name:"NAME"`
	Username string `name:"USERNAME"`
	Password string `name:"PASSWORD"`
	Port     string `name:"PORT"`
}

var Environment Config

func LoadEnvironment() (err error) {
	// err = godotenv.Load()
	// if err != nil {
	// return
	// }

	Environment.Database = &Database{}
	scanStruct(&Environment)

	return
}

func GetEnvironment() *Config {
	return &Environment
}

func scanStruct(envVar any, name ...string) {
	typeOfENV := reflect.TypeOf(envVar)
	valueOfENV := reflect.ValueOf(envVar)

	for i := 0; i < typeOfENV.Elem().NumField(); i++ {
		envName := typeOfENV.Elem().Field(i).Tag.Get("name")

		if len(name) > 0 {
			envName = name[0] + "_" + envName
		}

		value := os.Getenv(envName)

		if typeOfENV.Elem().Field(i).Type.Kind() == reflect.Pointer {
			temp := valueOfENV.Elem().Field(i).Interface()
			scanStruct(temp, envName)
			continue
		}

		valueOfENV.Elem().Field(i).SetString(value)
	}
}
