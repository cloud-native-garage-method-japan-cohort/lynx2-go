package title

import "testing"

func Test_title_Get(t *testing.T) {

	tests := []struct {
		name string
		tr   *title
		want string
	}{
		{
			name: "正常系",
			tr:   &title{},
			want: "Lynx Search 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &title{}
			if got := tr.Get(); got != tt.want {
				t.Errorf("title.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
