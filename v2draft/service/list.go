// github.com/moabukar/csvquery/service/list_service.go
package service

func List() error {
	rows, err := db.Query("select name from sqlite_master;")
	if err != nil {
		return
	}
	var tableRows [][]string

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return
		}

		tableRows = append(tableRows, []string{name})
	}

	tableRender([]string{"Tables"}, tableRows)

	return
}
