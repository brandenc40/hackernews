package hackernews

import (
	"reflect"
	"testing"
)

func TestBuildRequestURL(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Works properly",
			args: args{paths: []string{"user", "123"}},
			want: "https://hacker-news.firebaseio.com/v0/user/123.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildRequestURL(tt.args.paths...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildRequestURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	type args struct {
		itemID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{itemID: 8863},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetItem(tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Success",
			args:    args{userID: "jl"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetUser(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetTopStories(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "No error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTopStories()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopStories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestHydrateItems(t *testing.T) {
	type args struct {
		itemIDs []int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "No Error",
			args:    args{itemIDs: []int{8952, 9224}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HydrateItems(tt.args.itemIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("HydrateItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetMaxItem(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "No Error",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetMaxItem()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMaxItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
