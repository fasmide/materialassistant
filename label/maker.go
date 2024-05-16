package label

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"time"
	"unicode"

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

	regular := font.Face(128.0, canvas.Black)
	bold := fontBold.Face(200.0, canvas.Black)

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

func (m *Maker) MaterialSVG(who string, d time.Duration, tag *canvas.Canvas) (image.Image, error) {
	c := canvas.New(m.w, m.h)
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(m.w, m.h))
	m.logo.RenderViewTo(ctx, canvas.Identity.Translate(150, 0).Scale(1.5, 1.5).Rotate(90))

	ctx.DrawText(170, 350, canvas.NewTextLine(m.regular, "Medlem:", canvas.Left))
	ctx.DrawText(200, 260, canvas.NewTextLine(m.bold, ellipticalTruncate(who, 20), canvas.Left))

	ctx.DrawText(170, 180, canvas.NewTextLine(m.regular, fmt.Sprintf("Udløb %s, den:", durafmt.Parse(d).Format(m.durUnits)), canvas.Left))
	ctx.DrawText(200, 90, canvas.NewTextLine(m.bold, time.Now().Add(d).Format("2006-01-02 15:04:05"), canvas.Left))
	w, h := tag.Size()
	tag.RenderViewTo(ctx, canvas.Identity.Translate(m.w-(w+50), (m.h/2)-h/2))
	err := renderers.Write("output2.png", c, canvas.DPMM(1.0))

	return nil, err
}

func (m *Maker) TagUseAllowed() *canvas.Canvas {
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(2.3, 2.3)
	ctx.DrawText(0, 0, canvas.NewTextLine(m.bold, "use", canvas.Center))

	ctx = canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1.1, 1.1)
	ctx.DrawText(0, -70, canvas.NewTextLine(m.bold, "allowed", canvas.Center))

	c.Fit(0)
	return c
}

func (m *Maker) TagDoNotHack() *canvas.Canvas {
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1.45, 1.45)
	ctx.DrawText(0, 0, canvas.NewTextLine(m.bold, "do not", canvas.Center))

	ctx = canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1.7, 1.7)
	ctx.DrawText(0, -65, canvas.NewTextLine(m.bold, "hack", canvas.Center))

	c.Fit(0)
	return c
}

func (m *Maker) TagSkab() *canvas.Canvas {
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)
	ctx.Rotate(90)
	ctx.Scale(1, 1)
	ctx.DrawText(0, 65, canvas.NewTextLine(m.bold, "This skab", canvas.Center))

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
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}
