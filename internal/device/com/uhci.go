//*
// uHCI Micro Host Controller Interface
// ------------------------------------
//
// Copyright © 2025 dronectl

package com

import (
	"fmt"
	"log/slog"
	"net"

	v1 "github.com/dronectl/rdt/gen/raptor/v1"
	proto "google.golang.org/protobuf/proto"
)

type UHCICtx struct {
	logger *slog.Logger
	conn   *net.UDPConn
	addr   *net.UDPAddr
}

func (u *UHCICtx) startDiscoveryService() error {
	u.logger.Info("Starting UDP discovery service", "port", v1.UHCIPort_UHCI_PORT_DISCOVERY)
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", v1.UHCIPort_UHCI_PORT_DISCOVERY))
	if err != nil {
		u.logger.Error("Error resolving address:", "err", err)
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		u.logger.Error("Error creating connection:", "err", err)
		return err
	}
	defer conn.Close()
	fmt.Println("UDP server is up and listening on port 8080")

	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			u.logger.Warn("Error reading from UDP connection:", "error", err)
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
		u.logger.Error("Failed to read from connection:", "error", err)
		return
	}

	u.logger.Info("Received:", "buf", buf)
	req := &v1.UHCIBaseRequest{}
	if err := proto.Unmarshal(buf, req); err != nil {
		u.logger.Error("Failed to parse address book:", "error", err)
	}
	u.logger.Info("Marshalled: ", "req", req)
}

func (u *UHCICtx) startCommandService() error {
	u.logger.Info("Starting TCP command service")
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", v1.UHCIPort_UHCI_PORT_COMMAND))
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
		logger: slog.Default(),
	}
}
