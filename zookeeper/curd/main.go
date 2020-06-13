package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	host = []string{"192.168.31.205:2181"}
)

func main() {
	conn, _, err := zk.Connect(host, 5*time.Second)
	if err != nil {
		panic(err)
	}

	//add
	//if _, err = conn.Create("/test_tree1", []byte("hello world!"), 0, zk.WorldACL(zk.PermAll)); err != nil {
	//	fmt.Println("creat err :", err)
	//}

	//find

	nodeValue, dStat, err := conn.Get("/test_tree1")
	if err != nil {
		fmt.Println("get err :", err)
		return
	}
	fmt.Printf("nodeValue:%s, status: %s \n", nodeValue, dStat)

	//modify

	if _, err := conn.Set("/test_tree1", []byte("hello world2"), dStat.Version); err != nil {
		fmt.Println("update err : ", err)
	}

	nodeValue, dStat, err = conn.Get("/test_tree1")
	if err != nil {
		fmt.Println("get err :", err)
		return
	}
	fmt.Printf("nodeValue:%s, status: %s \n", nodeValue, dStat)

	//delete
	_, dstat, _ := conn.Get("/test_tree1")
	if err := conn.Delete("/test_tree1", dstat.Version); err != nil {
		fmt.Println("delete err :", err)
	}

	// exists
	hasNode, _, err := conn.Exists("test_tree1")
	if err != nil {
		fmt.Println("exits err :", err)
	}

	fmt.Println("has node :", hasNode)

	if _, err = conn.Create("/test_tree1", []byte("tree_content"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("creat err :", err)
	}

	if _, err = conn.Create("/test_tree1/sub", []byte("tree_content"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("creat err :", err)
	}

	if _, err = conn.Create("/test_tree1/sub1", []byte("tree_content1"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("creat err :", err)
	}

	childNodes, _, err := conn.Children("/test_tree1")
	if err != nil {
		fmt.Println("children err :", err)
	}

	fmt.Println("childNodes node :", childNodes)
}
