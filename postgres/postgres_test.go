package postgres

import (
	"net/http"
	"testing"

	"github.com/nuveo/prest/api"
	. "github.com/smartystreets/goconvey/convey"
)

func TestWhereByRequest(t *testing.T) {
	Convey("Where by request without paginate", t, func() {
		r, err := http.NewRequest("GET", "/databases?dbname=prest&test=cool", nil)
		So(err, ShouldBeNil)
		where := WhereByRequest(r)
		So(where, ShouldContainSubstring, "dbname='prest'")
		So(where, ShouldContainSubstring, "test='cool'")
		So(where, ShouldContainSubstring, "and")
	})
}

func TestConnection(t *testing.T) {
	Convey("Verify database connection", t, func() {
		sqlx := Conn()
		So(sqlx, ShouldNotBeNil)
		err := sqlx.Ping()
		So(err, ShouldBeNil)
	})
}

func TestQuery(t *testing.T) {
	Convey("Query execution", t, func() {
		sql := "SELECT schema_name FROM information_schema.schemata ORDER BY schema_name ASC"
		json, err := Query(sql)
		So(err, ShouldBeNil)
		So(len(json), ShouldBeGreaterThan, 0)
	})

	Convey("Query execution with params", t, func() {
		sql := "SELECT schema_name FROM information_schema.schemata WHERE schema_name = $1 ORDER BY schema_name ASC"
		json, err := Query(sql, "public")
		So(err, ShouldBeNil)
		So(len(json), ShouldBeGreaterThan, 0)
	})
}

func TestPaginateIfPossible(t *testing.T) {
	Convey("Paginate if possible", t, func() {
		r, err := http.NewRequest("GET", "/databases?dbname=prest&test=cool&_page=1&_page_size=20", nil)
		So(err, ShouldBeNil)
		where := PaginateIfPossible(r)
		So(where, ShouldContainSubstring, "LIMIT 20 OFFSET(1 - 1) * 20")
	})
}

func TestInsert(t *testing.T) {
	Convey("Insert data into a table", t, func() {
		r := api.Request{
			Data: map[string]string{
				"name": "prest",
			},
		}
		json, err := Insert("prest", "public", "test", r)
		So(err, ShouldBeNil)
		So(len(json), ShouldBeGreaterThan, 0)
	})
}

func TestDelete(t *testing.T) {
	Convey("Delete data from table", t, func() {
		json, err := Delete("prest", "public", "test", "name='nuveo'")
		So(err, ShouldBeNil)
		So(len(json), ShouldBeGreaterThan, 0)
	})
}

func TestUpdate(t *testing.T) {
	Convey("Update data into a table", t, func() {
		r := api.Request{
			Data: map[string]string{
				"name": "prest",
			},
		}
		json, err := Update("prest", "public", "test", "name='prest'", r)
		So(err, ShouldBeNil)
		So(len(json), ShouldBeGreaterThan, 0)
	})
}