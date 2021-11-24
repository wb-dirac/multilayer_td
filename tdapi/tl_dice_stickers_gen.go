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

// DiceStickersRegular represents TL type `diceStickersRegular#d3dfecce`.
type DiceStickersRegular struct {
	// The animated sticker with the dice animation
	Sticker Sticker
}

// DiceStickersRegularTypeID is TL type id of DiceStickersRegular.
const DiceStickersRegularTypeID = 0xd3dfecce

// construct implements constructor of DiceStickersClass.
func (d DiceStickersRegular) construct() DiceStickersClass { return &d }

// Ensuring interfaces in compile-time for DiceStickersRegular.
var (
	_ bin.Encoder     = &DiceStickersRegular{}
	_ bin.Decoder     = &DiceStickersRegular{}
	_ bin.BareEncoder = &DiceStickersRegular{}
	_ bin.BareDecoder = &DiceStickersRegular{}

	_ DiceStickersClass = &DiceStickersRegular{}
)

func (d *DiceStickersRegular) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Sticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DiceStickersRegular) String() string {
	if d == nil {
		return "DiceStickersRegular(nil)"
	}
	type Alias DiceStickersRegular
	return fmt.Sprintf("DiceStickersRegular%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DiceStickersRegular) TypeID() uint32 {
	return DiceStickersRegularTypeID
}

// TypeName returns name of type in TL schema.
func (*DiceStickersRegular) TypeName() string {
	return "diceStickersRegular"
}

// TypeInfo returns info about TL type.
func (d *DiceStickersRegular) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "diceStickersRegular",
		ID:   DiceStickersRegularTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DiceStickersRegular) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode diceStickersRegular#d3dfecce as nil")
	}
	b.PutID(DiceStickersRegularTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DiceStickersRegular) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode diceStickersRegular#d3dfecce as nil")
	}
	if err := d.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersRegular#d3dfecce: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DiceStickersRegular) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode diceStickersRegular#d3dfecce to nil")
	}
	if err := b.ConsumeID(DiceStickersRegularTypeID); err != nil {
		return fmt.Errorf("unable to decode diceStickersRegular#d3dfecce: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DiceStickersRegular) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode diceStickersRegular#d3dfecce to nil")
	}
	{
		if err := d.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode diceStickersRegular#d3dfecce: field sticker: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DiceStickersRegular) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode diceStickersRegular#d3dfecce as nil")
	}
	b.ObjStart()
	b.PutID("diceStickersRegular")
	b.FieldStart("sticker")
	if err := d.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersRegular#d3dfecce: field sticker: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DiceStickersRegular) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode diceStickersRegular#d3dfecce to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("diceStickersRegular"); err != nil {
				return fmt.Errorf("unable to decode diceStickersRegular#d3dfecce: %w", err)
			}
		case "sticker":
			if err := d.Sticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode diceStickersRegular#d3dfecce: field sticker: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSticker returns value of Sticker field.
func (d *DiceStickersRegular) GetSticker() (value Sticker) {
	return d.Sticker
}

// DiceStickersSlotMachine represents TL type `diceStickersSlotMachine#e9a28cac`.
type DiceStickersSlotMachine struct {
	// The animated sticker with the slot machine background. The background animation must
	// start playing after all reel animations finish
	Background Sticker
	// The animated sticker with the lever animation. The lever animation must play once in
	// the initial dice state
	Lever Sticker
	// The animated sticker with the left reel
	LeftReel Sticker
	// The animated sticker with the center reel
	CenterReel Sticker
	// The animated sticker with the right reel
	RightReel Sticker
}

// DiceStickersSlotMachineTypeID is TL type id of DiceStickersSlotMachine.
const DiceStickersSlotMachineTypeID = 0xe9a28cac

// construct implements constructor of DiceStickersClass.
func (d DiceStickersSlotMachine) construct() DiceStickersClass { return &d }

// Ensuring interfaces in compile-time for DiceStickersSlotMachine.
var (
	_ bin.Encoder     = &DiceStickersSlotMachine{}
	_ bin.Decoder     = &DiceStickersSlotMachine{}
	_ bin.BareEncoder = &DiceStickersSlotMachine{}
	_ bin.BareDecoder = &DiceStickersSlotMachine{}

	_ DiceStickersClass = &DiceStickersSlotMachine{}
)

func (d *DiceStickersSlotMachine) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Background.Zero()) {
		return false
	}
	if !(d.Lever.Zero()) {
		return false
	}
	if !(d.LeftReel.Zero()) {
		return false
	}
	if !(d.CenterReel.Zero()) {
		return false
	}
	if !(d.RightReel.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DiceStickersSlotMachine) String() string {
	if d == nil {
		return "DiceStickersSlotMachine(nil)"
	}
	type Alias DiceStickersSlotMachine
	return fmt.Sprintf("DiceStickersSlotMachine%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DiceStickersSlotMachine) TypeID() uint32 {
	return DiceStickersSlotMachineTypeID
}

// TypeName returns name of type in TL schema.
func (*DiceStickersSlotMachine) TypeName() string {
	return "diceStickersSlotMachine"
}

// TypeInfo returns info about TL type.
func (d *DiceStickersSlotMachine) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "diceStickersSlotMachine",
		ID:   DiceStickersSlotMachineTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Background",
			SchemaName: "background",
		},
		{
			Name:       "Lever",
			SchemaName: "lever",
		},
		{
			Name:       "LeftReel",
			SchemaName: "left_reel",
		},
		{
			Name:       "CenterReel",
			SchemaName: "center_reel",
		},
		{
			Name:       "RightReel",
			SchemaName: "right_reel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DiceStickersSlotMachine) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode diceStickersSlotMachine#e9a28cac as nil")
	}
	b.PutID(DiceStickersSlotMachineTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DiceStickersSlotMachine) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode diceStickersSlotMachine#e9a28cac as nil")
	}
	if err := d.Background.Encode(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field background: %w", err)
	}
	if err := d.Lever.Encode(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field lever: %w", err)
	}
	if err := d.LeftReel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field left_reel: %w", err)
	}
	if err := d.CenterReel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field center_reel: %w", err)
	}
	if err := d.RightReel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field right_reel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DiceStickersSlotMachine) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode diceStickersSlotMachine#e9a28cac to nil")
	}
	if err := b.ConsumeID(DiceStickersSlotMachineTypeID); err != nil {
		return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DiceStickersSlotMachine) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode diceStickersSlotMachine#e9a28cac to nil")
	}
	{
		if err := d.Background.Decode(b); err != nil {
			return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field background: %w", err)
		}
	}
	{
		if err := d.Lever.Decode(b); err != nil {
			return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field lever: %w", err)
		}
	}
	{
		if err := d.LeftReel.Decode(b); err != nil {
			return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field left_reel: %w", err)
		}
	}
	{
		if err := d.CenterReel.Decode(b); err != nil {
			return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field center_reel: %w", err)
		}
	}
	{
		if err := d.RightReel.Decode(b); err != nil {
			return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field right_reel: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DiceStickersSlotMachine) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode diceStickersSlotMachine#e9a28cac as nil")
	}
	b.ObjStart()
	b.PutID("diceStickersSlotMachine")
	b.FieldStart("background")
	if err := d.Background.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field background: %w", err)
	}
	b.FieldStart("lever")
	if err := d.Lever.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field lever: %w", err)
	}
	b.FieldStart("left_reel")
	if err := d.LeftReel.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field left_reel: %w", err)
	}
	b.FieldStart("center_reel")
	if err := d.CenterReel.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field center_reel: %w", err)
	}
	b.FieldStart("right_reel")
	if err := d.RightReel.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode diceStickersSlotMachine#e9a28cac: field right_reel: %w", err)
	}
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DiceStickersSlotMachine) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode diceStickersSlotMachine#e9a28cac to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("diceStickersSlotMachine"); err != nil {
				return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: %w", err)
			}
		case "background":
			if err := d.Background.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field background: %w", err)
			}
		case "lever":
			if err := d.Lever.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field lever: %w", err)
			}
		case "left_reel":
			if err := d.LeftReel.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field left_reel: %w", err)
			}
		case "center_reel":
			if err := d.CenterReel.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field center_reel: %w", err)
			}
		case "right_reel":
			if err := d.RightReel.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode diceStickersSlotMachine#e9a28cac: field right_reel: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBackground returns value of Background field.
func (d *DiceStickersSlotMachine) GetBackground() (value Sticker) {
	return d.Background
}

// GetLever returns value of Lever field.
func (d *DiceStickersSlotMachine) GetLever() (value Sticker) {
	return d.Lever
}

// GetLeftReel returns value of LeftReel field.
func (d *DiceStickersSlotMachine) GetLeftReel() (value Sticker) {
	return d.LeftReel
}

// GetCenterReel returns value of CenterReel field.
func (d *DiceStickersSlotMachine) GetCenterReel() (value Sticker) {
	return d.CenterReel
}

// GetRightReel returns value of RightReel field.
func (d *DiceStickersSlotMachine) GetRightReel() (value Sticker) {
	return d.RightReel
}

// DiceStickersClass represents DiceStickers generic type.
//
// Example:
//  g, err := tdapi.DecodeDiceStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.DiceStickersRegular: // diceStickersRegular#d3dfecce
//  case *tdapi.DiceStickersSlotMachine: // diceStickersSlotMachine#e9a28cac
//  default: panic(v)
//  }
type DiceStickersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() DiceStickersClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeDiceStickers implements binary de-serialization for DiceStickersClass.
func DecodeDiceStickers(buf *bin.Buffer) (DiceStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DiceStickersRegularTypeID:
		// Decoding diceStickersRegular#d3dfecce.
		v := DiceStickersRegular{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DiceStickersClass: %w", err)
		}
		return &v, nil
	case DiceStickersSlotMachineTypeID:
		// Decoding diceStickersSlotMachine#e9a28cac.
		v := DiceStickersSlotMachine{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DiceStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DiceStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONDiceStickers implements binary de-serialization for DiceStickersClass.
func DecodeTDLibJSONDiceStickers(buf tdjson.Decoder) (DiceStickersClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "diceStickersRegular":
		// Decoding diceStickersRegular#d3dfecce.
		v := DiceStickersRegular{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DiceStickersClass: %w", err)
		}
		return &v, nil
	case "diceStickersSlotMachine":
		// Decoding diceStickersSlotMachine#e9a28cac.
		v := DiceStickersSlotMachine{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DiceStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DiceStickersClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// DiceStickers boxes the DiceStickersClass providing a helper.
type DiceStickersBox struct {
	DiceStickers DiceStickersClass
}

// Decode implements bin.Decoder for DiceStickersBox.
func (b *DiceStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DiceStickersBox to nil")
	}
	v, err := DecodeDiceStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DiceStickers = v
	return nil
}

// Encode implements bin.Encode for DiceStickersBox.
func (b *DiceStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DiceStickers == nil {
		return fmt.Errorf("unable to encode DiceStickersClass as nil")
	}
	return b.DiceStickers.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for DiceStickersBox.
func (b *DiceStickersBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode DiceStickersBox to nil")
	}
	v, err := DecodeTDLibJSONDiceStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DiceStickers = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for DiceStickersBox.
func (b *DiceStickersBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.DiceStickers == nil {
		return fmt.Errorf("unable to encode DiceStickersClass as nil")
	}
	return b.DiceStickers.EncodeTDLibJSON(buf)
}
