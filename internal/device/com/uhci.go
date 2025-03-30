//*
// uHCI Micro Host Controller Interface
// ------------------------------------
//
// Copyright © 2025 dronectl

package com

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/dronectl/rdt/gen/raptor/v1"
    proto "google.golang.org/protobuf/proto"
)

type UHCICtx struct {
    logger *log.Logger
    conn *net.UDPConn
    addr *net.UDPAddr
}

func (u *UHCICtx) startDiscoveryService() error {
    u.logger.Printf("Starting UDP discovery service on port %d", v1.UHCIPort_UHCI_PORT_DISCOVERY)
    addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", v1.UHCIPort_UHCI_PORT_DISCOVERY))
    if err != nil {
        u.logger.Fatalln("Error resolving address:", err)
    }

    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        u.logger.Fatalln("Error creating connection:", err)
    }
    defer conn.Close()
    fmt.Println("UDP server is up and listening on port 8080")

    buffer := make([]byte, 1024)
    for {
        n, remoteAddr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println("Error reading from UDP connection:", err)
            continue
        }
        fmt.Printf("Received message from %s: %s\n", remoteAddr, string(buffer[:n]))
        _, err = conn.WriteToUDP([]byte("Message received"), remoteAddr)
        if err != nil {
            fmt.Println("Error responding to client:", err)
            continue
        }
    }
}

func (u *UHCICtx) handleConnection(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 1024)
    _, err := conn.Read(buf)
    if err != nil {
        u.logger.Println(err)
        return
    }

    u.logger.Printf("Received: %s", buf)
    req := &v1.UHCIBaseRequest{}
    if err := proto.Unmarshal(buf, req); err != nil {
        log.Println("Failed to parse address book:", err)
    }
    u.logger.Printf("Marshalled: %t", req)
}

func (u *UHCICtx) startCommandService() error {
    u.logger.Println("Starting TCP command service")
    // Listen for incoming connections on port 8080
    ln, err := net.Listen("tcp", fmt.Sprintf(":%u", v1.UHCIPort_UHCI_PORT_COMMAND))
    if err != nil {
        fmt.Println(err)
        return err
    }

    // Accept incoming connections and handle them
    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        // Handle the connection in a new goroutine
        go u.handleConnection(conn)
    }
}

func (u *UHCICtx) Start() error {
    go u.startDiscoveryService()
    u.startCommandService()
    return nil
}

func NewUHCI() *UHCICtx {
    return &UHCICtx{
        logger: log.New(os.Stdout, "UHCI: ", log.LstdFlags),
    }
}
