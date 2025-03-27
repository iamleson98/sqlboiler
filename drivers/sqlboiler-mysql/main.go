package main

import (
	"github.com/iamleson98/sqlboiler/v4/drivers"
	"github.com/iamleson98/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
