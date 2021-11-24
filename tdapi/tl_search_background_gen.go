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

// SearchBackgroundRequest represents TL type `searchBackground#80fb9121`.
type SearchBackgroundRequest struct {
	// The name of the background
	Name string
}

// SearchBackgroundRequestTypeID is TL type id of SearchBackgroundRequest.
const SearchBackgroundRequestTypeID = 0x80fb9121

// Ensuring interfaces in compile-time for SearchBackgroundRequest.
var (
	_ bin.Encoder     = &SearchBackgroundRequest{}
	_ bin.Decoder     = &SearchBackgroundRequest{}
	_ bin.BareEncoder = &SearchBackgroundRequest{}
	_ bin.BareDecoder = &SearchBackgroundRequest{}
)

func (s *SearchBackgroundRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchBackgroundRequest) String() string {
	if s == nil {
		return "SearchBackgroundRequest(nil)"
	}
	type Alias SearchBackgroundRequest
	return fmt.Sprintf("SearchBackgroundRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchBackgroundRequest) TypeID() uint32 {
	return SearchBackgroundRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchBackgroundRequest) TypeName() string {
	return "searchBackground"
}

// TypeInfo returns info about TL type.
func (s *SearchBackgroundRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchBackground",
		ID:   SearchBackgroundRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchBackgroundRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchBackground#80fb9121 as nil")
	}
	b.PutID(SearchBackgroundRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchBackgroundRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchBackground#80fb9121 as nil")
	}
	b.PutString(s.Name)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchBackgroundRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchBackground#80fb9121 to nil")
	}
	if err := b.ConsumeID(SearchBackgroundRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchBackground#80fb9121: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchBackgroundRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchBackground#80fb9121 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchBackground#80fb9121: field name: %w", err)
		}
		s.Name = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchBackgroundRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchBackground#80fb9121 as nil")
	}
	b.ObjStart()
	b.PutID("searchBackground")
	b.FieldStart("name")
	b.PutString(s.Name)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchBackgroundRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchBackground#80fb9121 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchBackground"); err != nil {
				return fmt.Errorf("unable to decode searchBackground#80fb9121: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchBackground#80fb9121: field name: %w", err)
			}
			s.Name = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (s *SearchBackgroundRequest) GetName() (value string) {
	return s.Name
}

// SearchBackground invokes method searchBackground#80fb9121 returning error if any.
func (c *Client) SearchBackground(ctx context.Context, name string) (*Background, error) {
	var result Background

	request := &SearchBackgroundRequest{
		Name: name,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
