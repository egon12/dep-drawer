package main

import (
	"reflect"
	"testing"
)

func TestGrouper(t *testing.T) {

	tests := []struct {
		name              string
		inputName         string
		inputDependencies map[string][]string
		want              map[string][]string
	}{
		{
			name:              "single",
			inputName:         "module_a",
			inputDependencies: map[string][]string{"module_a/submodule_a1": []string{"module_b", "module_c"}},
			want:              map[string][]string{"module_a": []string{"module_b", "module_c"}},
		},
		{
			name:      "single",
			inputName: "module_a",
			inputDependencies: map[string][]string{
				"module_a/submodule_a1": []string{"module_b", "module_c"},
				"module_a/submodule_a2": []string{"module_c", "module_d"},
			},
			want: map[string][]string{"module_a": []string{"module_b", "module_c", "module_d"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GroupBy(tt.inputName, tt.inputDependencies)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v\ngot: %v", tt.want, got)
			}
		})
	}
}
