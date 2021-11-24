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

// AccountGetWallPaperRequest represents TL type `account.getWallPaper#fc8ddbea`.
// Get info about a certain wallpaper
//
// See https://core.telegram.org/method/account.getWallPaper for reference.
type AccountGetWallPaperRequest struct {
	// The wallpaper to get info about
	Wallpaper InputWallPaperClass
}

// AccountGetWallPaperRequestTypeID is TL type id of AccountGetWallPaperRequest.
const AccountGetWallPaperRequestTypeID = 0xfc8ddbea

// Ensuring interfaces in compile-time for AccountGetWallPaperRequest.
var (
	_ bin.Encoder     = &AccountGetWallPaperRequest{}
	_ bin.Decoder     = &AccountGetWallPaperRequest{}
	_ bin.BareEncoder = &AccountGetWallPaperRequest{}
	_ bin.BareDecoder = &AccountGetWallPaperRequest{}
)

func (g *AccountGetWallPaperRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Wallpaper == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetWallPaperRequest) String() string {
	if g == nil {
		return "AccountGetWallPaperRequest(nil)"
	}
	type Alias AccountGetWallPaperRequest
	return fmt.Sprintf("AccountGetWallPaperRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetWallPaperRequest from given interface.
func (g *AccountGetWallPaperRequest) FillFrom(from interface {
	GetWallpaper() (value InputWallPaperClass)
}) {
	g.Wallpaper = from.GetWallpaper()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetWallPaperRequest) TypeID() uint32 {
	return AccountGetWallPaperRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetWallPaperRequest) TypeName() string {
	return "account.getWallPaper"
}

// TypeInfo returns info about TL type.
func (g *AccountGetWallPaperRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getWallPaper",
		ID:   AccountGetWallPaperRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Wallpaper",
			SchemaName: "wallpaper",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetWallPaperRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWallPaper#fc8ddbea as nil")
	}
	b.PutID(AccountGetWallPaperRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetWallPaperRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWallPaper#fc8ddbea as nil")
	}
	if g.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.getWallPaper#fc8ddbea: field wallpaper is nil")
	}
	if err := g.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getWallPaper#fc8ddbea: field wallpaper: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetWallPaperRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWallPaper#fc8ddbea to nil")
	}
	if err := b.ConsumeID(AccountGetWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWallPaper#fc8ddbea: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetWallPaperRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWallPaper#fc8ddbea to nil")
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getWallPaper#fc8ddbea: field wallpaper: %w", err)
		}
		g.Wallpaper = value
	}
	return nil
}

// GetWallpaper returns value of Wallpaper field.
func (g *AccountGetWallPaperRequest) GetWallpaper() (value InputWallPaperClass) {
	return g.Wallpaper
}

// AccountGetWallPaper invokes method account.getWallPaper#fc8ddbea returning error if any.
// Get info about a certain wallpaper
//
// See https://core.telegram.org/method/account.getWallPaper for reference.
func (c *Client) AccountGetWallPaper(ctx context.Context, wallpaper InputWallPaperClass) (WallPaperClass, error) {
	var result WallPaperBox

	request := &AccountGetWallPaperRequest{
		Wallpaper: wallpaper,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WallPaper, nil
}
