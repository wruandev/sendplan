package action

import (
	"bufio"
	"fmt"
	"os"
	"wruandev/sendplan/entity"
	"wruandev/sendplan/store"
	"wruandev/sendplan/utils"

	"github.com/urfave/cli/v2"
)

func AddPlan(ctx *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Println("  Add a plan to the list")

	companyName, err := utils.GetUserPlanInputString(reader, "Company name")

	if err != nil {
		utils.PrintError("Company name must be present")
		return err
	}

	position, err := utils.GetUserPlanInputString(reader, "Position title")

	if err != nil {
		utils.PrintError("Position title must be present")
		return err
	}

	postedDate, err := utils.GetUserPlanInputString(reader, "Posted date")

	if err != nil {
		utils.PrintError("Posted date must be present")
		return err
	}

	deadlineDate, err := utils.GetUserPlanInputString(reader, "Deadline date")

	if err != nil {
		utils.PrintError("Deadline date must be present")
		return err
	}

	source, err := utils.GetUserPlanInputString(reader, "Source")

	if err != nil {
		utils.PrintError("Source must be present")
		return err
	}

	applyPlan := entity.ApplyPlan{
		CompanyName:  companyName,
		Position:     position,
		PostedDate:   postedDate,
		DeadlineDate: deadlineDate,
		Source:       source,
	}

	appData, err := store.Storage.ReadData()

	if err != nil {
		fmt.Println("Error reading file data")
		return err
	}

	applyPlan.Id = appData.LastId + 1
	appData.LastId = applyPlan.Id

	appData.Data = append([]entity.ApplyPlan{applyPlan}, appData.Data...)

	if err := store.Storage.WriteData(appData); err != nil {
		fmt.Println("Failed to write to file data: ", err)
		return err
	}

	fmt.Println("")
	fmt.Println("  Plan successfully saved: ")

	utils.RenderTable([]entity.ApplyPlan{applyPlan})

	return nil
}
