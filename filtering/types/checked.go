package types

import (
	"google.golang.org/protobuf/reflect/protoreflect"

	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

var (
	// CheckedPrimitives map from proto field descriptor type to expr.Type.
	CheckedPrimitives = map[protoreflect.Kind]*exprpb.Type{
		protoreflect.BoolKind:     checkedBool,
		protoreflect.BytesKind:    checkedBytes,
		protoreflect.DoubleKind:   checkedDouble,
		protoreflect.FloatKind:    checkedDouble,
		protoreflect.Int32Kind:    checkedInt,
		protoreflect.Int64Kind:    checkedInt,
		protoreflect.Sint32Kind:   checkedInt,
		protoreflect.Sint64Kind:   checkedInt,
		protoreflect.Uint32Kind:   checkedUint,
		protoreflect.Uint64Kind:   checkedUint,
		protoreflect.Fixed32Kind:  checkedUint,
		protoreflect.Fixed64Kind:  checkedUint,
		protoreflect.Sfixed32Kind: checkedInt,
		protoreflect.Sfixed64Kind: checkedInt,
		protoreflect.StringKind:   checkedString}

	// CheckedWellKnowns map from qualified proto type name to expr.Type for
	// well-known proto types.
	CheckedWellKnowns = map[string]*exprpb.Type{
		// Wrapper types.
		"google.protobuf.BoolValue":   checkedWrap(checkedBool),
		"google.protobuf.BytesValue":  checkedWrap(checkedBytes),
		"google.protobuf.DoubleValue": checkedWrap(checkedDouble),
		"google.protobuf.FloatValue":  checkedWrap(checkedDouble),
		"google.protobuf.Int64Value":  checkedWrap(checkedInt),
		"google.protobuf.Int32Value":  checkedWrap(checkedInt),
		"google.protobuf.UInt64Value": checkedWrap(checkedUint),
		"google.protobuf.UInt32Value": checkedWrap(checkedUint),
		"google.protobuf.StringValue": checkedWrap(checkedString),
		// Well-known types.
		"google.protobuf.Any":       checkedAny,
		"google.protobuf.Duration":  checkedDuration,
		"google.protobuf.Timestamp": checkedTimestamp,
		// Json types.
		"google.protobuf.ListValue": checkedListDyn,
		"google.protobuf.NullValue": checkedNull,
		"google.protobuf.Struct":    checkedMapStringDyn,
		"google.protobuf.Value":     checkedDyn,
	}

	// common types
	checkedDyn = &exprpb.Type{TypeKind: &exprpb.Type_Dyn{Dyn: &emptypb.Empty{}}}
	// Wrapper and primitive types.
	checkedBool   = checkedPrimitive(exprpb.Type_BOOL)
	checkedBytes  = checkedPrimitive(exprpb.Type_BYTES)
	checkedDouble = checkedPrimitive(exprpb.Type_DOUBLE)
	checkedInt    = checkedPrimitive(exprpb.Type_INT64)
	checkedString = checkedPrimitive(exprpb.Type_STRING)
	checkedUint   = checkedPrimitive(exprpb.Type_UINT64)
	// Well-known type equivalents.
	checkedAny       = checkedWellKnown(exprpb.Type_ANY)
	checkedDuration  = checkedWellKnown(exprpb.Type_DURATION)
	checkedTimestamp = checkedWellKnown(exprpb.Type_TIMESTAMP)
	// Json-based type equivalents.
	checkedNull = &exprpb.Type{
		TypeKind: &exprpb.Type_Null{
			Null: structpb.NullValue_NULL_VALUE}}
	checkedListDyn = &exprpb.Type{
		TypeKind: &exprpb.Type_ListType_{
			ListType: &exprpb.Type_ListType{ElemType: checkedDyn}}}
	checkedMapStringDyn = &exprpb.Type{
		TypeKind: &exprpb.Type_MapType_{
			MapType: &exprpb.Type_MapType{
				KeyType:   checkedString,
				ValueType: checkedDyn}}}
)
