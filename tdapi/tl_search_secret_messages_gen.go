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

// SearchSecretMessagesRequest represents TL type `searchSecretMessages#cd2a4c9c`.
type SearchSecretMessagesRequest struct {
	// Identifier of the chat in which to search. Specify 0 to search in all secret chats
	ChatID int64
	// Query to search for. If empty, searchChatMessages should be used instead
	Query string
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get first chunk of results
	Offset string
	// The maximum number of messages to be returned; up to 100. For optimal performance, the
	// number of returned messages is chosen by TDLib and can be smaller than the specified
	// limit
	Limit int32
	// A filter for message content in the search results
	Filter SearchMessagesFilterClass
}

// SearchSecretMessagesRequestTypeID is TL type id of SearchSecretMessagesRequest.
const SearchSecretMessagesRequestTypeID = 0xcd2a4c9c

// Ensuring interfaces in compile-time for SearchSecretMessagesRequest.
var (
	_ bin.Encoder     = &SearchSecretMessagesRequest{}
	_ bin.Decoder     = &SearchSecretMessagesRequest{}
	_ bin.BareEncoder = &SearchSecretMessagesRequest{}
	_ bin.BareDecoder = &SearchSecretMessagesRequest{}
)

func (s *SearchSecretMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.Offset == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchSecretMessagesRequest) String() string {
	if s == nil {
		return "SearchSecretMessagesRequest(nil)"
	}
	type Alias SearchSecretMessagesRequest
	return fmt.Sprintf("SearchSecretMessagesRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchSecretMessagesRequest) TypeID() uint32 {
	return SearchSecretMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchSecretMessagesRequest) TypeName() string {
	return "searchSecretMessages"
}

// TypeInfo returns info about TL type.
func (s *SearchSecretMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchSecretMessages",
		ID:   SearchSecretMessagesRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchSecretMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchSecretMessages#cd2a4c9c as nil")
	}
	b.PutID(SearchSecretMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchSecretMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchSecretMessages#cd2a4c9c as nil")
	}
	b.PutLong(s.ChatID)
	b.PutString(s.Query)
	b.PutString(s.Offset)
	b.PutInt32(s.Limit)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchSecretMessages#cd2a4c9c: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchSecretMessages#cd2a4c9c: field filter: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchSecretMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchSecretMessages#cd2a4c9c to nil")
	}
	if err := b.ConsumeID(SearchSecretMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchSecretMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchSecretMessages#cd2a4c9c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field offset: %w", err)
		}
		s.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field filter: %w", err)
		}
		s.Filter = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchSecretMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchSecretMessages#cd2a4c9c as nil")
	}
	b.ObjStart()
	b.PutID("searchSecretMessages")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("query")
	b.PutString(s.Query)
	b.FieldStart("offset")
	b.PutString(s.Offset)
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.FieldStart("filter")
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchSecretMessages#cd2a4c9c: field filter is nil")
	}
	if err := s.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchSecretMessages#cd2a4c9c: field filter: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchSecretMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchSecretMessages#cd2a4c9c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchSecretMessages"); err != nil {
				return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: %w", err)
			}
		case "chat_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field chat_id: %w", err)
			}
			s.ChatID = value
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field query: %w", err)
			}
			s.Query = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field offset: %w", err)
			}
			s.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field limit: %w", err)
			}
			s.Limit = value
		case "filter":
			value, err := DecodeTDLibJSONSearchMessagesFilter(b)
			if err != nil {
				return fmt.Errorf("unable to decode searchSecretMessages#cd2a4c9c: field filter: %w", err)
			}
			s.Filter = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SearchSecretMessagesRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetQuery returns value of Query field.
func (s *SearchSecretMessagesRequest) GetQuery() (value string) {
	return s.Query
}

// GetOffset returns value of Offset field.
func (s *SearchSecretMessagesRequest) GetOffset() (value string) {
	return s.Offset
}

// GetLimit returns value of Limit field.
func (s *SearchSecretMessagesRequest) GetLimit() (value int32) {
	return s.Limit
}

// GetFilter returns value of Filter field.
func (s *SearchSecretMessagesRequest) GetFilter() (value SearchMessagesFilterClass) {
	return s.Filter
}

// SearchSecretMessages invokes method searchSecretMessages#cd2a4c9c returning error if any.
func (c *Client) SearchSecretMessages(ctx context.Context, request *SearchSecretMessagesRequest) (*FoundMessages, error) {
	var result FoundMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
