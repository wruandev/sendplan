package action

import (
	"errors"
	"fmt"
	"strconv"
	"wruandev/sendplan/entity"
	"wruandev/sendplan/store"
	"wruandev/sendplan/utils"

	"github.com/urfave/cli/v2"
)

func DeletePlan(ctx *cli.Context) error {

	appData, err := store.Storage.ReadData()

	if err != nil {
		fmt.Println("Error reading file data")
		return err
	}

	deletedID, err := strconv.Atoi(ctx.Args().Get(0))

	if err != nil {
		fmt.Println("ID should be present and number")

		return errors.New("ID should be present")
	}

	if len(appData.Data) == 0 {
		fmt.Println("You don't have any saved plan")

	} else {
		var filteredPlans []entity.ApplyPlan
		var deletedPlan entity.ApplyPlan

		for _, plan := range appData.Data {
			if plan.Id != deletedID {
				filteredPlans = append(filteredPlans, plan)
			} else {
				deletedPlan = plan
			}
		}

		if len(filteredPlans) == len(appData.Data) {
			fmt.Printf("  Plan with ID #%d not found \n", deletedPlan.Id)

			return errors.New("ID not found")
		}

		appData.Data = filteredPlans

		if err := store.Storage.WriteData(appData); err != nil {
			fmt.Println("Failed to write to file data: ", err)
			return err
		}

		fmt.Println("")
		fmt.Printf("  Plan with ID #%d successfully deleted \n", deletedPlan.Id)

		utils.RenderTable([]entity.ApplyPlan{deletedPlan})
	}

	return nil
}
