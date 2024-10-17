package main

import (
	"flag"
	"fmt"
	"piblog/src/migrate"
)

func handleMigrateCommand(args []string) {
	migrateCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	helpFlag := migrateCmd.Bool("h", false, "help for migrate command")
	tagFlag := migrateCmd.String("tag", "migrate", "the specified which new blog will have")
	dirFlag := migrateCmd.String("dir", "", "where many html files store")

	err := migrateCmd.Parse(args)
	if err != nil {
		panic(err)
	}
	if *helpFlag {
		migrateCmd.Usage()
	} else if *dirFlag != "" {
		err = migrate.HandleMigrate(*dirFlag, *tagFlag)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Invalid usage of upload command.")
		migrateCmd.Usage()
	}

	// content, err := os.ReadFile(filePath)
	// if err != nil {
	// 	return err
	// }
}
