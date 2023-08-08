package utils

import (
	"os"
	"wruandev/sendplan/entity"

	"github.com/jedib0t/go-pretty/v6/table"
)

func RenderTable(plans []entity.ApplyPlan) string {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Position Name", "Company Name", "Posted", "Deadline", "Source"})

	for _, plan := range plans {
		t.AppendRow(table.Row{plan.Id, plan.Position, plan.CompanyName, plan.PostedDate, plan.DeadlineDate, plan.Source})
		t.AppendSeparator()
	}

	return t.Render()
}
