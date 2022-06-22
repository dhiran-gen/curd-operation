package implement

import (
	"testing"
)

func TestEmp_Greet(t *testing.T) {
	type fields struct {
		Name string
		DOB  string
		Noye int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{Name: "Niraj"},
			want: "Hello and Welcome Niraj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Emp{
				Name: tt.fields.Name,
				DOB:  tt.fields.DOB,
				Noye: tt.fields.Noye,
			}
			if got := e.Greet(); got != tt.want {
				t.Errorf("Emp.Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmp_Age(t *testing.T) {
	type fields struct {
		Name string
		DOB  string
		Noye int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{DOB: "2000-02-04" },
			want: 22,
		},
		{
			name: "success",
			fields: fields{DOB: "2023-02-04" },
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Emp{
				Name: tt.fields.Name,
				DOB:  tt.fields.DOB,
				Noye: tt.fields.Noye,
			}
			if got := e.Age(); got != tt.want {
				t.Errorf("Emp.Age() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmp_SalaryBonus(t *testing.T) {
	type fields struct {
		Name string
		DOB  string
		Noye int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "success",
			fields: fields{Noye: 2},
			want: 50000,
		},
		{
			name: "success",
			fields: fields{Noye: 1},
			want: 20000,
		},
		{
			name: "success",
			fields: fields{Noye: 5},
			want: 100000,
		},
		{
			name: "success",
			fields: fields{Noye: 10},
			want: 2000000,
		},
		{
			name: "success",
			fields: fields{Noye: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Emp{
				Name: tt.fields.Name,
				DOB:  tt.fields.DOB,
				Noye: tt.fields.Noye,
			}
			if got := e.SalaryBonus(); got != tt.want {
				t.Errorf("Emp.SalaryBonus() = %v, want %v", got, tt.want)
			}
		})
	}
}
