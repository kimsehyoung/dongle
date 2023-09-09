package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account"},
	}
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id"),
		field.String("email").
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(320)",
			}),
		field.String("hashed_password").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(64)",
			}),
		field.String("name").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(32)",
			}),

		field.String("phone_number").
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
