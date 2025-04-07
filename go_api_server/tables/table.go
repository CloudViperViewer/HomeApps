/*
 *Set table interface
 */

/*
* Package Components:


* Interface:
* - Holds all functions to be globally used by all tables


* Functions:
* - TableFactory: takes a key and returns the table type
 */

package tables

// Holds all functions to be globally used by all tables and table types
type Table interface {
	GetDatabase() string
	GetTableName() string
}

// takes a key and returns the table type
func TableFactory(key string) Table {
	switch key {
	case "bank":
		return BankTable{}
	default:
		return nil
	}

}
