// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// Stickers represents TL type `stickers#83491d00`.
type Stickers struct {
	// List of stickers
	Stickers []Sticker
}

// StickersTypeID is TL type id of Stickers.
const StickersTypeID = 0x83491d00

// Ensuring interfaces in compile-time for Stickers.
var (
	_ bin.Encoder     = &Stickers{}
	_ bin.Decoder     = &Stickers{}
	_ bin.BareEncoder = &Stickers{}
	_ bin.BareDecoder = &Stickers{}
)

func (s *Stickers) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Stickers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Stickers) String() string {
	if s == nil {
		return "Stickers(nil)"
	}
	type Alias Stickers
	return fmt.Sprintf("Stickers%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Stickers) TypeID() uint32 {
	return StickersTypeID
}

// TypeName returns name of type in TL schema.
func (*Stickers) TypeName() string {
	return "stickers"
}

// TypeInfo returns info about TL type.
func (s *Stickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers",
		ID:   StickersTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickers",
			SchemaName: "stickers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Stickers) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers#83491d00 as nil")
	}
	b.PutID(StickersTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Stickers) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers#83491d00 as nil")
	}
	b.PutInt(len(s.Stickers))
	for idx, v := range s.Stickers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare stickers#83491d00: field stickers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *Stickers) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers#83491d00 to nil")
	}
	if err := b.ConsumeID(StickersTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers#83491d00: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Stickers) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers#83491d00 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickers#83491d00: field stickers: %w", err)
		}

		if headerLen > 0 {
			s.Stickers = make([]Sticker, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Sticker
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare stickers#83491d00: field stickers: %w", err)
			}
			s.Stickers = append(s.Stickers, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Stickers) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode stickers#83491d00 as nil")
	}
	b.ObjStart()
	b.PutID("stickers")
	b.FieldStart("stickers")
	b.ArrStart()
	for idx, v := range s.Stickers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode stickers#83491d00: field stickers element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Stickers) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode stickers#83491d00 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("stickers"); err != nil {
				return fmt.Errorf("unable to decode stickers#83491d00: %w", err)
			}
		case "stickers":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Sticker
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode stickers#83491d00: field stickers: %w", err)
				}
				s.Stickers = append(s.Stickers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode stickers#83491d00: field stickers: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStickers returns value of Stickers field.
func (s *Stickers) GetStickers() (value []Sticker) {
	return s.Stickers
}
