// Code generated by ent, DO NOT EDIT.

package ent

import (
	"nevissGo/ent/pixel"
	"nevissGo/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	pixelFields := schema.Pixel{}.Fields()
	_ = pixelFields
	// pixelDescUpdatedAt is the schema descriptor for updated_at field.
	pixelDescUpdatedAt := pixelFields[2].Descriptor()
	// pixel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pixel.DefaultUpdatedAt = pixelDescUpdatedAt.Default.(func() time.Time)
	// pixel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pixel.UpdateDefaultUpdatedAt = pixelDescUpdatedAt.UpdateDefault.(func() time.Time)
}
