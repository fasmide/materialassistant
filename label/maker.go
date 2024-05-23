package label

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"time"

	"github.com/fasmide/materialassistant/acs"
	"github.com/hako/durafmt"
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

//go:embed logo-simple.svg
var Logo []byte

//go:embed FreeSans.ttf
var FreeSans []byte

//go:embed FreeSansBold.ttf
var FreeSansBold []byte

type Maker struct {
	w float64
	h float64

	regular *canvas.FontFace
	bold    *canvas.FontFace
	logo    *canvas.Canvas

	durUnits durafmt.Units
}

func NewMaker(w, h float64) (*Maker, error) {
	font, err := canvas.LoadFont(FreeSans, 0, canvas.FontRegular)
	if err != nil {
		return nil, fmt.Errorf("unable to read font: %w", err)
	}

	fontBold, err := canvas.LoadFont(FreeSansBold, 0, canvas.FontRegular)
	if err != nil {
		return nil, fmt.Errorf("unable to read font: %w", err)
	}

	logo, err := canvas.ParseSVG(bytes.NewReader(Logo))
	if err != nil {
		return nil, fmt.Errorf("unable to parse logo: %w", err)
	}

	regular := font.Face(12, canvas.Black)
	bold := fontBold.Face(20, canvas.Black)

	units, err := durafmt.DefaultUnitsCoder.Decode("år:år,uge:uger,dag:dage,time:timer,minute:minutes,second:seconds,millisecond:millliseconds,microsecond:microsseconds")
	if err != nil {
		panic(err)
	}

	return &Maker{w: w, h: h,
		logo:     logo,
		regular:  regular,
		bold:     bold,
		durUnits: units,
	}, nil
}

func (m *Maker) DebugSVG() error {
	// 42 x 89mm custom, no scale seems to be the best settings
	c := canvas.New(89, 36)
	ctx := canvas.NewContext(c)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(89, 36))
	return renderers.Write("output2.png", c, canvas.DPMM(7.5))
}

func (m *Maker) MaterialSVG(who acs.Identity, d time.Duration, tag *canvas.Canvas) (image.Image, error) {
	c := canvas.New(m.w, m.h)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(m.w, m.h))

	m.logo.RenderViewTo(ctx, canvas.Identity.Translate(18, 0).Scale(0.13, 0.13).Rotate(90))

	ctx.DrawText(17.0, 30.0, canvas.NewTextLine(m.regular, fmt.Sprintf("Medlem %d:", who.ForLetID), canvas.Left))
	ctx.DrawText(17.0, 21.75, canvas.NewTextLine(m.bold, ellipticalTruncate(who.Name, 13), canvas.Left))

	ctx.DrawText(17.0, 14.0, canvas.NewTextLine(m.regular, fmt.Sprintf("Udløb %s, den:", durafmt.Parse(d).Format(m.durUnits)), canvas.Left))
	ctx.DrawText(17.0, 6.0, canvas.NewTextLine(m.bold, time.Now().Add(d).Format("2006-01-02"), canvas.Left))
	w, h := tag.Size()
	tag.RenderViewTo(ctx, canvas.Identity.Translate(m.w-(w), (m.h/2)-h/2))

	err := renderers.Write("output2.png", c, canvas.DPMM(8))

	return nil, err
}

func (m *Maker) TagUseAllowed() *canvas.Canvas {
	c := canvas.New(0, 0)
	ctx := canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(2.5, 2.5)
	ctx.DrawText(0, 3, canvas.NewTextLine(m.bold, "use", canvas.Center))

	ctx = canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1.2, 1.2)
	ctx.DrawText(0, 0, canvas.NewTextLine(m.bold, "allowed", canvas.Center))

	c.Fit(0)
	return c
}

func (m *Maker) TagDoNotHack() *canvas.Canvas {
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1.45, 1.45)
	ctx.DrawText(0, 7.5, canvas.NewTextLine(m.bold, "do not", canvas.Center))

	ctx = canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(2, 2)
	ctx.DrawText(0, 0, canvas.NewTextLine(m.bold, "hack", canvas.Center))

	c.Fit(0)
	return c
}

func (m *Maker) TagSkab() *canvas.Canvas {
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1, 1)
	ctx.DrawText(0, 7.5, canvas.NewTextLine(m.bold, "This skab", canvas.Center))

	ctx = canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(0.85, 0.85)
	ctx.DrawText(0, 0, canvas.NewTextLine(m.bold, "is my skab", canvas.Center))

	c.Fit(0)
	return c
}

func ellipticalTruncate(text string, maxLen int) string {
	lastSpaceIx := maxLen
	len := 0
	for range text {
		// if unicode.IsSpace(r) {
		//	lastSpaceIx = i
		// }
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}
