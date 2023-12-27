package random

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

type Node struct {
	node *snowflake.Node
}

func NewSnowflakeNode(node int64) (*Node, error) {
	n, err := snowflake.NewNode(node)
	if err != nil {
		return nil, fmt.Errorf("node creation err = %w", err)
	}

	return &Node{node: n}, nil
}

func (n *Node) GenerateInt64() int64 {
	return n.node.Generate().Int64()
}
