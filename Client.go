package LinkClient

import (
	"github.com/funny/link"
	"github.com/funny/link/codec"
	"log"
	"proto"
)



func Start() {
	json := codec.Json()
	json.Register(proto.AddReq{})
	json.Register(proto.AddRsp{})

	client, err := link.Dial("tcp", "0.0.0.0:8888", json, 0)
	checkErr(err)
	clientSessionLoop(client)
}
func clientSessionLoop(session *link.Session) {
	for i := 0; i < 10; i++ {
		err := session.Send(&proto.AddReq{
			i, i,
		})
		checkErr(err)
		log.Printf("Send: %d + %d", i, i)

		rsp, err := session.Receive()
		checkErr(err)
		log.Printf("Receive: %d", rsp.(*proto.AddRsp).C)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}