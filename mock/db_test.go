package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("Not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o2, o1)
	// m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)
	GetFromDB(m, "Sam")
	GetFromDB(m, "Tom")
	// if v := GetFormDB(m, "Tom"); v != -1 {
	// 	t.Fatal("expected -1, but got", v)
	// }
}
