package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

func (Video) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "video"},
	}
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.Int("account_id"),
		field.String("title").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(32)",
			}),
		field.String("url").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(64)",
			}).
			Unique(),
		field.Time("created_at").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Default(time.Now()),
	}
}

func (Video) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "title").
			Unique(),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subtitles", Subtitle.Type),
	}
}
