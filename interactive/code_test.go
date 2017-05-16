package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = [][]byte{
	[]byte(`{"type":"method","method":"auth","arguments":[1,1,"fc3f865c156f32cac0755cde007654a8"],"id":0}`),
	[]byte(`{"type":"reply","error":null,"id":0,"data":{"authenticated":true,"roles":["Owner"]}}`),
	[]byte(`{"type":"method","method":"msg","arguments":["Hello world :)"],"id":2}`),
	[]byte(`{"type":"reply","error":null,"id":2,"data":{"channel":1,"id":"6351f9e0-3bf2-11e6-a3b3-bdc62094c158","user_name":"connor","user_id":1,"user_roles":["Owner"],"message":{"message":[{"type":"text","data":"Hello world ","text":"Hello world!"}],"meta":{}}}}`),
}

func TestGzipRT(t *testing.T) {
	code := NewGzip()
	for _, x := range data {
		in, err := code.Encode(x)
		assert.Nil(t, err)
		out, err := code.Decode(in)
		assert.Nil(t, err)
		assert.Equal(t, x, out)
	}
}
func TestLZ4RT(t *testing.T) {
	code := NewLZ4()
	for _, x := range data {
		in, err := code.Encode(x)
		assert.Nil(t, err)
		out, err := code.Decode(in)
		assert.Nil(t, err)
		assert.Equal(t, x, out)
	}
}
