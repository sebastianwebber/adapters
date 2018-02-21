package fake

import (
	"net/http"
	"net/url"

	"github.com/prest/adapters"
)

// Fake implements a testable adapter
// it is usefull for mock and test stuff
type Fake struct{}

// TablePermissions get tables permissions based in prest configuration
func (adapter *Fake) TablePermissions(table string, op string) (out bool) {
	return
}

// GetScript do some stuff
func (adapter *Fake) GetScript(verb, folder, scriptName string) (script string, err error) {
	return
}

// ParseScript do some stuff
func (adapter *Fake) ParseScript(scriptPath string, queryURL url.Values) (sqlQuery string, values []interface{}, err error) {
	return
}

// ExecuteScripts do some stuff
func (adapter *Fake) ExecuteScripts(method, sql string, values []interface{}) (sc adapters.Scanner) {
	return
}

// WhereByRequest do some stuff
func (adapter *Fake) WhereByRequest(r *http.Request, initialPlaceholderID int) (whereSyntax string, values []interface{}, err error) {
	return
}

// DatabaseClause do some stuff
func (adapter *Fake) DatabaseClause(req *http.Request) (query string, hasCount bool) {
	return
}

// OrderByRequest do some stuff
func (adapter *Fake) OrderByRequest(r *http.Request) (values string, err error) {
	return
}

// PaginateIfPossible do some stuff
func (adapter *Fake) PaginateIfPossible(r *http.Request) (paginatedQuery string, err error) {
	return
}

// Query do some stuff
func (adapter *Fake) Query(SQL string, params ...interface{}) (sc adapters.Scanner) {
	return
}

// SchemaClause do some stuff
func (adapter *Fake) SchemaClause(req *http.Request) (query string, hasCount bool) {
	return
}

// FieldsPermissions do some stuff
func (adapter *Fake) FieldsPermissions(r *http.Request, table string, op string) (fields []string, err error) {
	return
}

// SelectFields do some stuff
func (adapter *Fake) SelectFields(fields []string) (sql string, err error) {
	return
}

// CountByRequest do some stuff
func (adapter *Fake) CountByRequest(req *http.Request) (countQuery string, err error) {
	return
}

// JoinByRequest do some stuff
func (adapter *Fake) JoinByRequest(r *http.Request) (values []string, err error) {
	return
}

// GroupByClause do some stuff
func (adapter *Fake) GroupByClause(r *http.Request) (groupBySQL string) {
	return
}

// QueryCount do some stuff
func (adapter *Fake) QueryCount(SQL string, params ...interface{}) (sc adapters.Scanner) {
	return
}

// ParseInsertRequest do some stuff
func (adapter *Fake) ParseInsertRequest(r *http.Request) (colsName string, colsValue string, values []interface{}, err error) {
	return
}

// Insert do some stuff
func (adapter *Fake) Insert(SQL string, params ...interface{}) (sc adapters.Scanner) {
	return
}

// Delete do some stuff
func (adapter *Fake) Delete(SQL string, params ...interface{}) (sc adapters.Scanner) {
	return
}

// SetByRequest do some stuff
func (adapter *Fake) SetByRequest(r *http.Request, initialPlaceholderID int) (setSyntax string, values []interface{}, err error) {
	return
}

// Update do some stuff
func (adapter *Fake) Update(SQL string, params ...interface{}) (sc adapters.Scanner) {
	return
}

// DistinctClause do some stuff
func (adapter *Fake) DistinctClause(r *http.Request) (distinctQuery string, err error) {
	return
}

// SetDatabase do some stuff
func (adapter *Fake) SetDatabase(name string) {
	return
}

// SelectSQL do some stuff
func (adapter *Fake) SelectSQL(selectStr string, database string, schema string, table string) (out string) {
	return
}

// InsertSQL do some stuff
func (adapter *Fake) InsertSQL(database string, schema string, table string, names string, placeholders string) (out string) {
	return
}

// DeleteSQL do some stuff
func (adapter *Fake) DeleteSQL(database string, schema string, table string) (out string) {
	return
}

// UpdateSQL do some stuff
func (adapter *Fake) UpdateSQL(database string, schema string, table string, setSyntax string) (out string) {
	return
}
