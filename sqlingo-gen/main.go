package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s username:password@/database\n", os.Args[0])
		return
	}

	dataSourceName := os.Args[1]

	code, err := generate("mysql", dataSourceName, nil)
	if err != nil {
		panic(err)
	}

	fmt.Print(code)
}
