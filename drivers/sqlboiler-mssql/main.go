package main

import (
	"github.com/iamleson98/sqlboiler/v4/drivers"
	"github.com/iamleson98/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
