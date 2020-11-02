package hackernews

import (
	"testing"
)

func Test_buildRequestURL(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0 args",
			args: args{[]string{}},
			want: "https://hacker-news.firebaseio.com/v0.json",
		},
		{
			name: "1 path arg",
			args: args{[]string{"path1"}},
			want: "https://hacker-news.firebaseio.com/v0/path1.json",
		},
		{
			name: "2 path args",
			args: args{[]string{"path1", "path2"}},
			want: "https://hacker-news.firebaseio.com/v0/path1/path2.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRequestURL(tt.args.paths...); got != tt.want {
				t.Errorf("buildRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_get(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no error",
			args:    args{"https://hacker-news.firebaseio.com/v0/item/8863.json"},
			wantErr: false,
		},
		{
			name:    "has error",
			args:    args{"thisisntarealurl"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
