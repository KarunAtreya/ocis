// Package maputil provide map data util functions. eg: convert, sub-value get, simple merge
package maputil

import (
	"reflect"
	"strings"

	"github.com/gookit/goutil/arrutil"
)

// Key, value sep char consts
const (
	ValSepStr  = ","
	ValSepChar = ','
	KeySepStr  = "."
	KeySepChar = '.'
)

// SimpleMerge simple merge two data map by string key.
// will merge the src to dst map
func SimpleMerge(src, dst map[string]any) map[string]any {
	if len(src) == 0 {
		return dst
	}

	if len(dst) == 0 {
		return src
	}

	for key, val := range src {
		dst[key] = val
	}
	return dst
}

// func DeepMerge(src, dst map[string]any, deep int) map[string]any { TODO
// }

// MergeSMap simple merge two string map. merge src to dst map
func MergeSMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	return MergeStringMap(src, dst, ignoreCase)
}

// MergeStringMap simple merge two string map. merge src to dst map
func MergeStringMap(src, dst map[string]string, ignoreCase bool) map[string]string {
	if len(src) == 0 {
		return dst
	}
	if len(dst) == 0 {
		return src
	}

	for k, v := range src {
		if ignoreCase {
			k = strings.ToLower(k)
		}

		dst[k] = v
	}
	return dst
}

// MakeByPath build new value by key names
//
// Example:
//
//	"site.info"
//	->
//	map[string]any {
//		site: {info: val}
//	}
//
//	// case 2, last key is slice:
//	"site.tags[1]"
//	->
//	map[string]any {
//		site: {tags: [val]}
//	}
func MakeByPath(path string, val any) (mp map[string]any) {
	return MakeByKeys(strings.Split(path, KeySepStr), val)
}

// MakeByKeys build new value by key names
//
// Example:
//
//	// case 1:
//	[]string{"site", "info"}
//	->
//	map[string]any {
//		site: {info: val}
//	}
//
//	// case 2, last key is slice:
//	[]string{"site", "tags[1]"}
//	->
//	map[string]any {
//		site: {tags: [val]}
//	}
func MakeByKeys(keys []string, val any) (mp map[string]any) {
	size := len(keys)

	// if last key contains slice index, make slice wrap the val
	lastKey := keys[size-1]
	if newK, idx, ok := parseArrKeyIndex(lastKey); ok {
		// valTyp := reflect.TypeOf(val)
		sliTyp := reflect.SliceOf(reflect.TypeOf(val))
		sliVal := reflect.MakeSlice(sliTyp, idx+1, idx+1)
		sliVal.Index(idx).Set(reflect.ValueOf(val))

		// update val and last key
		val = sliVal.Interface()
		keys[size-1] = newK
	}

	if size == 1 {
		return map[string]any{keys[0]: val}
	}

	// multi nodes
	arrutil.Reverse(keys)
	for _, p := range keys {
		if mp == nil {
			mp = map[string]any{p: val}
		} else {
			mp = map[string]any{p: mp}
		}
	}
	return
}
