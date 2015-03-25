package push

import (
	"fmt"
	"testing"
)

func Test_Audience(t *testing.T) {
	a := NewAudience()
	a.SetTag([]string{"深圳", "广州", "北京", "北京", "北京"})
	fmt.Println("SetTag:", a.Value())

	a.SetTagAnd([]string{"深圳", "女"})
	fmt.Println("SetTagAnd:", a.Value())

	a.SetAlias([]string{"4314", "892", "4531"})
	fmt.Println("SetAlias:", a.Value())

	a.SetRegistrationId([]string{"4312kjklfds2", "8914afd2", "45fdsa31"})
	fmt.Println("SetRegistrationId:", a.Value())

	a.All()
	fmt.Println("All:", a.Value())
}
