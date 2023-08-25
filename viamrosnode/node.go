package viamrosnode

import (
	"strconv"
	"strings"
	"sync"

	"github.com/bluenviron/goroslib/v2"
)

var lock *sync.Mutex
var i int
var nodes map[string]*goroslib.Node

func init() {
	lock = &sync.Mutex{}
	i = 0
	nodes = make(map[string]*goroslib.Node)
}

func GetInstance(primary string) (*goroslib.Node, error) {
	lock.Lock()
	defer lock.Unlock()
	node, ok := nodes[primary]
	if ok {
		return node, nil
	} else {
		node, err := goroslib.NewNode(goroslib.NodeConf{
			Name:          strings.Join([]string{primary, strconv.Itoa(i)}, ""),
			MasterAddress: primary,
		})
		if err != nil {
			return nil, err
		}

		nodes[primary] = node
		i = i + 1
		return node, nil
	}
}

func ShutdownNodes() {
	lock.Lock()
	defer lock.Unlock()
	for _, node := range nodes {
		node.Close()
	}
}
