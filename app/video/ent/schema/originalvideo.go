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

// OriginalVideo holds the schema definition for the OriginalVideo entity.
type OriginalVideo struct {
	ent.Schema
}

func (OriginalVideo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "original_video"},
	}
}

// Fields of the OriginalVideo.
func (OriginalVideo) Fields() []ent.Field {
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

func (OriginalVideo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_id", "title").
			Unique(),
	}
}

// Edges of the OriginalVideo.
func (OriginalVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subtitled_videos", SubtitledVideo.Type),
	}
}
