package controllers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetLimOffError(t *testing.T) {
	cases := []struct {
		name string
		lim  string
		off  string
	}{
		{
			name: "test_1",
			lim:  "q",
			off:  "1",
		},
		{
			name: "test_2",
			lim:  "1",
			off:  "q",
		},
		{
			name: "test_3",
			lim:  "afs",
			off:  "asdfg",
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			_, _, err := GetLimOff(tCase.lim, tCase.off)
			require.Error(t, err)
		})
	}
}

func TestGetLimOff(t *testing.T) {
	cases := []struct {
		name string
		lim  string
		off  string
	}{
		{
			name: "test_1",
			lim:  "",
			off:  "",
		},
		{
			name: "test_2",
			lim:  "1",
			off:  "1",
		},
		{
			name: "test_3",
			lim:  "25",
			off:  "",
		},
		{
			name: "test_4",
			lim:  "",
			off:  "34545",
		},
	}
	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			_, _, err := GetLimOff(tCase.lim, tCase.off)
			require.NoError(t, err)
		})
	}
}
