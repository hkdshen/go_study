package main

import (
	"fmt"
	"strconv"
)
//简单的hashmap
type HashMap struct {
	key string
	value string
	hashCode int
	next *HashMap
}

var table [16]*HashMap

func initTable() {
	for i := range table{
		table[i] = &HashMap{"","",i,nil}
	}
}

func getInstance() [16]*HashMap {
	if table[0] == nil {
		initTable()
	}
	return table
}

func genHashCode(k string) int{
	if len(k) == 0{
		return 0
	}
	var hashCode int = 0
	var lastIndex int = len(k) - 1
	for i := range k {
		if i == lastIndex {
			hashCode += int(k[i])
			break
		}
		//hashCode += (hashCode + int(k[i]))*31
		hashCode += hashCode*31 + int(k[i])
	}
	return hashCode
}

func indexTable(hashCode int) int{
	return hashCode%16
}

func indexNode(hashCode int) int {
	return hashCode>>4
}

func put(k string, v string) string {
	var hashCode = genHashCode(k)
	var thisNode = HashMap{k,v,hashCode,nil}

	var tableIndex = indexTable(hashCode)
	var nodeIndex = indexNode(hashCode)

	var headPtr = getInstance()
	var headNode = headPtr[tableIndex]

	if (*headNode).key == "" {
		*headNode = thisNode
		return ""
	}

	var lastNode *HashMap = headNode
	var nextNode *HashMap = (*headNode).next

	for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex){
		lastNode = nextNode
		nextNode = (*nextNode).next
	}
	if (*lastNode).hashCode == thisNode.hashCode {
		var oldValue string = lastNode.value
		lastNode.value = thisNode.value
		return oldValue
	}
	if lastNode.hashCode < thisNode.hashCode {
		lastNode.next = &thisNode
	}
	if nextNode != nil {
		thisNode.next = nextNode
	}
	return ""
}

func get(k string) string {
	var hashCode = genHashCode(k)
	var tableIndex = indexTable(hashCode)

	var headPtr = getInstance()
	var node *HashMap = headPtr[tableIndex]

	if (*node).key == k{
		return (*node).value
	}

	for (*node).next != nil {
		if k == (*node).key {
			return (*node).value
		}
		node = (*node).next
	}
	return ""
}
func convertToBin(num int) string {
	s := ""

	if num == 0 {
		return "0"
	}

	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		s = strconv.Itoa(lsb) + s
	}
	return s
}
func main() {
	getInstance()
	put("book","a_put")
	put("b","b_put")
	fmt.Println(get("book"))
	fmt.Println(get("b"))
	put("p","p_put")
	fmt.Println(get("p"))
}
