package iso316

import (
	"reflect"
	"sync"
	"testing"
)

func TestIso3166(t *testing.T) {
	tests := []struct {
		name string
		want ISO
	}{
		{
			name: "new iso3166",
			want: i,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := I3166(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Iso3166() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iso_ALPHA2(t *testing.T) {
	type fields struct {
		Alpha2 *alpha2
		Alpha3 *alpha3
	}
	tests := []struct {
		name   string
		fields fields
		want   Alpha2CI
	}{
		{
			name: "getting alpha2",
			fields: fields{
				Alpha2: &alpha2{},
			},
			want: &alpha2{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &iso{
				Alpha2: tt.fields.Alpha2,
				Alpha3: tt.fields.Alpha3,
			}
			if got := i.ALPHA2(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ALPHA2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_iso_ALPHA3(t *testing.T) {
	type fields struct {
		Alpha2 *alpha2
		Alpha3 *alpha3
	}
	tests := []struct {
		name   string
		fields fields
		want   Alpha3CI
	}{
		{
			name: "getting alpha3",
			fields: fields{
				Alpha3: &alpha3{},
			},
			want: &alpha3{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &iso{
				Alpha2: tt.fields.Alpha2,
				Alpha3: tt.fields.Alpha3,
			}
			if got := i.ALPHA3(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ALPHA3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alpha2_Exist(t *testing.T) {
	type fields struct {
		data map[string]*data
		mu   *sync.RWMutex
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "getting US",
			fields: fields{
				data: i.Alpha2.data,
				mu:   i.Alpha2.mu,
			},
			args: args{code: "US"},
			want: true,
		},
		{
			name: "getting us",
			fields: fields{
				data: i.Alpha2.data,
				mu:   i.Alpha2.mu,
			},
			args: args{code: "us"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := alpha2{
				data: tt.fields.data,
				mu:   tt.fields.mu,
			}
			if got := a.Exist(tt.args.code); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alpha3_Exist(t *testing.T) {
	type fields struct {
		data map[string]*data
		mu   *sync.RWMutex
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "getting USA",
			fields: fields{
				data: i.Alpha3.data,
				mu:   i.Alpha3.mu,
			},
			args: args{code: "USA"},
			want: true,
		},
		{
			name: "getting usa",
			fields: fields{
				data: i.Alpha3.data,
				mu:   i.Alpha3.mu,
			},
			args: args{code: "usa"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := alpha3{
				data: tt.fields.data,
				mu:   tt.fields.mu,
			}
			if got := a.Exist(tt.args.code); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alpha2_GetAlpha3(t *testing.T) {
	type fields struct {
		data map[string]*data
		mu   *sync.RWMutex
	}
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get alpha3 by alpha2 code",
			fields: fields{
				data: i.Alpha2.data,
				mu:   i.Alpha2.mu,
			},
			args:    args{code: "US"},
			want:    "USA",
			wantErr: false,
		},
		{
			name: "get alpha3 by invalid alpha2 code",
			fields: fields{
				data: i.Alpha2.data,
				mu:   i.Alpha2.mu,
			},
			args:    args{code: "us"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := alpha2{
				data: tt.fields.data,
				mu:   tt.fields.mu,
			}
			got, err := a.GetAlpha3(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAlpha3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAlpha3() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_alpha3_GetAlpha2(t *testing.T) {
	type fields struct {
		data map[string]*data
		mu   *sync.RWMutex
	}
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get alpha2 by alpha3 code",
			fields: fields{
				data: i.Alpha3.data,
				mu:   i.Alpha3.mu,
			},
			args:    args{code: "USA"},
			want:    "US",
			wantErr: false,
		},
		{
			name: "get alpha2 by invalid alpha3 code",
			fields: fields{
				data: i.Alpha3.data,
				mu:   i.Alpha3.mu,
			},
			args:    args{code: "usa"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := alpha3{
				data: tt.fields.data,
				mu:   tt.fields.mu,
			}
			got, err := a.GetAlpha2(tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAlpha2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAlpha2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
