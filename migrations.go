package hook360

import "embed"

//go:embed db/migrations/*.sql
var Migrations embed.FS
