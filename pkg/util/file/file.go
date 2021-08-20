package file

import (
	"strings"
)

var root string

//层级节点
type Node struct {
	Name     string  `json:"name"`
	Key      string  `json:"key"`
	Value    string  `json:"value"`
	Children []*Node `json:"children"`
}

func NewNode(path string, value string,key string) *Node {
	return &Node{
		Name:     path,
		Value:    value,
		Key:key,
		Children: []*Node{},
	}

}

func (n *Node)Insert(path string, value string){
	root = path
	n.insert(path,value,path)
}


//根据路径插入值
func (n *Node) insert(path string, value string,orginPath string) {
	if string(path[0])=="/"{
		path=path[1:]
	}
	ps := strings.Split(path, "/")
	length := len(ps)
	isExist := false
	for _, p := range n.Children {
		if ps[0] == p.Name {
			isExist = true
			if length==1{
				p.Value = value
				p.Key=orginPath
			}
			if length>1{
				p.insert(path[len(ps[0]):], value,orginPath)
				p.Key=orginPath[:strings.LastIndex(orginPath,ps[0])+len(ps[0])]
			}
		}
	}
	if !isExist {
		newNode := &Node{
			Name:     ps[0],
			Children: []*Node{},
		}
		if length == 1 {
			newNode.Value = value
			newNode.Key=orginPath
		} else if length > 1 {
			newNode.insert(path[len(ps[0]):], value,orginPath)
			newNode.Key=orginPath[:strings.LastIndex(orginPath,ps[0])+len(ps[0])]
		}
		n.Children = append(n.Children, newNode)
	}
	return
}
