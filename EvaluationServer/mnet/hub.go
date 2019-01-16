package mnet

import (
	"sync"

	"golang.org/x/net/websocket"
)

//Hub 集线器
type Hub struct {
	rwlock      *sync.RWMutex             //读写锁
	connectors  map[int64]*websocket.Conn //客户端集合
	clientclose chan int64                //客户端关闭通道
	exitC       chan struct{}             //退出通道
}

//Add 添加客户端
func (m *Hub) Add(id int64, c *websocket.Conn) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.connectors[id] = c
}

//Del 删除客户端
func (m *Hub) Del(id int64) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.connectors, id)
}

//Get 获得客户端
func (m *Hub) Get(id int64) *websocket.Conn {
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	if conn, OK := m.connectors[id]; OK {
		return conn
	}
	return nil
}

//Close self 关闭集线器
func (m *Hub) Close() {
	close(m.exitC)
	for _, v := range m.connectors {
		v.Close()
	}
}

//ClientClose 关闭客户端
func (m *Hub) ClientClose(id int64) {
	m.clientclose <- id
}

//Server hub服务
func (m *Hub) server() {
	go func() {
		select {
		case <-m.exitC:
			return
		case id := <-m.clientclose:
			m.connectors[id].Close()
			m.Del(id)
		}
	}()
}

// NewHub 创建集线器
func NewHub() *Hub {
	hub := &Hub{
		rwlock:      new(sync.RWMutex),
		connectors:  make(map[int64]*websocket.Conn),
		clientclose: make(chan int64, 100),
		exitC:       make(chan struct{}, 1),
	}
	hub.server()
	return hub
}
