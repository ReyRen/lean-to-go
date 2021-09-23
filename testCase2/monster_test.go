package testCase2

import "testing"

func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "zhangsan",
		Age:   10,
		Skill: "xxxxx",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() err, exptect=%v, but=%v", true, res)
	}
	t.Logf("monster.Store() successful...")
}
func TestReStore(t *testing.T) {
	var monster Monster // = var monster = &Monster{}
	res := monster.reStore()
	if !res {
		t.Fatalf("monster.reStore() err, exptect=%v, but=%v", true, res)
	}

	if monster.Name != "zhangsan" {
		t.Fatalf("monster.reStore() NAME err, exptect=%v, but=%v", "zhangsan", monster.Name)
	}
	t.Logf("monster.reStore() successful...")
}
