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

	r, _ := repo.New()
	ref := r.BranchRef("master")

	if ref != nil {
		fmt.Printf("hey: %v\n", ref)
	} else {
		fmt.Println("no ref")
	}

	commands.Execute()
}
