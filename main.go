package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Transaction struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Date     string  `json:"date"`
}

var transactionsFile = "transactions.json"

// Function to load transactions from a file
func loadTransactions() ([]Transaction, error) {
	file, err := os.ReadFile(transactionsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Transaction{}, nil // Return an empty slice if file doesn't exist
		}
		return nil, err
	}

	var transactions []Transaction
	err = json.Unmarshal(file, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// Function to save transactions to a file
func saveTransactions(transactions []Transaction) error {
	data, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(transactionsFile, data, 0644)
}

func main() {
	// Command-line flags
	category := flag.String("category", "", "Category of the transaction (e.g., 'Food', 'Utilities')")
	amount := flag.Float64("amount", 0, "Amount of the transaction")
	date := flag.String("date", "", "Date of the transaction (e.g., '2024-12-14')")

	flag.Parse()

	// Validate input
	if *category == "" || *amount <= 0 || *date == "" {
		fmt.Println("Error: All flags (category, amount, date) are required.")
		flag.Usage()
		return
	}

	// Load existing transactions from the file
	transactions, err := loadTransactions()
	if err != nil {
		fmt.Println("Error loading transactions:", err)
		return
	}

	// Create a new transaction
	transaction := Transaction{
		Category: *category,
		Amount:   *amount,
		Date:     *date,
	}

	// Add the new transaction to the slice
	transactions = append(transactions, transaction)

	// Save the updated list of transactions
	err = saveTransactions(transactions)
	if err != nil {
		fmt.Println("Error saving transactions:", err)
		return
	}

	// Success message
	fmt.Println("Transaction added successfully!")

	// Optionally, print all saved transactions
	fmt.Println("\nAll Transactions:")
	for _, t := range transactions {
		fmt.Printf("Category: %s, Amount: %.2f, Date: %s\n", t.Category, t.Amount, t.Date)
	}
}
