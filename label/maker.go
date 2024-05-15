package label

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

//go:embed logo-simple.svg
var Logo []byte

type Maker struct {
	MmWidth  int
	MmHeight int
}

func (m *Maker) MaterialSVG(who string) (image.Image, error) {
	font, err := canvas.LoadSystemFont("sans-serif", canvas.FontRegular)
	if err != nil {
		return nil, fmt.Errorf("unable to read font: %w", err)
	}
	cLogo, err := canvas.ParseSVG(bytes.NewReader(Logo))
	if err != nil {
		return nil, fmt.Errorf("unable to parse logo: %w", err)
	}

	c := canvas.New(float64(m.MmWidth), float64(m.MmHeight))
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(float64(m.MmWidth), float64(m.MmHeight)))
	cLogo.RenderViewTo(ctx, canvas.Identity.Translate(150, 0).Scale(1.5, 1.5).Rotate(90))

	face := font.Face(128.0, canvas.Black)
	ctx.DrawText(0, 0, canvas.NewTextLine(face, "Lorem ip  sum", canvas.Left))
	ctx.Push()
	err = renderers.Write("output2.png", c, canvas.DPMM(1.0))

	return nil, err
}
