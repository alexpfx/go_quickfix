package action

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_list_Query(t *testing.T) {
	type fields struct {
		actions []Item
	}
	type args struct {
		text string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Item
	}{
		{
			name: "t1",
			fields: fields{
				actions: []Item{
					CpfToNum(),
				},
			},
			args: args{
				text: "123.123.123-00",
			},
			want: []Item{
				CpfToNum(),
			},
		},
		{
			name: "t2",
			fields: fields{
				actions: []Item{
					CpfToNum(),
				},
			},
			args: args{
				text: "123.123.123-001",
			},
			want: []Item{

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := list{
				actions: tt.fields.actions,
			}
			got := a.Query(tt.args.text)
			bGot := []byte(fmt.Sprintf("%v", got))
			bWant := []byte(fmt.Sprintf("%v", tt.want))
			if !bytes.Equal(bGot, bWant) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}

		})
	}
}
