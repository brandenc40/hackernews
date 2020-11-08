package hackernews

import (
	"testing"
)

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

func TestGetUpates(t *testing.T) {
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
			_, err := GetUpates()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUpates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetStories(t *testing.T) {
	type args struct {
		storyType StoryType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "No Error",
			args:    args{storyType: StoriesTop},
			wantErr: false,
		},
		{
			name:    "No Error",
			args:    args{storyType: StoriesNew},
			wantErr: false,
		},
		{
			name:    "No Error",
			args:    args{storyType: StoriesBest},
			wantErr: false,
		},
		{
			name:    "No Error",
			args:    args{storyType: StoriesAsk},
			wantErr: false,
		},
		{
			name:    "No Error",
			args:    args{storyType: StoriesShow},
			wantErr: false,
		},
		{
			name:    "No Error",
			args:    args{storyType: StoriesJob},
			wantErr: false,
		},
		{
			name:    "No Error",
			args:    args{storyType: 6},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetStories(tt.args.storyType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetPaginatedStories(t *testing.T) {
	type args struct {
		storyType  StoryType
		limit      int
		pageNumber int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Request works properly",
			args:    args{StoriesBest, 10, 1},
			wantErr: false,
		},
		{
			name:    "Out of bounds no error",
			args:    args{StoriesBest, 50, 1000},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPaginatedStories(tt.args.storyType, tt.args.limit, tt.args.pageNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPaginatedStories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.limit != got.Limit {
				t.Error("GetPaginatedStories() returned an invalid length of pages")
			}
			if tt.args.pageNumber != got.PageNumber {
				t.Error("GetPaginatedStories() returned an invalid length of pages")
			}
		})
	}
}

func TestPaginatedStoriesResponse_HasNextpage(t *testing.T) {
	type fields struct {
		Stories      []Item
		Limit        int
		PageNumber   int
		TotalResults int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "Check for True",
			fields: fields{[]Item{}, 2, 3, 10},
			want:   true,
		},
		{
			name:   "Check for False",
			fields: fields{[]Item{}, 2, 5, 10},
			want:   false,
		},
		{
			name:   "Check for False with too large page num",
			fields: fields{[]Item{}, 2, 7, 10},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaginatedStoriesResponse{
				Stories:      tt.fields.Stories,
				Limit:        tt.fields.Limit,
				PageNumber:   tt.fields.PageNumber,
				TotalResults: tt.fields.TotalResults,
			}
			if got := p.HasNextpage(); got != tt.want {
				t.Errorf("PaginatedStoriesResponse.HasNextpage() = %v, want %v", got, tt.want)
			}
		})
	}
}
