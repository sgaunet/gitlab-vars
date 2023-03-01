package gitlabapi

import (
	"reflect"
	"testing"
)

func TestVariables_GetVarOfScope(t *testing.T) {
	type args struct {
		key   string
		scope string
	}

	vars := make(Variables, 0)
	vars = append(vars, Variable{Key: "key1", Value: "value1", EnvironmentScope: "preprod/*"})
	vars = append(vars, Variable{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"})
	tests := []struct {
		name    string
		v       Variables
		args    args
		want    Variable
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "GetVarOfScope",
			v:       vars,
			args:    args{key: "key1", scope: "preprod/1"},
			want:    Variable{Key: "key1", Value: "value1", EnvironmentScope: "preprod/*"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.v.GetVarOfScope(tt.args.key, tt.args.scope)
			if (err != nil) != tt.wantErr {
				t.Errorf("Variables.GetVarOfScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Variables.GetVarOfScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeVars(t *testing.T) {
	type args struct {
		parentVars []Variable
		childVars  []Variable
	}
	tests := []struct {
		name string
		args args
		want []Variable
	}{
		{
			name: "MergeVars",
			args: args{
				parentVars: []Variable{
					{Key: "key1", Value: "value1", EnvironmentScope: "preprod/*"},
					{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
				},
				childVars: []Variable{
					{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
					{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
				},
			},
			want: []Variable{
				{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
				{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
			},
		},
		{
			name: "MergeVars with empty parentVars",
			args: args{
				parentVars: []Variable{},
				childVars: []Variable{
					{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
					{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
				},
			},
			want: []Variable{
				{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
				{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
			},
		},
		{
			name: "MergeVars with empty childVars",
			args: args{
				parentVars: []Variable{
					{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
					{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
				},
				childVars: []Variable{},
			},
			want: []Variable{
				{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
				{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
			},
		},
		{
			name: "MergeVars without merge",
			args: args{
				parentVars: []Variable{
					{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
					{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
				},
				childVars: []Variable{
					{Key: "key3", Value: "value4", EnvironmentScope: "preprod/*"},
					{Key: "key4", Value: "value2", EnvironmentScope: "preprod/*"},
				},
			},
			want: []Variable{
				{Key: "key1", Value: "value4", EnvironmentScope: "preprod/*"},
				{Key: "key2", Value: "value2", EnvironmentScope: "preprod/*"},
				{Key: "key3", Value: "value4", EnvironmentScope: "preprod/*"},
				{Key: "key4", Value: "value2", EnvironmentScope: "preprod/*"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeVars(tt.args.parentVars, tt.args.childVars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeVars() = %v, want %v", got, tt.want)
			}
		})
	}
}
