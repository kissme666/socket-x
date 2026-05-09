package socks

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strconv"

	"github.com/kissme666/socketx/internal/config"
)

var ErrSocksCmdNotSupport = errors.New("socks cmd is not support")

type ServerV5 struct {
	config *config.SocksConfig
}

func NewServerV5(config *config.SocksConfig) *ServerV5 {
	return &ServerV5{config: config}
}

func (s *ServerV5) Run() error {
	s.serve()
	return nil
}

func (s *ServerV5) serve() {
	slog.Info("start socks5 server", "addr", s.config.Addr)
	l, err := net.Listen("tcp", s.config.Addr)
	if err != nil {
		slog.Error("socks5 listener failed", "err", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error("socks5 accept failed", "err", err)
			continue
		}
		go s.handle(conn)
	}
}

func (s *ServerV5) handle(left net.Conn) {
	defer left.Close()

	addr, err := s.socks(left)
	if err != nil {
		slog.Error("socks5 handshake failed", "err", err)
		return
	}

	right, err := net.Dial("tcp", addr)
	if err != nil {
		slog.Error("connect target failed", "err", err)
		return
	}
	defer right.Close()

	slog.Info("socks5 relay", "left", left.LocalAddr().String(), "right", addr)
	stopErrChan := make(chan error, 2)
	go func() {
		_, err1 := io.Copy(left, right)
		stopErrChan <- err1
	}()
	go func() {
		_, err2 := io.Copy(right, left)
		stopErrChan <- err2
	}()

	var stopErr error
	for i := 0; i < 2; i++ {
		err := <-stopErrChan
		stopErr = errors.Join(stopErr, err)
	}
	if stopErr != nil {
		slog.Error("relay error", "err", stopErr)
	}
}

func (s *ServerV5) socks(conn net.Conn) (string, error) {
	// handshake
	bs := make([]byte, 256)
	if _, err := conn.Read(bs); err != nil {
		slog.Error("socks5 handshake request not legal", "raw", fmt.Sprintf("%X", bs), "err", err)
		return "", err
	}

	if _, err := conn.Write([]byte{0x05, 0x00}); err != nil {
		slog.Error("socks5 handshake response failed", "raw", fmt.Sprintf("%X", bs), "err", err)
		return "", err
	}

	// request
	if _, err := conn.Read(bs); err != nil {
		slog.Error("socks5 request response failed", "raw", fmt.Sprintf("%X", bs), "err", err)
		return "", err
	}
	buf := bytes.NewBuffer(bs)
	buf.Next(1)
	cmd, err := buf.ReadByte()
	if err != nil {
		return "", err
	}
	switch cmd {
	case 0x01:
	default:
		return "", ErrSocksCmdNotSupport
	}
	buf.Next(1)
	atyp, err := buf.ReadByte()
	if err != nil {
		return "", err
	}

	// parse addr
	var addr string
	switch atyp {
	case 0x01:
		addr = net.IP(buf.Next(4)).String()
	case 0x03:
		dl, err := buf.ReadByte()
		if err != nil {
			return "", err
		}
		addr = string(buf.Next(int(dl)))
	case 0x04:
		addr = net.IP(buf.Next(16)).String()
	}
	port := int(binary.BigEndian.Uint16(buf.Next(2)))
	if port >= 1 && port <= 65535 {
		addr = addr + ":" + strconv.Itoa(port)
	}

	// resp parse addr complete and start relaying
	if _, err = conn.Write([]byte{5, 0, 0, 1, 0, 0, 0, 0, 0, 0}); err != nil {
		slog.Error("socks5 response failed", "err", err)
		return "", err
	}

	slog.Debug("socks request success", "addr", addr)
	return addr, nil
}
