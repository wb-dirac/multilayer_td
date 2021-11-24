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

// SetLogStreamRequest represents TL type `setLogStream#aeaff791`.
type SetLogStreamRequest struct {
	// New log stream
	LogStream LogStreamClass
}

// SetLogStreamRequestTypeID is TL type id of SetLogStreamRequest.
const SetLogStreamRequestTypeID = 0xaeaff791

// Ensuring interfaces in compile-time for SetLogStreamRequest.
var (
	_ bin.Encoder     = &SetLogStreamRequest{}
	_ bin.Decoder     = &SetLogStreamRequest{}
	_ bin.BareEncoder = &SetLogStreamRequest{}
	_ bin.BareDecoder = &SetLogStreamRequest{}
)

func (s *SetLogStreamRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.LogStream == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetLogStreamRequest) String() string {
	if s == nil {
		return "SetLogStreamRequest(nil)"
	}
	type Alias SetLogStreamRequest
	return fmt.Sprintf("SetLogStreamRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetLogStreamRequest) TypeID() uint32 {
	return SetLogStreamRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetLogStreamRequest) TypeName() string {
	return "setLogStream"
}

// TypeInfo returns info about TL type.
func (s *SetLogStreamRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setLogStream",
		ID:   SetLogStreamRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LogStream",
			SchemaName: "log_stream",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetLogStreamRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setLogStream#aeaff791 as nil")
	}
	b.PutID(SetLogStreamRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetLogStreamRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setLogStream#aeaff791 as nil")
	}
	if s.LogStream == nil {
		return fmt.Errorf("unable to encode setLogStream#aeaff791: field log_stream is nil")
	}
	if err := s.LogStream.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setLogStream#aeaff791: field log_stream: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetLogStreamRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setLogStream#aeaff791 to nil")
	}
	if err := b.ConsumeID(SetLogStreamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setLogStream#aeaff791: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetLogStreamRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setLogStream#aeaff791 to nil")
	}
	{
		value, err := DecodeLogStream(b)
		if err != nil {
			return fmt.Errorf("unable to decode setLogStream#aeaff791: field log_stream: %w", err)
		}
		s.LogStream = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetLogStreamRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setLogStream#aeaff791 as nil")
	}
	b.ObjStart()
	b.PutID("setLogStream")
	b.FieldStart("log_stream")
	if s.LogStream == nil {
		return fmt.Errorf("unable to encode setLogStream#aeaff791: field log_stream is nil")
	}
	if err := s.LogStream.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setLogStream#aeaff791: field log_stream: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetLogStreamRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setLogStream#aeaff791 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setLogStream"); err != nil {
				return fmt.Errorf("unable to decode setLogStream#aeaff791: %w", err)
			}
		case "log_stream":
			value, err := DecodeTDLibJSONLogStream(b)
			if err != nil {
				return fmt.Errorf("unable to decode setLogStream#aeaff791: field log_stream: %w", err)
			}
			s.LogStream = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLogStream returns value of LogStream field.
func (s *SetLogStreamRequest) GetLogStream() (value LogStreamClass) {
	return s.LogStream
}

// SetLogStream invokes method setLogStream#aeaff791 returning error if any.
func (c *Client) SetLogStream(ctx context.Context, logstream LogStreamClass) error {
	var ok Ok

	request := &SetLogStreamRequest{
		LogStream: logstream,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
