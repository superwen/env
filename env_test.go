package env

import (
	"testing"
)

func TestLoad(t *testing.T) {

}

func TestLoadMust(t *testing.T) {

}

func TestHas(t *testing.T) {
	Load()
	if actually := Has("TEST01"); actually != true {
		t.Errorf("Has: [%v], actually: [%v]", "TEST01", actually)
	}
}

func TestBoolDefault(t *testing.T) {
	Load()
	if actually := BoolDefault("TEST01", false); actually != true {
		t.Errorf("Has: [%v], actually: [%v]", "TEST01", actually)
	}
	if actually := BoolDefault("TEST02", false); actually != false {
		t.Errorf("Has: [%v], actually: [%v]", "TEST02", actually)
	}
	if actually := BoolDefault("TEST06", false); actually != false {
		t.Errorf("Has: [%v], actually: [%v]", "TEST06", actually)
	}
}

func TestStringDefault(t *testing.T) {

}

func TestStringMust(t *testing.T) {

}

func TestIntDefault(t *testing.T) {

}

func TestInt64Default(t *testing.T) {

}

func TestFloat32Default(t *testing.T) {

}

func TestFloat64Default(t *testing.T) {

}
