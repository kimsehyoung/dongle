// Code generated by ent, DO NOT EDIT.

package subtitlegen

import (
	"time"

	"github.com/kimsehyoung/dongle/app/subtitle/ent/schema"
	"github.com/kimsehyoung/dongle/app/subtitle/ent/subtitlegen/subtitle"
	"github.com/kimsehyoung/dongle/app/subtitle/ent/subtitlegen/video"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	subtitleFields := schema.Subtitle{}.Fields()
	_ = subtitleFields
	// subtitleDescCreatedAt is the schema descriptor for created_at field.
	subtitleDescCreatedAt := subtitleFields[3].Descriptor()
	// subtitle.DefaultCreatedAt holds the default value on creation for the created_at field.
	subtitle.DefaultCreatedAt = subtitleDescCreatedAt.Default.(time.Time)
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescCreatedAt is the schema descriptor for created_at field.
	videoDescCreatedAt := videoFields[3].Descriptor()
	// video.DefaultCreatedAt holds the default value on creation for the created_at field.
	video.DefaultCreatedAt = videoDescCreatedAt.Default.(time.Time)
}
