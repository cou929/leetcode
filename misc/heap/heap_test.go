package heap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHeap_Insert(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		h    *MinHeap
		args args
		want *MinHeap
	}{
		{
			name: "single",
			h:    &MinHeap{},
			args: args{1},
			want: &MinHeap{[]int{1}},
		},
		{
			name: "normal",
			h:    &MinHeap{[]int{1, 10, 11, 15}},
			args: args{3},
			want: &MinHeap{[]int{1, 3, 11, 15, 10}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Insert(tt.args.val)
			got := tt.h
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("heap.Insert diff (-got +want)\n%s", diff)
			}
		})
	}
}

func TestMinHeap_Pop(t *testing.T) {
	type fields struct {
		Tree []int
	}
	tests := []struct {
		name     string
		fields   fields
		want     int
		wantErr  bool
		wantHeap *MinHeap
	}{
		{
			name:     "normal",
			fields:   fields{[]int{3, 10, 11, 12}},
			want:     3,
			wantErr:  false,
			wantHeap: &MinHeap{[]int{10, 12, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MinHeap{
				Tree: tt.fields.Tree,
			}
			got, err := h.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("MinHeap.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MinHeap.Pop() = %v, want %v", got, tt.want)
			}
			if diff := cmp.Diff(h, tt.wantHeap); diff != "" {
				t.Errorf("heap.Pop diff (-got +want)\n%s", diff)
			}
		})
	}
}
