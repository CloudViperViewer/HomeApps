/*
 *Holds structures and methods for the bank table
 */

/*
* Package Components:


* Structures
* - Bank: Structure for the bank table


* Functions:
* - GetDatabase: Gets the table database
* - GetTableName: Gets the name of the table
 */

package tables

import "time"

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

func (b Bank) GetDatabase() string {
	return "finance"
}

func (b Bank) GetTableName() string {
	return "fin_ref_bank"
}
