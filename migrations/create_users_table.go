// migrations/create_users_table.go
package migrations

// import (
// 	"github.com/go-jet/jet/v2/migrations"
// 	"github.com/go-jet/jet/v2/postgres"
// )

// func init() {
// 	migrations.MustRegister(func(db migrations.Database) {
// 		// Up Migration: Create the 'users' table
// 		users := postgres.Table("users",
// 			postgres.Column("id", postgres.Serial, postgres.PrimaryKey),
// 			postgres.Column("name", postgres.VarChar(255).NotNull()),
// 			postgres.Column("email", postgres.VarChar(255).Unique().NotNull()),
// 			postgres.Column("created_at", postgres.Timestamp().Default(postgres.CurrentTimestamp)),
// 		)

// 		db.CreateTable(users)
// 	}, func(db migrations.Database) {
// 		// Down Migration: Drop the 'users' table
// 		db.DropTable("users")
// 	})
// }
