// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// VideoSize represents TL type `videoSize#de33b094`.
// Animated profile picture¹ in MPEG4 format
//
// Links:
//  1) https://core.telegram.org/api/files#animated-profile-pictures
//
// See https://core.telegram.org/constructor/videoSize for reference.
type VideoSize struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// u for animated profile pictures, and v for trimmed and downscaled video previews
	Type string
	// Video width
	W int
	// Video height
	H int
	// File size
	Size int
	// Timestamp that should be shown as static preview to the user (seconds)
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// VideoSizeTypeID is TL type id of VideoSize.
const VideoSizeTypeID = 0xde33b094

// Ensuring interfaces in compile-time for VideoSize.
var (
	_ bin.Encoder     = &VideoSize{}
	_ bin.Decoder     = &VideoSize{}
	_ bin.BareEncoder = &VideoSize{}
	_ bin.BareDecoder = &VideoSize{}
)

func (v *VideoSize) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.Type == "") {
		return false
	}
	if !(v.W == 0) {
		return false
	}
	if !(v.H == 0) {
		return false
	}
	if !(v.Size == 0) {
		return false
	}
	if !(v.VideoStartTs == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VideoSize) String() string {
	if v == nil {
		return "VideoSize(nil)"
	}
	type Alias VideoSize
	return fmt.Sprintf("VideoSize%+v", Alias(*v))
}

// FillFrom fills VideoSize from given interface.
func (v *VideoSize) FillFrom(from interface {
	GetType() (value string)
	GetW() (value int)
	GetH() (value int)
	GetSize() (value int)
	GetVideoStartTs() (value float64, ok bool)
}) {
	v.Type = from.GetType()
	v.W = from.GetW()
	v.H = from.GetH()
	v.Size = from.GetSize()
	if val, ok := from.GetVideoStartTs(); ok {
		v.VideoStartTs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VideoSize) TypeID() uint32 {
	return VideoSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*VideoSize) TypeName() string {
	return "videoSize"
}

// TypeInfo returns info about TL type.
func (v *VideoSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "videoSize",
		ID:   VideoSizeTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !v.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VideoSize) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSize#de33b094 as nil")
	}
	b.PutID(VideoSizeTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VideoSize) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSize#de33b094 as nil")
	}
	if !(v.VideoStartTs == 0) {
		v.Flags.Set(0)
	}
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSize#de33b094: field flags: %w", err)
	}
	b.PutString(v.Type)
	b.PutInt(v.W)
	b.PutInt(v.H)
	b.PutInt(v.Size)
	if v.Flags.Has(0) {
		b.PutDouble(v.VideoStartTs)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VideoSize) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSize#de33b094 to nil")
	}
	if err := b.ConsumeID(VideoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode videoSize#de33b094: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VideoSize) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSize#de33b094 to nil")
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field type: %w", err)
		}
		v.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field w: %w", err)
		}
		v.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field h: %w", err)
		}
		v.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field size: %w", err)
		}
		v.Size = value
	}
	if v.Flags.Has(0) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field video_start_ts: %w", err)
		}
		v.VideoStartTs = value
	}
	return nil
}

// GetType returns value of Type field.
func (v *VideoSize) GetType() (value string) {
	return v.Type
}

// GetW returns value of W field.
func (v *VideoSize) GetW() (value int) {
	return v.W
}

// GetH returns value of H field.
func (v *VideoSize) GetH() (value int) {
	return v.H
}

// GetSize returns value of Size field.
func (v *VideoSize) GetSize() (value int) {
	return v.Size
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (v *VideoSize) SetVideoStartTs(value float64) {
	v.Flags.Set(0)
	v.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (v *VideoSize) GetVideoStartTs() (value float64, ok bool) {
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.VideoStartTs, true
}
