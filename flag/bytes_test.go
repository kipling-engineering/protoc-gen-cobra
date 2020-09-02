package flag

import (
	"bytes"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestBytesBase64Var(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v []byte
	BytesBase64Var(fs, &v, "foo", "")
	assert.NoError(t, fs.Set("foo", "YWJjZA=="))
	assert.Equal(t, "abcd", string(v))
}

func TestBytesBase64VarStdIn(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v []byte
	BytesBase64Var(fs, &v, "foo", "")
	stdin = bytes.NewReader([]byte("abcd"))
	assert.NoError(t, fs.Set("foo", "-"))
	assert.Equal(t, "abcd", string(v))
}

func TestBytesBase64SliceVar(t *testing.T) {
	fs := &pflag.FlagSet{}
	var v [][]byte
	BytesBase64SliceVar(fs, &v, "foo", "")
	assert.NoError(t, fs.Set("foo", "YWJjZA=="))
	assert.NoError(t, fs.Set("foo", "MTIzNA=="))
	assert.Len(t, v, 2)
	assert.Equal(t, "abcd", string(v[0]))
	assert.Equal(t, "1234", string(v[1]))
}
