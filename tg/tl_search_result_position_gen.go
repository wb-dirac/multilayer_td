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

// SearchResultPosition represents TL type `searchResultPosition#7f648b67`.
//
// See https://core.telegram.org/constructor/searchResultPosition for reference.
type SearchResultPosition struct {
	// MsgID field of SearchResultPosition.
	MsgID int
	// Date field of SearchResultPosition.
	Date int
	// Offset field of SearchResultPosition.
	Offset int
}

// SearchResultPositionTypeID is TL type id of SearchResultPosition.
const SearchResultPositionTypeID = 0x7f648b67

// Ensuring interfaces in compile-time for SearchResultPosition.
var (
	_ bin.Encoder     = &SearchResultPosition{}
	_ bin.Decoder     = &SearchResultPosition{}
	_ bin.BareEncoder = &SearchResultPosition{}
	_ bin.BareDecoder = &SearchResultPosition{}
)

func (s *SearchResultPosition) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.MsgID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.Offset == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchResultPosition) String() string {
	if s == nil {
		return "SearchResultPosition(nil)"
	}
	type Alias SearchResultPosition
	return fmt.Sprintf("SearchResultPosition%+v", Alias(*s))
}

// FillFrom fills SearchResultPosition from given interface.
func (s *SearchResultPosition) FillFrom(from interface {
	GetMsgID() (value int)
	GetDate() (value int)
	GetOffset() (value int)
}) {
	s.MsgID = from.GetMsgID()
	s.Date = from.GetDate()
	s.Offset = from.GetOffset()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchResultPosition) TypeID() uint32 {
	return SearchResultPositionTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchResultPosition) TypeName() string {
	return "searchResultPosition"
}

// TypeInfo returns info about TL type.
func (s *SearchResultPosition) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchResultPosition",
		ID:   SearchResultPositionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchResultPosition) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchResultPosition#7f648b67 as nil")
	}
	b.PutID(SearchResultPositionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchResultPosition) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchResultPosition#7f648b67 as nil")
	}
	b.PutInt(s.MsgID)
	b.PutInt(s.Date)
	b.PutInt(s.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchResultPosition) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchResultPosition#7f648b67 to nil")
	}
	if err := b.ConsumeID(SearchResultPositionTypeID); err != nil {
		return fmt.Errorf("unable to decode searchResultPosition#7f648b67: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchResultPosition) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchResultPosition#7f648b67 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultPosition#7f648b67: field msg_id: %w", err)
		}
		s.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultPosition#7f648b67: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultPosition#7f648b67: field offset: %w", err)
		}
		s.Offset = value
	}
	return nil
}

// GetMsgID returns value of MsgID field.
func (s *SearchResultPosition) GetMsgID() (value int) {
	return s.MsgID
}

// GetDate returns value of Date field.
func (s *SearchResultPosition) GetDate() (value int) {
	return s.Date
}

// GetOffset returns value of Offset field.
func (s *SearchResultPosition) GetOffset() (value int) {
	return s.Offset
}
