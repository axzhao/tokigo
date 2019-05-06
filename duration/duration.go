package duration

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

const (
	hour = time.Minute * 60
	day  = time.Minute * 60 * 24
	year = 365 * day
)

type Duration struct {
	time.Duration

	// var b strings.Builder
	// var b bytes.Buffer below Go 1.10
	FormatBuffer bytes.Buffer
}

func (d Duration) Format() string {

	switch {
	case d.Duration >= year:
		d.formatYears()
		fallthrough
	case d.Duration >= day:
		d.formatDays()
		fallthrough
	case d.Duration >= hour:
		d.formatHours()
		fallthrough
	default:
		d.formatMinutes()
	}

	return d.FormatBuffer.String()
}

func (d *Duration) formatMinutes() {
	m, _ := strconv.Atoi(fmt.Sprintf("%.0f", d.Minutes()))
	if m == 0 {
		return
	}
	fmt.Fprintf(&d.FormatBuffer, "%dm", m)
}

func (d *Duration) formatHours() {
	if d.Hours() == 0 {
		return
	}
	hours := d.Duration / hour
	d.Duration -= hours * hour
	fmt.Fprintf(&d.FormatBuffer, "%dh", hours)
}

func (d *Duration) formatDays() {
	days := d.Duration / day
	d.Duration -= days * day
	fmt.Fprintf(&d.FormatBuffer, "%dd", days)
}

func (d *Duration) formatYears() {
	years := d.Duration / year
	fmt.Fprintf(&d.FormatBuffer, "%dy", years)
	d.Duration -= years * year
}
