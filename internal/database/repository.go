package database

import (
	"encoding/csv"
	"errors"
	"etcli/internal/models"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Repository interface {
	ReadExpenses() (map[int]models.Expense, error)
	WriteExpenses(expenses *map[int]models.Expense) error
}

type ExpenseRepository struct{}

func (r *ExpenseRepository) ReadExpenses() (map[int]models.Expense, error) {
	file, err := os.Open("./internal/database/expenses.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]models.Expense), nil
		}
		fmt.Println("error opening file:", err)
		return nil, err
	}
	defer file.Close()

	data := make(map[int]models.Expense)
	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading csv:", err)
			return nil, err
		}
		if len(record) < 4 {
			return nil, errors.New("not enough data in CSV record")
		}

		amount, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing amount '%s': %w", record[3], err)
		}
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing ID '%s': %w", record[0], err)
		}
		data[id] = models.Expense{
			Description: record[2],
			Amount:      amount,
			CreatedAt:   record[1],
		}
	}
	return data, nil
}

func (r *ExpenseRepository) WriteExpenses(expenses *map[int]models.Expense) error {
	_, err := os.Stat("./internal/database/expenses.csv")
	if os.IsNotExist(err) {
		file, err := os.Create("./internal/database/expenses.csv")
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
		file.Close()
	}

	file, err := os.OpenFile("./internal/database/expenses.csv", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for writing: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	for id, expense := range *expenses {
		record := []string{
			strconv.Itoa(id),
			expense.CreatedAt,
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
		}
		err := writer.Write(record)
		if err != nil {
			return fmt.Errorf("error writing to CSV: %w", err)
		}
	}
	writer.Flush()
	return nil
}
