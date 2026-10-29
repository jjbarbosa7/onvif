package onvif

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestParseXsdFloat(t *testing.T) {
	cases := []struct {
		in     string
		want   float64
		wantOK bool
	}{
		{"0.25", 0.25, true},
		{"-1", -1, true},
		{"1E2", 100, true},
		{"INF", 0, false},
		{"-INF", 0, false},
		{"inf", 0, false},
		{"+Inf", 0, false},
		{"Infinity", 0, false},
		{"NaN", 0, false},
		{"nan", 0, false},
		{"", 0, false},
		{"abc", 0, false},
	}
	for _, c := range cases {
		got, err := parseXsdFloat(c.in, 64)
		if got != c.want || (err == nil) != c.wantOK {
			t.Errorf("parseXsdFloat(%q) = %v, %v; want %v, ok=%v", c.in, got, err, c.want, c.wantOK)
		}
	}
}

// A camera reporting a preset position as NaN or INF (legal xs:float) left a
// preset JSON cannot encode, and with it the camera's whole stored status: the
// camera status list answered an empty body.
func TestDecodePresetsNonFinitePositionEncodes(t *testing.T) {
	body := `<tptz:GetPresetsResponse xmlns:tptz="http://www.onvif.org/ver20/ptz/wsdl">` +
		`<tptz:Preset token="1"><tt:Name>Home</tt:Name><tt:PTZPosition>` +
		`<tt:PanTilt x="0.5" y="-0.25" space="s"/><tt:Zoom x="0.1" space="s"/>` +
		`</tt:PTZPosition></tptz:Preset>` +
		`<tptz:Preset token="2"><tt:Name>Unset</tt:Name><tt:PTZPosition>` +
		`<tt:PanTilt x="NaN" y="INF" space="s"/><tt:Zoom x="-INF" space="s"/>` +
		`</tt:PTZPosition></tptz:Preset>` +
		`</tptz:GetPresetsResponse>`

	presets, err := (&Device{}).DecodePresets([]byte(fmt.Sprintf(soapEnvelope, body)))
	if err != nil {
		t.Fatalf("DecodePresets: %v", err)
	}
	if len(presets) != 2 {
		t.Fatalf("got %d presets, want 2", len(presets))
	}
	if p := presets[0].PTZPosition; p.PanTilt.X != 0.5 || p.PanTilt.Y != -0.25 || p.Zoom.X != 0.1 {
		t.Errorf("finite position = %+v, want x 0.5, y -0.25, zoom 0.1", p)
	}
	if p := presets[1].PTZPosition; p.PanTilt.X != 0 || p.PanTilt.Y != 0 || p.Zoom.X != 0 {
		t.Errorf("non-finite position = %+v, want zeros", p)
	}
	if _, err := json.Marshal(presets); err != nil {
		t.Errorf("presets do not encode: %v", err)
	}
}
