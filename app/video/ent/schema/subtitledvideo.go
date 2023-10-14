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

// SubtitledVideo holds the schema definition for the SubtitledVideo entity.
type SubtitledVideo struct {
	ent.Schema
}

func (SubtitledVideo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "subtitled_video"},
	}
}

// Fields of the SubtitledVideo.
func (SubtitledVideo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("original_video_id"),
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

// Edges of the SubtitledVideo.
func (SubtitledVideo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("original_video", OriginalVideo.Type).
			Ref("subtitled_videos").
			Unique().
			Required().
			Field("original_video_id"),
	}
}
