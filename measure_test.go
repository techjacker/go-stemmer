package main

import (
	"reflect"
	"testing"
)

func Test_isConsonant(t *testing.T) {
	tests := []struct {
		name string
		args []rune
		want []bool
	}{
		{
			name: "isConsonant",
			args: []rune("syzygy"),
			want: []bool{true, false, true, false, true, false},
		},
		{
			name: "isConsonant",
			args: []rune("toy"),
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args {
				if got := isConsonant(tt.args, i); got != tt.want[i] {
					t.Errorf("rune: %v, isConsonant() = %v, want %v", string(v), got, tt.want[i])
				}
			}
		})
	}
}

func Test_genCVSequence(t *testing.T) {
	type args struct {
		in [][]byte
	}
	type want struct {
		seq  [][]byte
		size int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "genCVSequence() measure = 0",
			args: args{
				in: [][]byte{[]byte("tr"), []byte("ee"), []byte("tree"), []byte("y"), []byte("by")},
			},
			want: want{
				seq:  [][]byte{[]byte("cc"), []byte("vv"), []byte("ccvv"), []byte("c"), []byte("cv")},
				size: 0,
			},
		},
		{
			name: "genCVSequence() measure = 1",
			args: args{
				in: [][]byte{[]byte("trouble"), []byte("oats"), []byte("trees"), []byte("ivy")},
			},
			want: want{
				seq:  [][]byte{[]byte("ccvvccv"), []byte("vvcc"), []byte("ccvvc"), []byte("vcv")},
				size: 1,
			},
		},
		{
			name: "genCVSequence() measure = 1",
			args: args{
				in: [][]byte{[]byte("troubles"), []byte("private"), []byte("oaten"), []byte("orrery")},
			},
			want: want{
				seq:  [][]byte{[]byte("ccvvccvc"), []byte("ccvcvcv"), []byte("vvcvc"), []byte("vccvcv")},
				size: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.args.in {
				got := genCVSequence(v)
				if !reflect.DeepEqual(got, tt.want.seq[i]) {
					t.Errorf("in: %v, genCVSequence() = %v, want %v", string(v), string(got), string(tt.want.seq[i]))
				}
				if m := measureSequence(got); m != tt.want.size {
					t.Errorf("in: %v, measureSequence() = %d, want %d", string(v), m, tt.want.size)
				}
			}
		})
	}
}
