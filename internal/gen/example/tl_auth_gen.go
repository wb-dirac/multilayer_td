// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Auth represents TL type `auth#f8bb4a38`.
//
// See https://localhost:80/doc/constructor/auth for reference.
type Auth struct {
	// Name field of Auth.
	Name string
}

// AuthTypeID is TL type id of Auth.
const AuthTypeID = 0xf8bb4a38

// construct implements constructor of AuthClass.
func (a Auth) construct() AuthClass { return &a }

// Ensuring interfaces in compile-time for Auth.
var (
	_ bin.Encoder     = &Auth{}
	_ bin.Decoder     = &Auth{}
	_ bin.BareEncoder = &Auth{}
	_ bin.BareDecoder = &Auth{}

	_ AuthClass = &Auth{}
)

func (a *Auth) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *Auth) String() string {
	if a == nil {
		return "Auth(nil)"
	}
	type Alias Auth
	return fmt.Sprintf("Auth%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Auth) TypeID() uint32 {
	return AuthTypeID
}

// TypeName returns name of type in TL schema.
func (*Auth) TypeName() string {
	return "auth"
}

// TypeInfo returns info about TL type.
func (a *Auth) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth",
		ID:   AuthTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *Auth) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth#f8bb4a38 as nil")
	}
	b.PutID(AuthTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *Auth) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode auth#f8bb4a38 as nil")
	}
	b.PutString(a.Name)
	return nil
}

// Decode implements bin.Decoder.
func (a *Auth) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth#f8bb4a38 to nil")
	}
	if err := b.ConsumeID(AuthTypeID); err != nil {
		return fmt.Errorf("unable to decode auth#f8bb4a38: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *Auth) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode auth#f8bb4a38 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth#f8bb4a38: field name: %w", err)
		}
		a.Name = value
	}
	return nil
}

// GetName returns value of Name field.
func (a *Auth) GetName() (value string) {
	return a.Name
}

// AuthPassword represents TL type `authPassword#29bacabb`.
//
// See https://localhost:80/doc/constructor/authPassword for reference.
type AuthPassword struct {
	// Name field of AuthPassword.
	Name string
	// Password field of AuthPassword.
	Password string
}

// AuthPasswordTypeID is TL type id of AuthPassword.
const AuthPasswordTypeID = 0x29bacabb

// construct implements constructor of AuthClass.
func (a AuthPassword) construct() AuthClass { return &a }

// Ensuring interfaces in compile-time for AuthPassword.
var (
	_ bin.Encoder     = &AuthPassword{}
	_ bin.Decoder     = &AuthPassword{}
	_ bin.BareEncoder = &AuthPassword{}
	_ bin.BareDecoder = &AuthPassword{}

	_ AuthClass = &AuthPassword{}
)

func (a *AuthPassword) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Name == "") {
		return false
	}
	if !(a.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AuthPassword) String() string {
	if a == nil {
		return "AuthPassword(nil)"
	}
	type Alias AuthPassword
	return fmt.Sprintf("AuthPassword%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthPassword) TypeID() uint32 {
	return AuthPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthPassword) TypeName() string {
	return "authPassword"
}

// TypeInfo returns info about TL type.
func (a *AuthPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authPassword",
		ID:   AuthPasswordTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AuthPassword) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authPassword#29bacabb as nil")
	}
	b.PutID(AuthPasswordTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AuthPassword) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authPassword#29bacabb as nil")
	}
	b.PutString(a.Name)
	b.PutString(a.Password)
	return nil
}

// Decode implements bin.Decoder.
func (a *AuthPassword) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authPassword#29bacabb to nil")
	}
	if err := b.ConsumeID(AuthPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode authPassword#29bacabb: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AuthPassword) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authPassword#29bacabb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authPassword#29bacabb: field name: %w", err)
		}
		a.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authPassword#29bacabb: field password: %w", err)
		}
		a.Password = value
	}
	return nil
}

// GetName returns value of Name field.
func (a *AuthPassword) GetName() (value string) {
	return a.Name
}

// GetPassword returns value of Password field.
func (a *AuthPassword) GetPassword() (value string) {
	return a.Password
}

// AuthClass represents Auth generic type.
//
// See https://localhost:80/doc/type/Auth for reference.
//
// Example:
//  g, err := td.DecodeAuth(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *td.Auth: // auth#f8bb4a38
//  case *td.AuthPassword: // authPassword#29bacabb
//  default: panic(v)
//  }
type AuthClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthClass

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

	// Name field of Auth.
	GetName() (value string)
}

// DecodeAuth implements binary de-serialization for AuthClass.
func DecodeAuth(buf *bin.Buffer) (AuthClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthTypeID:
		// Decoding auth#f8bb4a38.
		v := Auth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthClass: %w", err)
		}
		return &v, nil
	case AuthPasswordTypeID:
		// Decoding authPassword#29bacabb.
		v := AuthPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthClass: %w", bin.NewUnexpectedID(id))
	}
}

// Auth boxes the AuthClass providing a helper.
type AuthBox struct {
	Auth AuthClass
}

// Decode implements bin.Decoder for AuthBox.
func (b *AuthBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthBox to nil")
	}
	v, err := DecodeAuth(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Auth = v
	return nil
}

// Encode implements bin.Encode for AuthBox.
func (b *AuthBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Auth == nil {
		return fmt.Errorf("unable to encode AuthClass as nil")
	}
	return b.Auth.Encode(buf)
}
