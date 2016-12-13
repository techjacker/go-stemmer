package main

import (
	"reflect"
	"testing"
)

func TestCountStems(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{

			name: "Flowery text",
			args: args{
				in: []byte("Friends are friendlier friendlies that are friendly and classify the friendly classification class. Flowery flowers flow through following the flower flows"),
			},
			want: map[string]int{
				"following":      1,
				"flow":           2,
				"classification": 1,
				"class":          2,
				"flower":         3,
				"friend":         5,
			},
			wantErr: false,
		},
		{

			name: "Flowery text",
			args: args{
				in: []byte("Now is the winter of our discontent, Made glorious summer by this sun of York."),
			},
			want: map[string]int{
				"summer":  1,
				"sun":     1,
				"york":    1,
				"now":     1,
				"winter":  1,
				"discont": 1,
				"made":    1,
				"glori":   1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountStems(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountStems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountStems() = %v, want %v", got, tt.want)
			}
		})
	}
}
