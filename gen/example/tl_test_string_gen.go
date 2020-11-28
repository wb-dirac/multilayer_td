// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// TestString represents TL type `testString#fe56688c`.
type TestString struct {
	// String
	Value string
}

// TestStringTypeID is TL type id of TestString.
const TestStringTypeID = 0xfe56688c

// Encode implements bin.Encoder.
func (t *TestString) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testString#fe56688c as nil")
	}
	b.PutID(TestStringTypeID)
	b.PutString(t.Value)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestString) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testString#fe56688c to nil")
	}
	if err := b.ConsumeID(TestStringTypeID); err != nil {
		return fmt.Errorf("unable to decode testString#fe56688c: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode testString#fe56688c: field value: %w", err)
		}
		t.Value = value
	}
	return nil
}

// Ensuring interfaces in compile-time for TestString.
var (
	_ bin.Encoder = &TestString{}
	_ bin.Decoder = &TestString{}
)
