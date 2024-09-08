package pre_processor

import (
	"reflect"
	"testing"
)

func TestIndexValidator_ProcessBlock(t *testing.T) {
	type args struct {
		block *coreBlock.Block
	}
	tests := []struct {
		name  string
		args  args
		want  error
		want1 *[]coreBlock.Block
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iv := &IndexValidator{}
			got, got1 := iv.ProcessBlock(tt.args.block)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessBlock() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProcessBlock() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
