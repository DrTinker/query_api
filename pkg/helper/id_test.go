package helper

import (
	"testing"
)

func TestUUid(t *testing.T) {
	type args struct {
		name  string
		phone int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				name:  "DrTinker",
				phone: 13759896733,
			},
		},
		{
			name: "case2",
			args: args{
				name:  "meepo",
				phone: 13572851091,
			},
		},
		{
			name: "case3",
			args: args{
				name:  "lpl",
				phone: 13769752310,
			},
		},
		{
			name: "case4",
			args: args{
				name:  "lzc",
				phone: 18743256301,
			},
		},
		{
			name: "case4",
			args: args{
				name:  "pop",
				phone: 73154689751,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenUid(tt.args.name, tt.args.phone)
			t.Errorf("id: %v\n", got)
		})
	}
}
