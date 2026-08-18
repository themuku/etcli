package utils

import "etcli/internal/database"

func GenerateID(repository database.Repository) int {
	data, err := repository.ReadExpenses()
	if err != nil || len(data) == 0 {
		return 1
	}

	maxID := 0
	for id := range data {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}
