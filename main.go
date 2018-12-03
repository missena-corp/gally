package main

import (
	"fmt"

	"github.com/missena-corp/gally/commands"
	"github.com/missena-corp/gally/repo"
)

func main() {
	// var config config.Config

	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }

	// if err := viper.Unmarshal(&config); err != nil {
	// 	log.Fatalf("unable to decode into struct, %v", err)
	// }

	// log.Println(config.Scripts)

	files, err := repo.UpdatedFiles("master")

	if err != nil {
		fmt.Printf("hey: %v\n", err)
	} else {
		fmt.Println("oy: %v", files)
	}

	commands.Execute()
}
