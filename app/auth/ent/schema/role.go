package schema

import (
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "role"},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(16)",
			}).
			Validate(func(in string) error {
				roles := []string{"root", "admin", "user"}
				for _, role := range roles {
					if role == in {
						return nil
					}
				}
				return errors.New("Allowed roles: 'root', 'admin', or 'user'")
			}),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
	}
}
