/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package metrics

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_api "k8s.io/kubernetes/pkg/api"
	pkg4_resource "k8s.io/kubernetes/pkg/api/resource"
	pkg1_unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	pkg3_types "k8s.io/kubernetes/pkg/types"
	"reflect"
	"runtime"
	pkg5_inf "speter.net/go/exp/math/dec/inf"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg2_api.ObjectMeta
		var v1 pkg4_resource.Quantity
		var v2 pkg1_unversioned.TypeMeta
		var v3 pkg3_types.UID
		var v4 pkg5_inf.Dec
		var v5 time.Time
		_, _, _, _, _, _ = v0, v1, v2, v3, v4, v5
	}
}

func (x *RawNode) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [2]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(2)
			} else {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *RawNode) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym9 := z.DecBinary()
	_ = yym9
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct10 := r.ContainerType()
		if yyct10 == codecSelferValueTypeMap1234 {
			yyl10 := r.ReadMapStart()
			if yyl10 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl10, d)
			}
		} else if yyct10 == codecSelferValueTypeArray1234 {
			yyl10 := r.ReadArrayStart()
			if yyl10 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl10, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *RawNode) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys11Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys11Slc
	var yyhl11 bool = l >= 0
	for yyj11 := 0; ; yyj11++ {
		if yyhl11 {
			if yyj11 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys11Slc = r.DecodeBytes(yys11Slc, true, true)
		yys11 := string(yys11Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys11 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys11)
		} // end switch yys11
	} // end for yyj11
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *RawNode) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj14 int
	var yyb14 bool
	var yyhl14 bool = l >= 0
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	for {
		yyj14++
		if yyhl14 {
			yyb14 = yyj14 > l
		} else {
			yyb14 = r.CheckBreak()
		}
		if yyb14 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj14-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *RawPod) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym17 := z.EncBinary()
		_ = yym17
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep18 := !z.EncBinary()
			yy2arr18 := z.EncBasicHandle().StructToArray
			var yyq18 [2]bool
			_, _, _ = yysep18, yyq18, yy2arr18
			const yyr18 bool = false
			yyq18[0] = x.Kind != ""
			yyq18[1] = x.APIVersion != ""
			var yynn18 int
			if yyr18 || yy2arr18 {
				r.EncodeArrayStart(2)
			} else {
				yynn18 = 0
				for _, b := range yyq18 {
					if b {
						yynn18++
					}
				}
				r.EncodeMapStart(yynn18)
				yynn18 = 0
			}
			if yyr18 || yy2arr18 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq18[0] {
					yym20 := z.EncBinary()
					_ = yym20
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq18[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym21 := z.EncBinary()
					_ = yym21
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr18 || yy2arr18 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq18[1] {
					yym23 := z.EncBinary()
					_ = yym23
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq18[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym24 := z.EncBinary()
					_ = yym24
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr18 || yy2arr18 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *RawPod) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym25 := z.DecBinary()
	_ = yym25
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct26 := r.ContainerType()
		if yyct26 == codecSelferValueTypeMap1234 {
			yyl26 := r.ReadMapStart()
			if yyl26 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl26, d)
			}
		} else if yyct26 == codecSelferValueTypeArray1234 {
			yyl26 := r.ReadArrayStart()
			if yyl26 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl26, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *RawPod) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys27Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys27Slc
	var yyhl27 bool = l >= 0
	for yyj27 := 0; ; yyj27++ {
		if yyhl27 {
			if yyj27 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys27Slc = r.DecodeBytes(yys27Slc, true, true)
		yys27 := string(yys27Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys27 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		default:
			z.DecStructFieldNotFound(-1, yys27)
		} // end switch yys27
	} // end for yyj27
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *RawPod) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj30 int
	var yyb30 bool
	var yyhl30 bool = l >= 0
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj30++
	if yyhl30 {
		yyb30 = yyj30 > l
	} else {
		yyb30 = r.CheckBreak()
	}
	if yyb30 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	for {
		yyj30++
		if yyhl30 {
			yyb30 = yyj30 > l
		} else {
			yyb30 = r.CheckBreak()
		}
		if yyb30 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj30-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym33 := z.EncBinary()
		_ = yym33
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep34 := !z.EncBinary()
			yy2arr34 := z.EncBasicHandle().StructToArray
			var yyq34 [4]bool
			_, _, _ = yysep34, yyq34, yy2arr34
			const yyr34 bool = false
			yyq34[0] = x.Kind != ""
			yyq34[1] = x.APIVersion != ""
			yyq34[2] = true
			yyq34[3] = true
			var yynn34 int
			if yyr34 || yy2arr34 {
				r.EncodeArrayStart(4)
			} else {
				yynn34 = 0
				for _, b := range yyq34 {
					if b {
						yynn34++
					}
				}
				r.EncodeMapStart(yynn34)
				yynn34 = 0
			}
			if yyr34 || yy2arr34 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq34[0] {
					yym36 := z.EncBinary()
					_ = yym36
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq34[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym37 := z.EncBinary()
					_ = yym37
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr34 || yy2arr34 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq34[1] {
					yym39 := z.EncBinary()
					_ = yym39
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq34[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym40 := z.EncBinary()
					_ = yym40
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr34 || yy2arr34 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq34[2] {
					yy42 := &x.ObjectMeta
					yy42.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq34[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy43 := &x.ObjectMeta
					yy43.CodecEncodeSelf(e)
				}
			}
			if yyr34 || yy2arr34 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq34[3] {
					yy45 := &x.Metrics
					yy45.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq34[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metrics"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy46 := &x.Metrics
					yy46.CodecEncodeSelf(e)
				}
			}
			if yyr34 || yy2arr34 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym47 := z.DecBinary()
	_ = yym47
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct48 := r.ContainerType()
		if yyct48 == codecSelferValueTypeMap1234 {
			yyl48 := r.ReadMapStart()
			if yyl48 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl48, d)
			}
		} else if yyct48 == codecSelferValueTypeArray1234 {
			yyl48 := r.ReadArrayStart()
			if yyl48 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl48, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys49Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys49Slc
	var yyhl49 bool = l >= 0
	for yyj49 := 0; ; yyj49++ {
		if yyhl49 {
			if yyj49 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys49Slc = r.DecodeBytes(yys49Slc, true, true)
		yys49 := string(yys49Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys49 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_api.ObjectMeta{}
			} else {
				yyv52 := &x.ObjectMeta
				yyv52.CodecDecodeSelf(d)
			}
		case "metrics":
			if r.TryDecodeAsNil() {
				x.Metrics = Metrics{}
			} else {
				yyv53 := &x.Metrics
				yyv53.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys49)
		} // end switch yys49
	} // end for yyj49
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj54 int
	var yyb54 bool
	var yyhl54 bool = l >= 0
	yyj54++
	if yyhl54 {
		yyb54 = yyj54 > l
	} else {
		yyb54 = r.CheckBreak()
	}
	if yyb54 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj54++
	if yyhl54 {
		yyb54 = yyj54 > l
	} else {
		yyb54 = r.CheckBreak()
	}
	if yyb54 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj54++
	if yyhl54 {
		yyb54 = yyj54 > l
	} else {
		yyb54 = r.CheckBreak()
	}
	if yyb54 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_api.ObjectMeta{}
	} else {
		yyv57 := &x.ObjectMeta
		yyv57.CodecDecodeSelf(d)
	}
	yyj54++
	if yyhl54 {
		yyb54 = yyj54 > l
	} else {
		yyb54 = r.CheckBreak()
	}
	if yyb54 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Metrics = Metrics{}
	} else {
		yyv58 := &x.Metrics
		yyv58.CodecDecodeSelf(d)
	}
	for {
		yyj54++
		if yyhl54 {
			yyb54 = yyj54 > l
		} else {
			yyb54 = r.CheckBreak()
		}
		if yyb54 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj54-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *Pod) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym59 := z.EncBinary()
		_ = yym59
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep60 := !z.EncBinary()
			yy2arr60 := z.EncBasicHandle().StructToArray
			var yyq60 [4]bool
			_, _, _ = yysep60, yyq60, yy2arr60
			const yyr60 bool = false
			yyq60[0] = x.Kind != ""
			yyq60[1] = x.APIVersion != ""
			yyq60[2] = true
			yyq60[3] = true
			var yynn60 int
			if yyr60 || yy2arr60 {
				r.EncodeArrayStart(4)
			} else {
				yynn60 = 0
				for _, b := range yyq60 {
					if b {
						yynn60++
					}
				}
				r.EncodeMapStart(yynn60)
				yynn60 = 0
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq60[0] {
					yym62 := z.EncBinary()
					_ = yym62
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq60[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym63 := z.EncBinary()
					_ = yym63
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq60[1] {
					yym65 := z.EncBinary()
					_ = yym65
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq60[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym66 := z.EncBinary()
					_ = yym66
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq60[2] {
					yy68 := &x.ObjectMeta
					yy68.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq60[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy69 := &x.ObjectMeta
					yy69.CodecEncodeSelf(e)
				}
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq60[3] {
					yy71 := &x.Metrics
					yy71.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq60[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metrics"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy72 := &x.Metrics
					yy72.CodecEncodeSelf(e)
				}
			}
			if yyr60 || yy2arr60 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *Pod) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym73 := z.DecBinary()
	_ = yym73
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct74 := r.ContainerType()
		if yyct74 == codecSelferValueTypeMap1234 {
			yyl74 := r.ReadMapStart()
			if yyl74 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl74, d)
			}
		} else if yyct74 == codecSelferValueTypeArray1234 {
			yyl74 := r.ReadArrayStart()
			if yyl74 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl74, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *Pod) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys75Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys75Slc
	var yyhl75 bool = l >= 0
	for yyj75 := 0; ; yyj75++ {
		if yyhl75 {
			if yyj75 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys75Slc = r.DecodeBytes(yys75Slc, true, true)
		yys75 := string(yys75Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys75 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				x.Kind = string(r.DecodeString())
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				x.APIVersion = string(r.DecodeString())
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg2_api.ObjectMeta{}
			} else {
				yyv78 := &x.ObjectMeta
				yyv78.CodecDecodeSelf(d)
			}
		case "metrics":
			if r.TryDecodeAsNil() {
				x.Metrics = Metrics{}
			} else {
				yyv79 := &x.Metrics
				yyv79.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys75)
		} // end switch yys75
	} // end for yyj75
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *Pod) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj80 int
	var yyb80 bool
	var yyhl80 bool = l >= 0
	yyj80++
	if yyhl80 {
		yyb80 = yyj80 > l
	} else {
		yyb80 = r.CheckBreak()
	}
	if yyb80 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		x.Kind = string(r.DecodeString())
	}
	yyj80++
	if yyhl80 {
		yyb80 = yyj80 > l
	} else {
		yyb80 = r.CheckBreak()
	}
	if yyb80 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		x.APIVersion = string(r.DecodeString())
	}
	yyj80++
	if yyhl80 {
		yyb80 = yyj80 > l
	} else {
		yyb80 = r.CheckBreak()
	}
	if yyb80 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg2_api.ObjectMeta{}
	} else {
		yyv83 := &x.ObjectMeta
		yyv83.CodecDecodeSelf(d)
	}
	yyj80++
	if yyhl80 {
		yyb80 = yyj80 > l
	} else {
		yyb80 = r.CheckBreak()
	}
	if yyb80 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Metrics = Metrics{}
	} else {
		yyv84 := &x.Metrics
		yyv84.CodecDecodeSelf(d)
	}
	for {
		yyj80++
		if yyhl80 {
			yyb80 = yyj80 > l
		} else {
			yyb80 = r.CheckBreak()
		}
		if yyb80 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj80-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *Metrics) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym85 := z.EncBinary()
		_ = yym85
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep86 := !z.EncBinary()
			yy2arr86 := z.EncBasicHandle().StructToArray
			var yyq86 [3]bool
			_, _, _ = yysep86, yyq86, yy2arr86
			const yyr86 bool = false
			var yynn86 int
			if yyr86 || yy2arr86 {
				r.EncodeArrayStart(3)
			} else {
				yynn86 = 3
				for _, b := range yyq86 {
					if b {
						yynn86++
					}
				}
				r.EncodeMapStart(yynn86)
				yynn86 = 0
			}
			if yyr86 || yy2arr86 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy88 := &x.StartTime
				yym89 := z.EncBinary()
				_ = yym89
				if false {
				} else if z.HasExtensions() && z.EncExt(yy88) {
				} else if yym89 {
					z.EncBinaryMarshal(yy88)
				} else if !yym89 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy88)
				} else {
					z.EncFallback(yy88)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("start"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy90 := &x.StartTime
				yym91 := z.EncBinary()
				_ = yym91
				if false {
				} else if z.HasExtensions() && z.EncExt(yy90) {
				} else if yym91 {
					z.EncBinaryMarshal(yy90)
				} else if !yym91 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy90)
				} else {
					z.EncFallback(yy90)
				}
			}
			if yyr86 || yy2arr86 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy93 := &x.EndTime
				yym94 := z.EncBinary()
				_ = yym94
				if false {
				} else if z.HasExtensions() && z.EncExt(yy93) {
				} else if yym94 {
					z.EncBinaryMarshal(yy93)
				} else if !yym94 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy93)
				} else {
					z.EncFallback(yy93)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("end"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy95 := &x.EndTime
				yym96 := z.EncBinary()
				_ = yym96
				if false {
				} else if z.HasExtensions() && z.EncExt(yy95) {
				} else if yym96 {
					z.EncBinaryMarshal(yy95)
				} else if !yym96 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy95)
				} else {
					z.EncFallback(yy95)
				}
			}
			if yyr86 || yy2arr86 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Usage == nil {
					r.EncodeNil()
				} else {
					yym98 := z.EncBinary()
					_ = yym98
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Usage) {
					} else {
						h.encapi_ResourceList((pkg2_api.ResourceList)(x.Usage), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("usage"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Usage == nil {
					r.EncodeNil()
				} else {
					yym99 := z.EncBinary()
					_ = yym99
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Usage) {
					} else {
						h.encapi_ResourceList((pkg2_api.ResourceList)(x.Usage), e)
					}
				}
			}
			if yyr86 || yy2arr86 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *Metrics) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym100 := z.DecBinary()
	_ = yym100
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct101 := r.ContainerType()
		if yyct101 == codecSelferValueTypeMap1234 {
			yyl101 := r.ReadMapStart()
			if yyl101 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl101, d)
			}
		} else if yyct101 == codecSelferValueTypeArray1234 {
			yyl101 := r.ReadArrayStart()
			if yyl101 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl101, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *Metrics) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys102Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys102Slc
	var yyhl102 bool = l >= 0
	for yyj102 := 0; ; yyj102++ {
		if yyhl102 {
			if yyj102 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys102Slc = r.DecodeBytes(yys102Slc, true, true)
		yys102 := string(yys102Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys102 {
		case "start":
			if r.TryDecodeAsNil() {
				x.StartTime = pkg1_unversioned.Time{}
			} else {
				yyv103 := &x.StartTime
				yym104 := z.DecBinary()
				_ = yym104
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv103) {
				} else if yym104 {
					z.DecBinaryUnmarshal(yyv103)
				} else if !yym104 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv103)
				} else {
					z.DecFallback(yyv103, false)
				}
			}
		case "end":
			if r.TryDecodeAsNil() {
				x.EndTime = pkg1_unversioned.Time{}
			} else {
				yyv105 := &x.EndTime
				yym106 := z.DecBinary()
				_ = yym106
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv105) {
				} else if yym106 {
					z.DecBinaryUnmarshal(yyv105)
				} else if !yym106 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv105)
				} else {
					z.DecFallback(yyv105, false)
				}
			}
		case "usage":
			if r.TryDecodeAsNil() {
				x.Usage = nil
			} else {
				yyv107 := &x.Usage
				yyv107.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys102)
		} // end switch yys102
	} // end for yyj102
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *Metrics) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj108 int
	var yyb108 bool
	var yyhl108 bool = l >= 0
	yyj108++
	if yyhl108 {
		yyb108 = yyj108 > l
	} else {
		yyb108 = r.CheckBreak()
	}
	if yyb108 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.StartTime = pkg1_unversioned.Time{}
	} else {
		yyv109 := &x.StartTime
		yym110 := z.DecBinary()
		_ = yym110
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv109) {
		} else if yym110 {
			z.DecBinaryUnmarshal(yyv109)
		} else if !yym110 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv109)
		} else {
			z.DecFallback(yyv109, false)
		}
	}
	yyj108++
	if yyhl108 {
		yyb108 = yyj108 > l
	} else {
		yyb108 = r.CheckBreak()
	}
	if yyb108 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.EndTime = pkg1_unversioned.Time{}
	} else {
		yyv111 := &x.EndTime
		yym112 := z.DecBinary()
		_ = yym112
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv111) {
		} else if yym112 {
			z.DecBinaryUnmarshal(yyv111)
		} else if !yym112 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv111)
		} else {
			z.DecFallback(yyv111, false)
		}
	}
	yyj108++
	if yyhl108 {
		yyb108 = yyj108 > l
	} else {
		yyb108 = r.CheckBreak()
	}
	if yyb108 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Usage = nil
	} else {
		yyv113 := &x.Usage
		yyv113.CodecDecodeSelf(d)
	}
	for {
		yyj108++
		if yyhl108 {
			yyb108 = yyj108 > l
		} else {
			yyb108 = r.CheckBreak()
		}
		if yyb108 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj108-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encapi_ResourceList(v pkg2_api.ResourceList, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeMapStart(len(v))
	for yyk114, yyv114 := range v {
		z.EncSendContainerState(codecSelfer_containerMapKey1234)
		yym115 := z.EncBinary()
		_ = yym115
		if false {
		} else if z.HasExtensions() && z.EncExt(yyk114) {
		} else {
			r.EncodeString(codecSelferC_UTF81234, string(yyk114))
		}
		z.EncSendContainerState(codecSelfer_containerMapValue1234)
		yy116 := &yyv114
		yym117 := z.EncBinary()
		_ = yym117
		if false {
		} else if z.HasExtensions() && z.EncExt(yy116) {
		} else if !yym117 && z.IsJSONHandle() {
			z.EncJSONMarshal(yy116)
		} else {
			z.EncFallback(yy116)
		}
	}
	z.EncSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) decapi_ResourceList(v *pkg2_api.ResourceList, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv118 := *v
	yyl118 := r.ReadMapStart()
	yybh118 := z.DecBasicHandle()
	if yyv118 == nil {
		yyrl118, _ := z.DecInferLen(yyl118, yybh118.MaxInitLen, 40)
		yyv118 = make(map[pkg2_api.ResourceName]pkg4_resource.Quantity, yyrl118)
		*v = yyv118
	}
	var yymk118 pkg2_api.ResourceName
	var yymv118 pkg4_resource.Quantity
	var yymg118 bool
	if yybh118.MapValueReset {
		yymg118 = true
	}
	if yyl118 > 0 {
		for yyj118 := 0; yyj118 < yyl118; yyj118++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk118 = ""
			} else {
				yymk118 = pkg2_api.ResourceName(r.DecodeString())
			}

			if yymg118 {
				yymv118 = yyv118[yymk118]
			} else {
				yymv118 = pkg4_resource.Quantity{}
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv118 = pkg4_resource.Quantity{}
			} else {
				yyv120 := &yymv118
				yym121 := z.DecBinary()
				_ = yym121
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv120) {
				} else if !yym121 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv120)
				} else {
					z.DecFallback(yyv120, false)
				}
			}

			if yyv118 != nil {
				yyv118[yymk118] = yymv118
			}
		}
	} else if yyl118 < 0 {
		for yyj118 := 0; !r.CheckBreak(); yyj118++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk118 = ""
			} else {
				yymk118 = pkg2_api.ResourceName(r.DecodeString())
			}

			if yymg118 {
				yymv118 = yyv118[yymk118]
			} else {
				yymv118 = pkg4_resource.Quantity{}
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv118 = pkg4_resource.Quantity{}
			} else {
				yyv123 := &yymv118
				yym124 := z.DecBinary()
				_ = yym124
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv123) {
				} else if !yym124 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv123)
				} else {
					z.DecFallback(yyv123, false)
				}
			}

			if yyv118 != nil {
				yyv118[yymk118] = yymv118
			}
		}
	} // else len==0: TODO: Should we clear map entries?
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}
