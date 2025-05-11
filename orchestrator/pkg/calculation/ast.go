package calculation

import "github.com/SteeperMold/Calculator-go/orchestrator/internal/domain"

func IsLeaf(n *domain.Node) bool {
	return n != nil && n.Left == nil && n.Right == nil
}

func AreChildrenLeafs(n *domain.Node) bool {
	return n != nil && IsLeaf(n.Left) && IsLeaf(n.Right)
}

func getDeepestHelper(n *domain.Node, currentDepth int) (*domain.Node, int) {
	if n == nil {
		return nil, currentDepth - 1
	}

	if AreChildrenLeafs(n) && n.Status == domain.StatusInProgress {
		return n, currentDepth
	}

	leftNode, leftDepth := getDeepestHelper(n.Left, currentDepth+1)
	rightNode, rightDepth := getDeepestHelper(n.Right, currentDepth+1)

	if leftNode != nil && rightNode != nil {
		if leftDepth >= rightDepth {
			return leftNode, leftDepth
		}
		return rightNode, rightDepth
	}

	if leftNode != nil {
		return leftNode, leftDepth
	}

	if rightNode != nil {
		return rightNode, rightDepth
	}

	return nil, currentDepth - 1
}

func GetDeepestUnfinishedOperation(n *domain.Node) (*domain.Node, bool) {
	node, _ := getDeepestHelper(n, 1)
	if node == nil {
		return nil, false
	}
	return node, true
}

func GetNodeByID(n *domain.Node, nodeID int64) (*domain.Node, bool) {
	if n == nil {
		return nil, false
	}

	if n.Id == nodeID {
		return n, true
	}

	leftNode, leftExists := GetNodeByID(n.Left, nodeID)
	rightNode, rightExists := GetNodeByID(n.Right, nodeID)

	if leftExists {
		return leftNode, true
	}

	if rightExists {
		return rightNode, true
	}

	return nil, false
}
