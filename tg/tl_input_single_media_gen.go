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

// InputSingleMedia represents TL type `inputSingleMedia#1cc6e91f`.
// A single media in an album or grouped media¹ sent with messages.sendMultiMedia².
//
// Links:
//  1) https://core.telegram.org/api/files#albums-grouped-media
//  2) https://core.telegram.org/method/messages.sendMultiMedia
//
// See https://core.telegram.org/constructor/inputSingleMedia for reference.
type InputSingleMedia struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The media
	Media InputMediaClass
	// Unique client media ID required to prevent message resending
	RandomID int64
	// A caption for the media
	Message string
	// Message entities¹ for styled text
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// InputSingleMediaTypeID is TL type id of InputSingleMedia.
const InputSingleMediaTypeID = 0x1cc6e91f

// Ensuring interfaces in compile-time for InputSingleMedia.
var (
	_ bin.Encoder     = &InputSingleMedia{}
	_ bin.Decoder     = &InputSingleMedia{}
	_ bin.BareEncoder = &InputSingleMedia{}
	_ bin.BareDecoder = &InputSingleMedia{}
)

func (i *InputSingleMedia) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Media == nil) {
		return false
	}
	if !(i.RandomID == 0) {
		return false
	}
	if !(i.Message == "") {
		return false
	}
	if !(i.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputSingleMedia) String() string {
	if i == nil {
		return "InputSingleMedia(nil)"
	}
	type Alias InputSingleMedia
	return fmt.Sprintf("InputSingleMedia%+v", Alias(*i))
}

// FillFrom fills InputSingleMedia from given interface.
func (i *InputSingleMedia) FillFrom(from interface {
	GetMedia() (value InputMediaClass)
	GetRandomID() (value int64)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
}) {
	i.Media = from.GetMedia()
	i.RandomID = from.GetRandomID()
	i.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		i.Entities = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputSingleMedia) TypeID() uint32 {
	return InputSingleMediaTypeID
}

// TypeName returns name of type in TL schema.
func (*InputSingleMedia) TypeName() string {
	return "inputSingleMedia"
}

// TypeInfo returns info about TL type.
func (i *InputSingleMedia) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputSingleMedia",
		ID:   InputSingleMediaTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Media",
			SchemaName: "media",
		},
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !i.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputSingleMedia) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSingleMedia#1cc6e91f as nil")
	}
	b.PutID(InputSingleMediaTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputSingleMedia) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSingleMedia#1cc6e91f as nil")
	}
	if !(i.Entities == nil) {
		i.Flags.Set(0)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSingleMedia#1cc6e91f: field flags: %w", err)
	}
	if i.Media == nil {
		return fmt.Errorf("unable to encode inputSingleMedia#1cc6e91f: field media is nil")
	}
	if err := i.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputSingleMedia#1cc6e91f: field media: %w", err)
	}
	b.PutLong(i.RandomID)
	b.PutString(i.Message)
	if i.Flags.Has(0) {
		b.PutVectorHeader(len(i.Entities))
		for idx, v := range i.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode inputSingleMedia#1cc6e91f: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputSingleMedia#1cc6e91f: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSingleMedia) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSingleMedia#1cc6e91f to nil")
	}
	if err := b.ConsumeID(InputSingleMediaTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputSingleMedia) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSingleMedia#1cc6e91f to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: field media: %w", err)
		}
		i.Media = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: field random_id: %w", err)
		}
		i.RandomID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: field message: %w", err)
		}
		i.Message = value
	}
	if i.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: field entities: %w", err)
		}

		if headerLen > 0 {
			i.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputSingleMedia#1cc6e91f: field entities: %w", err)
			}
			i.Entities = append(i.Entities, value)
		}
	}
	return nil
}

// GetMedia returns value of Media field.
func (i *InputSingleMedia) GetMedia() (value InputMediaClass) {
	return i.Media
}

// GetRandomID returns value of RandomID field.
func (i *InputSingleMedia) GetRandomID() (value int64) {
	return i.RandomID
}

// GetMessage returns value of Message field.
func (i *InputSingleMedia) GetMessage() (value string) {
	return i.Message
}

// SetEntities sets value of Entities conditional field.
func (i *InputSingleMedia) SetEntities(value []MessageEntityClass) {
	i.Flags.Set(0)
	i.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (i *InputSingleMedia) GetEntities() (value []MessageEntityClass, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.Entities, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (i *InputSingleMedia) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return MessageEntityClassArray(i.Entities), true
}
