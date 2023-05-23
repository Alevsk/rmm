package obsidian

import (
	"fmt"

	"github.com/Alevsk/rmm/internal/mindmap"
	"github.com/google/uuid"
)

type ObsidianCanvas struct {
	Nodes []ObsidianNode `json:"nodes"`
	Edges []ObsidianEdge `json:"edges"`
}

type ObsidianNode struct {
	ID       string `json:"id"`
	NodeType string `json:"type"`
	Text     string `json:"text"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type ObsidianEdge struct {
	ID       string `json:"id"`
	FromNode string `json:"fromNode"`
	FromSide string `json:"fromSide"`
	ToNode   string `json:"toNode"`
	ToSide   string `json:"toSide"`
}

func generateUUID() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return id.String()
}

type nodeStackEntry struct {
	node       mindmap.Node
	parentNode *ObsidianNode
}

func calculateChildY(parentY, childNumber int) int {
	switch {
	case childNumber%2 == 0:
		return parentY + ((childNumber / 2) * 90)
	default:
		return parentY - ((childNumber / 2) * 90)
	}
}

func GenerateObsidianCanvas(tree mindmap.Node) ObsidianCanvas {
	canvas := ObsidianCanvas{}
	stack := []nodeStackEntry{}
	coordsXnY := make(map[int]map[int]bool)
	coordsYnX := make(map[int]map[int]bool)
	var parentNode *ObsidianNode
	root := mindmap.Node{
		"domains": tree,
	}
	stack = append(stack, nodeStackEntry{root, parentNode})

	for len(stack) > 0 {
		entry := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentNode := entry.node
		parentNode = entry.parentNode

		if len(currentNode) == 0 {
			continue
		}

		x, y, childCounter := 0, 0, 0
		coordsXnY[0] = make(map[int]bool)
		coordsXnY[0][0] = true
		coordsYnX[0] = make(map[int]bool)
		coordsYnX[0][0] = true

		if parentNode != nil {

			y = parentNode.Y
			x += parentNode.X + 350

			for {
				if _, ok := coordsXnY[x]; !ok {
					break
				}
				x += 350
			}

			coordsXnY[x] = make(map[int]bool)
			coordsXnY[x][y] = true
			coordsYnX[y] = make(map[int]bool)
			coordsYnX[y][x] = true

		}

		// Sort the keys of the nested map.
		keys := make([]string, 0, len(currentNode))
		for k := range currentNode {
			keys = append(keys, k)
		}

		for _, currentNodeKey := range keys {

			childNode := currentNode[currentNodeKey]
			nodeID := generateUUID()

			if parentNode != nil {
				edgeID := generateUUID()
				canvas.Edges = append(canvas.Edges, ObsidianEdge{
					ID:       edgeID,
					FromNode: parentNode.ID,
					FromSide: "right",
					ToNode:   nodeID,
					ToSide:   "left",
				})
				if len(currentNode) > 1 { // currentNode is the parentNode in the context of this loop
					y = calculateChildY(parentNode.Y, childCounter+1)
				}
			}
			currentObsidianNode := ObsidianNode{
				ID:       nodeID,
				NodeType: "text",
				Text:     fmt.Sprintf("%s (%d,%d)", currentNodeKey, x, y),
				Width:    250,
				Height:   60,
				X:        x,
				Y:        y,
			}
			canvas.Nodes = append(canvas.Nodes, currentObsidianNode)

			stack = append(stack, nodeStackEntry{childNode, &currentObsidianNode})

			childCounter++
		}
	}
	return canvas
}
