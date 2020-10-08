package output

import (
	. "errorGen/types"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func CreatePlot(Silly ErrorSet) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	xticks := plot.TimeTicks{Format: "2006-01-02"}

	p.Title.Text = Silly.SloDef.Name
	p.Legend.Top = true
	p.Legend.Left = false
	p.X.Label.Text = "Date"
	p.X.Tick.Marker = xticks
	p.Y.Label.Text = "Minutes"

	lines := make([]interface{}, 0)
	lines = append(lines, Silly.SloDef.Name)
	lines = append(lines, CreatePoints(Silly))
	lines = append(lines, "SLO in mins")
	lines = append(lines, AddSloLine(Silly))
	_ = plotutil.AddLinePoints(p, lines...)

	if err := p.Save(24*vg.Inch, 12*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func CreatePoints(Silly ErrorSet) plotter.XYs {
	pts := make([]plotter.XY, 0)

	var pt plotter.XY

	for i := range Silly.ErrorDays {
		t := Silly.ErrorDays[i].Date
		c := Silly.ErrorDays[i].ErrorBurnt

		pt.X = float64(t.Unix())
		pt.Y = c
		pts = append(pts, pt)
	}

	return pts
}

func AddSloLine(Silly ErrorSet) plotter.XYs {
	pts := make([]plotter.XY, 0)

	var pt plotter.XY

	for i := range Silly.ErrorDays {
		t := Silly.ErrorDays[i].Date
		c := (100 - Silly.SloDef.SloPrecent) * (float64(Silly.SloDef.Days) * 24 * 60)

		pt.X = float64(t.Unix())
		pt.Y = c
		pts = append(pts, pt)
	}

	return pts
}
