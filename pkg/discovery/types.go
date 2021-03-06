package discovery

import (
	"sync"
)

// Used for unmarshalling JSON output from the main API server
type composition struct {
	Kind        string   `yaml:"kind"`
	Plural      string   `yaml:"plural"`
	Endpoint    string   `yaml:"endpoint"`
	Composition []string `yaml:"composition"`
}

// Used for Final output
type Composition struct {
	Level     int
	Kind      string
	Name      string
	Namespace string
	Status    string
	Children  []Composition
}

// Used to store information queried from the main API server
type MetaDataAndOwnerReferences struct {
	MetaDataName             string
	Status                   string
	Namespace                string
	OwnerReferenceName       string
	OwnerReferenceKind       string
	OwnerReferenceAPIVersion string
}

// Used for intermediate storage -- probably can be combined/merged with
// type Provenance and/or type Composition
type CompositionTreeNode struct {
	Level     int
	ChildKind string
	Children  []MetaDataAndOwnerReferences
}

// Used for intermediate storage -- probably can be merged with Composition
type Compositions struct {
	Kind            string
	Name            string
	Namespace       string
	Status          string
	CompositionTree *[]CompositionTreeNode
}

// Used to hold entire composition Provenance of all the Kinds
type ClusterCompositions struct {
	clusterCompositions []Compositions
	mux                 sync.Mutex
}

var (
	TotalClusterCompositions ClusterCompositions
)

func init() {
	TotalClusterCompositions = ClusterCompositions{}
}
