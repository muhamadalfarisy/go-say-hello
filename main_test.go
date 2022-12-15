package go_say_hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	type args struct {
		nama string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test case 1:",
			args: args{
				nama: "1",
			},
			want: "Hello Worlds1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayHello(tt.args.nama); got != tt.want {
				t.Errorf("SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddOne(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{num: 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddOne(tt.args.num); got != tt.want {
				t.Errorf("AddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
