package srpc

func (conn *Conn) close() error {
	err := conn.Flush()
	if conn.parent != nil {
		conn.parent.callLock.Unlock()
	}
	return err
}

func (conn *Conn) getUsername() string {
	if conn.parent != nil {
		panic("cannot call GetUsername() for client connection")
	}
	return conn.username
}
