package treeview

import (
	"strings"

	"github.com/nomad-software/goat/image"
	"github.com/nomad-software/goat/internal/tk"
	"github.com/nomad-software/goat/internal/widget/ui/element"
	"github.com/nomad-software/goat/widget"
)

const (
	Type = "treeview"
)

// The treeview widget displays a hierarchical collection of items. Each item
// has a textual label, an optional image, and an optional list of data values.
//
// There are two varieties of columns. The first is the main tree view column
// that is present all the time. The second are data columns that can be added
// when needed. This widget only uses the tree view column.
//
// Each tree item has a list of tags, which can be used to associate event
// bindings and control their appearance. Treeview widgets support horizontal
// and vertical scrolling with the standard scroll commands.
//
// Virtual events that can also be bound to.
// <<TreeviewSelect>>
// <<TreeviewOpen>>
// <<TreeviewClose>>
//
// Reference: https://www.tcl.tk/man/tcl8.6/TkCmd/ttk_treeview.html
//
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=common/bind
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=common/height
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=common/padding
//go:generate go run ../../internal/tools/generate/main.go -recv=*TreeView -pkg=common/scrollbar
type TreeView struct {
	widget.Widget

	nodeRef map[string]*Node
	nodes   []*Node
}

// New creates a new tree view.
// See [option.selectionmode] for mode values.
func New(parent element.Element, selectionMode string) *TreeView {
	tree := &TreeView{
		nodeRef: make(map[string]*Node),
		nodes:   make([]*Node, 0),
	}
	tree.SetParent(parent)
	tree.SetType(Type)

	tk.Get().Eval("ttk::treeview %s -selectmode {%s}", tree.GetID(), selectionMode)

	return tree
}

// SetSelectionMode sets the selection mode of the nodes.
// See [option.selectionmode] for mode values.
func (el *TreeView) SetSelectionMode(mode string) {
	tk.Get().Eval("%s configure -selectmode {%s}", el.GetID(), mode)
}

// EnableHeading controls showing the heading.
func (el *TreeView) EnableHeading(enable bool) {
	if enable {
		tk.Get().Eval("%s configure -show {tree headings}", el.GetID())
	} else {
		tk.Get().Eval("%s configure -show {tree}", el.GetID())
	}
}

// SetHeading sets the heading.
// See [option.anchor] for anchor values.
func (el *TreeView) SetHeading(text, anchor string) {
	tk.Get().Eval("%s heading #0 -text {%s} -anchor {%s}", el.GetID(), text, anchor)
}

// SetHeadingImage sets the heading image to display at the right of the
// heading.
func (el *TreeView) SetHeadingImage(img *image.Image) {
	tk.Get().Eval("%s heading #0 -image %s", el.GetID(), img.GetID())
}

// SetMinWidth sets the width of the tree view.
func (el *TreeView) SetWidth(width int) {
	tk.Get().Eval("%s column #0 -width %d", el.GetID(), width)
}

// SetMinWidth sets the minimum width of the tree view.
func (el *TreeView) SetMinWidth(width int) {
	tk.Get().Eval("%s column #0 -minwidth %d", el.GetID(), width)
}

// SetStretch sets if the tree view stretches or not.
func (el *TreeView) SetStretch(stretch bool) {
	tk.Get().Eval("%s column #0 -stretch %v", el.GetID(), stretch)
}

// RegisterTag registers a tag to be used by nodes in the tree.
// See [option.color] for color names. Use color.Default for no color.
// A hexadecimal string can be used too. e.g. #FFFFFF.
func (el *TreeView) RegisterTag(name string, img *image.Image, foregroundColor, backgroundColor string) {
	tk.Get().Eval("%s tag configure {%s} -image %s -foreground {%s} -background {%s}", el.GetID(), name, img.GetID(), foregroundColor, backgroundColor)
}

// AddNode adds a node to the tree view.
func (el *TreeView) AddNode(text string, open bool, tags ...string) *Node {
	tagStr := strings.Join(tags, " ")
	tk.Get().Eval("%s insert {} end -text {%s} -open %v -tags [list %s]", el.GetID(), text, open, tagStr)
	nodeID := tk.Get().GetStrResult()

	node := &Node{
		nodes: make([]*Node, 0),
	}

	node.SetParent(el)
	node.SetID(nodeID)

	el.nodeRef[nodeID] = node
	el.nodes = append(el.nodes, node)

	return node
}

// GetNode gets a node by its index.
// This will return nil if index is out of bounds.
func (el *TreeView) GetNode(index int) *Node {
	if index < len(el.nodes) {
		return el.nodes[index]
	}

	return nil

}

// GetSelectedNode returns the first selected node.
// This will return nil if nothing is selected.
func (el *TreeView) GetSelectedNode() *Node {
	nodes := el.GetSelectedNodes()

	if len(nodes) > 0 {
		return nodes[0]
	}

	return nil
}

// GetSelectedNodes gets all the selected nodes as an slice.
func (el *TreeView) GetSelectedNodes() []*Node {
	tk.Get().Eval("%s selection", el.GetID())
	ids := tk.Get().GetStrSliceResult()

	result := make([]*Node, 0)

	for _, id := range ids {
		if node, ok := el.nodeRef[id]; ok {
			result = append(result, node)
		}
	}

	return result
}

// Clear clears the tree view.
func (el *TreeView) Clear() {
	tk.Get().Eval("%s children {}", el.GetID())
	tk.Get().Eval("%s delete [list %s]", el.GetID(), tk.Get().GetStrResult())

	clear(el.nodeRef)
	clear(el.nodes)
}
