package main

import (
	"flag"
	"fmt"
	"piblog/src/create"
	"piblog/src/upload"
)

func handleUploadCommand(args []string) {
	uploadCmd := flag.NewFlagSet("upload", flag.ExitOnError)
	helpFlag := uploadCmd.Bool("h", false, "help for upload command")
	createFlag := uploadCmd.String("dir", "", "create a new markdown file from template in specified path")
	uploadFlag := uploadCmd.String("u", "", "Upload file in specified path to blogspot")
	forceFlag := uploadCmd.Bool("f", false, "Whether create a new blog when meets same title")

	err := uploadCmd.Parse(args)
	if err != nil {
		panic(err)
	}
	if *helpFlag {
		uploadCmd.Usage()
	} else if *createFlag != "" {
		err = create.HandleCreateFile(*createFlag)
		if err != nil {
			panic(err)
		}
	} else if *uploadFlag != "" {
		// TODO:
		err = upload.HandleUploadBlog(*uploadFlag, *forceFlag)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		fmt.Println("Invalid usage of upload command.")
		uploadCmd.Usage()
	}

	// content, err := os.ReadFile(filePath)
	// if err != nil {
	// 	return err
	// }
}
