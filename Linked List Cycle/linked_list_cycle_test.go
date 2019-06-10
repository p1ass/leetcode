package Linked_List_Cycle

import "testing"

func Test_hasCycle(t *testing.T) {
	n := &ListNode{}
	n.Next = &ListNode{Next: n}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				head: n,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				head: &ListNode{
					Val:  0,
					Next: &ListNode{},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
