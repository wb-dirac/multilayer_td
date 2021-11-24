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

// InputThemeSettings represents TL type `inputThemeSettings#8fde504f`.
// Theme settings
//
// See https://core.telegram.org/constructor/inputThemeSettings for reference.
type InputThemeSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// MessageColorsAnimated field of InputThemeSettings.
	MessageColorsAnimated bool
	// Default theme on which this theme is based
	BaseTheme BaseThemeClass
	// Accent color, RGB24 format
	AccentColor int
	// OutboxAccentColor field of InputThemeSettings.
	//
	// Use SetOutboxAccentColor and GetOutboxAccentColor helpers.
	OutboxAccentColor int
	// MessageColors field of InputThemeSettings.
	//
	// Use SetMessageColors and GetMessageColors helpers.
	MessageColors []int
	// Wallpaper
	//
	// Use SetWallpaper and GetWallpaper helpers.
	Wallpaper InputWallPaperClass
	// Wallpaper settings
	//
	// Use SetWallpaperSettings and GetWallpaperSettings helpers.
	WallpaperSettings WallPaperSettings
}

// InputThemeSettingsTypeID is TL type id of InputThemeSettings.
const InputThemeSettingsTypeID = 0x8fde504f

// Ensuring interfaces in compile-time for InputThemeSettings.
var (
	_ bin.Encoder     = &InputThemeSettings{}
	_ bin.Decoder     = &InputThemeSettings{}
	_ bin.BareEncoder = &InputThemeSettings{}
	_ bin.BareDecoder = &InputThemeSettings{}
)

func (i *InputThemeSettings) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.MessageColorsAnimated == false) {
		return false
	}
	if !(i.BaseTheme == nil) {
		return false
	}
	if !(i.AccentColor == 0) {
		return false
	}
	if !(i.OutboxAccentColor == 0) {
		return false
	}
	if !(i.MessageColors == nil) {
		return false
	}
	if !(i.Wallpaper == nil) {
		return false
	}
	if !(i.WallpaperSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputThemeSettings) String() string {
	if i == nil {
		return "InputThemeSettings(nil)"
	}
	type Alias InputThemeSettings
	return fmt.Sprintf("InputThemeSettings%+v", Alias(*i))
}

// FillFrom fills InputThemeSettings from given interface.
func (i *InputThemeSettings) FillFrom(from interface {
	GetMessageColorsAnimated() (value bool)
	GetBaseTheme() (value BaseThemeClass)
	GetAccentColor() (value int)
	GetOutboxAccentColor() (value int, ok bool)
	GetMessageColors() (value []int, ok bool)
	GetWallpaper() (value InputWallPaperClass, ok bool)
	GetWallpaperSettings() (value WallPaperSettings, ok bool)
}) {
	i.MessageColorsAnimated = from.GetMessageColorsAnimated()
	i.BaseTheme = from.GetBaseTheme()
	i.AccentColor = from.GetAccentColor()
	if val, ok := from.GetOutboxAccentColor(); ok {
		i.OutboxAccentColor = val
	}

	if val, ok := from.GetMessageColors(); ok {
		i.MessageColors = val
	}

	if val, ok := from.GetWallpaper(); ok {
		i.Wallpaper = val
	}

	if val, ok := from.GetWallpaperSettings(); ok {
		i.WallpaperSettings = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputThemeSettings) TypeID() uint32 {
	return InputThemeSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputThemeSettings) TypeName() string {
	return "inputThemeSettings"
}

// TypeInfo returns info about TL type.
func (i *InputThemeSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputThemeSettings",
		ID:   InputThemeSettingsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageColorsAnimated",
			SchemaName: "message_colors_animated",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "BaseTheme",
			SchemaName: "base_theme",
		},
		{
			Name:       "AccentColor",
			SchemaName: "accent_color",
		},
		{
			Name:       "OutboxAccentColor",
			SchemaName: "outbox_accent_color",
			Null:       !i.Flags.Has(3),
		},
		{
			Name:       "MessageColors",
			SchemaName: "message_colors",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Wallpaper",
			SchemaName: "wallpaper",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "WallpaperSettings",
			SchemaName: "wallpaper_settings",
			Null:       !i.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputThemeSettings) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThemeSettings#8fde504f as nil")
	}
	b.PutID(InputThemeSettingsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputThemeSettings) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputThemeSettings#8fde504f as nil")
	}
	if !(i.MessageColorsAnimated == false) {
		i.Flags.Set(2)
	}
	if !(i.OutboxAccentColor == 0) {
		i.Flags.Set(3)
	}
	if !(i.MessageColors == nil) {
		i.Flags.Set(0)
	}
	if !(i.Wallpaper == nil) {
		i.Flags.Set(1)
	}
	if !(i.WallpaperSettings.Zero()) {
		i.Flags.Set(1)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputThemeSettings#8fde504f: field flags: %w", err)
	}
	if i.BaseTheme == nil {
		return fmt.Errorf("unable to encode inputThemeSettings#8fde504f: field base_theme is nil")
	}
	if err := i.BaseTheme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputThemeSettings#8fde504f: field base_theme: %w", err)
	}
	b.PutInt(i.AccentColor)
	if i.Flags.Has(3) {
		b.PutInt(i.OutboxAccentColor)
	}
	if i.Flags.Has(0) {
		b.PutVectorHeader(len(i.MessageColors))
		for _, v := range i.MessageColors {
			b.PutInt(v)
		}
	}
	if i.Flags.Has(1) {
		if i.Wallpaper == nil {
			return fmt.Errorf("unable to encode inputThemeSettings#8fde504f: field wallpaper is nil")
		}
		if err := i.Wallpaper.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputThemeSettings#8fde504f: field wallpaper: %w", err)
		}
	}
	if i.Flags.Has(1) {
		if err := i.WallpaperSettings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputThemeSettings#8fde504f: field wallpaper_settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputThemeSettings) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThemeSettings#8fde504f to nil")
	}
	if err := b.ConsumeID(InputThemeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputThemeSettings) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputThemeSettings#8fde504f to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field flags: %w", err)
		}
	}
	i.MessageColorsAnimated = i.Flags.Has(2)
	{
		value, err := DecodeBaseTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field base_theme: %w", err)
		}
		i.BaseTheme = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field accent_color: %w", err)
		}
		i.AccentColor = value
	}
	if i.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field outbox_accent_color: %w", err)
		}
		i.OutboxAccentColor = value
	}
	if i.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field message_colors: %w", err)
		}

		if headerLen > 0 {
			i.MessageColors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field message_colors: %w", err)
			}
			i.MessageColors = append(i.MessageColors, value)
		}
	}
	if i.Flags.Has(1) {
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field wallpaper: %w", err)
		}
		i.Wallpaper = value
	}
	if i.Flags.Has(1) {
		if err := i.WallpaperSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputThemeSettings#8fde504f: field wallpaper_settings: %w", err)
		}
	}
	return nil
}

// SetMessageColorsAnimated sets value of MessageColorsAnimated conditional field.
func (i *InputThemeSettings) SetMessageColorsAnimated(value bool) {
	if value {
		i.Flags.Set(2)
		i.MessageColorsAnimated = true
	} else {
		i.Flags.Unset(2)
		i.MessageColorsAnimated = false
	}
}

// GetMessageColorsAnimated returns value of MessageColorsAnimated conditional field.
func (i *InputThemeSettings) GetMessageColorsAnimated() (value bool) {
	return i.Flags.Has(2)
}

// GetBaseTheme returns value of BaseTheme field.
func (i *InputThemeSettings) GetBaseTheme() (value BaseThemeClass) {
	return i.BaseTheme
}

// GetAccentColor returns value of AccentColor field.
func (i *InputThemeSettings) GetAccentColor() (value int) {
	return i.AccentColor
}

// SetOutboxAccentColor sets value of OutboxAccentColor conditional field.
func (i *InputThemeSettings) SetOutboxAccentColor(value int) {
	i.Flags.Set(3)
	i.OutboxAccentColor = value
}

// GetOutboxAccentColor returns value of OutboxAccentColor conditional field and
// boolean which is true if field was set.
func (i *InputThemeSettings) GetOutboxAccentColor() (value int, ok bool) {
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.OutboxAccentColor, true
}

// SetMessageColors sets value of MessageColors conditional field.
func (i *InputThemeSettings) SetMessageColors(value []int) {
	i.Flags.Set(0)
	i.MessageColors = value
}

// GetMessageColors returns value of MessageColors conditional field and
// boolean which is true if field was set.
func (i *InputThemeSettings) GetMessageColors() (value []int, ok bool) {
	if !i.Flags.Has(0) {
		return value, false
	}
	return i.MessageColors, true
}

// SetWallpaper sets value of Wallpaper conditional field.
func (i *InputThemeSettings) SetWallpaper(value InputWallPaperClass) {
	i.Flags.Set(1)
	i.Wallpaper = value
}

// GetWallpaper returns value of Wallpaper conditional field and
// boolean which is true if field was set.
func (i *InputThemeSettings) GetWallpaper() (value InputWallPaperClass, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Wallpaper, true
}

// SetWallpaperSettings sets value of WallpaperSettings conditional field.
func (i *InputThemeSettings) SetWallpaperSettings(value WallPaperSettings) {
	i.Flags.Set(1)
	i.WallpaperSettings = value
}

// GetWallpaperSettings returns value of WallpaperSettings conditional field and
// boolean which is true if field was set.
func (i *InputThemeSettings) GetWallpaperSettings() (value WallPaperSettings, ok bool) {
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.WallpaperSettings, true
}
