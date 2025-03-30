
package com

import (
    "os"
    "log"
    "net"

    "github.com/dronectl/rdt/gen/raptor/v1"
)


type UHCICtx struct {
    logger *log.Logger
    conn *net.UDPConn
    addr *net.UDPAddr
}

func (u *UHCICtx) startDiscoveryService() error {
    u.logger.Println("Starting UDP discovery Service")
    for {
        buf := make([]byte, 1024)
        n, addr, err := u.conn.ReadFromUDP(buf)
        if err != nil {
            u.logger.Println("Error reading from UDP connection")
            continue
        }
    }
}

func (u *UHCICtx) startCommandService() error {
    // Listen for incoming connections on port 8080
    v1.CommandBaseRequest_DeviceMetadata
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Accept incoming connections and handle them
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Handle the connection in a new goroutine
        go handleConnection(conn)
    }
}

func (u *UHCICtx) Start() error {
    go u.startDiscoveryService()

    return nil
}

func NewUHCI() *UHCICtx {
    return &UHCICtx{
        logger: log.New(os.Stdout, "UHCI: ", log.LstdFlags),
    }
}
