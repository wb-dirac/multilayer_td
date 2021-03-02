// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// ResponseID represents TL type `responseID#85d7fd8b`.
//
// See https://localhost:80/doc/constructor/responseID for reference.
type ResponseID struct {
	// ID field of ResponseID.
	ID int32
}

// ResponseIDTypeID is TL type id of ResponseID.
const ResponseIDTypeID = 0x85d7fd8b

func (r *ResponseID) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResponseID) String() string {
	if r == nil {
		return "ResponseID(nil)"
	}
	type Alias ResponseID
	return fmt.Sprintf("ResponseID%+v", Alias(*r))
}

// FillFrom fills ResponseID from given interface.
func (r *ResponseID) FillFrom(from interface {
	GetID() (value int32)
}) {
	r.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResponseID) TypeID() uint32 {
	return ResponseIDTypeID
}

// TypeName returns name of type in TL schema.
func (*ResponseID) TypeName() string {
	return "responseID"
}

// TypeInfo returns info about TL type.
func (r *ResponseID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "responseID",
		ID:   ResponseIDTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResponseID) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode responseID#85d7fd8b as nil")
	}
	b.PutID(ResponseIDTypeID)
	b.PutInt32(r.ID)
	return nil
}

// GetID returns value of ID field.
func (r *ResponseID) GetID() (value int32) {
	return r.ID
}

// Decode implements bin.Decoder.
func (r *ResponseID) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode responseID#85d7fd8b to nil")
	}
	if err := b.ConsumeID(ResponseIDTypeID); err != nil {
		return fmt.Errorf("unable to decode responseID#85d7fd8b: %w", err)
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode responseID#85d7fd8b: field id: %w", err)
		}
		r.ID = value
	}
	return nil
}

// construct implements constructor of ResponseClass.
func (r ResponseID) construct() ResponseClass { return &r }

// Ensuring interfaces in compile-time for ResponseID.
var (
	_ bin.Encoder = &ResponseID{}
	_ bin.Decoder = &ResponseID{}

	_ ResponseClass = &ResponseID{}
)

// ResponseText represents TL type `responseText#cb0244f2`.
//
// See https://localhost:80/doc/constructor/responseText for reference.
type ResponseText struct {
	// Text field of ResponseText.
	Text string
}

// ResponseTextTypeID is TL type id of ResponseText.
const ResponseTextTypeID = 0xcb0244f2

func (r *ResponseText) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResponseText) String() string {
	if r == nil {
		return "ResponseText(nil)"
	}
	type Alias ResponseText
	return fmt.Sprintf("ResponseText%+v", Alias(*r))
}

// FillFrom fills ResponseText from given interface.
func (r *ResponseText) FillFrom(from interface {
	GetText() (value string)
}) {
	r.Text = from.GetText()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResponseText) TypeID() uint32 {
	return ResponseTextTypeID
}

// TypeName returns name of type in TL schema.
func (*ResponseText) TypeName() string {
	return "responseText"
}

// TypeInfo returns info about TL type.
func (r *ResponseText) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "responseText",
		ID:   ResponseTextTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResponseText) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode responseText#cb0244f2 as nil")
	}
	b.PutID(ResponseTextTypeID)
	b.PutString(r.Text)
	return nil
}

// GetText returns value of Text field.
func (r *ResponseText) GetText() (value string) {
	return r.Text
}

// Decode implements bin.Decoder.
func (r *ResponseText) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode responseText#cb0244f2 to nil")
	}
	if err := b.ConsumeID(ResponseTextTypeID); err != nil {
		return fmt.Errorf("unable to decode responseText#cb0244f2: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode responseText#cb0244f2: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// construct implements constructor of ResponseClass.
func (r ResponseText) construct() ResponseClass { return &r }

// Ensuring interfaces in compile-time for ResponseText.
var (
	_ bin.Encoder = &ResponseText{}
	_ bin.Decoder = &ResponseText{}

	_ ResponseClass = &ResponseText{}
)

// ResponseClass represents Response generic type.
//
// See https://localhost:80/doc/type/Response for reference.
//
// Example:
//  g, err := td.DecodeResponse(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *td.ResponseID: // responseID#85d7fd8b
//  case *td.ResponseText: // responseText#cb0244f2
//  default: panic(v)
//  }
type ResponseClass interface {
	bin.Encoder
	bin.Decoder
	construct() ResponseClass

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

// DecodeResponse implements binary de-serialization for ResponseClass.
func DecodeResponse(buf *bin.Buffer) (ResponseClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ResponseIDTypeID:
		// Decoding responseID#85d7fd8b.
		v := ResponseID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResponseClass: %w", err)
		}
		return &v, nil
	case ResponseTextTypeID:
		// Decoding responseText#cb0244f2.
		v := ResponseText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ResponseClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ResponseClass: %w", bin.NewUnexpectedID(id))
	}
}

// Response boxes the ResponseClass providing a helper.
type ResponseBox struct {
	Response ResponseClass
}

// Decode implements bin.Decoder for ResponseBox.
func (b *ResponseBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ResponseBox to nil")
	}
	v, err := DecodeResponse(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Response = v
	return nil
}

// Encode implements bin.Encode for ResponseBox.
func (b *ResponseBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Response == nil {
		return fmt.Errorf("unable to encode ResponseClass as nil")
	}
	return b.Response.Encode(buf)
}

// ResponseClassArray is adapter for slice of ResponseClass.
type ResponseClassArray []ResponseClass

// Sort sorts slice of ResponseClass.
func (s ResponseClassArray) Sort(less func(a, b ResponseClass) bool) ResponseClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ResponseClass.
func (s ResponseClassArray) SortStable(less func(a, b ResponseClass) bool) ResponseClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ResponseClass.
func (s ResponseClassArray) Retain(keep func(x ResponseClass) bool) ResponseClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ResponseClassArray) First() (v ResponseClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ResponseClassArray) Last() (v ResponseClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ResponseClassArray) PopFirst() (v ResponseClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ResponseClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ResponseClassArray) Pop() (v ResponseClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsResponseID returns copy with only ResponseID constructors.
func (s ResponseClassArray) AsResponseID() (to ResponseIDArray) {
	for _, elem := range s {
		value, ok := elem.(*ResponseID)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsResponseText returns copy with only ResponseText constructors.
func (s ResponseClassArray) AsResponseText() (to ResponseTextArray) {
	for _, elem := range s {
		value, ok := elem.(*ResponseText)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ResponseIDArray is adapter for slice of ResponseID.
type ResponseIDArray []ResponseID

// Sort sorts slice of ResponseID.
func (s ResponseIDArray) Sort(less func(a, b ResponseID) bool) ResponseIDArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ResponseID.
func (s ResponseIDArray) SortStable(less func(a, b ResponseID) bool) ResponseIDArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ResponseID.
func (s ResponseIDArray) Retain(keep func(x ResponseID) bool) ResponseIDArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ResponseIDArray) First() (v ResponseID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ResponseIDArray) Last() (v ResponseID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ResponseIDArray) PopFirst() (v ResponseID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ResponseID
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ResponseIDArray) Pop() (v ResponseID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// ResponseTextArray is adapter for slice of ResponseText.
type ResponseTextArray []ResponseText

// Sort sorts slice of ResponseText.
func (s ResponseTextArray) Sort(less func(a, b ResponseText) bool) ResponseTextArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ResponseText.
func (s ResponseTextArray) SortStable(less func(a, b ResponseText) bool) ResponseTextArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ResponseText.
func (s ResponseTextArray) Retain(keep func(x ResponseText) bool) ResponseTextArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s ResponseTextArray) First() (v ResponseText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ResponseTextArray) Last() (v ResponseText, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ResponseTextArray) PopFirst() (v ResponseText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ResponseText
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ResponseTextArray) Pop() (v ResponseText, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
