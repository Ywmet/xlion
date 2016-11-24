package main

import (
    "im/proto"
	//"im/http_server"
    "fmt"
    "net"
	"encoding/json"
    //"os"
	"io"
	"net/http"

)


var g_userlist map[string]string

func get_user_list(rw http.ResponseWriter, req *http.Request) {
	//io.WriteString(rw, "hello widuu")
	fmt.Print("================get_user_list====================")
	for k, v := range g_userlist {
		//var data string
		data := fmt.Sprintf("k=%v, v=%v\n", k, v)
		io.WriteString(rw, data)
	}
}

func start_http_server () {
	http.HandleFunc("/", get_user_list)  //设定访问的路径
	http.ListenAndServe(":44444", nil) //设定端口和handler
}

func InitGlobal() {
	g_userlist = make(map[string]string)
}

func reader(readerChannel chan []byte) {
    for {
        select {
        case data := <-readerChannel:
            fmt.Print(len(data))
            fmt.Print(string(data))
			body := string(data)

			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(body), &dat); err == nil {
				//fmt.Println(dat)
				//fmt.Println(dat["session"])

				if (dat["session"] != nil) {
					fmt.Print(dat["session"])
					g_userlist[dat["session"].(string)] = "online"
					//push session into user map
				}
			}
		}
	}
}

func handleConnection(conn net.Conn) {

    tmpBuffer := make([]byte, 0)

    readerChannel := make(chan []byte, 16)

    go reader(readerChannel)

    buffer := make([]byte, 1024)
    for {
    n, err := conn.Read(buffer)
    if err != nil {
    fmt.Print(conn.RemoteAddr().String(), " connection error: ", err)
    return
    }

    tmpBuffer = proto.Depack(append(tmpBuffer, buffer[:n]...), readerChannel)
    }
    defer conn.Close()
}

func main() {
	InitGlobal()
	go start_http_server()

    netListen, err := net.Listen("tcp", "localhost:6060")
    //CheckError(err)
    fmt.Print("hi", err)

    //defer netListen.Close()

    //Log("Waiting for clients")
    for {
        conn, err := netListen.Accept()
        if err != nil {
            continue
        }

        //timeouSec :=10  
        //conn.  
        //Log(conn.RemoteAddr().String(), " tcp connect success")
        fmt.Print("connect success", err)
        go handleConnection(conn)
    }
}
