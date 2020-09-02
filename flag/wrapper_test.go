package flag

import (
	"bytes"
	"testing"

	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestBytesBase64WrapperVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v *wrappers.BytesValue
	BytesBase64WrapperVar(fs, &v, "foo", "")
	assert.NoError(t, fs.Set("foo", "YWJjZA=="))
	assert.NotNil(t, v)
	assert.Equal(t, "abcd", string(v.Value))
}

func TestBytesBase64WrapperVarStdIn(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v *wrappers.BytesValue
	BytesBase64WrapperVar(fs, &v, "foo", "")
	stdin = bytes.NewReader([]byte("abcd"))
	assert.NoError(t, fs.Set("foo", "-"))
	assert.NotNil(t, v)
	assert.Equal(t, "abcd", string(v.Value))
}

func TestBytesBase64WrapperSliceVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v []*wrappers.BytesValue
	BytesBase64WrapperSliceVar(fs, &v, "foo", "")
	assert.NoError(t, fs.Set("foo", "YWJjZA=="))
	assert.NoError(t, fs.Set("foo", "MTIzNA=="))
	assert.Len(t, v, 2)
	assert.Equal(t, "abcd", string(v[0].Value))
	assert.Equal(t, "1234", string(v[1].Value))
}
