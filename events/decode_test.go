package events

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	type Foo struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}
	assert := assert.New(t)
	jsonHeader := http.Header{}
	jsonHeader.Set("Content-Type", "application/json")
	var jsonOut Foo
	assert.NoError(UnmarshalJSON(bytes.NewBuffer([]byte(`{"a": "a", "b": "b", "c": "c"}`)), &jsonOut))
	assert.Equal(Foo{A: "a", B: "b", C: "c"}, jsonOut)
	formHeader := http.Header{}
	formHeader.Set("Content-Type", "application/x-www-form-urlencoded")
	var formOut Foo
	assert.NoError(UnmarshalForm(bytes.NewBuffer([]byte(`a=a&b=b&c=c`)), &formOut))
	assert.Equal(Foo{A: "a", B: "b", C: "c"}, formOut)
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	b.ReportAllocs()
	type Foo struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}
	reader := bytes.NewReader([]byte(`{"a": "a", "b": "b", "c": "c"}`))
	var jsonOut Foo
	for n := 0; n < b.N; n++ {
		reader.Seek(0, io.SeekStart)
		UnmarshalJSON(reader, &jsonOut)
	}
}

func BenchmarkUnmarshalForm(b *testing.B) {
	b.ReportAllocs()
	type Foo struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}
	reader := bytes.NewReader([]byte(`a=a&b=b&c=c`))
	var formOut Foo
	for n := 0; n < b.N; n++ {
		reader.Seek(0, io.SeekStart)
		UnmarshalForm(reader, &formOut)
	}
}
