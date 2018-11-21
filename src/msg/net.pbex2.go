// Code generated by protoc-gen-pbex2-go. DO NOT EDIT IT!!!
// source: net.proto

/*
It has these top-level messages:
	S2C_DisConn
	C2S_Ping
	S2C_Pong
*/

package msg

import "sync"
import protocol "github.com/trist725/mgsu/network/protocol/protobuf/v2"

var _ *sync.Pool
var _ = protocol.PH

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_DisConn] begin
func (m *S2C_DisConn) ResetEx() {
	m.UserId = 0

}

func (m S2C_DisConn) Clone() *S2C_DisConn {
	n, ok := g_S2C_DisConn_Pool.Get().(*S2C_DisConn)
	if !ok || n == nil {
		n = &S2C_DisConn{}
	}

	n.UserId = m.UserId

	return n
}

func Clone_S2C_DisConn_Slice(dst []*S2C_DisConn, src []*S2C_DisConn) []*S2C_DisConn {
	for _, i := range dst {
		Put_S2C_DisConn(i)
	}
	dst = []*S2C_DisConn{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_DisConn() *S2C_DisConn {
	m := &S2C_DisConn{}
	return m
}

var g_S2C_DisConn_Pool = sync.Pool{}

func Get_S2C_DisConn() *S2C_DisConn {
	m, ok := g_S2C_DisConn_Pool.Get().(*S2C_DisConn)
	if !ok {
		m = New_S2C_DisConn()
	} else {
		if m == nil {
			m = New_S2C_DisConn()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_DisConn(i interface{}) {
	if m, ok := i.(*S2C_DisConn); ok && m != nil {
		g_S2C_DisConn_Pool.Put(i)
	}
}

// message [S2C_DisConn] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_Ping] begin
func (m *C2S_Ping) ResetEx() {

}

func (m C2S_Ping) Clone() *C2S_Ping {
	n, ok := g_C2S_Ping_Pool.Get().(*C2S_Ping)
	if !ok || n == nil {
		n = &C2S_Ping{}
	}

	return n
}

func Clone_C2S_Ping_Slice(dst []*C2S_Ping, src []*C2S_Ping) []*C2S_Ping {
	for _, i := range dst {
		Put_C2S_Ping(i)
	}
	dst = []*C2S_Ping{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_C2S_Ping() *C2S_Ping {
	m := &C2S_Ping{}
	return m
}

var g_C2S_Ping_Pool = sync.Pool{}

func Get_C2S_Ping() *C2S_Ping {
	m, ok := g_C2S_Ping_Pool.Get().(*C2S_Ping)
	if !ok {
		m = New_C2S_Ping()
	} else {
		if m == nil {
			m = New_C2S_Ping()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_Ping(i interface{}) {
	if m, ok := i.(*C2S_Ping); ok && m != nil {
		g_C2S_Ping_Pool.Put(i)
	}
}

// message [C2S_Ping] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_Pong] begin
func (m *S2C_Pong) ResetEx() {

}

func (m S2C_Pong) Clone() *S2C_Pong {
	n, ok := g_S2C_Pong_Pool.Get().(*S2C_Pong)
	if !ok || n == nil {
		n = &S2C_Pong{}
	}

	return n
}

func Clone_S2C_Pong_Slice(dst []*S2C_Pong, src []*S2C_Pong) []*S2C_Pong {
	for _, i := range dst {
		Put_S2C_Pong(i)
	}
	dst = []*S2C_Pong{}

	for _, i := range src {
		dst = append(dst, i.Clone())
	}

	return dst
}

func New_S2C_Pong() *S2C_Pong {
	m := &S2C_Pong{}
	return m
}

var g_S2C_Pong_Pool = sync.Pool{}

func Get_S2C_Pong() *S2C_Pong {
	m, ok := g_S2C_Pong_Pool.Get().(*S2C_Pong)
	if !ok {
		m = New_S2C_Pong()
	} else {
		if m == nil {
			m = New_S2C_Pong()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_Pong(i interface{}) {
	if m, ok := i.(*S2C_Pong); ok && m != nil {
		g_S2C_Pong_Pool.Put(i)
	}
}

// message [S2C_Pong] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////