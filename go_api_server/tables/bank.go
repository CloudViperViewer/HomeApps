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
	"time"

	"github.com/CloudViperViewer/HomeApps/utils"
)

// Constants that hold the database and table names
const (
	finRefBankTable = "fin_ref_bank"
	BankTableKey    = "Bank"
)

// Struct representing the database table
type Bank struct {
	BankID       int       `db:"bank_id" json:"bankId,omitempty"`
	BankName     string    `db:"bank_name" json:"bankName,omitempty" binding:"required"`
	DisplayOrder int       `db:"display_order" json:"displayOrder,omitempty"`
	CreatedBy    string    `db:"created_by" json:"createdBy,omitempty" binding:"required"`
	CreatedOn    time.Time `db:"created_on" json:"createdOn,omitempty" binding:"required"`
	UpdatedBy    string    `db:"updated_by" json:"updatedBy,omitempty"`
	UpdatedOn    time.Time `db:"updated_on" json:"updatedOn,omitempty"`
	IsActive     bool      `db:"is_active" json:"isActive,omitempty"`
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
		utils.LogError(utils.ServiceDatabaseApi, "", "Append failed: value is not of type Bank or *Bank")

	}
}

// Returns all the rows of the table
func (b *BankTable) GetRows() any {
	return b.rows
}
