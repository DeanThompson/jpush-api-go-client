package push

import (
	"testing"

	"github.com/jukylin/jpush-api-go-client/common"
)

func Test_has(t *testing.T) {
	p := NewPlatform()

	// 不用 Add 方法，直接赋值
	p.value = []string{"all"}

	if !p.has("ios") {
		t.Error("`all` platform should contain `ios`")
	}

	p.value = []string{"ios", "android"}
	if !p.has("android") {
		t.Error("platform should has `android`")
	}
	if p.has("winphone") {
		t.Error("platform should not has `winphone`")
	}
}

func Test_Add(t *testing.T) {
	p := NewPlatform()

	err := p.Add("invalid")
	if err == nil || err != common.ErrInvalidPlatform {
		t.Error("p.Add should return ErrInvalidPlatform")
	}

	ps := []string{"ios", "android"}
	err = p.Add(ps...)
	if err != nil {
		t.Error("p.Add should return no error")
	}

	if !common.EqualStringSlice(p.value, ps) {
		t.Error("platforms not the same as added")
	}
}

func Test_Value(t *testing.T) {
	p := NewPlatform()

	p.Add("all", "ios")
	v, ok := p.Value().(string)
	if !ok {
		t.Error("p.Value should return a string type")
	}
	if v != "all" {
		t.Error("p.Value should return `all`")
	}

	p.value = nil

	added := []string{"winphone", "android"}
	p.Add(added...)
	ps, ok := p.Value().([]string)
	if !ok {
		t.Error("p.Value should return a slice of string")
	}
	if !common.EqualStringSlice(added, ps) {
		t.Errorf("p.Value should return %v", added)
	}
}

func Test_All(t *testing.T) {
	p := NewPlatform()

	p.Add("ios", "android")

	p.All()

	if !common.EqualStringSlice([]string{"all"}, p.value) {
		t.Error("All() does not work correctly")
	}

	v, ok := p.Value().(string)
	if !ok {
		t.Error("p.Value should return a string type")
	}
	if v != "all" {
		t.Error("p.Value should return `all`")
	}
}
