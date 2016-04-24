// +build netgo

package gl

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/webgl"
	"honnef.co/go/js/dom"
)

type Texture struct{ int }
type Buffer struct{ *js.Object }
type FrameBuffer struct{ int }
type RenderBuffer struct{ int }
type Program struct{ int }
type UniformLocation struct{ int }
type Shader struct{ int }
type Context struct {
	webgl.Context
}

func NewContext(canvas *js.Object) (gl *webgl.Context) {
	attrs := webgl.DefaultAttributes()
	//attrs.Alpha = false
	var err error
	gl, err = webgl.NewContext(canvas, attrs)
	if err != nil {
		dom.GetWindow().Alert("Error: " + err.Error())
	}
	return
}
