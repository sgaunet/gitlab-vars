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
