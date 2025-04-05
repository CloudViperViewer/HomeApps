package tables

import "time"

//Database and table constants
const (
	Database  = "finance"
	TableName = "fin_ref_bank"
)

//Struct representing the database table
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
