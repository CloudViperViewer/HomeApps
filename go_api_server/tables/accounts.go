/*
 *Holds structures and methods for the account table
 */

/*
* Package Components:

* Constants
* - financeDatabase: name of the database of the table
* - finRefBankTable: name of the table for bank

* Structures
* - Account: Structure for the account table
* - AccountTable: defines a slice of Accounts


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
	"github.com/shopspring/decimal"
)

// Constants that hold the database and table names
const (
	finAccountTable = "fin_accounts"
	AccountTableKey = "Account"
)

// Struct representing the database table
type Account struct {
	AccountID     int                 `db:"account_id" json:"accountId,omitempty"`
	AccountName   string              `db:"account_name" json:"accountName,omitempty"`
	AccountTypeID int                 `db:"account_type_id" json:"accountTypeId,omitempty"`
	BankID        int                 `db:"bank_id" json:"bankId,omitempty"`
	AccountNumber string              `db:"account_number" json:"accountNumber,omitempty"`
	BSB           string              `db:"bsb" json:"bsb,omitempty"`
	Balance       decimal.NullDecimal `db:"balance" json:"balance,omitempty"`
	CreatedBy     string              `db:"created_by" json:"createdBy,omitempty"`
	CreatedOn     time.Time           `db:"created_on" json:"createdOn,omitempty"`
	UpdatedBy     string              `db:"updated_by" json:"updatedBy,omitempty"`
	UpdatedOn     time.Time           `db:"updated_on" json:"updatedOn,omitempty"`
	IsActive      bool                `db:"is_active" json:"isActive,omitempty"`
}

// Defines a slice of Accounts
type AccountTable struct {
	rows []Account
}

// Get account database name
func (a *AccountTable) GetDatabase() string {
	return financeDatabase
}

// Get account table name
func (a *AccountTable) GetTableName() string {
	return finAccountTable
}

// Returns the rows for the table
func (a *AccountTable) GetBaseTableStruct() any {
	return &Account{}
}

// adds data to the list
func (a *AccountTable) Append(value any) {

	switch v := value.(type) {
	case Account:
		a.rows = append(a.rows, v)
	case *Account:
		a.rows = append(a.rows, *v)
	default:
		utils.LogError(utils.ServiceDatabaseApi, "", "Append failed: value is not of type Account or *Account")
	}
}

// Returns all the rows of the table
func (a *AccountTable) GetRows() any {
	return a.rows
}
