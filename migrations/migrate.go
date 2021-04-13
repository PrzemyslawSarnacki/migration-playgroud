package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

var migrations = []*gormigrate.Migration{
	initialMigration,
	timestampMigration,
}
