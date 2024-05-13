package namer

import "testing"

func Test_UserName(t *testing.T) {
	type fields struct {
		LastName  string
		FirstName string
		Age       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test1", fields{"田中", "裕二", 35}, "田中裕二,35歳"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				u := &User{
					LastName:  tt.fields.LastName,
					FirstName: tt.fields.FirstName,
					Age:       tt.fields.Age,
				}
				if got := u.UserName(); got != tt.want {
					t.Errorf("User.UserName() = %v, want %v", got, tt.want)
				}
			})
		})
	}
}
