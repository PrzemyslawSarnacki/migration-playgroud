package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
)

var Migrations = []*gormigrate.Migration{
	initialMigration,
	timestampMigration,
}
