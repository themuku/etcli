package commands

import (
	"etcli/internal/database"
	"fmt"

	"github.com/spf13/cobra"
)

func DeleteCommand(cmd *cobra.Command, args []string, repository database.Repository) {
	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		fmt.Println("Error reading ID:", err)
		return
	}

	data, err := repository.ReadExpenses()
	if err != nil {
		fmt.Println("Error reading expenses:", err)
		return
	}

	if _, exists := data[id]; !exists {
		fmt.Printf("Expense with ID %d not found\n", id)
		return
	}

	delete(data, id)
	err = repository.WriteExpenses(&data)
	if err != nil {
		fmt.Println("Error writing expenses:", err)
		return
	}

	fmt.Printf("Deleted expense with ID %d\n", id)
}
