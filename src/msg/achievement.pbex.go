// Code generated by protoc-gen-pbex-go. DO NOT EDIT IT!!!
// source: achievement.proto

package msg

import (
	json "encoding/json"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	sync "sync"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [Achievement] begin
func (m *Achievement) ResetEx() {

	m.Id = 0

	m.Type = 0

}

func (m Achievement) Clone() *Achievement {
	n, ok := g_Achievement_Pool.Get().(*Achievement)
	if !ok || n == nil {
		n = &Achievement{}
	}

	n.Id = m.Id

	n.Type = m.Type

	return n
}

func Clone_Achievement_Slice(dst []*Achievement, src []*Achievement) []*Achievement {
	for _, i := range dst {
		Put_Achievement(i)
	}
	if len(src) > 0 {
		dst = make([]*Achievement, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*Achievement{}
		dst = nil
	}
	return dst
}

func (m Achievement) JsonString() string {
	ba, _ := json.Marshal(m)
	return "Achievement:" + string(ba)
}

func New_Achievement() *Achievement {
	m := &Achievement{}
	return m
}

var g_Achievement_Pool = sync.Pool{}

func Get_Achievement() *Achievement {
	m, ok := g_Achievement_Pool.Get().(*Achievement)
	if !ok {
		m = New_Achievement()
	} else {
		if m == nil {
			m = New_Achievement()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_Achievement(i interface{}) {
	if m, ok := i.(*Achievement); ok && m != nil {
		g_Achievement_Pool.Put(i)
	}
}

// message [Achievement] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [C2S_GetCompletedAchievements] begin
func (m *C2S_GetCompletedAchievements) ResetEx() {

}

func (m C2S_GetCompletedAchievements) Clone() *C2S_GetCompletedAchievements {
	n, ok := g_C2S_GetCompletedAchievements_Pool.Get().(*C2S_GetCompletedAchievements)
	if !ok || n == nil {
		n = &C2S_GetCompletedAchievements{}
	}

	return n
}

func Clone_C2S_GetCompletedAchievements_Slice(dst []*C2S_GetCompletedAchievements, src []*C2S_GetCompletedAchievements) []*C2S_GetCompletedAchievements {
	for _, i := range dst {
		Put_C2S_GetCompletedAchievements(i)
	}
	if len(src) > 0 {
		dst = make([]*C2S_GetCompletedAchievements, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*C2S_GetCompletedAchievements{}
		dst = nil
	}
	return dst
}

func (m C2S_GetCompletedAchievements) JsonString() string {
	ba, _ := json.Marshal(m)
	return "C2S_GetCompletedAchievements:" + string(ba)
}

func New_C2S_GetCompletedAchievements() *C2S_GetCompletedAchievements {
	m := &C2S_GetCompletedAchievements{}
	return m
}

var g_C2S_GetCompletedAchievements_Pool = sync.Pool{}

func Get_C2S_GetCompletedAchievements() *C2S_GetCompletedAchievements {
	m, ok := g_C2S_GetCompletedAchievements_Pool.Get().(*C2S_GetCompletedAchievements)
	if !ok {
		m = New_C2S_GetCompletedAchievements()
	} else {
		if m == nil {
			m = New_C2S_GetCompletedAchievements()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_C2S_GetCompletedAchievements(i interface{}) {
	if m, ok := i.(*C2S_GetCompletedAchievements); ok && m != nil {
		g_C2S_GetCompletedAchievements_Pool.Put(i)
	}
}

// message [C2S_GetCompletedAchievements] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// message [S2C_GetCompletedAchievements] begin
func (m *S2C_GetCompletedAchievements) ResetEx() {

	for _, i := range m.Achievements {
		Put_Achievement(i)
	}

	//m.Achievements = []*Achievement{}
	m.Achievements = nil

}

func (m S2C_GetCompletedAchievements) Clone() *S2C_GetCompletedAchievements {
	n, ok := g_S2C_GetCompletedAchievements_Pool.Get().(*S2C_GetCompletedAchievements)
	if !ok || n == nil {
		n = &S2C_GetCompletedAchievements{}
	}

	if len(m.Achievements) > 0 {
		n.Achievements = make([]*Achievement, len(m.Achievements))
		for i, e := range m.Achievements {

			if e != nil {
				n.Achievements[i] = e.Clone()
			}

		}
	} else {
		//n.Achievements = []*Achievement{}
		n.Achievements = nil
	}

	return n
}

func Clone_S2C_GetCompletedAchievements_Slice(dst []*S2C_GetCompletedAchievements, src []*S2C_GetCompletedAchievements) []*S2C_GetCompletedAchievements {
	for _, i := range dst {
		Put_S2C_GetCompletedAchievements(i)
	}
	if len(src) > 0 {
		dst = make([]*S2C_GetCompletedAchievements, len(src))
		for i, e := range src {
			if e != nil {
				dst[i] = e.Clone()
			}
		}
	} else {
		//dst = []*S2C_GetCompletedAchievements{}
		dst = nil
	}
	return dst
}

func (m S2C_GetCompletedAchievements) JsonString() string {
	ba, _ := json.Marshal(m)
	return "S2C_GetCompletedAchievements:" + string(ba)
}

func New_S2C_GetCompletedAchievements() *S2C_GetCompletedAchievements {
	m := &S2C_GetCompletedAchievements{}
	return m
}

var g_S2C_GetCompletedAchievements_Pool = sync.Pool{}

func Get_S2C_GetCompletedAchievements() *S2C_GetCompletedAchievements {
	m, ok := g_S2C_GetCompletedAchievements_Pool.Get().(*S2C_GetCompletedAchievements)
	if !ok {
		m = New_S2C_GetCompletedAchievements()
	} else {
		if m == nil {
			m = New_S2C_GetCompletedAchievements()
		} else {
			m.ResetEx()
		}
	}
	return m
}

func Put_S2C_GetCompletedAchievements(i interface{}) {
	if m, ok := i.(*S2C_GetCompletedAchievements); ok && m != nil {
		g_S2C_GetCompletedAchievements_Pool.Put(i)
	}
}

// message [S2C_GetCompletedAchievements] end
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////