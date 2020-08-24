package flag

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestReflectMapVar(t *testing.T) {
	testCases := []struct {
		parser func(string) (interface{}, error)
		target interface{}
		val1   string
		val2   string
		want1  interface{}
		want2  interface{}
	}{
		{ParseBool, &map[bool]bool{}, "false=false,true=true", "false=true", &map[bool]bool{false: false, true: true}, &map[bool]bool{false: true, true: true}},
		{ParseInt32, &map[int32]int32{}, "1=2,3=4", "1=1,5=6", &map[int32]int32{1: 2, 3: 4}, &map[int32]int32{1: 1, 3: 4, 5: 6}},
		{ParseInt64, &map[int64]int64{}, "1=2,3=4", "1=1,5=6", &map[int64]int64{1: 2, 3: 4}, &map[int64]int64{1: 1, 3: 4, 5: 6}},
		{ParseUint32, &map[uint32]uint32{}, "1=2,3=4", "1=1,5=6", &map[uint32]uint32{1: 2, 3: 4}, &map[uint32]uint32{1: 1, 3: 4, 5: 6}},
		{ParseUint64, &map[uint64]uint64{}, "1=2,3=4", "1=1,5=6", &map[uint64]uint64{1: 2, 3: 4}, &map[uint64]uint64{1: 1, 3: 4, 5: 6}},
		{ParseFloat32, &map[float32]float32{}, "1=2,3=4", "1=1,5=6", &map[float32]float32{1: 2, 3: 4}, &map[float32]float32{1: 1, 3: 4, 5: 6}},
		{ParseFloat64, &map[float64]float64{}, "1=2,3=4", "1=1,5=6", &map[float64]float64{1: 2, 3: 4}, &map[float64]float64{1: 1, 3: 4, 5: 6}},
		{ParseString, &map[string]string{}, "a=b,c=d", "a=x,e=f", &map[string]string{"a": "b", "c": "d"}, &map[string]string{"a": "x", "c": "d", "e": "f"}},
		{ParseBytesBase64, &map[string][]byte{}, "a=Yg,c=ZA", "a=eA,e=Zg", &map[string][]byte{"a": []byte("b"), "c": []byte("d")}, &map[string][]byte{"a": []byte("x"), "c": []byte("d"), "e": []byte("f")}},
	}
	for _, tc := range testCases {
		fs := &pflag.FlagSet{}
		keyParser := tc.parser
		if _, ok := tc.target.(*map[string][]byte); ok {
			keyParser = ParseString
		}
		ReflectMapVar(fs, keyParser, tc.parser, "map", tc.target, "foo", "")
		assert.NoError(t, fs.Set("foo", tc.val1))
		assert.Equal(t, tc.want1, tc.target)
		assert.NoError(t, fs.Set("foo", tc.val2))
		assert.Equal(t, tc.want2, tc.target)
		assert.Equal(t, "map", fs.Lookup("foo").Value.Type())
	}
}
