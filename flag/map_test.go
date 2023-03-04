package flag

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestMapVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	p := make(map[string]float32)
	MapVar[string, float32](fs, ParseStringE, ParseFloat32E, &p, "foo", "")
	assert.NoError(t, fs.Set("foo", "a=1.1,b=2.2"))
	assert.Equal(t, map[string]float32{"a": 1.1, "b": 2.2}, p)
	assert.NoError(t, fs.Set("foo", "c=3.3,d=4.4"))
	assert.Equal(t, map[string]float32{"a": 1.1, "b": 2.2, "c": 3.3, "d": 4.4}, p)
	v := fs.Lookup("foo").Value
	assert.Equal(t, "map", v.Type())
	assert.Equal(t, "{}", v.String())
}

func TestMapVarJSON(t *testing.T) {
	fs := &pflag.FlagSet{}
	p := make(map[string]float32)
	MapVar[string, float32](fs, ParseStringE, ParseFloat32E, &p, "foo", "")
	assert.NoError(t, fs.Set("foo", `{"a":1.1,"b":2.2}`))
	assert.Equal(t, map[string]float32{"a": 1.1, "b": 2.2}, p)
	assert.NoError(t, fs.Set("foo", `{"c":3.3,"d":4.4}`))
	assert.Equal(t, map[string]float32{"a": 1.1, "b": 2.2, "c": 3.3, "d": 4.4}, p)
	v := fs.Lookup("foo").Value
	assert.Equal(t, "map", v.Type())
	assert.Equal(t, "{}", v.String())
}
