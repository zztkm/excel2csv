package main

import "testing"

func TestCsvFileName(t *testing.T) {
	type args struct {
		prefix string
		sheet  string
		suffix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no suffix",
			args: args{prefix: "test", sheet: "sheet1", suffix: ""},
			want: "test_sheet1.csv",
		},
		{
			name: "no prefix",
			args: args{prefix: "", sheet: "sheet1", suffix: "test"},
			want: "sheet1_test.csv",
		},
		{
			name: "no prefix and suffix",
			args: args{prefix: "", sheet: "sheet1", suffix: ""},
			want: "sheet1.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := csvFileName(tt.args.prefix, tt.args.sheet, tt.args.suffix); got != tt.want {
				t.Errorf("csvFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}
