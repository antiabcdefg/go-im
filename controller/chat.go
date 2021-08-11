package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go-im/args"
	"gopkg.in/fatih/set.v0"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

//读写锁
var rwlocker sync.RWMutex

// ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isvalid := checkToken(userId, token)

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return isvalid
	}}).Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	node := &Node{Conn: conn, DataQueue: make(chan []byte, 50), GroupSets: set.New(set.ThreadSafe)}

	comIds := contactService.SearchComunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()

	go sendproc(node)
	go recvproc(node)

}

//发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

//接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//dispatch(data)
		broadMsg(data)
		fmt.Printf("recv<=%s\n", data)
	}
}

//处理消息
func dispatch(data []byte) {
	msg := args.MessageArg{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	switch msg.Cmd {
	case args.CMD_SINGLE_MSG:
		sendMsg(msg.DstId, data)
	case args.CMD_ROOM_MSG:
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.DstId) {
				v.DataQueue <- data
			}
		}
	case args.CMD_HEART:

	}
}

func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

func AddGroupId(userId, gid int64) {
	rwlocker.Lock()
	node, ok := clientMap[userId]
	if ok {
		node.GroupSets.Add(gid)
	}
	clientMap[userId] = node
	rwlocker.Unlock()
}

var udpsendchan chan []byte = make(chan []byte, 1024)

func init() {
	go udpsendproc()
	go udprecvproc()
}

func broadMsg(data []byte) {
	udpsendchan <- data
}

func udpsendproc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.IPv4(192, 168, 3, 255), Port: 3000})
	defer con.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		select {
		case data := <-udpsendchan:
			_, err = con.Write(data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

func udprecvproc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 3000})
	defer con.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Printf(string(n))
		dispatch(buf[0:n])
	}
}
