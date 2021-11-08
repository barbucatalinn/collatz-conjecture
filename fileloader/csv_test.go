//go:build !integration
// +build !integration

package fileloader

import (
	"reflect"
	"testing"
)

func TestReadCSVFile(t *testing.T) {
	t.Parallel()

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				filename: "../testdata/test.csv",
			},
			want: [][]string{
				[]string{"2"},
				[]string{"4"},
				[]string{"6"},
				[]string{"8"},
				[]string{"10"},
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				filename: "../testdata/test-ne.csv",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCSVFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCSVFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCSVFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
