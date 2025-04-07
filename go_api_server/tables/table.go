/*
 *Set table interface
 */

/*
* Package Components:

* Interface:
* - Holds all functions to be globally used by all tables
 */

package tables

type Table interface {
	GetDatabase() string
	GetTableName() string
}
