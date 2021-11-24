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

// MessagesChats represents TL type `messages.chats#64ff9fd5`.
// List of chats with auxiliary data.
//
// See https://core.telegram.org/constructor/messages.chats for reference.
type MessagesChats struct {
	// List of chats
	Chats []ChatClass
}

// MessagesChatsTypeID is TL type id of MessagesChats.
const MessagesChatsTypeID = 0x64ff9fd5

// construct implements constructor of MessagesChatsClass.
func (c MessagesChats) construct() MessagesChatsClass { return &c }

// Ensuring interfaces in compile-time for MessagesChats.
var (
	_ bin.Encoder     = &MessagesChats{}
	_ bin.Decoder     = &MessagesChats{}
	_ bin.BareEncoder = &MessagesChats{}
	_ bin.BareDecoder = &MessagesChats{}

	_ MessagesChatsClass = &MessagesChats{}
)

func (c *MessagesChats) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesChats) String() string {
	if c == nil {
		return "MessagesChats(nil)"
	}
	type Alias MessagesChats
	return fmt.Sprintf("MessagesChats%+v", Alias(*c))
}

// FillFrom fills MessagesChats from given interface.
func (c *MessagesChats) FillFrom(from interface {
	GetChats() (value []ChatClass)
}) {
	c.Chats = from.GetChats()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesChats) TypeID() uint32 {
	return MessagesChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesChats) TypeName() string {
	return "messages.chats"
}

// TypeInfo returns info about TL type.
func (c *MessagesChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.chats",
		ID:   MessagesChatsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesChats) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chats#64ff9fd5 as nil")
	}
	b.PutID(MessagesChatsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesChats) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chats#64ff9fd5 as nil")
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chats#64ff9fd5: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chats#64ff9fd5: field chats element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesChats) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chats#64ff9fd5 to nil")
	}
	if err := b.ConsumeID(MessagesChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.chats#64ff9fd5: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesChats) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chats#64ff9fd5 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chats#64ff9fd5: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chats#64ff9fd5: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	return nil
}

// GetChats returns value of Chats field.
func (c *MessagesChats) GetChats() (value []ChatClass) {
	return c.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *MessagesChats) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// MessagesChatsSlice represents TL type `messages.chatsSlice#9cd81144`.
// Partial list of chats, more would have to be fetched with pagination¹
//
// Links:
//  1) https://core.telegram.org/api/offsets
//
// See https://core.telegram.org/constructor/messages.chatsSlice for reference.
type MessagesChatsSlice struct {
	// Total number of results that were found server-side (not all are included in chats)
	Count int
	// Chats
	Chats []ChatClass
}

// MessagesChatsSliceTypeID is TL type id of MessagesChatsSlice.
const MessagesChatsSliceTypeID = 0x9cd81144

// construct implements constructor of MessagesChatsClass.
func (c MessagesChatsSlice) construct() MessagesChatsClass { return &c }

// Ensuring interfaces in compile-time for MessagesChatsSlice.
var (
	_ bin.Encoder     = &MessagesChatsSlice{}
	_ bin.Decoder     = &MessagesChatsSlice{}
	_ bin.BareEncoder = &MessagesChatsSlice{}
	_ bin.BareDecoder = &MessagesChatsSlice{}

	_ MessagesChatsClass = &MessagesChatsSlice{}
)

func (c *MessagesChatsSlice) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Count == 0) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesChatsSlice) String() string {
	if c == nil {
		return "MessagesChatsSlice(nil)"
	}
	type Alias MessagesChatsSlice
	return fmt.Sprintf("MessagesChatsSlice%+v", Alias(*c))
}

// FillFrom fills MessagesChatsSlice from given interface.
func (c *MessagesChatsSlice) FillFrom(from interface {
	GetCount() (value int)
	GetChats() (value []ChatClass)
}) {
	c.Count = from.GetCount()
	c.Chats = from.GetChats()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesChatsSlice) TypeID() uint32 {
	return MessagesChatsSliceTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesChatsSlice) TypeName() string {
	return "messages.chatsSlice"
}

// TypeInfo returns info about TL type.
func (c *MessagesChatsSlice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.chatsSlice",
		ID:   MessagesChatsSliceTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesChatsSlice) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatsSlice#9cd81144 as nil")
	}
	b.PutID(MessagesChatsSliceTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesChatsSlice) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatsSlice#9cd81144 as nil")
	}
	b.PutInt(c.Count)
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatsSlice#9cd81144: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatsSlice#9cd81144: field chats element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesChatsSlice) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatsSlice#9cd81144 to nil")
	}
	if err := b.ConsumeID(MessagesChatsSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.chatsSlice#9cd81144: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesChatsSlice) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatsSlice#9cd81144 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatsSlice#9cd81144: field count: %w", err)
		}
		c.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatsSlice#9cd81144: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatsSlice#9cd81144: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (c *MessagesChatsSlice) GetCount() (value int) {
	return c.Count
}

// GetChats returns value of Chats field.
func (c *MessagesChatsSlice) GetChats() (value []ChatClass) {
	return c.Chats
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *MessagesChatsSlice) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// MessagesChatsClass represents messages.Chats generic type.
//
// See https://core.telegram.org/type/messages.Chats for reference.
//
// Example:
//  g, err := tg.DecodeMessagesChats(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesChats: // messages.chats#64ff9fd5
//  case *tg.MessagesChatsSlice: // messages.chatsSlice#9cd81144
//  default: panic(v)
//  }
type MessagesChatsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesChatsClass

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

	// List of chats
	GetChats() (value []ChatClass)
	// List of chats
	MapChats() (value ChatClassArray)
}

// DecodeMessagesChats implements binary de-serialization for MessagesChatsClass.
func DecodeMessagesChats(buf *bin.Buffer) (MessagesChatsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesChatsTypeID:
		// Decoding messages.chats#64ff9fd5.
		v := MessagesChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesChatsClass: %w", err)
		}
		return &v, nil
	case MessagesChatsSliceTypeID:
		// Decoding messages.chatsSlice#9cd81144.
		v := MessagesChatsSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesChatsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesChatsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesChats boxes the MessagesChatsClass providing a helper.
type MessagesChatsBox struct {
	Chats MessagesChatsClass
}

// Decode implements bin.Decoder for MessagesChatsBox.
func (b *MessagesChatsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesChatsBox to nil")
	}
	v, err := DecodeMessagesChats(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Chats = v
	return nil
}

// Encode implements bin.Encode for MessagesChatsBox.
func (b *MessagesChatsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Chats == nil {
		return fmt.Errorf("unable to encode MessagesChatsClass as nil")
	}
	return b.Chats.Encode(buf)
}
