package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync"
	"sync/atomic"
)

type EchoServer struct {
	Addr     string
	SerialId atomic.Int64

	ln     net.Listener
	conns  sync.Map
	joinCh chan struct{}

	actors sync.Map
}

func NewExampleServer() *EchoServer {
	return &EchoServer{}
}

func (es *EchoServer) Init() error {
	ln, err := net.Listen("tcp", es.Addr)
	if err != nil {
		return err
	}

	es.ln = ln
	es.joinCh = make(chan struct{})
	return nil
}

func (es *EchoServer) Run() {
acceptLoop:
	for {
		conn, err := es.ln.Accept()
		if err != nil {
			break acceptLoop
		}
		go es.handleConn(conn)
	}

	es.conns.Range(func(key, value interface{}) bool {
		conn := value.(*net.Conn)
		(*conn).Close()
		return true
	})
	close(es.joinCh)
}

func (es *EchoServer) handleConn(conn net.Conn) {
	defer conn.Close()

	id := es.SerialId.Add(1)
	p := NewPlayer(int(id))
	act := NewActor(p)
	es.actors.Store(id, act)

	buffer := make([]byte, 1024)
	for {
		n, err := io.ReadFull(conn, buffer[:EchoMessageLen])
		if n != EchoMessageLen || err != nil {
			panic(err)
		}
		es.dipatcherMessage(act, buffer[:EchoMessageLen])
		conn.Write(buffer[:EchoMessageLen])
	}
}

func (es *EchoServer) dipatcherMessage(actor *Actor, msg []byte) {
	n := rand.Intn(2)
	if n == 0 {
		actor.Heal()
	} else {
		actor.Attack()
	}
	fmt.Println(msg)
}

func (es *EchoServer) Stop() {
	es.ln.Close()
	<-es.joinCh
}
