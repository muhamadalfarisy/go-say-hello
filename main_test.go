package go_say_hello

import "testing"

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
			name: "test case 1",
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
