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

// GlobalPrivacySettings represents TL type `globalPrivacySettings#bea2f424`.
// Global privacy settings
//
// See https://core.telegram.org/constructor/globalPrivacySettings for reference.
type GlobalPrivacySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to archive and mute new chats from non-contacts
	//
	// Use SetArchiveAndMuteNewNoncontactPeers and GetArchiveAndMuteNewNoncontactPeers helpers.
	ArchiveAndMuteNewNoncontactPeers bool
}

// GlobalPrivacySettingsTypeID is TL type id of GlobalPrivacySettings.
const GlobalPrivacySettingsTypeID = 0xbea2f424

// Ensuring interfaces in compile-time for GlobalPrivacySettings.
var (
	_ bin.Encoder     = &GlobalPrivacySettings{}
	_ bin.Decoder     = &GlobalPrivacySettings{}
	_ bin.BareEncoder = &GlobalPrivacySettings{}
	_ bin.BareDecoder = &GlobalPrivacySettings{}
)

func (g *GlobalPrivacySettings) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GlobalPrivacySettings) String() string {
	if g == nil {
		return "GlobalPrivacySettings(nil)"
	}
	type Alias GlobalPrivacySettings
	return fmt.Sprintf("GlobalPrivacySettings%+v", Alias(*g))
}

// FillFrom fills GlobalPrivacySettings from given interface.
func (g *GlobalPrivacySettings) FillFrom(from interface {
	GetArchiveAndMuteNewNoncontactPeers() (value bool, ok bool)
}) {
	if val, ok := from.GetArchiveAndMuteNewNoncontactPeers(); ok {
		g.ArchiveAndMuteNewNoncontactPeers = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GlobalPrivacySettings) TypeID() uint32 {
	return GlobalPrivacySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*GlobalPrivacySettings) TypeName() string {
	return "globalPrivacySettings"
}

// TypeInfo returns info about TL type.
func (g *GlobalPrivacySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "globalPrivacySettings",
		ID:   GlobalPrivacySettingsTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ArchiveAndMuteNewNoncontactPeers",
			SchemaName: "archive_and_mute_new_noncontact_peers",
			Null:       !g.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GlobalPrivacySettings) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#bea2f424 as nil")
	}
	b.PutID(GlobalPrivacySettingsTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GlobalPrivacySettings) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#bea2f424 as nil")
	}
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode globalPrivacySettings#bea2f424: field flags: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutBool(g.ArchiveAndMuteNewNoncontactPeers)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GlobalPrivacySettings) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#bea2f424 to nil")
	}
	if err := b.ConsumeID(GlobalPrivacySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GlobalPrivacySettings) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#bea2f424 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: field flags: %w", err)
		}
	}
	if g.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#bea2f424: field archive_and_mute_new_noncontact_peers: %w", err)
		}
		g.ArchiveAndMuteNewNoncontactPeers = value
	}
	return nil
}

// SetArchiveAndMuteNewNoncontactPeers sets value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) SetArchiveAndMuteNewNoncontactPeers(value bool) {
	g.Flags.Set(0)
	g.ArchiveAndMuteNewNoncontactPeers = value
}

// GetArchiveAndMuteNewNoncontactPeers returns value of ArchiveAndMuteNewNoncontactPeers conditional field and
// boolean which is true if field was set.
func (g *GlobalPrivacySettings) GetArchiveAndMuteNewNoncontactPeers() (value bool, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.ArchiveAndMuteNewNoncontactPeers, true
}
