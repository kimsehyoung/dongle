// Code generated by ent, DO NOT EDIT.

package videodbgen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/originalvideo"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/subtitledvideo"
)

// SubtitledVideo is the model entity for the SubtitledVideo schema.
type SubtitledVideo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OriginalVideoID holds the value of the "original_video_id" field.
	OriginalVideoID int `json:"original_video_id,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SubtitledVideoQuery when eager-loading is set.
	Edges        SubtitledVideoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SubtitledVideoEdges holds the relations/edges for other nodes in the graph.
type SubtitledVideoEdges struct {
	// OriginalVideo holds the value of the original_video edge.
	OriginalVideo *OriginalVideo `json:"original_video,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OriginalVideoOrErr returns the OriginalVideo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SubtitledVideoEdges) OriginalVideoOrErr() (*OriginalVideo, error) {
	if e.loadedTypes[0] {
		if e.OriginalVideo == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: originalvideo.Label}
		}
		return e.OriginalVideo, nil
	}
	return nil, &NotLoadedError{edge: "original_video"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SubtitledVideo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case subtitledvideo.FieldID, subtitledvideo.FieldOriginalVideoID:
			values[i] = new(sql.NullInt64)
		case subtitledvideo.FieldURL:
			values[i] = new(sql.NullString)
		case subtitledvideo.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SubtitledVideo fields.
func (sv *SubtitledVideo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subtitledvideo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sv.ID = int(value.Int64)
		case subtitledvideo.FieldOriginalVideoID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field original_video_id", values[i])
			} else if value.Valid {
				sv.OriginalVideoID = int(value.Int64)
			}
		case subtitledvideo.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				sv.URL = value.String
			}
		case subtitledvideo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sv.CreatedAt = value.Time
			}
		default:
			sv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SubtitledVideo.
// This includes values selected through modifiers, order, etc.
func (sv *SubtitledVideo) Value(name string) (ent.Value, error) {
	return sv.selectValues.Get(name)
}

// QueryOriginalVideo queries the "original_video" edge of the SubtitledVideo entity.
func (sv *SubtitledVideo) QueryOriginalVideo() *OriginalVideoQuery {
	return NewSubtitledVideoClient(sv.config).QueryOriginalVideo(sv)
}

// Update returns a builder for updating this SubtitledVideo.
// Note that you need to call SubtitledVideo.Unwrap() before calling this method if this SubtitledVideo
// was returned from a transaction, and the transaction was committed or rolled back.
func (sv *SubtitledVideo) Update() *SubtitledVideoUpdateOne {
	return NewSubtitledVideoClient(sv.config).UpdateOne(sv)
}

// Unwrap unwraps the SubtitledVideo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sv *SubtitledVideo) Unwrap() *SubtitledVideo {
	_tx, ok := sv.config.driver.(*txDriver)
	if !ok {
		panic("videodbgen: SubtitledVideo is not a transactional entity")
	}
	sv.config.driver = _tx.drv
	return sv
}

// String implements the fmt.Stringer.
func (sv *SubtitledVideo) String() string {
	var builder strings.Builder
	builder.WriteString("SubtitledVideo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sv.ID))
	builder.WriteString("original_video_id=")
	builder.WriteString(fmt.Sprintf("%v", sv.OriginalVideoID))
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(sv.URL)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(sv.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// SubtitledVideos is a parsable slice of SubtitledVideo.
type SubtitledVideos []*SubtitledVideo