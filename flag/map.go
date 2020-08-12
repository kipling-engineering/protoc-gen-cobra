package flag

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

func ReflectMapVar(fs *pflag.FlagSet, keyParser, valParser func(val string) (interface{}, error), p interface{}, name, usage string) {
	v := reflect.ValueOf(p)
	if !v.IsValid() || v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Map {
		panic("must be a pointer to a map")
	}
	fs.Var(&reflectMapValue{value: v, keyParser: keyParser, valParser: valParser}, name, usage)
}

func ParseBool(val string) (interface{}, error) { return strconv.ParseBool(val) }

func ParseInt32(val string) (interface{}, error) {
	if i, err := strconv.ParseInt(val, 10, 32); err != nil {
		return nil, err
	} else {
		return int32(i), nil
	}
}

func ParseInt64(val string) (interface{}, error) { return strconv.ParseInt(val, 10, 64) }

func ParseUint32(val string) (interface{}, error) {
	if i, err := strconv.ParseUint(val, 10, 32); err != nil {
		return nil, err
	} else {
		return uint32(i), nil
	}
}

func ParseUint64(val string) (interface{}, error) { return strconv.ParseUint(val, 10, 64) }

func ParseFloat32(val string) (interface{}, error) {
	if i, err := strconv.ParseFloat(val, 32); err != nil {
		return nil, err
	} else {
		return float32(i), nil
	}
}

func ParseFloat64(val string) (interface{}, error) { return strconv.ParseFloat(val, 64) }

func ParseString(val string) (interface{}, error) { return val, nil }

func ParseBytesBase64(val string) (interface{}, error) { return base64.StdEncoding.DecodeString(val) }

type reflectMapValue struct {
	value     reflect.Value
	changed   bool
	keyParser func(val string) (interface{}, error)
	valParser func(val string) (interface{}, error)
}

func (s *reflectMapValue) Set(val string) error {
	ss := strings.Split(val, ",")
	v := s.value.Elem()
	out := reflect.MakeMapWithSize(v.Type(), len(ss))
	for _, pair := range ss {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("%s must be formatted as key=value", pair)
		}
		if k, err := s.keyParser(kv[0]); err != nil {
			return err
		} else if v, err := s.valParser(kv[1]); err != nil {
			return err
		} else {
			out.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
	}
	if !s.changed {
		v.Set(out)
		s.changed = true
	} else {
		iter := out.MapRange()
		for iter.Next() {
			v.SetMapIndex(iter.Key(), iter.Value())
		}
	}
	return nil
}

func (s *reflectMapValue) Type() string { return "ReflectMap" }

func (s *reflectMapValue) String() string { return "[]" }
