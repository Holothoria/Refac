package main





type Hub struct {
	rooms              map[strings]*Room
	register           chan *Client
	unregister         chan *Client
	broadcast          chan *Message
}


func newHub()  *hub {
	return &Hub{
		rooms:           make(map[string]*Room),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		broadcast:       make(chan*Message),
	}
}



func (h*Hub) Run() {
	for {
		select {
		case client := <-h.register:
		 //add client to their rrom
		 
		 if _, ok := h.rooms[client.roomName]; !ok {
			 h.rooms[client.roomName] = &Room{
				 Name: client.roomName,
		 }
	 }
	 case client := <-h.unregister:
	 //remove client from room
	 
	 delete(h.rooms, client.roomName)
	 
	 case message := <-h.broadcast:
	 //broadcast to everyone in room
 }
}

}
