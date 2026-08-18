package commands

import (
	"etcli/internal/database"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func SummaryCommand(cmd *cobra.Command, args []string, repository database.Repository) {
	// Implementation for showing a summary of expenses for a given month
	month, err := cmd.Flags().GetString("month")
	if err != nil {
		fmt.Println("Error reading month:", err)
		return
	}

	data, err := repository.ReadExpenses()
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}

	var total float64
	for _, expense := range data {
		t, err := time.Parse("2006-01-02", expense.CreatedAt)
		if err != nil {
			continue
		}

		if month == "" || t.Format("01") == month || fmt.Sprintf("%d", int(t.Month())) == month {
			total += expense.Amount
		}
	}

	if month != "" {
		fmt.Printf("Total expenses for %s: $%.2f\n", month, total)
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}
