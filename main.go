package main

import (
	"log"
	"os"
	"wruandev/sendplan/action"
	"wruandev/sendplan/config"
	"wruandev/sendplan/store"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "sendplan",
		Usage: "Save your job applying plan for easy access",
		Commands: []*cli.Command{
			{
				Name:   "add",
				Usage:  "Add job applying plan to the list",
				Action: action.AddPlan,
			},
			{
				Name:   "list",
				Usage:  "Display list of all saved plans",
				Action: action.ListPlan,
			},
			{
				Name:   "delete",
				Usage:  "Delete a saved plan by its ID",
				Action: action.DeletePlan,
			},
			{
				Name:   "check",
				Usage:  "Check all saved plan that almost reach deadline",
				Action: action.CheckPlan,
			},
		},
		Before: func(ctx *cli.Context) error {

			jsonFileStorage := store.JsonFileStorage{
				FolderPath: config.AppDataFolderPath,
				FilePath:   config.AppDataFilePath,
			}

			store.SetStorage(jsonFileStorage)

			if err := store.Storage.InitData(); err != nil {
				log.Fatal(err)

				return err
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
