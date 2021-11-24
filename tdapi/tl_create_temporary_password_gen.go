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

// CreateTemporaryPasswordRequest represents TL type `createTemporaryPassword#9f0d6f86`.
type CreateTemporaryPasswordRequest struct {
	// Persistent user password
	Password string
	// Time during which the temporary password will be valid, in seconds; should be between
	// 60 and 86400
	ValidFor int32
}

// CreateTemporaryPasswordRequestTypeID is TL type id of CreateTemporaryPasswordRequest.
const CreateTemporaryPasswordRequestTypeID = 0x9f0d6f86

// Ensuring interfaces in compile-time for CreateTemporaryPasswordRequest.
var (
	_ bin.Encoder     = &CreateTemporaryPasswordRequest{}
	_ bin.Decoder     = &CreateTemporaryPasswordRequest{}
	_ bin.BareEncoder = &CreateTemporaryPasswordRequest{}
	_ bin.BareDecoder = &CreateTemporaryPasswordRequest{}
)

func (c *CreateTemporaryPasswordRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Password == "") {
		return false
	}
	if !(c.ValidFor == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateTemporaryPasswordRequest) String() string {
	if c == nil {
		return "CreateTemporaryPasswordRequest(nil)"
	}
	type Alias CreateTemporaryPasswordRequest
	return fmt.Sprintf("CreateTemporaryPasswordRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateTemporaryPasswordRequest) TypeID() uint32 {
	return CreateTemporaryPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateTemporaryPasswordRequest) TypeName() string {
	return "createTemporaryPassword"
}

// TypeInfo returns info about TL type.
func (c *CreateTemporaryPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createTemporaryPassword",
		ID:   CreateTemporaryPasswordRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Password",
			SchemaName: "password",
		},
		{
			Name:       "ValidFor",
			SchemaName: "valid_for",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateTemporaryPasswordRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createTemporaryPassword#9f0d6f86 as nil")
	}
	b.PutID(CreateTemporaryPasswordRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateTemporaryPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createTemporaryPassword#9f0d6f86 as nil")
	}
	b.PutString(c.Password)
	b.PutInt32(c.ValidFor)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateTemporaryPasswordRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createTemporaryPassword#9f0d6f86 to nil")
	}
	if err := b.ConsumeID(CreateTemporaryPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createTemporaryPassword#9f0d6f86: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateTemporaryPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createTemporaryPassword#9f0d6f86 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode createTemporaryPassword#9f0d6f86: field password: %w", err)
		}
		c.Password = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode createTemporaryPassword#9f0d6f86: field valid_for: %w", err)
		}
		c.ValidFor = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateTemporaryPasswordRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createTemporaryPassword#9f0d6f86 as nil")
	}
	b.ObjStart()
	b.PutID("createTemporaryPassword")
	b.FieldStart("password")
	b.PutString(c.Password)
	b.FieldStart("valid_for")
	b.PutInt32(c.ValidFor)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateTemporaryPasswordRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createTemporaryPassword#9f0d6f86 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createTemporaryPassword"); err != nil {
				return fmt.Errorf("unable to decode createTemporaryPassword#9f0d6f86: %w", err)
			}
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode createTemporaryPassword#9f0d6f86: field password: %w", err)
			}
			c.Password = value
		case "valid_for":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode createTemporaryPassword#9f0d6f86: field valid_for: %w", err)
			}
			c.ValidFor = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPassword returns value of Password field.
func (c *CreateTemporaryPasswordRequest) GetPassword() (value string) {
	return c.Password
}

// GetValidFor returns value of ValidFor field.
func (c *CreateTemporaryPasswordRequest) GetValidFor() (value int32) {
	return c.ValidFor
}

// CreateTemporaryPassword invokes method createTemporaryPassword#9f0d6f86 returning error if any.
func (c *Client) CreateTemporaryPassword(ctx context.Context, request *CreateTemporaryPasswordRequest) (*TemporaryPasswordState, error) {
	var result TemporaryPasswordState

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
