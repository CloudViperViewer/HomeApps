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

const (
	financeDatabase = "finance"
)

// Holds all functions to be globally used by all tables and table types
type Table interface {
	GetDatabase() string
	GetTableName() string
	GetBaseTableStruct() any
	Append(value any)
	GetRows() any
}

// takes a key and returns the table type
func TableFactory(key string) (Table, error) {
	switch key {
	case BankTableKey:
		return &BankTable{}, nil
	case AccountTableKey:
		return &AccountTable{}, nil
	default:
		return nil, fmt.Errorf("unsupported table key: %s", key)
	}

}
