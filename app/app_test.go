//go:build !integration
// +build !integration

package app

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *App
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				filename: "../testdata/test.csv",
			},
			want:    &App{data: []uint64{2, 4, 6, 8, 10}},
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
			got, err := New(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_loadData(t *testing.T) {
	t.Parallel()

	type fields struct {
		data []uint64
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				data: []uint64{2, 4, 6, 8, 10},
			},
			args: args{
				filename: "../testdata/test.csv",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				data: nil,
			},
			args: args{
				filename: "../testdata/test-ne.csv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				data: tt.fields.data,
			}
			if err := app.loadData(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("loadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_CalculateCollatzConjectureSteps(t *testing.T) {
	type fields struct {
		data []uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   map[uint64]int
	}{
		{
			name: "92",
			fields: fields{
				data: []uint64{92},
			},
			want: map[uint64]int{92: 19},
		},
		{
			name: "98",
			fields: fields{
				data: []uint64{98},
			},
			want: map[uint64]int{98: 27},
		},
		{
			name: "92, 98",
			fields: fields{
				data: []uint64{92, 98},
			},
			want: map[uint64]int{92: 19, 98: 27},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				data: tt.fields.data,
			}
			if got := app.CalculateCollatzConjectureSteps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateCollatzConjectureSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
