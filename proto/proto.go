
package proto

import (
    "bytes"
    "encoding/binary"
)
const (
    ConstHeader         = "Headers"
    ConstHeaderLength   = 7
    ConstMLength = 4
)

//封包
func Enpack(message []byte) []byte {
    return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

//解包
func Depack(buffer []byte, readerChannel chan []byte) []byte {
    length := len(buffer)

    var i int
    for i = 0; i < length; i = i + 1 {
        if length < i+ConstHeaderLength+ConstMLength {
            break
        }
        if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
            messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
            //if length < i+ConstHeaderLength+ConstLength+messageLength {
            //    break
            //}
            data := buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]
            readerChannel <- data

        }
    }

    if i == length {
        return make([]byte, 0)
    }
    return buffer[i:]
}

//整形转换成字节
func IntToBytes(n int) []byte {
    x := int32(n)

    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.BigEndian, x)
    return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
    bytesBuffer := bytes.NewBuffer(b)

    var x int32
    binary.Read(bytesBuffer, binary.BigEndian, &x)

    return int(x)
}
//package protocal
//
//import (
//    "fmt"
//)
//
//const COMMAND (
//    ConstLogin = 100
//    ConstLogout = 101
//    ConstMsgSingle = 201
//    ConstMsgGroup = 202
//    ConstMsgBroad = 203
//)
//
//type Msg_Header {
//    cmd int
//    length int
//}
//
//
//func login (user string, passwd string) int {
//    msg := &Msg_Header {
//        ConstLogin,
//        0,
//    }
//    
//    fmt.Print("login", msg)
//    //p1 := &Person{ 
//    //    "zhangsan", 
//    //    25, 
//    //    []string{"lisi", "wangwu"}, 
//    //    "Jinlin China", 
//    //} 
//    //p, err := json.Marshal(p1) 
//    return 0
//}
//
