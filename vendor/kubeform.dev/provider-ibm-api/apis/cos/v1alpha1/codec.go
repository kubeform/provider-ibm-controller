/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecActivityTracking{}).Type1()):  BucketSpecActivityTrackingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecArchiveRule{}).Type1()):       BucketSpecArchiveRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecMetricsMonitoring{}).Type1()): BucketSpecMetricsMonitoringCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecObjectVersioning{}).Type1()):  BucketSpecObjectVersioningCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecRetentionRule{}).Type1()):     BucketSpecRetentionRuleCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecActivityTracking{}).Type1()):  BucketSpecActivityTrackingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecArchiveRule{}).Type1()):       BucketSpecArchiveRuleCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecMetricsMonitoring{}).Type1()): BucketSpecMetricsMonitoringCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecObjectVersioning{}).Type1()):  BucketSpecObjectVersioningCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecRetentionRule{}).Type1()):     BucketSpecRetentionRuleCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type BucketSpecActivityTrackingCodec struct {
}

func (BucketSpecActivityTrackingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BucketSpecActivityTracking)(ptr) == nil
}

func (BucketSpecActivityTrackingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BucketSpecActivityTracking)(ptr)
	var objs []BucketSpecActivityTracking
	if obj != nil {
		objs = []BucketSpecActivityTracking{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecActivityTracking{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BucketSpecActivityTrackingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BucketSpecActivityTracking)(ptr) = BucketSpecActivityTracking{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BucketSpecActivityTracking

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecActivityTracking{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BucketSpecActivityTracking)(ptr) = objs[0]
			} else {
				*(*BucketSpecActivityTracking)(ptr) = BucketSpecActivityTracking{}
			}
		} else {
			*(*BucketSpecActivityTracking)(ptr) = BucketSpecActivityTracking{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BucketSpecActivityTracking

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecActivityTracking{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BucketSpecActivityTracking)(ptr) = obj
		} else {
			*(*BucketSpecActivityTracking)(ptr) = BucketSpecActivityTracking{}
		}
	default:
		iter.ReportError("decode BucketSpecActivityTracking", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BucketSpecArchiveRuleCodec struct {
}

func (BucketSpecArchiveRuleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BucketSpecArchiveRule)(ptr) == nil
}

func (BucketSpecArchiveRuleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BucketSpecArchiveRule)(ptr)
	var objs []BucketSpecArchiveRule
	if obj != nil {
		objs = []BucketSpecArchiveRule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecArchiveRule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BucketSpecArchiveRuleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BucketSpecArchiveRule)(ptr) = BucketSpecArchiveRule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BucketSpecArchiveRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecArchiveRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BucketSpecArchiveRule)(ptr) = objs[0]
			} else {
				*(*BucketSpecArchiveRule)(ptr) = BucketSpecArchiveRule{}
			}
		} else {
			*(*BucketSpecArchiveRule)(ptr) = BucketSpecArchiveRule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BucketSpecArchiveRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecArchiveRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BucketSpecArchiveRule)(ptr) = obj
		} else {
			*(*BucketSpecArchiveRule)(ptr) = BucketSpecArchiveRule{}
		}
	default:
		iter.ReportError("decode BucketSpecArchiveRule", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BucketSpecMetricsMonitoringCodec struct {
}

func (BucketSpecMetricsMonitoringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BucketSpecMetricsMonitoring)(ptr) == nil
}

func (BucketSpecMetricsMonitoringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BucketSpecMetricsMonitoring)(ptr)
	var objs []BucketSpecMetricsMonitoring
	if obj != nil {
		objs = []BucketSpecMetricsMonitoring{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecMetricsMonitoring{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BucketSpecMetricsMonitoringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BucketSpecMetricsMonitoring)(ptr) = BucketSpecMetricsMonitoring{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BucketSpecMetricsMonitoring

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecMetricsMonitoring{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BucketSpecMetricsMonitoring)(ptr) = objs[0]
			} else {
				*(*BucketSpecMetricsMonitoring)(ptr) = BucketSpecMetricsMonitoring{}
			}
		} else {
			*(*BucketSpecMetricsMonitoring)(ptr) = BucketSpecMetricsMonitoring{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BucketSpecMetricsMonitoring

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecMetricsMonitoring{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BucketSpecMetricsMonitoring)(ptr) = obj
		} else {
			*(*BucketSpecMetricsMonitoring)(ptr) = BucketSpecMetricsMonitoring{}
		}
	default:
		iter.ReportError("decode BucketSpecMetricsMonitoring", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BucketSpecObjectVersioningCodec struct {
}

func (BucketSpecObjectVersioningCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BucketSpecObjectVersioning)(ptr) == nil
}

func (BucketSpecObjectVersioningCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BucketSpecObjectVersioning)(ptr)
	var objs []BucketSpecObjectVersioning
	if obj != nil {
		objs = []BucketSpecObjectVersioning{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecObjectVersioning{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BucketSpecObjectVersioningCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BucketSpecObjectVersioning)(ptr) = BucketSpecObjectVersioning{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BucketSpecObjectVersioning

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecObjectVersioning{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BucketSpecObjectVersioning)(ptr) = objs[0]
			} else {
				*(*BucketSpecObjectVersioning)(ptr) = BucketSpecObjectVersioning{}
			}
		} else {
			*(*BucketSpecObjectVersioning)(ptr) = BucketSpecObjectVersioning{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BucketSpecObjectVersioning

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecObjectVersioning{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BucketSpecObjectVersioning)(ptr) = obj
		} else {
			*(*BucketSpecObjectVersioning)(ptr) = BucketSpecObjectVersioning{}
		}
	default:
		iter.ReportError("decode BucketSpecObjectVersioning", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BucketSpecRetentionRuleCodec struct {
}

func (BucketSpecRetentionRuleCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BucketSpecRetentionRule)(ptr) == nil
}

func (BucketSpecRetentionRuleCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BucketSpecRetentionRule)(ptr)
	var objs []BucketSpecRetentionRule
	if obj != nil {
		objs = []BucketSpecRetentionRule{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecRetentionRule{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BucketSpecRetentionRuleCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BucketSpecRetentionRule)(ptr) = BucketSpecRetentionRule{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BucketSpecRetentionRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecRetentionRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BucketSpecRetentionRule)(ptr) = objs[0]
			} else {
				*(*BucketSpecRetentionRule)(ptr) = BucketSpecRetentionRule{}
			}
		} else {
			*(*BucketSpecRetentionRule)(ptr) = BucketSpecRetentionRule{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BucketSpecRetentionRule

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BucketSpecRetentionRule{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BucketSpecRetentionRule)(ptr) = obj
		} else {
			*(*BucketSpecRetentionRule)(ptr) = BucketSpecRetentionRule{}
		}
	default:
		iter.ReportError("decode BucketSpecRetentionRule", "unexpected JSON type")
	}
}
