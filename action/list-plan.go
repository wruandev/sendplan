package action

import (
	"fmt"
	"wruandev/sendplan/store"
	"wruandev/sendplan/utils"

	"github.com/urfave/cli/v2"
)

func ListPlan(ctx *cli.Context) error {

	appData, err := store.Storage.ReadData()

	if err != nil {
		fmt.Println("Error reading file data")
		return err
	}

	if len(appData.Data) == 0 {
		fmt.Println("You don't have any saved plan")

	} else {
		utils.RenderTable(appData.Data)

	}

	return nil
}
