package merkel_test


import (
	"github.com/Shubham-Rasal/blockchain/merkel"
	"testing"
)


func TestInit(t *testing.T) {
	m := merkel.MerkelTree{}
	values := []string{"a", "b", "c", "d", "e", "f", "g"}
	m.Init(values)
	if len(m.Levels) != 1 {
		t.Error("Init failed")
	}
	if len(m.Levels[0]) != 7 {
		t.Error("Init failed")
	}
}

func TestBuild(t *testing.T) {
	m := merkel.MerkelTree{}
	values := []string{"a", "b", "c", "d", "e", "f", "g"}
	m.Init(values)
	m.Build()
	if len(m.Levels) != 2 {
		t.Error("Build failed")
	}
	if len(m.Levels[1]) != 4 {
		t.Error("Build failed")
	}
}

func TestInsert(t *testing.T) {
	m := merkel.MerkelTree{}
	values := []string{"a", "b", "c", "d", "e", "f", "g"}
	m.Init(values)
	m.Build()
	m.Insert("h")
	m.Insert("i")
	if len(m.Levels) != 3 {
		t.Error("Insert failed")
	}
	if len(m.Levels[2]) != 2 {
		t.Error("Insert failed")
	}
}

func TestDelete(t *testing.T) {
	m := merkel.MerkelTree{}
	values := []string{"a", "b", "c", "d", "e", "f", "g"}
	m.Init(values)
	m.Build()
	m.Delete("g")
	if len(m.Levels) != 2 {
		t.Error("Delete failed")
	}
	if len(m.Levels[1]) != 4 {
		t.Error("Delete failed")
	}
}



