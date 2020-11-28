// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// TextEntity represents TL type `textEntity#8bab99a8`.
type TextEntity struct {
	// Offset of the entity, in UTF-16 code units
	Offset int32
	// Length of the entity, in UTF-16 code units
	Length int32
	// Type of the entity
	Type TextEntityTypeClass
}

// TextEntityTypeID is TL type id of TextEntity.
const TextEntityTypeID = 0x8bab99a8

// Encode implements bin.Encoder.
func (t *TextEntity) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEntity#8bab99a8 as nil")
	}
	b.PutID(TextEntityTypeID)
	b.PutInt32(t.Offset)
	b.PutInt32(t.Length)
	if t.Type == nil {
		return fmt.Errorf("unable to encode textEntity#8bab99a8: field type is nil")
	}
	if err := t.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textEntity#8bab99a8: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextEntity) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEntity#8bab99a8 to nil")
	}
	if err := b.ConsumeID(TextEntityTypeID); err != nil {
		return fmt.Errorf("unable to decode textEntity#8bab99a8: %w", err)
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode textEntity#8bab99a8: field offset: %w", err)
		}
		t.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode textEntity#8bab99a8: field length: %w", err)
		}
		t.Length = value
	}
	{
		value, err := DecodeTextEntityType(b)
		if err != nil {
			return fmt.Errorf("unable to decode textEntity#8bab99a8: field type: %w", err)
		}
		t.Type = value
	}
	return nil
}

// Ensuring interfaces in compile-time for TextEntity.
var (
	_ bin.Encoder = &TextEntity{}
	_ bin.Decoder = &TextEntity{}
)
