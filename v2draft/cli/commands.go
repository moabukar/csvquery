// github.com/moabukar/csvquery/cli/commands.go
package cli

import (
	"fmt"
	"strconv"
	"strings"
)

func ExportCMD(words []string) {
	if ok := validateCommandParamters(2, len(words)-1); !ok {
		return
	}

	err := createFileFromTable(words[1], words[2])
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("file exported")
}

func ShowCMD(words []string) {
	if ok := validateCommandParamters(2, len(words)-1); !ok {
		return
	}

	n, _ := strconv.Atoi(words[2])
	columns, rows, err := tableSelect(words[1], n)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tableRender(columns, rows)
}

func ListCMD() {
	err := service.List()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func LoadCMD(words []string) {
	tablename := words[0]
	err := loadFiles(tablename, words[3])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func SelectCMD(words []string) {
	err := createTableFromQuery(words[0], strings.Join(words[2:], " "))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func SelectAppendCMD(words []string) {
	err := appendToTableFromQuery(words[0], strings.Join(words[2:], " "))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
