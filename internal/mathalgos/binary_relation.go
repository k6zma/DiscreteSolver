package mathalgos

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/k6zma/DiscreteSolver/internal/api/models"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

type BinaryRelationGraph struct {
	graph    *simple.DirectedGraph
	elements map[string]int64
	relation [][2]string
}

func NewBinaryRelationGraph(model models.BinaryRelationModel) *BinaryRelationGraph {
	graph := simple.NewDirectedGraph()
	elements := make(map[string]int64)

	for _, element := range model.SetOfElements {
		node := graph.NewNode()
		graph.AddNode(node)
		elements[element] = node.ID()
	}

	for _, pair := range model.BinaryRelation {
		from := elements[pair[0]]
		to := elements[pair[1]]

		if from != to {
			graph.SetEdge(graph.NewEdge(graph.Node(from), graph.Node(to)))
		}
	}

	return &BinaryRelationGraph{
		graph:    graph,
		elements: elements,
		relation: model.BinaryRelation,
	}
}

func (g *BinaryRelationGraph) DOTID() string {
	return "BinaryRelationGraph"
}

func (g *BinaryRelationGraph) Attributes() []encoding.Attribute {
	return []encoding.Attribute{
		{Key: "label", Value: "Binary Relation Graph"},
		{Key: "labelloc", Value: "t"},
	}
}

func (g *BinaryRelationGraph) NodeAttributes(id int64) []encoding.Attribute {
	for key, nodeID := range g.elements {
		if nodeID == id {
			return []encoding.Attribute{
				{Key: "label", Value: key},
				{Key: "color", Value: "skyblue"},
				{Key: "style", Value: "filled"},
			}
		}
	}

	return nil
}

func (g *BinaryRelationGraph) EdgeAttributes(e simple.Edge) []encoding.Attribute {
	return []encoding.Attribute{
		{Key: "color", Value: "black"},
	}
}

func (g *BinaryRelationGraph) GenerateImage() ([]byte, error) {
	dotData, err := dot.Marshal(g.graph, "BinaryRelationGraph", "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal graph to DOT: %v", err)
	}

	cmd := exec.Command("dot", "-Tpng")
	cmd.Stdin = bytes.NewReader(dotData)

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to generate graph image: %v", err)
	}

	return out.Bytes(), nil
}

func checkReflexiveProperty(elements map[string]struct{}, relation map[[2]string]struct{}) map[string]bool {
	isReflexive := true
	isAntireflexive := true

	for e := range elements {
		pair := [2]string{e, e}
		_, exists := relation[pair]

		if exists {
			isAntireflexive = false
		} else {
			isReflexive = false
		}
	}

	isNonreflexive := !isReflexive && !isAntireflexive

	return map[string]bool{
		"Рефлексивно":     isReflexive,
		"Антирефлексивно": isAntireflexive,
		"Нерефлексивно":   isNonreflexive,
	}
}

func checkSymmetryProperties(relation map[[2]string]struct{}) map[string]bool {
	isSymmetric := true
	isAsymmetric := true
	isAntisymmetric := true

	for pair := range relation {
		reversePair := [2]string{pair[1], pair[0]}
		_, exists := relation[reversePair]

		if exists {
			isAsymmetric = false
		} else {
			isSymmetric = false
		}

		if pair[0] != pair[1] && exists {
			isAntisymmetric = false
		}
	}

	isNonsymmetric := !isSymmetric && !isAntisymmetric

	return map[string]bool{
		"Симметрично":     isSymmetric,
		"Асимметрично":    isAsymmetric,
		"Антисимметрично": isAntisymmetric,
		"Несимметрично":   isNonsymmetric,
	}
}

func checkTransitivityProperties(relation map[[2]string]struct{}) map[string]bool {
	isTransitive := true
	isAntitransitive := true

	for pair1 := range relation {
		for pair2 := range relation {
			if pair1[1] == pair2[0] {
				newPair := [2]string{pair1[0], pair2[1]}
				_, exists := relation[newPair]
				if !exists {
					isTransitive = false
				} else {
					isAntitransitive = false
				}
			}
		}
	}

	isNontransitive := !isTransitive && !isAntitransitive

	return map[string]bool{
		"Транзитивно":     isTransitive,
		"Антитранзитивно": isAntitransitive,
		"Нетранзитивно":   isNontransitive,
	}
}

func GetRelationProperties(model models.BinaryRelationModel) []string {
	elements := make(map[string]struct{})
	for _, e := range model.SetOfElements {
		elements[e] = struct{}{}
	}

	if len(elements) == 0 {
		for _, pair := range model.BinaryRelation {
			elements[pair[0]] = struct{}{}
			elements[pair[1]] = struct{}{}
		}
	}

	relation := make(map[[2]string]struct{})
	for _, pair := range model.BinaryRelation {
		relation[[2]string{pair[0], pair[1]}] = struct{}{}
	}

	reflexiveProperties := checkReflexiveProperty(elements, relation)
	symmetryProperties := checkSymmetryProperties(relation)
	transitivityProperties := checkTransitivityProperties(relation)

	propertiesList := []string{}
	for property, isTrue := range reflexiveProperties {
		if isTrue {
			propertiesList = append(propertiesList, property)
		}
	}
	for property, isTrue := range symmetryProperties {
		if isTrue {
			propertiesList = append(propertiesList, property)
		}
	}
	for property, isTrue := range transitivityProperties {
		if isTrue {
			propertiesList = append(propertiesList, property)
		}
	}

	return propertiesList
}
