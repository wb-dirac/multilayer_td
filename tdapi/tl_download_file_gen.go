// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// DownloadFileRequest represents TL type `downloadFile#be50685a`.
type DownloadFileRequest struct {
	// Identifier of the file to download
	FileID int32
	// Priority of the download (1-32). The higher the priority, the earlier the file will be
	// downloaded. If the priorities of two files are equal, then the last one for which
	// downloadFile was called will be downloaded first
	Priority int32
	// The starting position from which the file should be downloaded
	Offset int32
	// Number of bytes which should be downloaded starting from the "offset" position before
	// the download will be automatically canceled; use 0 to download without a limit
	Limit int32
	// If false, this request returns file state just after the download has been started. If
	// true, this request returns file state only after
	Synchronous bool
}

// DownloadFileRequestTypeID is TL type id of DownloadFileRequest.
const DownloadFileRequestTypeID = 0xbe50685a

// Ensuring interfaces in compile-time for DownloadFileRequest.
var (
	_ bin.Encoder     = &DownloadFileRequest{}
	_ bin.Decoder     = &DownloadFileRequest{}
	_ bin.BareEncoder = &DownloadFileRequest{}
	_ bin.BareDecoder = &DownloadFileRequest{}
)

func (d *DownloadFileRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.FileID == 0) {
		return false
	}
	if !(d.Priority == 0) {
		return false
	}
	if !(d.Offset == 0) {
		return false
	}
	if !(d.Limit == 0) {
		return false
	}
	if !(d.Synchronous == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DownloadFileRequest) String() string {
	if d == nil {
		return "DownloadFileRequest(nil)"
	}
	type Alias DownloadFileRequest
	return fmt.Sprintf("DownloadFileRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DownloadFileRequest) TypeID() uint32 {
	return DownloadFileRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DownloadFileRequest) TypeName() string {
	return "downloadFile"
}

// TypeInfo returns info about TL type.
func (d *DownloadFileRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "downloadFile",
		ID:   DownloadFileRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileID",
			SchemaName: "file_id",
		},
		{
			Name:       "Priority",
			SchemaName: "priority",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Synchronous",
			SchemaName: "synchronous",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DownloadFileRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode downloadFile#be50685a as nil")
	}
	b.PutID(DownloadFileRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DownloadFileRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode downloadFile#be50685a as nil")
	}
	b.PutInt32(d.FileID)
	b.PutInt32(d.Priority)
	b.PutInt32(d.Offset)
	b.PutInt32(d.Limit)
	b.PutBool(d.Synchronous)
	return nil
}

// Decode implements bin.Decoder.
func (d *DownloadFileRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode downloadFile#be50685a to nil")
	}
	if err := b.ConsumeID(DownloadFileRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode downloadFile#be50685a: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DownloadFileRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode downloadFile#be50685a to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadFile#be50685a: field file_id: %w", err)
		}
		d.FileID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadFile#be50685a: field priority: %w", err)
		}
		d.Priority = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadFile#be50685a: field offset: %w", err)
		}
		d.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode downloadFile#be50685a: field limit: %w", err)
		}
		d.Limit = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode downloadFile#be50685a: field synchronous: %w", err)
		}
		d.Synchronous = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DownloadFileRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode downloadFile#be50685a as nil")
	}
	b.ObjStart()
	b.PutID("downloadFile")
	b.FieldStart("file_id")
	b.PutInt32(d.FileID)
	b.FieldStart("priority")
	b.PutInt32(d.Priority)
	b.FieldStart("offset")
	b.PutInt32(d.Offset)
	b.FieldStart("limit")
	b.PutInt32(d.Limit)
	b.FieldStart("synchronous")
	b.PutBool(d.Synchronous)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DownloadFileRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode downloadFile#be50685a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("downloadFile"); err != nil {
				return fmt.Errorf("unable to decode downloadFile#be50685a: %w", err)
			}
		case "file_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadFile#be50685a: field file_id: %w", err)
			}
			d.FileID = value
		case "priority":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadFile#be50685a: field priority: %w", err)
			}
			d.Priority = value
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadFile#be50685a: field offset: %w", err)
			}
			d.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode downloadFile#be50685a: field limit: %w", err)
			}
			d.Limit = value
		case "synchronous":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode downloadFile#be50685a: field synchronous: %w", err)
			}
			d.Synchronous = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFileID returns value of FileID field.
func (d *DownloadFileRequest) GetFileID() (value int32) {
	return d.FileID
}

// GetPriority returns value of Priority field.
func (d *DownloadFileRequest) GetPriority() (value int32) {
	return d.Priority
}

// GetOffset returns value of Offset field.
func (d *DownloadFileRequest) GetOffset() (value int32) {
	return d.Offset
}

// GetLimit returns value of Limit field.
func (d *DownloadFileRequest) GetLimit() (value int32) {
	return d.Limit
}

// GetSynchronous returns value of Synchronous field.
func (d *DownloadFileRequest) GetSynchronous() (value bool) {
	return d.Synchronous
}

// DownloadFile invokes method downloadFile#be50685a returning error if any.
func (c *Client) DownloadFile(ctx context.Context, request *DownloadFileRequest) (*File, error) {
	var result File

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
