package controller

import (
	"reflect"
	"testing"
)

func Test_getNodesToBeInformed(t *testing.T) {
	type args struct {
		receivedNodes []string
		savedNodes    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNodesToBeInformed(tt.args.receivedNodes, tt.args.savedNodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNodesToBeInformed() = %v, want %v", got, tt.want)
			}
		})
	}
}
