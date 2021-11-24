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

// SearchInstalledStickerSetsRequest represents TL type `searchInstalledStickerSets#2899d990`.
type SearchInstalledStickerSetsRequest struct {
	// Pass true to return mask sticker sets; pass false to return ordinary sticker sets
	IsMasks bool
	// Query to search for
	Query string
	// The maximum number of sticker sets to return
	Limit int32
}

// SearchInstalledStickerSetsRequestTypeID is TL type id of SearchInstalledStickerSetsRequest.
const SearchInstalledStickerSetsRequestTypeID = 0x2899d990

// Ensuring interfaces in compile-time for SearchInstalledStickerSetsRequest.
var (
	_ bin.Encoder     = &SearchInstalledStickerSetsRequest{}
	_ bin.Decoder     = &SearchInstalledStickerSetsRequest{}
	_ bin.BareEncoder = &SearchInstalledStickerSetsRequest{}
	_ bin.BareDecoder = &SearchInstalledStickerSetsRequest{}
)

func (s *SearchInstalledStickerSetsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.IsMasks == false) {
		return false
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchInstalledStickerSetsRequest) String() string {
	if s == nil {
		return "SearchInstalledStickerSetsRequest(nil)"
	}
	type Alias SearchInstalledStickerSetsRequest
	return fmt.Sprintf("SearchInstalledStickerSetsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchInstalledStickerSetsRequest) TypeID() uint32 {
	return SearchInstalledStickerSetsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchInstalledStickerSetsRequest) TypeName() string {
	return "searchInstalledStickerSets"
}

// TypeInfo returns info about TL type.
func (s *SearchInstalledStickerSetsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchInstalledStickerSets",
		ID:   SearchInstalledStickerSetsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsMasks",
			SchemaName: "is_masks",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchInstalledStickerSetsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchInstalledStickerSets#2899d990 as nil")
	}
	b.PutID(SearchInstalledStickerSetsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchInstalledStickerSetsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchInstalledStickerSets#2899d990 as nil")
	}
	b.PutBool(s.IsMasks)
	b.PutString(s.Query)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchInstalledStickerSetsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchInstalledStickerSets#2899d990 to nil")
	}
	if err := b.ConsumeID(SearchInstalledStickerSetsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchInstalledStickerSetsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchInstalledStickerSets#2899d990 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: field is_masks: %w", err)
		}
		s.IsMasks = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchInstalledStickerSetsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchInstalledStickerSets#2899d990 as nil")
	}
	b.ObjStart()
	b.PutID("searchInstalledStickerSets")
	b.FieldStart("is_masks")
	b.PutBool(s.IsMasks)
	b.FieldStart("query")
	b.PutString(s.Query)
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchInstalledStickerSetsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchInstalledStickerSets#2899d990 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchInstalledStickerSets"); err != nil {
				return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: %w", err)
			}
		case "is_masks":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: field is_masks: %w", err)
			}
			s.IsMasks = value
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: field query: %w", err)
			}
			s.Query = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchInstalledStickerSets#2899d990: field limit: %w", err)
			}
			s.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsMasks returns value of IsMasks field.
func (s *SearchInstalledStickerSetsRequest) GetIsMasks() (value bool) {
	return s.IsMasks
}

// GetQuery returns value of Query field.
func (s *SearchInstalledStickerSetsRequest) GetQuery() (value string) {
	return s.Query
}

// GetLimit returns value of Limit field.
func (s *SearchInstalledStickerSetsRequest) GetLimit() (value int32) {
	return s.Limit
}

// SearchInstalledStickerSets invokes method searchInstalledStickerSets#2899d990 returning error if any.
func (c *Client) SearchInstalledStickerSets(ctx context.Context, request *SearchInstalledStickerSetsRequest) (*StickerSets, error) {
	var result StickerSets

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
