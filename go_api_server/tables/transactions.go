/*
 *Holds structures and methods for the Transaction table
 */

/*
* Package Components:

* Constants
* - financeDatabase: name of the database of the table
* - finTransactionTable: name of the table for Transaction

* Structures
* - Transaction: Structure for the Transaction table
* - TransactionTable: defines a slice of Transactions


* Functions:
* - GetDatabase: Gets the table database
* - GetTableName: Gets the name of the table
* - GetBaseTableStruct: Returns the base struct for scanning
* - GetRows: Returns appended rows
* - Append: adds data to the list
 */

package tables

import (
	"database/sql"
	"time"

	"github.com/CloudViperViewer/HomeApps/utils"
	"github.com/shopspring/decimal"
)

// Constants that hold the database and table names
const (
	finTransactionTable  = "fin_transaction"
	TransactionTableKey  = "Transaction"
	TransactonPrimaryKey = "TransactionId"
)

// Struct representing the database table
type Transaction struct {
	TransactionId      int             `db:"transaction_id" json:"transactionId,omitempty"`
	AccountId          int             `db:"account_id" json:"accountId,omitempty" binding:"required"`
	TransactionTypeId  int             `db:"transaction_type_id" json:"transactionTypeId,omitempty" binding:"required"`
	Value              decimal.Decimal `db:"value" json:"value,omitempty" binding:"required"`
	RecurringPaymentId sql.NullInt16   `db:"recurring_payment_id" json:"recurringPaymentId,omitempty"`
	OnOffBillId        sql.NullInt16   `db:"on_off_bill_id" json:"onOffBillId,omitempty"`
	ViaPaypal          bool            `db:"via_paypal" json:"viaPaypal"`
	DateTime           time.Time       `db:"date_time" json:"dateTime,omitempty" binding:"required"`
	TransactionWith    string          `db:"transaction_with" json:"transactionWith,omitempty"`
	CreatedBy          string          `db:"created_by" json:"createdBy,omitempty" binding:"required"`
	CreatedOn          time.Time       `db:"created_on" json:"createdOn,omitempty" binding:"required"`
	UpdatedBy          string          `db:"updated_by" json:"updatedBy,omitempty"`
	UpdatedOn          time.Time       `db:"updated_on" json:"updatedOn,omitempty"`
	IsActive           bool            `db:"is_active" json:"isActive"`
}

// Defines a slice of Transactions
type TransactionTable struct {
	rows []Transaction
}

// Get Transaction database name
func (t *TransactionTable) GetDatabase() string {
	return financeDatabase
}

// Get Transaction table name
func (t *TransactionTable) GetTableName() string {
	return finTransactionTable
}

// Get Transaction Primary key
func (t *TransactionTable) GetPrimaryKey() string {
	return TransactonPrimaryKey
}

// Returns the rows for the table
func (t *TransactionTable) GetBaseTableStruct() any {
	return &Transaction{}
}

// adds data to the list
func (t *TransactionTable) Append(value any) {

	switch v := value.(type) {
	case Transaction:
		t.rows = append(t.rows, v)
	case *Transaction:
		t.rows = append(t.rows, *v)
	default:
		utils.LogError(utils.ServiceDatabaseApi, "", "Append failed: value is not of type Transaction or *Transaction")

	}
}

// Returns all the rows of the table
func (t *TransactionTable) GetRows() any {
	return t.rows
}
