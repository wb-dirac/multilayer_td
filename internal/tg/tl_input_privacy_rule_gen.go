// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// InputPrivacyValueAllowContacts represents TL type `inputPrivacyValueAllowContacts#d09e07b`.
type InputPrivacyValueAllowContacts struct {
}

// InputPrivacyValueAllowContactsTypeID is TL type id of InputPrivacyValueAllowContacts.
const InputPrivacyValueAllowContactsTypeID = 0xd09e07b

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowContacts#d09e07b as nil")
	}
	b.PutID(InputPrivacyValueAllowContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowContacts#d09e07b to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowContacts#d09e07b: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowContacts) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowContacts.
var (
	_ bin.Encoder = &InputPrivacyValueAllowContacts{}
	_ bin.Decoder = &InputPrivacyValueAllowContacts{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowContacts{}
)

// InputPrivacyValueAllowAll represents TL type `inputPrivacyValueAllowAll#184b35ce`.
type InputPrivacyValueAllowAll struct {
}

// InputPrivacyValueAllowAllTypeID is TL type id of InputPrivacyValueAllowAll.
const InputPrivacyValueAllowAllTypeID = 0x184b35ce

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowAll) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowAll#184b35ce as nil")
	}
	b.PutID(InputPrivacyValueAllowAllTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowAll) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowAll#184b35ce to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowAll#184b35ce: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowAll) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowAll.
var (
	_ bin.Encoder = &InputPrivacyValueAllowAll{}
	_ bin.Decoder = &InputPrivacyValueAllowAll{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowAll{}
)

// InputPrivacyValueAllowUsers represents TL type `inputPrivacyValueAllowUsers#131cc67f`.
type InputPrivacyValueAllowUsers struct {
	// Users field of InputPrivacyValueAllowUsers.
	Users []InputUserClass
}

// InputPrivacyValueAllowUsersTypeID is TL type id of InputPrivacyValueAllowUsers.
const InputPrivacyValueAllowUsersTypeID = 0x131cc67f

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowUsers#131cc67f as nil")
	}
	b.PutID(InputPrivacyValueAllowUsersTypeID)
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode inputPrivacyValueAllowUsers#131cc67f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPrivacyValueAllowUsers#131cc67f: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowUsers#131cc67f to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueAllowUsers#131cc67f: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowUsers) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowUsers.
var (
	_ bin.Encoder = &InputPrivacyValueAllowUsers{}
	_ bin.Decoder = &InputPrivacyValueAllowUsers{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowUsers{}
)

// InputPrivacyValueDisallowContacts represents TL type `inputPrivacyValueDisallowContacts#ba52007`.
type InputPrivacyValueDisallowContacts struct {
}

// InputPrivacyValueDisallowContactsTypeID is TL type id of InputPrivacyValueDisallowContacts.
const InputPrivacyValueDisallowContactsTypeID = 0xba52007

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowContacts#ba52007 as nil")
	}
	b.PutID(InputPrivacyValueDisallowContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowContacts#ba52007 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowContacts#ba52007: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowContacts) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowContacts.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowContacts{}
	_ bin.Decoder = &InputPrivacyValueDisallowContacts{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowContacts{}
)

// InputPrivacyValueDisallowAll represents TL type `inputPrivacyValueDisallowAll#d66b66c9`.
type InputPrivacyValueDisallowAll struct {
}

// InputPrivacyValueDisallowAllTypeID is TL type id of InputPrivacyValueDisallowAll.
const InputPrivacyValueDisallowAllTypeID = 0xd66b66c9

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowAll) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowAll#d66b66c9 as nil")
	}
	b.PutID(InputPrivacyValueDisallowAllTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowAll) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowAll#d66b66c9 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowAll#d66b66c9: %w", err)
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowAll) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowAll.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowAll{}
	_ bin.Decoder = &InputPrivacyValueDisallowAll{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowAll{}
)

// InputPrivacyValueDisallowUsers represents TL type `inputPrivacyValueDisallowUsers#90110467`.
type InputPrivacyValueDisallowUsers struct {
	// Users field of InputPrivacyValueDisallowUsers.
	Users []InputUserClass
}

// InputPrivacyValueDisallowUsersTypeID is TL type id of InputPrivacyValueDisallowUsers.
const InputPrivacyValueDisallowUsersTypeID = 0x90110467

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowUsers) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowUsers#90110467 as nil")
	}
	b.PutID(InputPrivacyValueDisallowUsersTypeID)
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode inputPrivacyValueDisallowUsers#90110467: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputPrivacyValueDisallowUsers#90110467: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowUsers) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowUsers#90110467 to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueDisallowUsers#90110467: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowUsers) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowUsers.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowUsers{}
	_ bin.Decoder = &InputPrivacyValueDisallowUsers{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowUsers{}
)

// InputPrivacyValueAllowChatParticipants represents TL type `inputPrivacyValueAllowChatParticipants#4c81c1ba`.
type InputPrivacyValueAllowChatParticipants struct {
	// Chats field of InputPrivacyValueAllowChatParticipants.
	Chats []int
}

// InputPrivacyValueAllowChatParticipantsTypeID is TL type id of InputPrivacyValueAllowChatParticipants.
const InputPrivacyValueAllowChatParticipantsTypeID = 0x4c81c1ba

// Encode implements bin.Encoder.
func (i *InputPrivacyValueAllowChatParticipants) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueAllowChatParticipants#4c81c1ba as nil")
	}
	b.PutID(InputPrivacyValueAllowChatParticipantsTypeID)
	b.PutVectorHeader(len(i.Chats))
	for _, v := range i.Chats {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueAllowChatParticipants) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueAllowChatParticipants#4c81c1ba to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueAllowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#4c81c1ba: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#4c81c1ba: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueAllowChatParticipants#4c81c1ba: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueAllowChatParticipants) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueAllowChatParticipants.
var (
	_ bin.Encoder = &InputPrivacyValueAllowChatParticipants{}
	_ bin.Decoder = &InputPrivacyValueAllowChatParticipants{}

	_ InputPrivacyRuleClass = &InputPrivacyValueAllowChatParticipants{}
)

// InputPrivacyValueDisallowChatParticipants represents TL type `inputPrivacyValueDisallowChatParticipants#d82363af`.
type InputPrivacyValueDisallowChatParticipants struct {
	// Chats field of InputPrivacyValueDisallowChatParticipants.
	Chats []int
}

// InputPrivacyValueDisallowChatParticipantsTypeID is TL type id of InputPrivacyValueDisallowChatParticipants.
const InputPrivacyValueDisallowChatParticipantsTypeID = 0xd82363af

// Encode implements bin.Encoder.
func (i *InputPrivacyValueDisallowChatParticipants) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPrivacyValueDisallowChatParticipants#d82363af as nil")
	}
	b.PutID(InputPrivacyValueDisallowChatParticipantsTypeID)
	b.PutVectorHeader(len(i.Chats))
	for _, v := range i.Chats {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPrivacyValueDisallowChatParticipants) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPrivacyValueDisallowChatParticipants#d82363af to nil")
	}
	if err := b.ConsumeID(InputPrivacyValueDisallowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#d82363af: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#d82363af: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode inputPrivacyValueDisallowChatParticipants#d82363af: field chats: %w", err)
			}
			i.Chats = append(i.Chats, value)
		}
	}
	return nil
}

// construct implements constructor of InputPrivacyRuleClass.
func (i InputPrivacyValueDisallowChatParticipants) construct() InputPrivacyRuleClass { return &i }

// Ensuring interfaces in compile-time for InputPrivacyValueDisallowChatParticipants.
var (
	_ bin.Encoder = &InputPrivacyValueDisallowChatParticipants{}
	_ bin.Decoder = &InputPrivacyValueDisallowChatParticipants{}

	_ InputPrivacyRuleClass = &InputPrivacyValueDisallowChatParticipants{}
)

// InputPrivacyRuleClass represents InputPrivacyRule generic type.
//
// Example:
//  g, err := DecodeInputPrivacyRule(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputPrivacyValueAllowContacts: // inputPrivacyValueAllowContacts#d09e07b
//  case *InputPrivacyValueAllowAll: // inputPrivacyValueAllowAll#184b35ce
//  case *InputPrivacyValueAllowUsers: // inputPrivacyValueAllowUsers#131cc67f
//  case *InputPrivacyValueDisallowContacts: // inputPrivacyValueDisallowContacts#ba52007
//  case *InputPrivacyValueDisallowAll: // inputPrivacyValueDisallowAll#d66b66c9
//  case *InputPrivacyValueDisallowUsers: // inputPrivacyValueDisallowUsers#90110467
//  case *InputPrivacyValueAllowChatParticipants: // inputPrivacyValueAllowChatParticipants#4c81c1ba
//  case *InputPrivacyValueDisallowChatParticipants: // inputPrivacyValueDisallowChatParticipants#d82363af
//  default: panic(v)
//  }
type InputPrivacyRuleClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPrivacyRuleClass
}

// DecodeInputPrivacyRule implements binary de-serialization for InputPrivacyRuleClass.
func DecodeInputPrivacyRule(buf *bin.Buffer) (InputPrivacyRuleClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPrivacyValueAllowContactsTypeID:
		// Decoding inputPrivacyValueAllowContacts#d09e07b.
		v := InputPrivacyValueAllowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowAllTypeID:
		// Decoding inputPrivacyValueAllowAll#184b35ce.
		v := InputPrivacyValueAllowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowUsersTypeID:
		// Decoding inputPrivacyValueAllowUsers#131cc67f.
		v := InputPrivacyValueAllowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowContactsTypeID:
		// Decoding inputPrivacyValueDisallowContacts#ba52007.
		v := InputPrivacyValueDisallowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowAllTypeID:
		// Decoding inputPrivacyValueDisallowAll#d66b66c9.
		v := InputPrivacyValueDisallowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowUsersTypeID:
		// Decoding inputPrivacyValueDisallowUsers#90110467.
		v := InputPrivacyValueDisallowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueAllowChatParticipantsTypeID:
		// Decoding inputPrivacyValueAllowChatParticipants#4c81c1ba.
		v := InputPrivacyValueAllowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	case InputPrivacyValueDisallowChatParticipantsTypeID:
		// Decoding inputPrivacyValueDisallowChatParticipants#d82363af.
		v := InputPrivacyValueDisallowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPrivacyRuleClass: %w", bin.NewUnexpectedID(id))
	}
}