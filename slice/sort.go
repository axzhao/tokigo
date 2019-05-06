package slice

import (
	"reflect"
	"sort"
)

// Int64s sorts a slice of int64s in increasing order.
func Int64s(a []int64) {
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
}

type sortable struct {
	sort  string
	value reflect.Value
}

// Sort sorts a slice of Struct with key in increasing order.
func Sort(data interface{}, key string) {
	sort.Sort(sortable{key, reflect.ValueOf(data)})
}

func (s sortable) Len() int { return s.value.Len() }

func (s sortable) Less(i, j int) bool {
	iv := s.fieldValue(i, s.sort)
	jv := s.fieldValue(j, s.sort)

	if iv != nil && jv != nil {
		switch v := iv.(type) {
		case int:
			return v < jv.(int)
		default:
			return false
		}
	} else if iv == nil {
		return false
	} else if jv == nil {
		return true
	} else {
		return false
	}
}

func (s sortable) Swap(i, j int) {
	iv := s.value.Index(i)
	jv := s.value.Index(j)
	tmp := iv.Interface()
	iv.Set(jv)
	jv.Set(reflect.ValueOf(tmp))
}

func (s sortable) fieldValue(index int, name string) interface{} {
	if index >= s.value.Len() {
		return nil
	}
	v := s.value.Index(index)
	if v.Kind() == reflect.Interface {
		v = reflect.ValueOf(v.Interface())
	}
	switch v.Kind() {
	case reflect.Struct:
		return v.FieldByName(name).Interface()
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(name)).Interface()
	default:
		panic(`unsupported kind: ` + v.Kind().String())
	}
}
