package actions

import (
	"github.com/c0rby/dcfs/internal/cli"
	"github.com/c0rby/dcfs/pkg/dcfs"
)

func Tree(path string) error {
	trees, err := dcfs.ListTrees(path)
	if err != nil {
		return err
	}

	for _, t := range trees {
		if err := cli.PrintTree(t, 0); err != nil {
			return err
		}
	}
	return nil
}
