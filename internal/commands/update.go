package commands

import (
	"etcli/internal/database"
	"fmt"

	"github.com/spf13/cobra"
)

func UpdateCommand(cmd *cobra.Command, args []string, repository database.Repository) error {
	// Implementation for updating an expense with the given ID, description, and amount

	id, err := cmd.Flags().GetInt("id")
	if err != nil {
		return fmt.Errorf("Error reading ID: %v", err)
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return fmt.Errorf("Error reading description: %v", err)
	}
	amount, err := cmd.Flags().GetFloat64("amount")
	if err != nil {
		return fmt.Errorf("Error reading amount: %v", err)
	}

	data, err := repository.ReadExpenses()
	if err != nil {
		return err
	}
	if _, exists := data[id]; !exists {
		return fmt.Errorf("Expense with ID %d not found", id)
	}

	item := data[id]
	if description != "" {
		item.Description = description
	}
	if amount != 0 {
		item.Amount = amount
	}
	data[id] = item

	err = repository.WriteExpenses(&data)
	if err != nil {
		return err
	}
	return nil
}
