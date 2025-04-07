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

import (
	"fmt"
)

// Holds all functions to be globally used by all tables and table types
type Table interface {
	GetDatabase() string
	GetTableName() string
}

// takes a key and returns the table type
func TableFactory(key string) (Table, error) {
	switch key {
	case BankTableKey:
		return BankTable{}, nil
	default:
		return nil, fmt.Errorf("unsupported table key: %s", key)
	}

}
