package action

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"wruandev/sendplan/entity"
	"wruandev/sendplan/store"
	"wruandev/sendplan/utils"

	"github.com/gen2brain/beeep"
	"github.com/urfave/cli/v2"
)

func CheckPlan(ctx *cli.Context) error {
	appData, err := store.Storage.ReadData()

	if err != nil {
		fmt.Println("Error reading file data")
		return err
	}

	if len(appData.Data) == 0 {
		fmt.Println("You don't have any saved plan")
		return nil
	}

	var filteredPlan []entity.ApplyPlan

	for _, plan := range appData.Data {
		dateSplit := strings.Split(plan.DeadlineDate, "/")

		deadlineDay, _ := strconv.Atoi(dateSplit[0])
		deadlineMonth, _ := strconv.Atoi(dateSplit[1])
		deadlineYear, _ := strconv.Atoi(dateSplit[2])

		deadlineDate := time.Date(deadlineYear, time.Month(deadlineMonth), deadlineDay, 0, 0, 0, 0, time.Local)
		currentDate := time.Now().In(time.Local)

		difference := deadlineDate.Sub(currentDate)

		if int64(difference.Hours()/24) <= 3 {
			filteredPlan = append(filteredPlan, plan)
		}
	}

	if len(filteredPlan) > 0 {
		fmt.Println("")
		fmt.Println("Plans that near deadline!")

		for _, plan := range filteredPlan {

			notifBody := fmt.Sprintf(
				"you haven't send your application to %s for position %s yet. Send it as soon as possible!",
				plan.CompanyName,
				plan.Position,
			)

			err := beeep.Alert("You have unfinished job", notifBody, "./assets/image.jpg")
			if err != nil {
				panic(err)
			}
		}

		utils.RenderTable(filteredPlan)

	} else {
		fmt.Println("")
		fmt.Println("All plan is good")
	}

	return nil
}
