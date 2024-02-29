package memory

import (
	jgf "github.com/converged-computing/jsongraph-go/jsongraph/v2/graph"
)

// Generate a new resource from a JGF node
func NewResource(node jgf.Node) *Resource {

	// We assume the node has a type for metadata
	resourceType := "resource"
	typ, err := node.Metadata.GetStringElement("type")
	if err == nil {
		resourceType = typ
	}

	resourceSize := int32(1)
	size, err := node.Metadata.GetIntElement("size")
	if err == nil {
		resourceSize = size
	}

	resourceUnit := ""
	unit, err := node.Metadata.GetStringElement("unit")
	if err == nil {
		resourceUnit = unit
	}

	return &Resource{
		Size: resourceSize,
		Unit: resourceUnit,
		Type: resourceType,
	}
}
