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

// ChangeStickerSetRequest represents TL type `changeStickerSet#1ac8a5ed`.
type ChangeStickerSetRequest struct {
	// Identifier of the sticker set
	SetID int64
	// The new value of is_installed
	IsInstalled bool
	// The new value of is_archived. A sticker set can't be installed and archived
	// simultaneously
	IsArchived bool
}

// ChangeStickerSetRequestTypeID is TL type id of ChangeStickerSetRequest.
const ChangeStickerSetRequestTypeID = 0x1ac8a5ed

// Ensuring interfaces in compile-time for ChangeStickerSetRequest.
var (
	_ bin.Encoder     = &ChangeStickerSetRequest{}
	_ bin.Decoder     = &ChangeStickerSetRequest{}
	_ bin.BareEncoder = &ChangeStickerSetRequest{}
	_ bin.BareDecoder = &ChangeStickerSetRequest{}
)

func (c *ChangeStickerSetRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.SetID == 0) {
		return false
	}
	if !(c.IsInstalled == false) {
		return false
	}
	if !(c.IsArchived == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChangeStickerSetRequest) String() string {
	if c == nil {
		return "ChangeStickerSetRequest(nil)"
	}
	type Alias ChangeStickerSetRequest
	return fmt.Sprintf("ChangeStickerSetRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChangeStickerSetRequest) TypeID() uint32 {
	return ChangeStickerSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChangeStickerSetRequest) TypeName() string {
	return "changeStickerSet"
}

// TypeInfo returns info about TL type.
func (c *ChangeStickerSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "changeStickerSet",
		ID:   ChangeStickerSetRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SetID",
			SchemaName: "set_id",
		},
		{
			Name:       "IsInstalled",
			SchemaName: "is_installed",
		},
		{
			Name:       "IsArchived",
			SchemaName: "is_archived",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChangeStickerSetRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode changeStickerSet#1ac8a5ed as nil")
	}
	b.PutID(ChangeStickerSetRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChangeStickerSetRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode changeStickerSet#1ac8a5ed as nil")
	}
	b.PutLong(c.SetID)
	b.PutBool(c.IsInstalled)
	b.PutBool(c.IsArchived)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChangeStickerSetRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode changeStickerSet#1ac8a5ed to nil")
	}
	if err := b.ConsumeID(ChangeStickerSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChangeStickerSetRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode changeStickerSet#1ac8a5ed to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: field set_id: %w", err)
		}
		c.SetID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: field is_installed: %w", err)
		}
		c.IsInstalled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: field is_archived: %w", err)
		}
		c.IsArchived = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChangeStickerSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode changeStickerSet#1ac8a5ed as nil")
	}
	b.ObjStart()
	b.PutID("changeStickerSet")
	b.FieldStart("set_id")
	b.PutLong(c.SetID)
	b.FieldStart("is_installed")
	b.PutBool(c.IsInstalled)
	b.FieldStart("is_archived")
	b.PutBool(c.IsArchived)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChangeStickerSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode changeStickerSet#1ac8a5ed to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("changeStickerSet"); err != nil {
				return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: %w", err)
			}
		case "set_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: field set_id: %w", err)
			}
			c.SetID = value
		case "is_installed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: field is_installed: %w", err)
			}
			c.IsInstalled = value
		case "is_archived":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode changeStickerSet#1ac8a5ed: field is_archived: %w", err)
			}
			c.IsArchived = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSetID returns value of SetID field.
func (c *ChangeStickerSetRequest) GetSetID() (value int64) {
	return c.SetID
}

// GetIsInstalled returns value of IsInstalled field.
func (c *ChangeStickerSetRequest) GetIsInstalled() (value bool) {
	return c.IsInstalled
}

// GetIsArchived returns value of IsArchived field.
func (c *ChangeStickerSetRequest) GetIsArchived() (value bool) {
	return c.IsArchived
}

// ChangeStickerSet invokes method changeStickerSet#1ac8a5ed returning error if any.
func (c *Client) ChangeStickerSet(ctx context.Context, request *ChangeStickerSetRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
