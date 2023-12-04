package radix

import (
    "bytes"
    "encoding/json"
)
 
type RadixNode struct {
    Val      string
    IsEnd    bool
    Children map[string]*RadixNode
}
 
const RootVal = "#"
 
type RadixTree struct {
    Root *RadixNode
}
 
func NewRadixTree() *RadixTree {
    return &RadixTree{
        Root: &RadixNode{
            Val:      RootVal,
            Children: make(map[string]*RadixNode),
        },
    }
}
 
func (t *RadixTree) Insert(str string) {
    if len(str) == 0 {
        return
    }
    t.insertRecursively(t.Root, str)
}
 
func (t *RadixTree) insertAtNode(node *RadixNode, str string) {
    preFixIndex := prefixIdx(str, node.Val)
 
    if preFixIndex < len(node.Val)-1 {
        k := node.Val[preFixIndex+1:]
        node.Val = node.Val[0 : preFixIndex+1]
        if len(node.Children) == 0 {
            node.Children[k] = &RadixNode{
                Val:   k,
                IsEnd: true,
            }
        } else {
            newNode := &RadixNode{
                Val:      k,
                Children: node.Children,
                IsEnd:    node.IsEnd,
            }
            node.Children = map[string]*RadixNode{}
            node.Children[k] = newNode
        }
        node.IsEnd = preFixIndex == len(str)-1
    }
 
    if preFixIndex < len(str)-1 {
        t.insertRecursively(node, str[preFixIndex+1:])
    }
}
 
func (t *RadixTree) insertRecursively(node *RadixNode, str string) {
    for _, child := range node.Children {
        if str[0] == child.Val[0] {
            t.insertAtNode(child, str)
            return
        }
    }
    node.Children[str] = &RadixNode{
        Val:      str,
        Children: map[string]*RadixNode{},
        IsEnd:    true,
    }
}
 
func (t *RadixTree) Find(str string) bool {
    return t.findRecursively(t.Root, str)
}
 
func (t *RadixTree) findAtNode(node *RadixNode, str string) bool {
    preFixIndex := prefixIdx(str, node.Val)
    if preFixIndex < 0 {
        // 没有公共前缀，不在当前节点上
        return false
    }
    if preFixIndex == len(str)-1 && node.IsEnd {
        return true
    } else {
        return t.findRecursively(node, str[preFixIndex+1:])
    }
}
 
func (t *RadixTree) findRecursively(node *RadixNode, str string) bool {
    for _, child := range node.Children {
        if t.findAtNode(child, str) {
            return true
        }
    }
    return false
}
 
func (t *RadixTree) Delete(str string) bool {
    return t.deleteRecursively(t.Root, str)
}
 
func (t *RadixTree) deleteAtNode(parentNode *RadixNode, node *RadixNode, str string) bool {
    preFixIndex := prefixIdx(str, node.Val)
    if preFixIndex < 0 {
        return false
    }
    if preFixIndex == len(str)-1 && node.Val == str && node.IsEnd {
        if len(node.Children) == 0 {
            delete(parentNode.Children, node.Val)
            if len(parentNode.Children) == 1 && !parentNode.isRoot() {
                var theChild *RadixNode
                for _, child := range parentNode.Children {
                    if child.Val != str {
                        theChild = child
                        break
                    }
                }
                parentNode.Val += theChild.Val
                parentNode.IsEnd = theChild.IsEnd
                parentNode.Children = theChild.Children
            }
        } else {
            node.IsEnd = false
            if len(node.Children) == 1 {
                var theChild *RadixNode
                for _, child := range node.Children {
                    theChild = child
                    break
                }
                node.Val += theChild.Val
                node.IsEnd = theChild.IsEnd
                node.Children = theChild.Children
            }
        }
        return true
    } else {
        return t.deleteRecursively(node, str[preFixIndex+1:])
    }
}
 
func (t *RadixTree) deleteRecursively(node *RadixNode, str string) bool {
    for _, child := range node.Children {
        if t.deleteAtNode(node, child, str) {
            return true
        }
    }
    return false
}
 
func (t *RadixTree) Equals(other *RadixTree) bool {
    return t.Root.equal(other.Root)
}
 
func (n *RadixNode) equal(node *RadixNode) bool {
    if n == nil || node == nil {
        return n == nil && node == nil
    }
    if n.Val != node.Val {
        return false
    }
    if len(n.Children) != len(node.Children) {
        return false
    }
    for _, child := range n.Children {
        if child2, ok := node.Children[child.Val]; ok {
            if !child.equal(child2) {
                return false
            }
        } else {
            return false
        }
    }
    return true
}
 
func (t *RadixTree) String() string {
    b, _ := json.Marshal(*t)
    var out bytes.Buffer
    _ = json.Indent(&out, b, "", "  ")
    return out.String()
}
 
func (t *RadixTree) Clone() *RadixTree {
    return &RadixTree{Root: t.Root.clone()}
}
 
func (n *RadixNode) clone() *RadixNode {
    if n == nil {
        return nil
    }
    newRoot := &RadixNode{
        Val:      n.Val,
        IsEnd:    n.IsEnd,
        Children: map[string]*RadixNode{},
    }
    for _, child := range n.Children {
        newRoot.Children[child.Val] = child.clone()
    }
    return newRoot
}
 
func (n *RadixNode) isRoot() bool {
    return n.Val == RootVal
}
 
func prefixIdx(str1, str2 string) int {
    idx := 0
    for ; idx < len(str1) && idx < len(str2); idx++ {
        if str1[idx] != str2[idx] {
            break
        }
    }
    return idx - 1
}