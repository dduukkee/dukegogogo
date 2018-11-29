package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestGetMachineGame(t *testing.T) {
	type args struct {
		MID string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"111", args{"111"}, "YYYYYYY"},
		{"222", args{"222"}, "RRRRRRRr"},
		{"333", args{"333"}, "RRRRRRRr"},
		{"444", args{"444"}, "RRRRRRRr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMachineGame(tt.args.MID); got != tt.want {
				t.Errorf("GetMachineGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMachineGame(t *testing.T) {
	type args struct {
		UserID string
	}
	tests := []struct {
		name string
		args args
		want []MachineInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMachineGame(tt.args.UserID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMachineGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
