package codeext

import (
	"testing"

	"github.com/fluentassert/verify"
)

type setIfObj struct {
	value int
}

func TestTernTrue(t *testing.T) {
	res := Tern(true, 1, 2)
	verify.Number(res).Equal(1).Require(t)
}

func TestTernFalse(t *testing.T) {
	res := Tern(false, 1, 2)
	verify.Number(res).Equal(2).Require(t)
}

func TestTernExprTrue(t *testing.T) {
	res := Tern(1 > 0, 1, 2)
	verify.Number(res).Equal(1).Require(t)
}

func TestTernExprFalse(t *testing.T) {
	res := Tern(1 < 0, 1, 2)
	verify.Number(res).Equal(2).Require(t)
}

func TestSetIfTrue(t *testing.T) {
	obj := setIfObj{
		value: 0,
	}
	SetIf(&obj.value, true, 1)
	verify.Number(obj.value).Equal(1).Require(t)
}

func TestSetIfFalse(t *testing.T) {
	obj := setIfObj{
		value: 0,
	}
	SetIf(&obj.value, false, 1)
	verify.Number(obj.value).Equal(0).Require(t)
}

func TestSetIfExprTrue(t *testing.T) {
	obj := setIfObj{
		value: 0,
	}
	SetIf(&obj.value, 1 > 0, 1)
	verify.Number(obj.value).Equal(1).Require(t)
}

func TestSetIfExprFalse(t *testing.T) {
	obj := setIfObj{
		value: 0,
	}
	SetIf(&obj.value, 1 < 0, 1)
	verify.Number(obj.value).Equal(0).Require(t)
}
