package dcfs

import (
	"os"
	"path/filepath"

	"github.com/pkg/xattr"
)

const (
	XattrKeyParentId = "user.ocis.parentid"
	XattrKeyName     = "user.ocis.name"
)

type Node struct {
	Name     string
	ParentId string
	Path     string
	Children []Node
}

func ListRoots(path string) ([]Node, error) {
	path = filepath.Clean(path)
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var roots []Node
	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		pIdval, err := xattr.Get(entryPath, XattrKeyParentId)
		if err != nil {
			return nil, err
		}
		if string(pIdval) != "root" {
			continue
		}
		root, err := Read(entryPath)
		if err != nil {
			return nil, err
		}
		roots = append(roots, root)
	}
	return roots, nil
}

func Read(path string) (Node, error) {
	pIdval, err := xattr.Get(path, XattrKeyParentId)
	if err != nil {
		return Node{}, err
	}

	nval, err := xattr.Get(path, XattrKeyName)
	if err != nil {
		return Node{}, err
	}
	return Node{Path: path, ParentId: string(pIdval), Name: string(nval)}, nil
}

func ListChildren(n Node) ([]Node, error) {
	info, err := os.Stat(n.Path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, nil
	}

	entries, err := os.ReadDir(n.Path)
	if err != nil {
		return nil, err
	}
	var children []Node
	for _, entry := range entries {
		entryPath := filepath.Join(n.Path, entry.Name())
		nval, err := xattr.Get(entryPath, XattrKeyName)
		if err != nil {
			return nil, err
		}
		child := Node{Path: entryPath, Name: string(nval)}
		if cs, err := ListChildren(child); err != nil {
			return nil, err
		} else {
			child.Children = cs
		}
		children = append(children, child)
	}

	return children, nil
}

func ListTrees(path string) ([]Node, error) {
	roots, err := ListRoots(path)
	if err != nil {
		return nil, err
	}

	for i := range roots {
		if children, err := ListChildren(roots[i]); err != nil {
			return nil, err
		} else {
			roots[i].Children = children
		}
	}
	return roots, nil
}
