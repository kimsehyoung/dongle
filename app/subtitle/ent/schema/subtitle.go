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

// Subtitle holds the schema definition for the Subtitle entity.
type Subtitle struct {
	ent.Schema
}

func (Subtitle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "subtitle"},
	}
}

// Fields of the Subtitle.
func (Subtitle) Fields() []ent.Field {
	return []ent.Field{
		field.Int("video_id"),
		field.String("language").
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(16)",
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

// Edges of the Subtitle.
func (Subtitle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Video.Type).
			Ref("subtitles").
			Unique().
			Required().
			Field("video_id"),
	}
}
