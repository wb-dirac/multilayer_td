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

// NetworkTypeNone represents TL type `networkTypeNone#8a7a5f11`.
type NetworkTypeNone struct {
}

// NetworkTypeNoneTypeID is TL type id of NetworkTypeNone.
const NetworkTypeNoneTypeID = 0x8a7a5f11

// construct implements constructor of NetworkTypeClass.
func (n NetworkTypeNone) construct() NetworkTypeClass { return &n }

// Ensuring interfaces in compile-time for NetworkTypeNone.
var (
	_ bin.Encoder     = &NetworkTypeNone{}
	_ bin.Decoder     = &NetworkTypeNone{}
	_ bin.BareEncoder = &NetworkTypeNone{}
	_ bin.BareDecoder = &NetworkTypeNone{}

	_ NetworkTypeClass = &NetworkTypeNone{}
)

func (n *NetworkTypeNone) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkTypeNone) String() string {
	if n == nil {
		return "NetworkTypeNone(nil)"
	}
	type Alias NetworkTypeNone
	return fmt.Sprintf("NetworkTypeNone%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkTypeNone) TypeID() uint32 {
	return NetworkTypeNoneTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkTypeNone) TypeName() string {
	return "networkTypeNone"
}

// TypeInfo returns info about TL type.
func (n *NetworkTypeNone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkTypeNone",
		ID:   NetworkTypeNoneTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkTypeNone) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeNone#8a7a5f11 as nil")
	}
	b.PutID(NetworkTypeNoneTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkTypeNone) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeNone#8a7a5f11 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkTypeNone) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeNone#8a7a5f11 to nil")
	}
	if err := b.ConsumeID(NetworkTypeNoneTypeID); err != nil {
		return fmt.Errorf("unable to decode networkTypeNone#8a7a5f11: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkTypeNone) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeNone#8a7a5f11 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkTypeNone) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeNone#8a7a5f11 as nil")
	}
	b.ObjStart()
	b.PutID("networkTypeNone")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkTypeNone) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeNone#8a7a5f11 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkTypeNone"); err != nil {
				return fmt.Errorf("unable to decode networkTypeNone#8a7a5f11: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// NetworkTypeMobile represents TL type `networkTypeMobile#30d46e4f`.
type NetworkTypeMobile struct {
}

// NetworkTypeMobileTypeID is TL type id of NetworkTypeMobile.
const NetworkTypeMobileTypeID = 0x30d46e4f

// construct implements constructor of NetworkTypeClass.
func (n NetworkTypeMobile) construct() NetworkTypeClass { return &n }

// Ensuring interfaces in compile-time for NetworkTypeMobile.
var (
	_ bin.Encoder     = &NetworkTypeMobile{}
	_ bin.Decoder     = &NetworkTypeMobile{}
	_ bin.BareEncoder = &NetworkTypeMobile{}
	_ bin.BareDecoder = &NetworkTypeMobile{}

	_ NetworkTypeClass = &NetworkTypeMobile{}
)

func (n *NetworkTypeMobile) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkTypeMobile) String() string {
	if n == nil {
		return "NetworkTypeMobile(nil)"
	}
	type Alias NetworkTypeMobile
	return fmt.Sprintf("NetworkTypeMobile%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkTypeMobile) TypeID() uint32 {
	return NetworkTypeMobileTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkTypeMobile) TypeName() string {
	return "networkTypeMobile"
}

// TypeInfo returns info about TL type.
func (n *NetworkTypeMobile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkTypeMobile",
		ID:   NetworkTypeMobileTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkTypeMobile) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeMobile#30d46e4f as nil")
	}
	b.PutID(NetworkTypeMobileTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkTypeMobile) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeMobile#30d46e4f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkTypeMobile) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeMobile#30d46e4f to nil")
	}
	if err := b.ConsumeID(NetworkTypeMobileTypeID); err != nil {
		return fmt.Errorf("unable to decode networkTypeMobile#30d46e4f: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkTypeMobile) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeMobile#30d46e4f to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkTypeMobile) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeMobile#30d46e4f as nil")
	}
	b.ObjStart()
	b.PutID("networkTypeMobile")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkTypeMobile) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeMobile#30d46e4f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkTypeMobile"); err != nil {
				return fmt.Errorf("unable to decode networkTypeMobile#30d46e4f: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// NetworkTypeMobileRoaming represents TL type `networkTypeMobileRoaming#aa7496f0`.
type NetworkTypeMobileRoaming struct {
}

// NetworkTypeMobileRoamingTypeID is TL type id of NetworkTypeMobileRoaming.
const NetworkTypeMobileRoamingTypeID = 0xaa7496f0

// construct implements constructor of NetworkTypeClass.
func (n NetworkTypeMobileRoaming) construct() NetworkTypeClass { return &n }

// Ensuring interfaces in compile-time for NetworkTypeMobileRoaming.
var (
	_ bin.Encoder     = &NetworkTypeMobileRoaming{}
	_ bin.Decoder     = &NetworkTypeMobileRoaming{}
	_ bin.BareEncoder = &NetworkTypeMobileRoaming{}
	_ bin.BareDecoder = &NetworkTypeMobileRoaming{}

	_ NetworkTypeClass = &NetworkTypeMobileRoaming{}
)

func (n *NetworkTypeMobileRoaming) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkTypeMobileRoaming) String() string {
	if n == nil {
		return "NetworkTypeMobileRoaming(nil)"
	}
	type Alias NetworkTypeMobileRoaming
	return fmt.Sprintf("NetworkTypeMobileRoaming%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkTypeMobileRoaming) TypeID() uint32 {
	return NetworkTypeMobileRoamingTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkTypeMobileRoaming) TypeName() string {
	return "networkTypeMobileRoaming"
}

// TypeInfo returns info about TL type.
func (n *NetworkTypeMobileRoaming) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkTypeMobileRoaming",
		ID:   NetworkTypeMobileRoamingTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkTypeMobileRoaming) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeMobileRoaming#aa7496f0 as nil")
	}
	b.PutID(NetworkTypeMobileRoamingTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkTypeMobileRoaming) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeMobileRoaming#aa7496f0 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkTypeMobileRoaming) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeMobileRoaming#aa7496f0 to nil")
	}
	if err := b.ConsumeID(NetworkTypeMobileRoamingTypeID); err != nil {
		return fmt.Errorf("unable to decode networkTypeMobileRoaming#aa7496f0: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkTypeMobileRoaming) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeMobileRoaming#aa7496f0 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkTypeMobileRoaming) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeMobileRoaming#aa7496f0 as nil")
	}
	b.ObjStart()
	b.PutID("networkTypeMobileRoaming")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkTypeMobileRoaming) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeMobileRoaming#aa7496f0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkTypeMobileRoaming"); err != nil {
				return fmt.Errorf("unable to decode networkTypeMobileRoaming#aa7496f0: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// NetworkTypeWiFi represents TL type `networkTypeWiFi#da37e13a`.
type NetworkTypeWiFi struct {
}

// NetworkTypeWiFiTypeID is TL type id of NetworkTypeWiFi.
const NetworkTypeWiFiTypeID = 0xda37e13a

// construct implements constructor of NetworkTypeClass.
func (n NetworkTypeWiFi) construct() NetworkTypeClass { return &n }

// Ensuring interfaces in compile-time for NetworkTypeWiFi.
var (
	_ bin.Encoder     = &NetworkTypeWiFi{}
	_ bin.Decoder     = &NetworkTypeWiFi{}
	_ bin.BareEncoder = &NetworkTypeWiFi{}
	_ bin.BareDecoder = &NetworkTypeWiFi{}

	_ NetworkTypeClass = &NetworkTypeWiFi{}
)

func (n *NetworkTypeWiFi) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkTypeWiFi) String() string {
	if n == nil {
		return "NetworkTypeWiFi(nil)"
	}
	type Alias NetworkTypeWiFi
	return fmt.Sprintf("NetworkTypeWiFi%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkTypeWiFi) TypeID() uint32 {
	return NetworkTypeWiFiTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkTypeWiFi) TypeName() string {
	return "networkTypeWiFi"
}

// TypeInfo returns info about TL type.
func (n *NetworkTypeWiFi) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkTypeWiFi",
		ID:   NetworkTypeWiFiTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkTypeWiFi) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeWiFi#da37e13a as nil")
	}
	b.PutID(NetworkTypeWiFiTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkTypeWiFi) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeWiFi#da37e13a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkTypeWiFi) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeWiFi#da37e13a to nil")
	}
	if err := b.ConsumeID(NetworkTypeWiFiTypeID); err != nil {
		return fmt.Errorf("unable to decode networkTypeWiFi#da37e13a: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkTypeWiFi) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeWiFi#da37e13a to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkTypeWiFi) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeWiFi#da37e13a as nil")
	}
	b.ObjStart()
	b.PutID("networkTypeWiFi")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkTypeWiFi) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeWiFi#da37e13a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkTypeWiFi"); err != nil {
				return fmt.Errorf("unable to decode networkTypeWiFi#da37e13a: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// NetworkTypeOther represents TL type `networkTypeOther#73c2879b`.
type NetworkTypeOther struct {
}

// NetworkTypeOtherTypeID is TL type id of NetworkTypeOther.
const NetworkTypeOtherTypeID = 0x73c2879b

// construct implements constructor of NetworkTypeClass.
func (n NetworkTypeOther) construct() NetworkTypeClass { return &n }

// Ensuring interfaces in compile-time for NetworkTypeOther.
var (
	_ bin.Encoder     = &NetworkTypeOther{}
	_ bin.Decoder     = &NetworkTypeOther{}
	_ bin.BareEncoder = &NetworkTypeOther{}
	_ bin.BareDecoder = &NetworkTypeOther{}

	_ NetworkTypeClass = &NetworkTypeOther{}
)

func (n *NetworkTypeOther) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkTypeOther) String() string {
	if n == nil {
		return "NetworkTypeOther(nil)"
	}
	type Alias NetworkTypeOther
	return fmt.Sprintf("NetworkTypeOther%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkTypeOther) TypeID() uint32 {
	return NetworkTypeOtherTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkTypeOther) TypeName() string {
	return "networkTypeOther"
}

// TypeInfo returns info about TL type.
func (n *NetworkTypeOther) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkTypeOther",
		ID:   NetworkTypeOtherTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkTypeOther) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeOther#73c2879b as nil")
	}
	b.PutID(NetworkTypeOtherTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkTypeOther) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeOther#73c2879b as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkTypeOther) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeOther#73c2879b to nil")
	}
	if err := b.ConsumeID(NetworkTypeOtherTypeID); err != nil {
		return fmt.Errorf("unable to decode networkTypeOther#73c2879b: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkTypeOther) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeOther#73c2879b to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *NetworkTypeOther) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkTypeOther#73c2879b as nil")
	}
	b.ObjStart()
	b.PutID("networkTypeOther")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *NetworkTypeOther) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode networkTypeOther#73c2879b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("networkTypeOther"); err != nil {
				return fmt.Errorf("unable to decode networkTypeOther#73c2879b: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// NetworkTypeClass represents NetworkType generic type.
//
// Example:
//  g, err := tdapi.DecodeNetworkType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.NetworkTypeNone: // networkTypeNone#8a7a5f11
//  case *tdapi.NetworkTypeMobile: // networkTypeMobile#30d46e4f
//  case *tdapi.NetworkTypeMobileRoaming: // networkTypeMobileRoaming#aa7496f0
//  case *tdapi.NetworkTypeWiFi: // networkTypeWiFi#da37e13a
//  case *tdapi.NetworkTypeOther: // networkTypeOther#73c2879b
//  default: panic(v)
//  }
type NetworkTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() NetworkTypeClass

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

// DecodeNetworkType implements binary de-serialization for NetworkTypeClass.
func DecodeNetworkType(buf *bin.Buffer) (NetworkTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case NetworkTypeNoneTypeID:
		// Decoding networkTypeNone#8a7a5f11.
		v := NetworkTypeNone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case NetworkTypeMobileTypeID:
		// Decoding networkTypeMobile#30d46e4f.
		v := NetworkTypeMobile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case NetworkTypeMobileRoamingTypeID:
		// Decoding networkTypeMobileRoaming#aa7496f0.
		v := NetworkTypeMobileRoaming{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case NetworkTypeWiFiTypeID:
		// Decoding networkTypeWiFi#da37e13a.
		v := NetworkTypeWiFi{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case NetworkTypeOtherTypeID:
		// Decoding networkTypeOther#73c2879b.
		v := NetworkTypeOther{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONNetworkType implements binary de-serialization for NetworkTypeClass.
func DecodeTDLibJSONNetworkType(buf tdjson.Decoder) (NetworkTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "networkTypeNone":
		// Decoding networkTypeNone#8a7a5f11.
		v := NetworkTypeNone{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case "networkTypeMobile":
		// Decoding networkTypeMobile#30d46e4f.
		v := NetworkTypeMobile{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case "networkTypeMobileRoaming":
		// Decoding networkTypeMobileRoaming#aa7496f0.
		v := NetworkTypeMobileRoaming{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case "networkTypeWiFi":
		// Decoding networkTypeWiFi#da37e13a.
		v := NetworkTypeWiFi{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	case "networkTypeOther":
		// Decoding networkTypeOther#73c2879b.
		v := NetworkTypeOther{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NetworkTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// NetworkType boxes the NetworkTypeClass providing a helper.
type NetworkTypeBox struct {
	NetworkType NetworkTypeClass
}

// Decode implements bin.Decoder for NetworkTypeBox.
func (b *NetworkTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode NetworkTypeBox to nil")
	}
	v, err := DecodeNetworkType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NetworkType = v
	return nil
}

// Encode implements bin.Encode for NetworkTypeBox.
func (b *NetworkTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.NetworkType == nil {
		return fmt.Errorf("unable to encode NetworkTypeClass as nil")
	}
	return b.NetworkType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for NetworkTypeBox.
func (b *NetworkTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode NetworkTypeBox to nil")
	}
	v, err := DecodeTDLibJSONNetworkType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NetworkType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for NetworkTypeBox.
func (b *NetworkTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.NetworkType == nil {
		return fmt.Errorf("unable to encode NetworkTypeClass as nil")
	}
	return b.NetworkType.EncodeTDLibJSON(buf)
}
