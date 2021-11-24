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

// BotCommandScopeDefault represents TL type `botCommandScopeDefault#2f6cb2ab`.
//
// See https://core.telegram.org/constructor/botCommandScopeDefault for reference.
type BotCommandScopeDefault struct {
}

// BotCommandScopeDefaultTypeID is TL type id of BotCommandScopeDefault.
const BotCommandScopeDefaultTypeID = 0x2f6cb2ab

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeDefault) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeDefault.
var (
	_ bin.Encoder     = &BotCommandScopeDefault{}
	_ bin.Decoder     = &BotCommandScopeDefault{}
	_ bin.BareEncoder = &BotCommandScopeDefault{}
	_ bin.BareDecoder = &BotCommandScopeDefault{}

	_ BotCommandScopeClass = &BotCommandScopeDefault{}
)

func (b *BotCommandScopeDefault) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeDefault) String() string {
	if b == nil {
		return "BotCommandScopeDefault(nil)"
	}
	type Alias BotCommandScopeDefault
	return fmt.Sprintf("BotCommandScopeDefault%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeDefault) TypeID() uint32 {
	return BotCommandScopeDefaultTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeDefault) TypeName() string {
	return "botCommandScopeDefault"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeDefault) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeDefault",
		ID:   BotCommandScopeDefaultTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeDefault) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeDefault#2f6cb2ab as nil")
	}
	buf.PutID(BotCommandScopeDefaultTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeDefault) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeDefault#2f6cb2ab as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeDefault) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeDefault#2f6cb2ab to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeDefaultTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeDefault#2f6cb2ab: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeDefault) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeDefault#2f6cb2ab to nil")
	}
	return nil
}

// BotCommandScopeUsers represents TL type `botCommandScopeUsers#3c4f04d8`.
//
// See https://core.telegram.org/constructor/botCommandScopeUsers for reference.
type BotCommandScopeUsers struct {
}

// BotCommandScopeUsersTypeID is TL type id of BotCommandScopeUsers.
const BotCommandScopeUsersTypeID = 0x3c4f04d8

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeUsers) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeUsers.
var (
	_ bin.Encoder     = &BotCommandScopeUsers{}
	_ bin.Decoder     = &BotCommandScopeUsers{}
	_ bin.BareEncoder = &BotCommandScopeUsers{}
	_ bin.BareDecoder = &BotCommandScopeUsers{}

	_ BotCommandScopeClass = &BotCommandScopeUsers{}
)

func (b *BotCommandScopeUsers) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeUsers) String() string {
	if b == nil {
		return "BotCommandScopeUsers(nil)"
	}
	type Alias BotCommandScopeUsers
	return fmt.Sprintf("BotCommandScopeUsers%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeUsers) TypeID() uint32 {
	return BotCommandScopeUsersTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeUsers) TypeName() string {
	return "botCommandScopeUsers"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeUsers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeUsers",
		ID:   BotCommandScopeUsersTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeUsers) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeUsers#3c4f04d8 as nil")
	}
	buf.PutID(BotCommandScopeUsersTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeUsers) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeUsers#3c4f04d8 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeUsers) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeUsers#3c4f04d8 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeUsers#3c4f04d8: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeUsers) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeUsers#3c4f04d8 to nil")
	}
	return nil
}

// BotCommandScopeChats represents TL type `botCommandScopeChats#6fe1a881`.
//
// See https://core.telegram.org/constructor/botCommandScopeChats for reference.
type BotCommandScopeChats struct {
}

// BotCommandScopeChatsTypeID is TL type id of BotCommandScopeChats.
const BotCommandScopeChatsTypeID = 0x6fe1a881

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeChats) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeChats.
var (
	_ bin.Encoder     = &BotCommandScopeChats{}
	_ bin.Decoder     = &BotCommandScopeChats{}
	_ bin.BareEncoder = &BotCommandScopeChats{}
	_ bin.BareDecoder = &BotCommandScopeChats{}

	_ BotCommandScopeClass = &BotCommandScopeChats{}
)

func (b *BotCommandScopeChats) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeChats) String() string {
	if b == nil {
		return "BotCommandScopeChats(nil)"
	}
	type Alias BotCommandScopeChats
	return fmt.Sprintf("BotCommandScopeChats%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeChats) TypeID() uint32 {
	return BotCommandScopeChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeChats) TypeName() string {
	return "botCommandScopeChats"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeChats",
		ID:   BotCommandScopeChatsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeChats) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChats#6fe1a881 as nil")
	}
	buf.PutID(BotCommandScopeChatsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeChats) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChats#6fe1a881 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeChats) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChats#6fe1a881 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeChats#6fe1a881: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeChats) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChats#6fe1a881 to nil")
	}
	return nil
}

// BotCommandScopeChatAdmins represents TL type `botCommandScopeChatAdmins#b9aa606a`.
//
// See https://core.telegram.org/constructor/botCommandScopeChatAdmins for reference.
type BotCommandScopeChatAdmins struct {
}

// BotCommandScopeChatAdminsTypeID is TL type id of BotCommandScopeChatAdmins.
const BotCommandScopeChatAdminsTypeID = 0xb9aa606a

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeChatAdmins) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeChatAdmins.
var (
	_ bin.Encoder     = &BotCommandScopeChatAdmins{}
	_ bin.Decoder     = &BotCommandScopeChatAdmins{}
	_ bin.BareEncoder = &BotCommandScopeChatAdmins{}
	_ bin.BareDecoder = &BotCommandScopeChatAdmins{}

	_ BotCommandScopeClass = &BotCommandScopeChatAdmins{}
)

func (b *BotCommandScopeChatAdmins) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeChatAdmins) String() string {
	if b == nil {
		return "BotCommandScopeChatAdmins(nil)"
	}
	type Alias BotCommandScopeChatAdmins
	return fmt.Sprintf("BotCommandScopeChatAdmins%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeChatAdmins) TypeID() uint32 {
	return BotCommandScopeChatAdminsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeChatAdmins) TypeName() string {
	return "botCommandScopeChatAdmins"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeChatAdmins) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeChatAdmins",
		ID:   BotCommandScopeChatAdminsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeChatAdmins) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatAdmins#b9aa606a as nil")
	}
	buf.PutID(BotCommandScopeChatAdminsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeChatAdmins) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatAdmins#b9aa606a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeChatAdmins) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatAdmins#b9aa606a to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeChatAdminsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeChatAdmins#b9aa606a: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeChatAdmins) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatAdmins#b9aa606a to nil")
	}
	return nil
}

// BotCommandScopePeer represents TL type `botCommandScopePeer#db9d897d`.
//
// See https://core.telegram.org/constructor/botCommandScopePeer for reference.
type BotCommandScopePeer struct {
	// Peer field of BotCommandScopePeer.
	Peer InputPeerClass
}

// BotCommandScopePeerTypeID is TL type id of BotCommandScopePeer.
const BotCommandScopePeerTypeID = 0xdb9d897d

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopePeer) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopePeer.
var (
	_ bin.Encoder     = &BotCommandScopePeer{}
	_ bin.Decoder     = &BotCommandScopePeer{}
	_ bin.BareEncoder = &BotCommandScopePeer{}
	_ bin.BareDecoder = &BotCommandScopePeer{}

	_ BotCommandScopeClass = &BotCommandScopePeer{}
)

func (b *BotCommandScopePeer) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopePeer) String() string {
	if b == nil {
		return "BotCommandScopePeer(nil)"
	}
	type Alias BotCommandScopePeer
	return fmt.Sprintf("BotCommandScopePeer%+v", Alias(*b))
}

// FillFrom fills BotCommandScopePeer from given interface.
func (b *BotCommandScopePeer) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	b.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopePeer) TypeID() uint32 {
	return BotCommandScopePeerTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopePeer) TypeName() string {
	return "botCommandScopePeer"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopePeer) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopePeer",
		ID:   BotCommandScopePeerTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopePeer) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopePeer#db9d897d as nil")
	}
	buf.PutID(BotCommandScopePeerTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopePeer) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopePeer#db9d897d as nil")
	}
	if b.Peer == nil {
		return fmt.Errorf("unable to encode botCommandScopePeer#db9d897d: field peer is nil")
	}
	if err := b.Peer.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botCommandScopePeer#db9d897d: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopePeer) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopePeer#db9d897d to nil")
	}
	if err := buf.ConsumeID(BotCommandScopePeerTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopePeer#db9d897d: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopePeer) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopePeer#db9d897d to nil")
	}
	{
		value, err := DecodeInputPeer(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopePeer#db9d897d: field peer: %w", err)
		}
		b.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (b *BotCommandScopePeer) GetPeer() (value InputPeerClass) {
	return b.Peer
}

// BotCommandScopePeerAdmins represents TL type `botCommandScopePeerAdmins#3fd863d1`.
//
// See https://core.telegram.org/constructor/botCommandScopePeerAdmins for reference.
type BotCommandScopePeerAdmins struct {
	// Peer field of BotCommandScopePeerAdmins.
	Peer InputPeerClass
}

// BotCommandScopePeerAdminsTypeID is TL type id of BotCommandScopePeerAdmins.
const BotCommandScopePeerAdminsTypeID = 0x3fd863d1

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopePeerAdmins) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopePeerAdmins.
var (
	_ bin.Encoder     = &BotCommandScopePeerAdmins{}
	_ bin.Decoder     = &BotCommandScopePeerAdmins{}
	_ bin.BareEncoder = &BotCommandScopePeerAdmins{}
	_ bin.BareDecoder = &BotCommandScopePeerAdmins{}

	_ BotCommandScopeClass = &BotCommandScopePeerAdmins{}
)

func (b *BotCommandScopePeerAdmins) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopePeerAdmins) String() string {
	if b == nil {
		return "BotCommandScopePeerAdmins(nil)"
	}
	type Alias BotCommandScopePeerAdmins
	return fmt.Sprintf("BotCommandScopePeerAdmins%+v", Alias(*b))
}

// FillFrom fills BotCommandScopePeerAdmins from given interface.
func (b *BotCommandScopePeerAdmins) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
}) {
	b.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopePeerAdmins) TypeID() uint32 {
	return BotCommandScopePeerAdminsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopePeerAdmins) TypeName() string {
	return "botCommandScopePeerAdmins"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopePeerAdmins) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopePeerAdmins",
		ID:   BotCommandScopePeerAdminsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopePeerAdmins) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopePeerAdmins#3fd863d1 as nil")
	}
	buf.PutID(BotCommandScopePeerAdminsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopePeerAdmins) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopePeerAdmins#3fd863d1 as nil")
	}
	if b.Peer == nil {
		return fmt.Errorf("unable to encode botCommandScopePeerAdmins#3fd863d1: field peer is nil")
	}
	if err := b.Peer.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botCommandScopePeerAdmins#3fd863d1: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopePeerAdmins) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopePeerAdmins#3fd863d1 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopePeerAdminsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopePeerAdmins#3fd863d1: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopePeerAdmins) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopePeerAdmins#3fd863d1 to nil")
	}
	{
		value, err := DecodeInputPeer(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopePeerAdmins#3fd863d1: field peer: %w", err)
		}
		b.Peer = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (b *BotCommandScopePeerAdmins) GetPeer() (value InputPeerClass) {
	return b.Peer
}

// BotCommandScopePeerUser represents TL type `botCommandScopePeerUser#a1321f3`.
//
// See https://core.telegram.org/constructor/botCommandScopePeerUser for reference.
type BotCommandScopePeerUser struct {
	// Peer field of BotCommandScopePeerUser.
	Peer InputPeerClass
	// UserID field of BotCommandScopePeerUser.
	UserID InputUserClass
}

// BotCommandScopePeerUserTypeID is TL type id of BotCommandScopePeerUser.
const BotCommandScopePeerUserTypeID = 0xa1321f3

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopePeerUser) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopePeerUser.
var (
	_ bin.Encoder     = &BotCommandScopePeerUser{}
	_ bin.Decoder     = &BotCommandScopePeerUser{}
	_ bin.BareEncoder = &BotCommandScopePeerUser{}
	_ bin.BareDecoder = &BotCommandScopePeerUser{}

	_ BotCommandScopeClass = &BotCommandScopePeerUser{}
)

func (b *BotCommandScopePeerUser) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Peer == nil) {
		return false
	}
	if !(b.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopePeerUser) String() string {
	if b == nil {
		return "BotCommandScopePeerUser(nil)"
	}
	type Alias BotCommandScopePeerUser
	return fmt.Sprintf("BotCommandScopePeerUser%+v", Alias(*b))
}

// FillFrom fills BotCommandScopePeerUser from given interface.
func (b *BotCommandScopePeerUser) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetUserID() (value InputUserClass)
}) {
	b.Peer = from.GetPeer()
	b.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopePeerUser) TypeID() uint32 {
	return BotCommandScopePeerUserTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopePeerUser) TypeName() string {
	return "botCommandScopePeerUser"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopePeerUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopePeerUser",
		ID:   BotCommandScopePeerUserTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopePeerUser) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopePeerUser#a1321f3 as nil")
	}
	buf.PutID(BotCommandScopePeerUserTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopePeerUser) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopePeerUser#a1321f3 as nil")
	}
	if b.Peer == nil {
		return fmt.Errorf("unable to encode botCommandScopePeerUser#a1321f3: field peer is nil")
	}
	if err := b.Peer.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botCommandScopePeerUser#a1321f3: field peer: %w", err)
	}
	if b.UserID == nil {
		return fmt.Errorf("unable to encode botCommandScopePeerUser#a1321f3: field user_id is nil")
	}
	if err := b.UserID.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode botCommandScopePeerUser#a1321f3: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopePeerUser) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopePeerUser#a1321f3 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopePeerUserTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopePeerUser#a1321f3: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopePeerUser) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopePeerUser#a1321f3 to nil")
	}
	{
		value, err := DecodeInputPeer(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopePeerUser#a1321f3: field peer: %w", err)
		}
		b.Peer = value
	}
	{
		value, err := DecodeInputUser(buf)
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopePeerUser#a1321f3: field user_id: %w", err)
		}
		b.UserID = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (b *BotCommandScopePeerUser) GetPeer() (value InputPeerClass) {
	return b.Peer
}

// GetUserID returns value of UserID field.
func (b *BotCommandScopePeerUser) GetUserID() (value InputUserClass) {
	return b.UserID
}

// BotCommandScopeClass represents BotCommandScope generic type.
//
// See https://core.telegram.org/type/BotCommandScope for reference.
//
// Example:
//  g, err := tg.DecodeBotCommandScope(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.BotCommandScopeDefault: // botCommandScopeDefault#2f6cb2ab
//  case *tg.BotCommandScopeUsers: // botCommandScopeUsers#3c4f04d8
//  case *tg.BotCommandScopeChats: // botCommandScopeChats#6fe1a881
//  case *tg.BotCommandScopeChatAdmins: // botCommandScopeChatAdmins#b9aa606a
//  case *tg.BotCommandScopePeer: // botCommandScopePeer#db9d897d
//  case *tg.BotCommandScopePeerAdmins: // botCommandScopePeerAdmins#3fd863d1
//  case *tg.BotCommandScopePeerUser: // botCommandScopePeerUser#a1321f3
//  default: panic(v)
//  }
type BotCommandScopeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BotCommandScopeClass

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
}

// DecodeBotCommandScope implements binary de-serialization for BotCommandScopeClass.
func DecodeBotCommandScope(buf *bin.Buffer) (BotCommandScopeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BotCommandScopeDefaultTypeID:
		// Decoding botCommandScopeDefault#2f6cb2ab.
		v := BotCommandScopeDefault{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeUsersTypeID:
		// Decoding botCommandScopeUsers#3c4f04d8.
		v := BotCommandScopeUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeChatsTypeID:
		// Decoding botCommandScopeChats#6fe1a881.
		v := BotCommandScopeChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeChatAdminsTypeID:
		// Decoding botCommandScopeChatAdmins#b9aa606a.
		v := BotCommandScopeChatAdmins{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopePeerTypeID:
		// Decoding botCommandScopePeer#db9d897d.
		v := BotCommandScopePeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopePeerAdminsTypeID:
		// Decoding botCommandScopePeerAdmins#3fd863d1.
		v := BotCommandScopePeerAdmins{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopePeerUserTypeID:
		// Decoding botCommandScopePeerUser#a1321f3.
		v := BotCommandScopePeerUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", bin.NewUnexpectedID(id))
	}
}

// BotCommandScope boxes the BotCommandScopeClass providing a helper.
type BotCommandScopeBox struct {
	BotCommandScope BotCommandScopeClass
}

// Decode implements bin.Decoder for BotCommandScopeBox.
func (b *BotCommandScopeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotCommandScopeBox to nil")
	}
	v, err := DecodeBotCommandScope(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotCommandScope = v
	return nil
}

// Encode implements bin.Encode for BotCommandScopeBox.
func (b *BotCommandScopeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BotCommandScope == nil {
		return fmt.Errorf("unable to encode BotCommandScopeClass as nil")
	}
	return b.BotCommandScope.Encode(buf)
}
