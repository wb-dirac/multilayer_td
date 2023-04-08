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

// PhotosUpdateProfilePhotoRequest represents TL type `photos.updateProfilePhoto#1c3d5956`.
// Installs a previously uploaded photo as a profile photo.
//
// See https://core.telegram.org/method/photos.updateProfilePhoto for reference.
type PhotosUpdateProfilePhotoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, the chosen profile photo will be shown to users that can't display your main
	// profile photo due to your privacy settings.
	Fallback bool
	// Input photo
	ID InputPhotoClass
}

// PhotosUpdateProfilePhotoRequestTypeID is TL type id of PhotosUpdateProfilePhotoRequest.
const PhotosUpdateProfilePhotoRequestTypeID = 0x1c3d5956

// Ensuring interfaces in compile-time for PhotosUpdateProfilePhotoRequest.
var (
	_ bin.Encoder     = &PhotosUpdateProfilePhotoRequest{}
	_ bin.Decoder     = &PhotosUpdateProfilePhotoRequest{}
	_ bin.BareEncoder = &PhotosUpdateProfilePhotoRequest{}
	_ bin.BareDecoder = &PhotosUpdateProfilePhotoRequest{}
)

func (u *PhotosUpdateProfilePhotoRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Fallback == false) {
		return false
	}
	if !(u.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *PhotosUpdateProfilePhotoRequest) String() string {
	if u == nil {
		return "PhotosUpdateProfilePhotoRequest(nil)"
	}
	type Alias PhotosUpdateProfilePhotoRequest
	return fmt.Sprintf("PhotosUpdateProfilePhotoRequest%+v", Alias(*u))
}

// FillFrom fills PhotosUpdateProfilePhotoRequest from given interface.
func (u *PhotosUpdateProfilePhotoRequest) FillFrom(from interface {
	GetFallback() (value bool)
	GetID() (value InputPhotoClass)
}) {
	u.Fallback = from.GetFallback()
	u.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosUpdateProfilePhotoRequest) TypeID() uint32 {
	return PhotosUpdateProfilePhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosUpdateProfilePhotoRequest) TypeName() string {
	return "photos.updateProfilePhoto"
}

// TypeInfo returns info about TL type.
func (u *PhotosUpdateProfilePhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.updateProfilePhoto",
		ID:   PhotosUpdateProfilePhotoRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Fallback",
			SchemaName: "fallback",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *PhotosUpdateProfilePhotoRequest) SetFlags() {
	if !(u.Fallback == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *PhotosUpdateProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.updateProfilePhoto#1c3d5956 as nil")
	}
	b.PutID(PhotosUpdateProfilePhotoRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *PhotosUpdateProfilePhotoRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode photos.updateProfilePhoto#1c3d5956 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.updateProfilePhoto#1c3d5956: field flags: %w", err)
	}
	if u.ID == nil {
		return fmt.Errorf("unable to encode photos.updateProfilePhoto#1c3d5956: field id is nil")
	}
	if err := u.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.updateProfilePhoto#1c3d5956: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *PhotosUpdateProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.updateProfilePhoto#1c3d5956 to nil")
	}
	if err := b.ConsumeID(PhotosUpdateProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.updateProfilePhoto#1c3d5956: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *PhotosUpdateProfilePhotoRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode photos.updateProfilePhoto#1c3d5956 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photos.updateProfilePhoto#1c3d5956: field flags: %w", err)
		}
	}
	u.Fallback = u.Flags.Has(0)
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.updateProfilePhoto#1c3d5956: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// SetFallback sets value of Fallback conditional field.
func (u *PhotosUpdateProfilePhotoRequest) SetFallback(value bool) {
	if value {
		u.Flags.Set(0)
		u.Fallback = true
	} else {
		u.Flags.Unset(0)
		u.Fallback = false
	}
}

// GetFallback returns value of Fallback conditional field.
func (u *PhotosUpdateProfilePhotoRequest) GetFallback() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// GetID returns value of ID field.
func (u *PhotosUpdateProfilePhotoRequest) GetID() (value InputPhotoClass) {
	if u == nil {
		return
	}
	return u.ID
}

// GetIDAsNotEmpty returns mapped value of ID field.
func (u *PhotosUpdateProfilePhotoRequest) GetIDAsNotEmpty() (*InputPhoto, bool) {
	return u.ID.AsNotEmpty()
}

// PhotosUpdateProfilePhoto invokes method photos.updateProfilePhoto#1c3d5956 returning error if any.
// Installs a previously uploaded photo as a profile photo.
//
// Possible errors:
//
//	400 ALBUM_PHOTOS_TOO_MANY: You have uploaded too many profile photos, delete some before retrying.
//	400 FILE_PARTS_INVALID: The number of file parts is invalid.
//	400 IMAGE_PROCESS_FAILED: Failure while processing image.
//	400 LOCATION_INVALID: The provided location is invalid.
//	400 PHOTO_CROP_SIZE_SMALL: Photo is too small.
//	400 PHOTO_EXT_INVALID: The extension of the photo is invalid.
//	400 PHOTO_ID_INVALID: Photo ID invalid.
//
// See https://core.telegram.org/method/photos.updateProfilePhoto for reference.
func (c *Client) PhotosUpdateProfilePhoto(ctx context.Context, request *PhotosUpdateProfilePhotoRequest) (*PhotosPhoto, error) {
	var result PhotosPhoto

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
