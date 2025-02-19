// Package endpointf comment
// This file was generated by tars2go 1.2.1
// Generated from EndpointF.tars
package endpointf

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// EndpointF struct implement
type EndpointF struct {
	Host        string `json:"host"`
	Port        int32  `json:"port"`
	Timeout     int32  `json:"timeout"`
	Istcp       int32  `json:"istcp"`
	Grid        int32  `json:"grid"`
	Groupworkid int32  `json:"groupworkid"`
	Grouprealid int32  `json:"grouprealid"`
	SetId       string `json:"setId"`
	Qos         int32  `json:"qos"`
	BakFlag     int32  `json:"bakFlag"`
	Weight      int32  `json:"weight"`
	WeightType  int32  `json:"weightType"`
	AuthType    int32  `json:"authType"`
}

func (st *EndpointF) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *EndpointF) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Host, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Port, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Timeout, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Istcp, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Grid, 4, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Groupworkid, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Grouprealid, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SetId, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Qos, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.BakFlag, 9, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Weight, 11, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.WeightType, 12, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.AuthType, 13, false)
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
func (st *EndpointF) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require EndpointF, but not exist. tag %d", tag)
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
func (st *EndpointF) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Host, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Port, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Timeout, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Istcp, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Grid, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Groupworkid, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Grouprealid, 6)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SetId, 7)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Qos, 8)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.BakFlag, 9)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Weight, 11)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.WeightType, 12)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.AuthType, 13)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *EndpointF) WriteBlock(buf *codec.Buffer, tag byte) error {
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
