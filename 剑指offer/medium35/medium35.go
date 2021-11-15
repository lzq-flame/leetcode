/**
 * @Author: liuzhiqi
 * @Data: 2021/11/14 09:47
 */

//复杂链表的复制
package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cacheNode := map[*Node]*Node{}
	var deepCopy func(node *Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, has := cacheNode[node]; has {
			return n
		}
		newNode := &Node{Val: node.Val}
		cacheNode[node] = newNode
		node.Next = deepCopy(node.Next)
		node.Random = deepCopy(node.Random)
		return newNode
	}
	return deepCopy(head)
}
