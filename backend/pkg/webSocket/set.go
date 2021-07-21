package webSocket

type Set struct {
	connections map[*Connection]bool
}

func (set *Set) Add(conn *Connection) {
	set.connections[conn] = true
}

func (set *Set) Remove(conn *Connection) {
	delete(set.connections, conn)
}

func (set *Set) Send(message *Message) {
	for conn := range set.connections {
		conn.send <- message.Text
	}
}

func (set Set) Length() int {
	return len(set.connections)
}

func NewSet() *Set {
	return &Set{
		connections: map[*Connection]bool{},
	}
}
