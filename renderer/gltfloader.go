package renderer

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/g3n/engine/core"
	"github.com/g3n/engine/loader/gltf"
)

// nameChildren names all gltf nodes by path
func (a *RenderingApp) nameChildren(p string, n core.INode) {
	node := n.GetNode()
	node.SetName(p)
	a.nodeBuffer[p] = node
	for _, child := range node.Children() {
		idx := node.ChildIndex(child)
		title := p + "/" + strconv.Itoa(idx)
		a.nameChildren(title, child)
	}
}

// loadScene loads a gltf file
func (a *RenderingApp) loadScene(fpath string) error {
	a.respondToClient("loading", fpath)
	// Checks file extension
	ext := filepath.Ext(fpath)
	var g *gltf.GLTF
	var err error

	// Parses file
	if ext == ".gltf" {
		g, err = gltf.ParseJSON(fpath)
	} else if ext == ".glb" {
		g, err = gltf.ParseBin(fpath)
	} else {
		return fmt.Errorf("unrecognized file extension:%s", ext)
	}

	if err != nil {
		return err
	}

	defaultSceneIdx := 0
	if g.Scene != nil {
		defaultSceneIdx = *g.Scene
	}

	// Create default scene
	n, err := g.LoadScene(defaultSceneIdx)
	if err != nil {
		return err
	}

	a.scene.Add(n)
	root := a.scene.ChildIndex(n)
	a.nameChildren("/"+strconv.Itoa(root), n)
	a.respondToClient("loaded", fpath)
	return nil
}
