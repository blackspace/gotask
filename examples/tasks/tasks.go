package tasks

import (
	"github.com/fogleman/gg"
	"bytes"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type HelloWorld struct {}

func (r *HelloWorld)Exec() interface{} {
	return "Hello World"
}

func NewHelloWorld() *HelloWorld {
	return &HelloWorld{}
}


////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type Image struct {}

func (i *Image)Exec() interface{} {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}

	buf:=bytes.NewBuffer(make([]byte,0,1<<10))

	dc.EncodePNG(buf)

	return buf.Bytes()
}

func NewImage() *Image {
	return &Image{}
}