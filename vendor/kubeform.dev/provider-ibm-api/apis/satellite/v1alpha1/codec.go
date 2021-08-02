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
		jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosConfig{}).Type1()):      LocationSpecCosConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosCredentials{}).Type1()): LocationSpecCosCredentialsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosConfig{}).Type1()):      LocationSpecCosConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosCredentials{}).Type1()): LocationSpecCosCredentialsCodec{},
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
type LocationSpecCosConfigCodec struct {
}

func (LocationSpecCosConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LocationSpecCosConfig)(ptr) == nil
}

func (LocationSpecCosConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LocationSpecCosConfig)(ptr)
	var objs []LocationSpecCosConfig
	if obj != nil {
		objs = []LocationSpecCosConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LocationSpecCosConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LocationSpecCosConfig)(ptr) = LocationSpecCosConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LocationSpecCosConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LocationSpecCosConfig)(ptr) = objs[0]
			} else {
				*(*LocationSpecCosConfig)(ptr) = LocationSpecCosConfig{}
			}
		} else {
			*(*LocationSpecCosConfig)(ptr) = LocationSpecCosConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LocationSpecCosConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LocationSpecCosConfig)(ptr) = obj
		} else {
			*(*LocationSpecCosConfig)(ptr) = LocationSpecCosConfig{}
		}
	default:
		iter.ReportError("decode LocationSpecCosConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type LocationSpecCosCredentialsCodec struct {
}

func (LocationSpecCosCredentialsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*LocationSpecCosCredentials)(ptr) == nil
}

func (LocationSpecCosCredentialsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*LocationSpecCosCredentials)(ptr)
	var objs []LocationSpecCosCredentials
	if obj != nil {
		objs = []LocationSpecCosCredentials{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosCredentials{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (LocationSpecCosCredentialsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*LocationSpecCosCredentials)(ptr) = LocationSpecCosCredentials{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []LocationSpecCosCredentials

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosCredentials{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*LocationSpecCosCredentials)(ptr) = objs[0]
			} else {
				*(*LocationSpecCosCredentials)(ptr) = LocationSpecCosCredentials{}
			}
		} else {
			*(*LocationSpecCosCredentials)(ptr) = LocationSpecCosCredentials{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj LocationSpecCosCredentials

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(LocationSpecCosCredentials{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*LocationSpecCosCredentials)(ptr) = obj
		} else {
			*(*LocationSpecCosCredentials)(ptr) = LocationSpecCosCredentials{}
		}
	default:
		iter.ReportError("decode LocationSpecCosCredentials", "unexpected JSON type")
	}
}
