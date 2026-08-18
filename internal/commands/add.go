package commands

import (
	"etcli/internal/database"
	"etcli/internal/models"
	"etcli/internal/utils"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func AddCommand(cmd *cobra.Command, args []string, repository database.Repository) {
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	amount, err := cmd.Flags().GetFloat64("amount")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data, err := repository.ReadExpenses()
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}

	id := utils.GenerateID(repository)
	data[id] = models.Expense{
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now().Format("2006-01-02"),
	}
	err = repository.WriteExpenses(&data)
	if err != nil {
		fmt.Println("Error writing expenses:", err)
		return
	}

	fmt.Printf("Adding expense: %s - $%.2f\n", description, amount)
}
