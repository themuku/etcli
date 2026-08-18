package main

import (
	"etcli/internal/commands"
	"etcli/internal/database"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	repository := &database.ExpenseRepository{}

	rootCommand := cobra.Command{
		Use:   "etcli",
		Short: "A simple CLI for tracking expenses",
		Long:  "etcli is a command-line application that allows users to track their expenses easily and efficiently.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	addCommand := cobra.Command{
		Use:   "add",
		Short: "Add a new expense",
		Run: func(cmd *cobra.Command, args []string) {
			commands.AddCommand(cmd, args, repository)
		},
	}
	addCommand.Flags().StringP("description", "d", "", "Description of the expense")
	addCommand.Flags().Float64P("amount", "a", 0.0, "Amount of the expense")

	updateCommand := cobra.Command{
		Use:   "update",
		Short: "Update an existing expense",
		Run: func(cmd *cobra.Command, args []string) {
			commands.UpdateCommand(cmd, args, repository)
		},
	}
	updateCommand.Flags().IntP("id", "i", 0, "ID of the expense to update")
	updateCommand.Flags().StringP("description", "d", "", "Description of the expense")
	updateCommand.Flags().Float64P("amount", "a", 0.0, "Amount of the expense")

	deleteCommand := cobra.Command{
		Use:   "delete",
		Short: "Delete an expense",
		Run: func(cmd *cobra.Command, args []string) {
			commands.DeleteCommand(cmd, args, repository)
		},
	}
	deleteCommand.Flags().IntP("id", "i", 0, "ID of the expense to delete")

	listCommand := cobra.Command{
		Use:   "list",
		Short: "List all expenses",
		Run: func(cmd *cobra.Command, args []string) {
			commands.ListCommand(cmd, args, repository)
		},
	}

	summaryCommand := cobra.Command{
		Use:   "summary",
		Short: "Show a summary of expenses",
		Run: func(cmd *cobra.Command, args []string) {
			commands.SummaryCommand(cmd, args, repository)
		},
	}
	summaryCommand.Flags().StringP("month", "m", "", "Month for the summary (format: MM)")

	rootCommand.AddCommand(&addCommand, &updateCommand, &deleteCommand, &listCommand, &summaryCommand)

	if err := rootCommand.Execute(); err != nil {
		err = fmt.Errorf("Error: %v\n", err)
		fmt.Println(err)
		os.Exit(1)
	}
}
