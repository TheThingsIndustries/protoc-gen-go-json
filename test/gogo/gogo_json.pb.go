// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v0.0.0-dev
// - protoc             v3.9.1
// source: gogo.proto

package test

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "github.com/TheThingsIndustries/protoc-gen-go-json/test/types"
)

// MarshalProtoJSON marshals the MessageWithGoGoOptions message to JSON.
func (x *MessageWithGoGoOptions) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.EUIWithCustomName) > 0 || s.HasField("eui_with_custom_name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("eui_with_custom_name")
		s.WriteBytes(x.EUIWithCustomName)
	}
	if x.EUIWithCustomNameAndType != nil || s.HasField("eui_with_custom_name_and_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("eui_with_custom_name_and_type")
		x.EUIWithCustomNameAndType.MarshalProtoJSON(s.WithField("eui_with_custom_name_and_type"))
	}
	if len(x.NonNullableEUIWithCustomNameAndType) > 0 || s.HasField("non_nullable_eui_with_custom_name_and_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("non_nullable_eui_with_custom_name_and_type")
		x.NonNullableEUIWithCustomNameAndType.MarshalProtoJSON(s.WithField("non_nullable_eui_with_custom_name_and_type"))
	}
	if len(x.EUIsWithCustomNameAndType) > 0 || s.HasField("euis_with_custom_name_and_type") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("euis_with_custom_name_and_type")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.EUIsWithCustomNameAndType {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("euis_with_custom_name_and_type"))
		}
		s.WriteArrayEnd()
	}
	if x.Duration != nil || s.HasField("duration") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("duration")
		if x.Duration == nil {
			s.WriteNil()
		} else {
			s.WriteDuration(*x.Duration)
		}
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("non_nullable_duration")
		s.WriteDuration(x.NonNullableDuration)
	}
	if x.Timestamp != nil || s.HasField("timestamp") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("timestamp")
		if x.Timestamp == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.Timestamp)
		}
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("non_nullable_timestamp")
		s.WriteTime(x.NonNullableTimestamp)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the MessageWithGoGoOptions to JSON.
func (x *MessageWithGoGoOptions) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the MessageWithGoGoOptions message from JSON.
func (x *MessageWithGoGoOptions) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "eui_with_custom_name", "euiWithCustomName":
			s.AddField("eui_with_custom_name")
			x.EUIWithCustomName = s.ReadBytes()
		case "eui_with_custom_name_and_type", "euiWithCustomNameAndType":
			s.AddField("eui_with_custom_name_and_type")
			if s.ReadNil() {
				x.EUIWithCustomNameAndType = nil
				return
			}
			x.EUIWithCustomNameAndType = &types.EUI64{}
			x.EUIWithCustomNameAndType.UnmarshalProtoJSON(s.WithField("eui_with_custom_name_and_type", false))
		case "non_nullable_eui_with_custom_name_and_type", "nonNullableEuiWithCustomNameAndType":
			s.AddField("non_nullable_eui_with_custom_name_and_type")
			x.NonNullableEUIWithCustomNameAndType.UnmarshalProtoJSON(s.WithField("non_nullable_eui_with_custom_name_and_type", false))
		case "euis_with_custom_name_and_type", "euisWithCustomNameAndType":
			s.AddField("euis_with_custom_name_and_type")
			if s.ReadNil() {
				x.EUIsWithCustomNameAndType = nil
				return
			}
			s.ReadArray(func() {
				var v types.EUI64
				v.UnmarshalProtoJSON(s.WithField("euis_with_custom_name_and_type", false))
				x.EUIsWithCustomNameAndType = append(x.EUIsWithCustomNameAndType, v)
			})
		case "duration":
			s.AddField("duration")
			if s.ReadNil() {
				x.Duration = nil
				return
			}
			v := s.ReadDuration()
			if s.Err() != nil {
				return
			}
			x.Duration = v
		case "non_nullable_duration", "nonNullableDuration":
			s.AddField("non_nullable_duration")
			v := s.ReadDuration()
			if s.Err() != nil {
				return
			}
			x.NonNullableDuration = *v
		case "timestamp":
			s.AddField("timestamp")
			if s.ReadNil() {
				x.Timestamp = nil
				return
			}
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.Timestamp = v
		case "non_nullable_timestamp", "nonNullableTimestamp":
			s.AddField("non_nullable_timestamp")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.NonNullableTimestamp = *v
		}
	})
}

// UnmarshalJSON unmarshals the MessageWithGoGoOptions from JSON.
func (x *MessageWithGoGoOptions) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the SubMessage message to JSON.
func (x *SubMessage) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Field != "" || s.HasField("field") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field")
		s.WriteString(x.Field)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the SubMessage to JSON.
func (x *SubMessage) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SubMessage message from JSON.
func (x *SubMessage) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "field":
			s.AddField("field")
			x.Field = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the SubMessage from JSON.
func (x *SubMessage) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the MessageWithNullable message to JSON.
func (x *MessageWithNullable) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("sub")
		x.Sub.MarshalProtoJSON(s.WithField("sub"))
	}
	if len(x.Subs) > 0 || s.HasField("subs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("subs")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Subs {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("subs"))
		}
		s.WriteArrayEnd()
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("other_sub")
		// NOTE: SubMessageWithoutMarshalers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.OtherSub)
	}
	if len(x.OtherSubs) > 0 || s.HasField("other_subs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("other_subs")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.OtherSubs {
			s.WriteMoreIf(&wroteElement)
			// NOTE: SubMessageWithoutMarshalers does not seem to implement MarshalProtoJSON.
			gogo.MarshalMessage(s, &element)
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the MessageWithNullable to JSON.
func (x *MessageWithNullable) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the MessageWithNullable message from JSON.
func (x *MessageWithNullable) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "sub":
			x.Sub.UnmarshalProtoJSON(s.WithField("sub", true))
		case "subs":
			s.AddField("subs")
			if s.ReadNil() {
				x.Subs = nil
				return
			}
			s.ReadArray(func() {
				v := SubMessage{}
				v.UnmarshalProtoJSON(s.WithField("subs", false))
				if s.Err() != nil {
					return
				}
				x.Subs = append(x.Subs, v)
			})
		case "other_sub", "otherSub":
			s.AddField("other_sub")
			// NOTE: SubMessageWithoutMarshalers does not seem to implement UnmarshalProtoJSON.
			var v SubMessageWithoutMarshalers
			gogo.UnmarshalMessage(s, &v)
			x.OtherSub = v
		case "other_subs", "otherSubs":
			s.AddField("other_subs")
			if s.ReadNil() {
				x.OtherSubs = nil
				return
			}
			s.ReadArray(func() {
				// NOTE: SubMessageWithoutMarshalers does not seem to implement UnmarshalProtoJSON.
				var v SubMessageWithoutMarshalers
				gogo.UnmarshalMessage(s, &v)
				x.OtherSubs = append(x.OtherSubs, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the MessageWithNullable from JSON.
func (x *MessageWithNullable) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the MessageWithEmbedded message to JSON.
func (x *MessageWithEmbedded) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.SubMessage != nil || s.HasField("sub") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("sub")
		x.SubMessage.MarshalProtoJSON(s.WithField("sub"))
	}
	if x.SubMessageWithoutMarshalers != nil || s.HasField("other_sub") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("other_sub")
		// NOTE: SubMessageWithoutMarshalers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.SubMessageWithoutMarshalers)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the MessageWithEmbedded to JSON.
func (x *MessageWithEmbedded) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the MessageWithEmbedded message from JSON.
func (x *MessageWithEmbedded) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "sub":
			if s.ReadNil() {
				x.SubMessage = nil
				return
			}
			x.SubMessage = &SubMessage{}
			x.SubMessage.UnmarshalProtoJSON(s.WithField("sub", true))
		case "other_sub", "otherSub":
			s.AddField("other_sub")
			if s.ReadNil() {
				x.SubMessageWithoutMarshalers = nil
				return
			}
			// NOTE: SubMessageWithoutMarshalers does not seem to implement UnmarshalProtoJSON.
			var v SubMessageWithoutMarshalers
			gogo.UnmarshalMessage(s, &v)
			x.SubMessageWithoutMarshalers = &v
		}
	})
}

// UnmarshalJSON unmarshals the MessageWithEmbedded from JSON.
func (x *MessageWithEmbedded) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the MessageWithNullableEmbedded message to JSON.
func (x *MessageWithNullableEmbedded) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("sub")
		x.SubMessage.MarshalProtoJSON(s.WithField("sub"))
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("other_sub")
		// NOTE: SubMessageWithoutMarshalers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.SubMessageWithoutMarshalers)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the MessageWithNullableEmbedded to JSON.
func (x *MessageWithNullableEmbedded) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the MessageWithNullableEmbedded message from JSON.
func (x *MessageWithNullableEmbedded) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "sub":
			x.SubMessage.UnmarshalProtoJSON(s.WithField("sub", true))
		case "other_sub", "otherSub":
			s.AddField("other_sub")
			// NOTE: SubMessageWithoutMarshalers does not seem to implement UnmarshalProtoJSON.
			var v SubMessageWithoutMarshalers
			gogo.UnmarshalMessage(s, &v)
			x.SubMessageWithoutMarshalers = v
		}
	})
}

// UnmarshalJSON unmarshals the MessageWithNullableEmbedded from JSON.
func (x *MessageWithNullableEmbedded) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
