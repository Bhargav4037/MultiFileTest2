package maindb

import (
	"fmt"

	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

func Check() {
	i:=2
	if i==1 {
	pgxv4.RegisterDriver("ghitnj")
	fmt.Println("Hello")
		}
}
