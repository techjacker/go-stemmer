package main

import (
	"reflect"
	"testing"
)

func Test_stemmer_replaceSuffix_Run(t *testing.T) {
	type args struct {
		in [][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Following combinations",
			args: args{
				in: [][]byte{[]byte("following")},
			},
			want: []byte("following"),
		},
		{
			name: "Friend combinations",
			args: args{
				in: [][]byte{[]byte("Friends"), []byte("friendlier"), []byte("friendlies"), []byte("friendly")},
			},
			want: []byte("friend"),
		},
		{
			name: "Class combinations",
			args: args{
				in: [][]byte{[]byte("class"), []byte("classes"), []byte("classify")},
			},
			want: []byte("class"),
		},
		{
			name: "Classification combinations",
			args: args{
				in: [][]byte{[]byte("classification")},
			},
			want: []byte("classification"),
		},
		{
			name: "Flower combinations",
			args: args{
				in: [][]byte{[]byte("Flowery"), []byte("flowers"), []byte("flower")},
			},
			want: []byte("flower"),
		},
		{
			name: "Flow combinations",
			args: args{
				in: [][]byte{[]byte("flow"), []byte("flows")},
			},
			want: []byte("flow"),
		},
		{
			name: "Stop words should return empty byte slice",
			args: args{
				in: [][]byte{[]byte("through"), []byte("are"), []byte("and"), []byte("then"), []byte("the")},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, in := range tt.args.in {
				// don't need to for Run() - should do it itself
				if got := Stem(in); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Stem(%v) got = %q, want %q", string(in), got, tt.want)
				}
			}
		})
	}
}

func Test_isStopWord(t *testing.T) {
	type args struct {
		in [][]byte
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Removes stop words",
			args: args{
				in: [][]byte{[]byte("through"), []byte("are"), []byte("and"), []byte("then"), []byte("the"), []byte("notastopword"), []byte("long")},
			},
			want: []bool{true, true, true, true, true, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args.in {
				if got := isStopWord(v); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("isStopWord(%s) = %v, want: %v", string(v), got, tt.want[i])
				}
			}
		})
	}
}
