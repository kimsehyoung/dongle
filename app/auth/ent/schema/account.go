package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("role_id"),
		field.String("login_id").
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(32)",
			}),
		field.String("hashed_password").
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(32)",
			}),
		field.String("name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(32)",
			}),
		field.String("email").
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(320)",
			}),
		field.String("phone_number").
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(32)",
			}),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Default(time.Now()),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("accounts").
			Unique().
			Required().
			Field("role_id"),
	}
}
