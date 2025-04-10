/*
 *Holds structures and methods for the bank table
 */

/*
* Package Components:

* Constants
* - financeDatabase: name of the database of the table
* - finRefBankTable: name of the table for bank

* Structures
* - Bank: Structure for the bank table
* - BankTable: defines a slice of Banks


* Functions:
* - GetDatabase: Gets the table database
* - GetTableName: Gets the name of the table
* - GetAllData: Returns the rows for the table
* - Append: adds data to the list
 */

package tables

import (
	"log"
	"time"
)

// Constants that hold the database and table names
const (
	financeDatabase = "finance"
	finRefBankTable = "fin_ref_bank"
	BankTableKey    = "bank"
)

// Struct representing the database table
type Bank struct {
	BankID       int       `db:"bank_id"`
	BankName     string    `db:"bank_name"`
	DisplayOrder int       `db:"display_order"`
	CreatedBy    string    `db:"created_by"`
	CreatedOn    time.Time `db:"created_on"`
	UpdatedBy    string    `db:"updated_by"`
	UpdatedOn    time.Time `db:"updated_on"`
	IsActive     bool      `db:"is_active"`
}

// Defines a slice of Banks
type BankTable struct {
	rows []Bank
}

// Get bank database name
func (b *BankTable) GetDatabase() string {
	return financeDatabase
}

// Get bank table name
func (b *BankTable) GetTableName() string {
	return finRefBankTable
}

// Returns the rows for the table
func (b *BankTable) GetBaseTableStruct() any {
	return &Bank{}
}

// adds data to the list
func (b *BankTable) Append(value any) {

	switch v := value.(type) {
	case Bank:
		b.rows = append(b.rows, v)
	case *Bank:
		b.rows = append(b.rows, *v)
	default:
		log.Printf("Append failed: value is not of type Bank or *Bank")
	}
}

// Returns all the rows of the table
func (b *BankTable) GetRows() any {
	return b.rows
}
