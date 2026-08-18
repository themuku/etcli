package commands

import (
	"etcli/internal/database"
	"etcli/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

func ListCommand(cmd *cobra.Command, args []string, repository database.Repository) {
	data, err := repository.ReadExpenses()
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}

	if len(data) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	utils.PrintExpenseTable(data)
}
