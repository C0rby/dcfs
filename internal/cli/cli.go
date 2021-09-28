package cli

import (
	"fmt"

	"github.com/c0rby/dcfs/pkg/dcfs"
)

func PrintTree(n dcfs.Node, depth int) error {
	for i := 0; i < depth; i++ {
		fmt.Print("  ")
	}
	fmt.Println(n.Name)
	for _, child := range n.Children {
		if err := PrintTree(child, depth+1); err != nil {
			return err
		}
	}
	return nil
}
