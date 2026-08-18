# etcli

A lightweight command-line expense tracker written in Go using Cobra.

## Features

- **Add Expenses**: Track new expenses with a description and amount.
- **List Expenses**: Display all logged expenses in a clean, formatted table.
- **Update Expenses**: Modify existing expense descriptions or amounts by ID.
- **Delete Expenses**: Remove unwanted entries by ID.
- **Expense Summary**: View overall expenses or filter totals by month (`MM`).

---

## Installation

### Prerequisites

- Go `1.20+` installed on your system.

### Build from source

1. Clone the repository:

   ```bash
   git clone https://github.com/themuku/etcli.git
   cd etcli
   ```

2. Build the binary:
   ```bash
   go build -o etcli cmd/main.go
   ```

---

## Usage

```bash
etcli [command] [flags]
```

### Commands

| Command   | Description                                 | Example                         |
| :-------- | :------------------------------------------ | :------------------------------ |
| `add`     | Add a new expense                           | `etcli add -d "Coffee" -a 4.50` |
| `list`    | List all saved expenses                     | `etcli list`                    |
| `update`  | Update an existing expense by ID            | `etcli update -i 1 -a 5.00`     |
| `delete`  | Delete an expense by ID                     | `etcli delete -i 1`             |
| `summary` | View total expenses (optional month filter) | `etcli summary -m 08`           |

---

## Flags Summary

- **`-d, --description`** _(string)_: Expense description
- **`-a, --amount`** _(float)_: Expense amount
- **`-i, --id`** _(int)_: Target expense ID
- **`-m, --month`** _(string)_: Month for summary filtering (e.g., `08` for August)

---

## Data Storage

Expenses are stored locally in an `expenses.csv` file created automatically in the execution directory.

---

## License

[MIT](LICENSE)
