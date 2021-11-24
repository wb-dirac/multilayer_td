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

// MessagesSearchCounterVector is a box for Vector<messages.SearchCounter>
type MessagesSearchCounterVector struct {
	// Elements of Vector<messages.SearchCounter>
	Elems []MessagesSearchCounter
}

// MessagesSearchCounterVectorTypeID is TL type id of MessagesSearchCounterVector.
const MessagesSearchCounterVectorTypeID = bin.TypeVector

// Ensuring interfaces in compile-time for MessagesSearchCounterVector.
var (
	_ bin.Encoder     = &MessagesSearchCounterVector{}
	_ bin.Decoder     = &MessagesSearchCounterVector{}
	_ bin.BareEncoder = &MessagesSearchCounterVector{}
	_ bin.BareDecoder = &MessagesSearchCounterVector{}
)

func (vec *MessagesSearchCounterVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *MessagesSearchCounterVector) String() string {
	if vec == nil {
		return "MessagesSearchCounterVector(nil)"
	}
	type Alias MessagesSearchCounterVector
	return fmt.Sprintf("MessagesSearchCounterVector%+v", Alias(*vec))
}

// FillFrom fills MessagesSearchCounterVector from given interface.
func (vec *MessagesSearchCounterVector) FillFrom(from interface {
	GetElems() (value []MessagesSearchCounter)
}) {
	vec.Elems = from.GetElems()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchCounterVector) TypeID() uint32 {
	return MessagesSearchCounterVectorTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchCounterVector) TypeName() string {
	return ""
}

// TypeInfo returns info about TL type.
func (vec *MessagesSearchCounterVector) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "",
		ID:   MessagesSearchCounterVectorTypeID,
	}
	if vec == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elems",
			SchemaName: "Elems",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (vec *MessagesSearchCounterVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<messages.SearchCounter> as nil")
	}

	return vec.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (vec *MessagesSearchCounterVector) EncodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<messages.SearchCounter> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<messages.SearchCounter>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *MessagesSearchCounterVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<messages.SearchCounter> to nil")
	}

	return vec.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (vec *MessagesSearchCounterVector) DecodeBare(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<messages.SearchCounter> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<messages.SearchCounter>: field Elems: %w", err)
		}

		if headerLen > 0 {
			vec.Elems = make([]MessagesSearchCounter, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessagesSearchCounter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<messages.SearchCounter>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *MessagesSearchCounterVector) GetElems() (value []MessagesSearchCounter) {
	return vec.Elems
}
