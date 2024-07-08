package maindb

import (
	"fmt"

	_ "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func Check() {
	// pgxv4.RegisterDriver("ghitnj")
	fmt.Println("Hello")
}
