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

type UHCIHandle struct {
	logger *slog.Logger
	conn   *net.UDPConn
	addr   *net.UDPAddr
}

func (u *UHCIHandle) startDiscoveryService() error {
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
	u.logger.Debug("UDP server is up and listening", "port", v1.UHCIPort_UHCI_PORT_DISCOVERY)
	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			u.logger.Error("Error reading from UDP connection:", "error", err)
			continue
		}
		u.logger.Info("Received message", "bytes", n, "addr", remoteAddr, "buf", buffer)
        req := &v1.UHCIBaseRequest{}
        resp := &v1.UHCIBaseResponse{}
        if err := proto.Unmarshal(buffer, req); err != nil {
            u.logger.Error("Failed to unmarshal request", "error", err)
            resp.Status = v1.UHCIStatus_UHCI_STATUS_DECODE_ERR
        } else {
            resp.Status = v1.UHCIStatus_UHCI_STATUS_OK
            resp.ResponseMux = &v1.UHCIBaseResponse_Uhci{
                Uhci: &v1.UHCIProtocolResponse{
                    Status: v1.UHCIProtocolStatus_UHCI_PROTOCOL_STATUS_OK,
                    ResponseMux: &v1.UHCIProtocolResponse_Discovery{
                        Discovery: &v1.UHCIDiscoveryResponse{
                            Device: &v1.DeviceMetadata{
                                Uuid: 12345678,
                                FirmwareVersion: &v1.Version{
                                    Major: 1,
                                    Minor: 0,
                                    Patch: 0,
                                },
                                HardwareVersion: &v1.Version{
                                    Major: 1,
                                    Minor: 0,
                                    Patch: 0,
                                },
                                DigitalTwin: true,
                            },
                        },
                    },
                },
            }
        }
        respBuffer, err := proto.Marshal(resp)
        if err != nil {
            u.logger.Error("Failed to marshal response", "error", err)
        }
		_, err = conn.WriteToUDP(respBuffer, remoteAddr)
		if err != nil {
			u.logger.Error("Error responding to client", "error", err)
		}
	}
}

func (u *UHCIHandle) handleConnection(conn net.Conn) {
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

func (u *UHCIHandle) startCommandService() error {
	u.logger.Info("Starting TCP command service")
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", v1.UHCIPort_UHCI_PORT_COMMAND))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			u.logger.Error("Failed to accept connection", "error", err, "conn", conn)
			continue
		}
		go u.handleConnection(conn)
	}
}

func (u *UHCIHandle) Start() {
	go u.startDiscoveryService()
	go u.startCommandService()
}

func NewUHCIHandle() *UHCIHandle {
	return &UHCIHandle{
		logger: slog.Default(),
	}
}
