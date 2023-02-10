package flag

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestSliceVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	var p []string
	SliceVar[string](fs, ParseStringE, &p, "foo", "")
	assert.NoError(t, fs.Set("foo", "one,two"))
	assert.Equal(t, []string{"one", "two"}, p)
	assert.NoError(t, fs.Set("foo", "three,four"))
	assert.Equal(t, []string{"one", "two", "three", "four"}, p)
	v := fs.Lookup("foo").Value
	assert.Equal(t, "slice", v.Type())
	assert.Equal(t, "[]", v.String())
}
