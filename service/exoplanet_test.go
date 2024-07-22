package service

import (
	"NTTHomeTestDemo/model"
	"reflect"
	"testing"
)

const (
	validId   = "6e5d848c-6b78-44cc-9af1-c00b34982700"
	invalidId = "6e5d848c-6b78-44cc-9af1-c00b3498270"
)

var dummydata = map[string]model.Exoplanet{validId: {ID: validId, Name: "XO-2N b", Description: "Test",
	Distance: 200, Radius: 1.6, Mass: 0.62, Type: "GasGiant"}}

func TestExoplanetService_EstimateFuel(t *testing.T) {
	type fields struct {
		Exoplanets map[string]model.Exoplanet
	}
	type args struct {
		id       string
		noOfCrew int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "Planet Exist in dababse",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id:       validId,
				noOfCrew: 2,
			},
			want:    10485.760000000004,
			wantErr: false,
		},
		{
			name: "Planet not Exist in dababse",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id:       invalidId,
				noOfCrew: 2,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exo := &ExoplanetService{
				Exoplanets: tt.fields.Exoplanets,
			}
			got, err := exo.EstimateFuel(tt.args.id, tt.args.noOfCrew)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExoplanetService.EstimateFuel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExoplanetService.EstimateFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExoplanetService_DeleteExoplanet(t *testing.T) {
	type fields struct {
		Exoplanets map[string]model.Exoplanet
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Exoplanet
		wantErr bool
	}{
		{
			name: "Positive case ",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id: validId,
			},
			want:    dummydata[validId],
			wantErr: false,
		},
		{
			name: "Nagative case ",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id: invalidId,
			},
			want:    model.Exoplanet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exo := &ExoplanetService{
				Exoplanets: tt.fields.Exoplanets,
			}
			got, err := exo.DeleteExoplanet(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExoplanetService.DeleteExoplanet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExoplanetService.DeleteExoplanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExoplanetService_UpdateExoplanet(t *testing.T) {
	type fields struct {
		Exoplanets map[string]model.Exoplanet
	}
	type args struct {
		id     string
		planet model.Exoplanet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Exoplanet
		wantErr bool
	}{
		{
			name: "Positive case ",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id:     validId,
				planet: dummydata[validId],
			},
			want:    dummydata[validId],
			wantErr: false,
		},
		{
			name: "Nagative case",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id:     invalidId,
				planet: dummydata[validId],
			},
			want:    model.Exoplanet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exo := &ExoplanetService{
				Exoplanets: tt.fields.Exoplanets,
			}
			got, err := exo.UpdateExoplanet(tt.args.id, tt.args.planet)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExoplanetService.UpdateExoplanet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExoplanetService.UpdateExoplanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExoplanetService_GetExoplanet(t *testing.T) {
	type fields struct {
		Exoplanets map[string]model.Exoplanet
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Exoplanet
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Positive case",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id: validId,
			},
			want:    dummydata[validId],
			wantErr: false,
		},
		{
			name: "Nagative case",
			fields: fields{
				Exoplanets: dummydata,
			},
			args: args{
				id: invalidId,
			},
			want:    model.Exoplanet{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exo := &ExoplanetService{
				Exoplanets: tt.fields.Exoplanets,
			}
			got, err := exo.GetExoplanet(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExoplanetService.GetExoplanet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExoplanetService.GetExoplanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExoplanetService_ListExoplanet(t *testing.T) {
	type fields struct {
		Exoplanets map[string]model.Exoplanet
	}
	tests := []struct {
		name   string
		fields fields
		want   []model.Exoplanet
	}{
		{
			name: "Positive Test case ",
			fields: fields{
				Exoplanets: dummydata,
			},
			want: []model.Exoplanet{dummydata[validId]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exo := &ExoplanetService{
				Exoplanets: tt.fields.Exoplanets,
			}
			if got := exo.ListExoplanet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExoplanetService.ListExoplanet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExoplanetService_CreateExoplanet(t *testing.T) {
	type fields struct {
		Exoplanets map[string]model.Exoplanet
	}
	type args struct {
		data model.Exoplanet
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Positive case ",
			fields: fields{
				Exoplanets: dummydata,
			},
			args:    args{data: dummydata[validId]},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exo := &ExoplanetService{
				Exoplanets: tt.fields.Exoplanets,
			}
			_, err := exo.CreateExoplanet(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExoplanetService.CreateExoplanet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
