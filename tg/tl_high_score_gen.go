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

// HighScore represents TL type `highScore#73a379eb`.
// Game highscore
//
// See https://core.telegram.org/constructor/highScore for reference.
type HighScore struct {
	// Position in highscore list
	Pos int
	// User ID
	UserID int64
	// Score
	Score int
}

// HighScoreTypeID is TL type id of HighScore.
const HighScoreTypeID = 0x73a379eb

// Ensuring interfaces in compile-time for HighScore.
var (
	_ bin.Encoder     = &HighScore{}
	_ bin.Decoder     = &HighScore{}
	_ bin.BareEncoder = &HighScore{}
	_ bin.BareDecoder = &HighScore{}
)

func (h *HighScore) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Pos == 0) {
		return false
	}
	if !(h.UserID == 0) {
		return false
	}
	if !(h.Score == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *HighScore) String() string {
	if h == nil {
		return "HighScore(nil)"
	}
	type Alias HighScore
	return fmt.Sprintf("HighScore%+v", Alias(*h))
}

// FillFrom fills HighScore from given interface.
func (h *HighScore) FillFrom(from interface {
	GetPos() (value int)
	GetUserID() (value int64)
	GetScore() (value int)
}) {
	h.Pos = from.GetPos()
	h.UserID = from.GetUserID()
	h.Score = from.GetScore()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HighScore) TypeID() uint32 {
	return HighScoreTypeID
}

// TypeName returns name of type in TL schema.
func (*HighScore) TypeName() string {
	return "highScore"
}

// TypeInfo returns info about TL type.
func (h *HighScore) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "highScore",
		ID:   HighScoreTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pos",
			SchemaName: "pos",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (h *HighScore) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode highScore#73a379eb as nil")
	}
	b.PutID(HighScoreTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *HighScore) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode highScore#73a379eb as nil")
	}
	b.PutInt(h.Pos)
	b.PutLong(h.UserID)
	b.PutInt(h.Score)
	return nil
}

// Decode implements bin.Decoder.
func (h *HighScore) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode highScore#73a379eb to nil")
	}
	if err := b.ConsumeID(HighScoreTypeID); err != nil {
		return fmt.Errorf("unable to decode highScore#73a379eb: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *HighScore) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode highScore#73a379eb to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode highScore#73a379eb: field pos: %w", err)
		}
		h.Pos = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode highScore#73a379eb: field user_id: %w", err)
		}
		h.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode highScore#73a379eb: field score: %w", err)
		}
		h.Score = value
	}
	return nil
}

// GetPos returns value of Pos field.
func (h *HighScore) GetPos() (value int) {
	return h.Pos
}

// GetUserID returns value of UserID field.
func (h *HighScore) GetUserID() (value int64) {
	return h.UserID
}

// GetScore returns value of Score field.
func (h *HighScore) GetScore() (value int) {
	return h.Score
}
