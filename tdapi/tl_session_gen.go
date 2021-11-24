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

// Session represents TL type `session#727950d8`.
type Session struct {
	// Session identifier
	ID int64
	// True, if this session is the current session
	IsCurrent bool
	// True, if a password is needed to complete authorization of the session
	IsPasswordPending bool
	// Telegram API identifier, as provided by the application
	APIID int32
	// Name of the application, as provided by the application
	ApplicationName string
	// The version of the application, as provided by the application
	ApplicationVersion string
	// True, if the application is an official application or uses the api_id of an official
	// application
	IsOfficialApplication bool
	// Model of the device the application has been run or is running on, as provided by the
	// application
	DeviceModel string
	// Operating system the application has been run or is running on, as provided by the
	// application
	Platform string
	// Version of the operating system the application has been run or is running on, as
	// provided by the application
	SystemVersion string
	// Point in time (Unix timestamp) when the user has logged in
	LogInDate int32
	// Point in time (Unix timestamp) when the session was last used
	LastActiveDate int32
	// IP address from which the session was created, in human-readable format
	IP string
	// A two-letter country code for the country from which the session was created, based on
	// the IP address
	Country string
	// Region code from which the session was created, based on the IP address
	Region string
}

// SessionTypeID is TL type id of Session.
const SessionTypeID = 0x727950d8

// Ensuring interfaces in compile-time for Session.
var (
	_ bin.Encoder     = &Session{}
	_ bin.Decoder     = &Session{}
	_ bin.BareEncoder = &Session{}
	_ bin.BareDecoder = &Session{}
)

func (s *Session) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.IsCurrent == false) {
		return false
	}
	if !(s.IsPasswordPending == false) {
		return false
	}
	if !(s.APIID == 0) {
		return false
	}
	if !(s.ApplicationName == "") {
		return false
	}
	if !(s.ApplicationVersion == "") {
		return false
	}
	if !(s.IsOfficialApplication == false) {
		return false
	}
	if !(s.DeviceModel == "") {
		return false
	}
	if !(s.Platform == "") {
		return false
	}
	if !(s.SystemVersion == "") {
		return false
	}
	if !(s.LogInDate == 0) {
		return false
	}
	if !(s.LastActiveDate == 0) {
		return false
	}
	if !(s.IP == "") {
		return false
	}
	if !(s.Country == "") {
		return false
	}
	if !(s.Region == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Session) String() string {
	if s == nil {
		return "Session(nil)"
	}
	type Alias Session
	return fmt.Sprintf("Session%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Session) TypeID() uint32 {
	return SessionTypeID
}

// TypeName returns name of type in TL schema.
func (*Session) TypeName() string {
	return "session"
}

// TypeInfo returns info about TL type.
func (s *Session) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "session",
		ID:   SessionTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "IsCurrent",
			SchemaName: "is_current",
		},
		{
			Name:       "IsPasswordPending",
			SchemaName: "is_password_pending",
		},
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "ApplicationName",
			SchemaName: "application_name",
		},
		{
			Name:       "ApplicationVersion",
			SchemaName: "application_version",
		},
		{
			Name:       "IsOfficialApplication",
			SchemaName: "is_official_application",
		},
		{
			Name:       "DeviceModel",
			SchemaName: "device_model",
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
		{
			Name:       "SystemVersion",
			SchemaName: "system_version",
		},
		{
			Name:       "LogInDate",
			SchemaName: "log_in_date",
		},
		{
			Name:       "LastActiveDate",
			SchemaName: "last_active_date",
		},
		{
			Name:       "IP",
			SchemaName: "ip",
		},
		{
			Name:       "Country",
			SchemaName: "country",
		},
		{
			Name:       "Region",
			SchemaName: "region",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Session) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode session#727950d8 as nil")
	}
	b.PutID(SessionTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Session) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode session#727950d8 as nil")
	}
	b.PutLong(s.ID)
	b.PutBool(s.IsCurrent)
	b.PutBool(s.IsPasswordPending)
	b.PutInt32(s.APIID)
	b.PutString(s.ApplicationName)
	b.PutString(s.ApplicationVersion)
	b.PutBool(s.IsOfficialApplication)
	b.PutString(s.DeviceModel)
	b.PutString(s.Platform)
	b.PutString(s.SystemVersion)
	b.PutInt32(s.LogInDate)
	b.PutInt32(s.LastActiveDate)
	b.PutString(s.IP)
	b.PutString(s.Country)
	b.PutString(s.Region)
	return nil
}

// Decode implements bin.Decoder.
func (s *Session) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode session#727950d8 to nil")
	}
	if err := b.ConsumeID(SessionTypeID); err != nil {
		return fmt.Errorf("unable to decode session#727950d8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Session) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode session#727950d8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field is_current: %w", err)
		}
		s.IsCurrent = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field is_password_pending: %w", err)
		}
		s.IsPasswordPending = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field api_id: %w", err)
		}
		s.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field application_name: %w", err)
		}
		s.ApplicationName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field application_version: %w", err)
		}
		s.ApplicationVersion = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field is_official_application: %w", err)
		}
		s.IsOfficialApplication = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field device_model: %w", err)
		}
		s.DeviceModel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field platform: %w", err)
		}
		s.Platform = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field system_version: %w", err)
		}
		s.SystemVersion = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field log_in_date: %w", err)
		}
		s.LogInDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field last_active_date: %w", err)
		}
		s.LastActiveDate = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field ip: %w", err)
		}
		s.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field country: %w", err)
		}
		s.Country = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode session#727950d8: field region: %w", err)
		}
		s.Region = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Session) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode session#727950d8 as nil")
	}
	b.ObjStart()
	b.PutID("session")
	b.FieldStart("id")
	b.PutLong(s.ID)
	b.FieldStart("is_current")
	b.PutBool(s.IsCurrent)
	b.FieldStart("is_password_pending")
	b.PutBool(s.IsPasswordPending)
	b.FieldStart("api_id")
	b.PutInt32(s.APIID)
	b.FieldStart("application_name")
	b.PutString(s.ApplicationName)
	b.FieldStart("application_version")
	b.PutString(s.ApplicationVersion)
	b.FieldStart("is_official_application")
	b.PutBool(s.IsOfficialApplication)
	b.FieldStart("device_model")
	b.PutString(s.DeviceModel)
	b.FieldStart("platform")
	b.PutString(s.Platform)
	b.FieldStart("system_version")
	b.PutString(s.SystemVersion)
	b.FieldStart("log_in_date")
	b.PutInt32(s.LogInDate)
	b.FieldStart("last_active_date")
	b.PutInt32(s.LastActiveDate)
	b.FieldStart("ip")
	b.PutString(s.IP)
	b.FieldStart("country")
	b.PutString(s.Country)
	b.FieldStart("region")
	b.PutString(s.Region)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Session) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode session#727950d8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("session"); err != nil {
				return fmt.Errorf("unable to decode session#727950d8: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field id: %w", err)
			}
			s.ID = value
		case "is_current":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field is_current: %w", err)
			}
			s.IsCurrent = value
		case "is_password_pending":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field is_password_pending: %w", err)
			}
			s.IsPasswordPending = value
		case "api_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field api_id: %w", err)
			}
			s.APIID = value
		case "application_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field application_name: %w", err)
			}
			s.ApplicationName = value
		case "application_version":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field application_version: %w", err)
			}
			s.ApplicationVersion = value
		case "is_official_application":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field is_official_application: %w", err)
			}
			s.IsOfficialApplication = value
		case "device_model":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field device_model: %w", err)
			}
			s.DeviceModel = value
		case "platform":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field platform: %w", err)
			}
			s.Platform = value
		case "system_version":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field system_version: %w", err)
			}
			s.SystemVersion = value
		case "log_in_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field log_in_date: %w", err)
			}
			s.LogInDate = value
		case "last_active_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field last_active_date: %w", err)
			}
			s.LastActiveDate = value
		case "ip":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field ip: %w", err)
			}
			s.IP = value
		case "country":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field country: %w", err)
			}
			s.Country = value
		case "region":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode session#727950d8: field region: %w", err)
			}
			s.Region = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *Session) GetID() (value int64) {
	return s.ID
}

// GetIsCurrent returns value of IsCurrent field.
func (s *Session) GetIsCurrent() (value bool) {
	return s.IsCurrent
}

// GetIsPasswordPending returns value of IsPasswordPending field.
func (s *Session) GetIsPasswordPending() (value bool) {
	return s.IsPasswordPending
}

// GetAPIID returns value of APIID field.
func (s *Session) GetAPIID() (value int32) {
	return s.APIID
}

// GetApplicationName returns value of ApplicationName field.
func (s *Session) GetApplicationName() (value string) {
	return s.ApplicationName
}

// GetApplicationVersion returns value of ApplicationVersion field.
func (s *Session) GetApplicationVersion() (value string) {
	return s.ApplicationVersion
}

// GetIsOfficialApplication returns value of IsOfficialApplication field.
func (s *Session) GetIsOfficialApplication() (value bool) {
	return s.IsOfficialApplication
}

// GetDeviceModel returns value of DeviceModel field.
func (s *Session) GetDeviceModel() (value string) {
	return s.DeviceModel
}

// GetPlatform returns value of Platform field.
func (s *Session) GetPlatform() (value string) {
	return s.Platform
}

// GetSystemVersion returns value of SystemVersion field.
func (s *Session) GetSystemVersion() (value string) {
	return s.SystemVersion
}

// GetLogInDate returns value of LogInDate field.
func (s *Session) GetLogInDate() (value int32) {
	return s.LogInDate
}

// GetLastActiveDate returns value of LastActiveDate field.
func (s *Session) GetLastActiveDate() (value int32) {
	return s.LastActiveDate
}

// GetIP returns value of IP field.
func (s *Session) GetIP() (value string) {
	return s.IP
}

// GetCountry returns value of Country field.
func (s *Session) GetCountry() (value string) {
	return s.Country
}

// GetRegion returns value of Region field.
func (s *Session) GetRegion() (value string) {
	return s.Region
}
