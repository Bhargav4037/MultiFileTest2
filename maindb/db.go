package maindb

import "cloud.google.com/go/cloudsqlconn/postgres/pgxv4"

func Check() {
	pgxv4.RegisterDriver("ghitnj")
}
