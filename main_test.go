package main

import "testing"

func TestSub(t *testing.T) {
	type args []int
	tests := []struct {
		name    string
		args    args
		wantSub int
	}{
		// TODO: Add test cases.
		{"sub-1", args{1, 2, 3, 4}, -8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSub := Sub(tt.args...); gotSub != tt.wantSub {
				t.Errorf("Sub() = %v, want %v", gotSub, tt.wantSub)
			}
		})
	}
}
