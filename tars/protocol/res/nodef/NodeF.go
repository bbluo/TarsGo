// Package nodef comment
// This file was generated by tars2go 1.2.1
// Generated from NodeF.tars
package nodef

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// ServerInfo struct implement
type ServerInfo struct {
	Application string `json:"application"`
	ServerName  string `json:"serverName"`
	Pid         int32  `json:"pid"`
	Adapter     string `json:"adapter"`
}

func (st *ServerInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ServerInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Application, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.ServerName, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Pid, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Adapter, 3, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *ServerInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ServerInfo, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *ServerInfo) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Application, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.ServerName, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Pid, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Adapter, 3)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ServerInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}
