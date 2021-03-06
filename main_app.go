package main

import(
    "flag"
    "log"
    "os"
    "github.com/conorbrady/go-chat-server/tcpserver"
    "github.com/conorbrady/go-chat-server/protocol"
    "github.com/conorbrady/go-chat-server/chat"
    )

func main (){

    port := flag.Int("port",-1,"Port to listen on")
    threadCount := flag.Int("threadCount", 100, "Available thread count")

    flag.Parse()

    if *port<0 {
        log.Fatal("Must pass port via -port x flag")
    }

    tcpServer := tcpserver.MakeTCPServer(os.Getenv("IP_ADDRESS"),*port,*threadCount)

    chat := chat.MakeChat()

    tcpServer.AddProtocol(protocol.MakeHelo(os.Getenv("IP_ADDRESS"),*port,4))
    tcpServer.AddProtocol(protocol.MakeChatJoinProtocol(chat,4))
    tcpServer.AddProtocol(protocol.MakeChatLeaveProtocol(chat,4))
    tcpServer.AddProtocol(protocol.MakeChatMessageProtocol(chat,4))
    tcpServer.AddProtocol(protocol.MakeDisconnectProtocol(chat,1))

    tcpServer.BlockingRun()
}
