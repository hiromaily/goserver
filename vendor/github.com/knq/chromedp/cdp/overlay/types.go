package overlay

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// HighlightConfig configuration data for the highlighting of page elements.
type HighlightConfig struct {
	ShowInfo           bool      `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowRulers         bool      `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines bool      `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial  bool      `json:"displayAsMaterial,omitempty"`
	ContentColor       *cdp.RGBA `json:"contentColor,omitempty"`     // The content box highlight fill color (default: transparent).
	PaddingColor       *cdp.RGBA `json:"paddingColor,omitempty"`     // The padding highlight fill color (default: transparent).
	BorderColor        *cdp.RGBA `json:"borderColor,omitempty"`      // The border highlight fill color (default: transparent).
	MarginColor        *cdp.RGBA `json:"marginColor,omitempty"`      // The margin highlight fill color (default: transparent).
	EventTargetColor   *cdp.RGBA `json:"eventTargetColor,omitempty"` // The event target element highlight fill color (default: transparent).
	ShapeColor         *cdp.RGBA `json:"shapeColor,omitempty"`       // The shape outside fill color (default: transparent).
	ShapeMarginColor   *cdp.RGBA `json:"shapeMarginColor,omitempty"` // The shape margin fill color (default: transparent).
	SelectorList       string    `json:"selectorList,omitempty"`     // Selectors to highlight relevant nodes.
	CSSGridColor       *cdp.RGBA `json:"cssGridColor,omitempty"`     // The grid layout color (default: transparent).
}

// InspectMode [no description].
type InspectMode string

// String returns the InspectMode as string value.
func (t InspectMode) String() string {
	return string(t)
}

// InspectMode values.
const (
	InspectModeSearchForNode        InspectMode = "searchForNode"
	InspectModeSearchForUAShadowDOM InspectMode = "searchForUAShadowDOM"
	InspectModeNone                 InspectMode = "none"
)

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t InspectMode) MarshalEasyJSON(out *jwriter.Writer) {
	out.String(string(t))
}

// MarshalJSON satisfies json.Marshaler.
func (t InspectMode) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *InspectMode) UnmarshalEasyJSON(in *jlexer.Lexer) {
	switch InspectMode(in.String()) {
	case InspectModeSearchForNode:
		*t = InspectModeSearchForNode
	case InspectModeSearchForUAShadowDOM:
		*t = InspectModeSearchForUAShadowDOM
	case InspectModeNone:
		*t = InspectModeNone

	default:
		in.AddError(errors.New("unknown InspectMode value"))
	}
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *InspectMode) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
