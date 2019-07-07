// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"html/template"
	"strings"
	"time"
)

// Struct type ViewHelper -
type ViewHelper struct{}

func (this *ViewHelper) Flash(msg ...interface{}) (flash template.HTML) {
	m := msg[0].(*map[string]string)
	if m != nil {
		fMsg := *m
		flash = template.HTML("<div class='alert alert-" + fMsg["type"] + " role='alert'>" +
			"<span>" + fMsg["message"] + "</span>" +
			"</div>")
	}
	return
}

// FormatDate method - format a date
func (this *ViewHelper) FormatDate(t time.Time) string {
	layout := "2006-01-01"
	return t.Format(layout)
}

// ToLower method -
func (this *ViewHelper) ToLower(text string) string {
	return strings.ToLower(text)
}

// ToUpper method -
func (this *ViewHelper) ToUpper(text string) string {
	return strings.ToUpper(text)
}

// UCFirst method -
func (this *ViewHelper) UCFirst(text string) string {
	return strings.Title(text)
}
