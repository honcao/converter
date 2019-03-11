// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.
package main

import (
	"log"
	"reflect"
)

// DeepAssignment dst and src should be the same type in different API version
// dst should be pointer type
func DeepAssignment(dst, src interface{}) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Fail to covert object", r)
		}
	}()
	dstValue := reflect.ValueOf(dst)
	srcValue := reflect.ValueOf(src)
	if dstValue.Kind() != reflect.Ptr {
		log.Fatal("The dst must be Ptr")
	}
	dstValue = dstValue.Elem()
	deepAssignmentInternal(dstValue, srcValue, 0, "")
}

func deepAssignmentInternal(dstValue, srcValue reflect.Value, depth int, path string) {
	if dstValue.CanSet() {
		switch srcValue.Kind() {
		case reflect.Bool:
			dstValue.SetBool(srcValue.Bool())
		case reflect.String:
			dstValue.SetString(srcValue.String())
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			dstValue.SetInt(srcValue.Int())
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			dstValue.SetUint(srcValue.Uint())
		case reflect.Float64, reflect.Float32:
			dstValue.SetFloat(srcValue.Float())
		case reflect.Complex64, reflect.Complex128:
			dstValue.SetComplex(srcValue.Complex())
		case reflect.Ptr:
			if !srcValue.IsNil() {
				d := reflect.New(dstValue.Type().Elem())
				dstValue.Set(d)
				deepAssignmentInternal(dstValue.Elem(), srcValue.Elem(), depth+1, "")
			}
		case reflect.Slice:
			if !srcValue.IsNil() {
				d := reflect.MakeSlice(dstValue.Type(), srcValue.Len(), srcValue.Cap())
				dstValue.Set(d)
				for i := 0; i < srcValue.Len(); i++ {
					v := dstValue.Index(i)
					deepAssignmentInternal(v, srcValue.Index(i), depth+1, "")
					v.Set(v)
				}
			}
		case reflect.Array:
			if !srcValue.IsNil() {
				d := reflect.New(dstValue.Type()).Elem()
				dstValue.Set(d)
				for i := 0; i < srcValue.Len(); i++ {
					v := reflect.New(srcValue.Index(i).Type()).Elem()
					deepAssignmentInternal(v, srcValue.Index(i), depth+1, "")
					d.Index(i).Set(v)
				}
			}

		case reflect.Map:
			if !srcValue.IsNil() {
				d := reflect.MakeMap(dstValue.Type())
				for _, key := range srcValue.MapKeys() {
					v := reflect.New(srcValue.MapIndex(key).Type()).Elem()
					deepAssignmentInternal(v, srcValue.MapIndex(key), depth+1, "")
					d.SetMapIndex(key, v)
				}
				dstValue.Set(d)
			}
		case reflect.Struct:
			for i := 0; i < srcValue.NumField(); i++ {
				srcField := srcValue.Field(i)
				dstField := dstValue.FieldByName(srcValue.Type().Field(i).Name)
				if dstField.IsValid() && dstField.CanAddr() && dstField.CanSet() {
					deepAssignmentInternal(dstField, srcField, depth+1, "")
				}
			}
		default:
		}
	}
}
