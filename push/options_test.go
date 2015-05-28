package push

import (
	"encoding/json"
	"testing"
)

func Test_Option(t *testing.T) {
	option := NewOptions()
	data, _ := json.Marshal(option)
	if string(data) != `{"apns_production":false}` {
		t.Error("apns_production should be false by default ")
	}

	option.ApnsProduction = true
	data, _ = json.Marshal(option)
	if string(data) != `{"apns_production":true}` {
		t.Error("apns_production should be true after setted")
	}
}
