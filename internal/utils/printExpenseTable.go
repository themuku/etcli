package utils

import (
	"etcli/internal/models"
	"fmt"
	"os"

	"text/tabwriter"
)

func PrintExpenseTable(expenses map[int]models.Expense) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	_, err := fmt.Fprintln(w, "ID\tDATE\tDESCRIPTION\tAMOUNT")
	if err != nil {
		return
	}

	_, err = fmt.Fprintln(w, "--\t----\t-----------\t------")
	if err != nil {
		return
	}

	for id, expense := range expenses {

		_, err := fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n",
			id,
			expense.CreatedAt,
			expense.Description,
			expense.Amount,
		)
		if err != nil {
			errMsg := fmt.Errorf("error writing to tabwriter: %v", err)
			fmt.Println(errMsg)
		}
	}

	err = w.Flush()
	if err != nil {
		return
	}

}
