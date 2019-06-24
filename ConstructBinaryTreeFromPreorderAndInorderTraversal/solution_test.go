package ConstructBinaryTreeFromPreorderAndInorderTraversal

import (
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				preorder: []int{1, 2, 3},
				inorder:  []int{2, 3, 1},
			},
			want: &TreeNode{
				Val:   0,
				Left:  &TreeNode{},
				Right: &TreeNode{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
