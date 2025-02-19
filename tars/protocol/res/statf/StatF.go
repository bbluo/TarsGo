// Package statf comment
// This file was generated by tars2go 1.2.1
// Generated from StatF.tars
package statf

import (
	"fmt"

	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// StatMicMsgHead struct implement
type StatMicMsgHead struct {
	MasterName    string `json:"masterName"`
	SlaveName     string `json:"slaveName"`
	InterfaceName string `json:"interfaceName"`
	MasterIp      string `json:"masterIp"`
	SlaveIp       string `json:"slaveIp"`
	SlavePort     int32  `json:"slavePort"`
	ReturnValue   int32  `json:"returnValue"`
	SlaveSetName  string `json:"slaveSetName"`
	SlaveSetArea  string `json:"slaveSetArea"`
	SlaveSetID    string `json:"slaveSetID"`
	TarsVersion   string `json:"tarsVersion"`
}

func (st *StatMicMsgHead) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *StatMicMsgHead) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.MasterName, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveName, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.InterfaceName, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.MasterIp, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveIp, 4, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.SlavePort, 5, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.ReturnValue, 6, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveSetName, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveSetArea, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveSetID, 9, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.TarsVersion, 10, false)
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
func (st *StatMicMsgHead) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require StatMicMsgHead, but not exist. tag %d", tag)
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
func (st *StatMicMsgHead) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.MasterName, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveName, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.InterfaceName, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.MasterIp, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveIp, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.SlavePort, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.ReturnValue, 6)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveSetName, 7)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveSetArea, 8)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveSetID, 9)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.TarsVersion, 10)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *StatMicMsgHead) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// StatMicMsgBody struct implement
type StatMicMsgBody struct {
	Count         int32           `json:"count"`
	TimeoutCount  int32           `json:"timeoutCount"`
	ExecCount     int32           `json:"execCount"`
	IntervalCount map[int32]int32 `json:"intervalCount"`
	TotalRspTime  int64           `json:"totalRspTime"`
	MaxRspTime    int32           `json:"maxRspTime"`
	MinRspTime    int32           `json:"minRspTime"`
}

func (st *StatMicMsgBody) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *StatMicMsgBody) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Count, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.TimeoutCount, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.ExecCount, 2, true)
	if err != nil {
		return err
	}

	_, err = readBuf.SkipTo(codec.MAP, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&length, 0, true)
	if err != nil {
		return err
	}

	st.IntervalCount = make(map[int32]int32)
	for i0, e0 := int32(0), length; i0 < e0; i0++ {
		var k0 int32
		var v0 int32

		err = readBuf.ReadInt32(&k0, 0, false)
		if err != nil {
			return err
		}

		err = readBuf.ReadInt32(&v0, 1, false)
		if err != nil {
			return err
		}

		st.IntervalCount[k0] = v0
	}

	err = readBuf.ReadInt64(&st.TotalRspTime, 4, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.MaxRspTime, 5, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.MinRspTime, 6, true)
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
func (st *StatMicMsgBody) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require StatMicMsgBody, but not exist. tag %d", tag)
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
func (st *StatMicMsgBody) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Count, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.TimeoutCount, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.ExecCount, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.IntervalCount)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.IntervalCount {

		err = buf.WriteInt32(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteInt32(v1, 1)
		if err != nil {
			return err
		}

	}

	err = buf.WriteInt64(st.TotalRspTime, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.MaxRspTime, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.MinRspTime, 6)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *StatMicMsgBody) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// StatSampleMsg struct implement
type StatSampleMsg struct {
	Unid          string `json:"unid"`
	MasterName    string `json:"masterName"`
	SlaveName     string `json:"slaveName"`
	InterfaceName string `json:"interfaceName"`
	MasterIp      string `json:"masterIp"`
	SlaveIp       string `json:"slaveIp"`
	Depth         int32  `json:"depth"`
	Width         int32  `json:"width"`
	ParentWidth   int32  `json:"parentWidth"`
}

func (st *StatSampleMsg) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *StatSampleMsg) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Unid, 0, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.MasterName, 1, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveName, 2, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.InterfaceName, 3, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.MasterIp, 4, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.SlaveIp, 5, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Depth, 6, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Width, 7, true)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.ParentWidth, 8, true)
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
func (st *StatSampleMsg) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require StatSampleMsg, but not exist. tag %d", tag)
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
func (st *StatSampleMsg) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Unid, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.MasterName, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveName, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.InterfaceName, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.MasterIp, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.SlaveIp, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Depth, 6)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Width, 7)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.ParentWidth, 8)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *StatSampleMsg) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// ProxyInfo struct implement
type ProxyInfo struct {
	BFromClient bool `json:"bFromClient"`
}

func (st *ProxyInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *ProxyInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadBool(&st.BFromClient, 0, true)
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
func (st *ProxyInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require ProxyInfo, but not exist. tag %d", tag)
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
func (st *ProxyInfo) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteBool(st.BFromClient, 0)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *ProxyInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
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
