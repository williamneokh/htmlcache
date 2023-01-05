package htmlcache

import (
	"testing"
)

func Test_cache_CreateTemplateCache(t *testing.T) {
	type fields struct {
		PagePath   string
		LayoutPath string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{{
		name: "Test 1",
		fields: fields{
			PagePath:   "./templates/*.page.gohtml",
			LayoutPath: "./templates/*.layout.gohtml",
		},
		want:    true,
		wantErr: false,
	},
		{
			name: "Test 2",
			fields: fields{
				PagePath:   "",
				LayoutPath: "./templates/*.layout.gohtml",
			},
			want:    true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := cache{
				PagePath:   tt.fields.PagePath,
				LayoutPath: tt.fields.LayoutPath,
			}
			got, err := c.CreateTemplateCache()

			if (err != nil) != tt.wantErr {
				t.Errorf("cache.CreateTemplateCache() error = %v, wantErr %v", (err != nil), tt.wantErr)
				return
			}
			if (len(got) > 0) != tt.want {

				t.Errorf("cache.CreateTemplateCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
