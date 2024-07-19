package proto

import (
	"fmt"
	"math"
)

const Undefined = 0xFFFFFFFF

const (
	FormatUint8     = 0
	FormatInt16LE   = 3
	FormatInt16BE   = 4
	FormatFloat32LE = 5
	FormatFloat32BE = 6
	FormatInt32LE   = 7
	FormatInt32BE   = 8
)

type Channel byte

func (c Channel) String() string {
	if s, ok := channels[c]; ok {
		return s
	}
	return fmt.Sprintf("%d", c)
}

const (
	ChannelMono           Channel = 0
	ChannelLeft           Channel = 1
	ChannelRight          Channel = 2
	ChannelCenter         Channel = 3
	ChannelFrontLeft      Channel = 1
	ChannelFrontRight     Channel = 2
	ChannelFrontCenter    Channel = 3
	ChannelRearCenter     Channel = 4
	ChannelRearLeft       Channel = 5
	ChannelRearRight      Channel = 6
	ChannelLFE            Channel = 7
	ChannelLeftCenter     Channel = 8
	ChannelRightCenter    Channel = 9
	ChannelLeftSide       Channel = 10
	ChannelRightSide      Channel = 11
	ChannelAux0           Channel = 12
	ChannelAux31          Channel = 43
	ChannelTopCenter      Channel = 44
	ChannelTopFrontLeft   Channel = 45
	ChannelTopFrontRight  Channel = 46
	ChannelTopFrontCenter Channel = 47
	ChannelTopRearLeft    Channel = 48
	ChannelTopRearRight   Channel = 49
	ChannelTopRearCenter  Channel = 50
)

var channels = map[Channel]string{
	ChannelMono:           "mono",
	ChannelFrontLeft:      "front-left",
	ChannelFrontRight:     "front-right",
	ChannelFrontCenter:    "front-center",
	ChannelRearCenter:     "rear-center",
	ChannelRearLeft:       "rear-left",
	ChannelRearRight:      "rear-right",
	ChannelLFE:            "LFE",
	ChannelLeftCenter:     "left-center",
	ChannelRightCenter:    "right-center",
	ChannelLeftSide:       "left-side",
	ChannelRightSide:      "right-side",
	ChannelAux0:           "aux0",
	ChannelAux31:          "aux31",
	ChannelTopCenter:      "top-center",
	ChannelTopFrontLeft:   "top-front-left",
	ChannelTopFrontRight:  "top-front-right",
	ChannelTopFrontCenter: "top-front-center",
	ChannelTopRearLeft:    "top-rear-left",
	ChannelTopRearRight:   "top-rear-right",
	ChannelTopRearCenter:  "top-rear-center",
}

const (
	EncodingPCM = 1
)

type SampleSpec struct {
	Format   byte
	Channels byte
	Rate     uint32
}

type Microseconds uint64

type ChannelMap []Channel

type ChannelVolumes []uint32

type Time struct {
	Seconds      uint32
	Microseconds uint32
}

type Volume uint32

const (
	// Muted (minimal valid) volume (0%, -inf dB)
	VolumeMuted Volume = 0
	// Normal volume (100%, 0 dB)
	VolumeNorm Volume = 0x10000
	// Maximum valid volume we can store.
	VolumeMax Volume = math.MaxUint32 / 2
	// Special 'invalid' volume.
	VolumeInvalid Volume = math.MaxUint32
)

type FormatInfo struct {
	Encoding   byte
	Properties PropList
}

type SubscriptionMask uint32

const (
	SubscriptionMaskSink SubscriptionMask = 1 << iota
	SubscriptionMaskSource
	SubscriptionMaskSinkInput
	SubscriptionMaskSourceInput
	SubscriptionMaskModule
	SubscriptionMaskClient
	SubscriptionMaskSampleCache
	SubscriptionMaskServer
	SubscriptionMaskAutoload
	SubscriptionMaskCard

	SubscriptionMaskNull SubscriptionMask = 0
	SubscriptionMaskAll  SubscriptionMask = 0x02ff
)

type SubscriptionEventType uint32

const (
	EventSink SubscriptionEventType = iota
	EventSource
	EventSinkSinkInput
	EventSinkSourceOutput
	EventModule
	EventClient
	EventSampleCache
	EventServer
	EventAutoload
	EventCard
	EventFacilityMask SubscriptionEventType = 0xf

	EventNew      SubscriptionEventType = 0x0000
	EventChange   SubscriptionEventType = 0x0010
	EventRemove   SubscriptionEventType = 0x0020
	EventTypeMask SubscriptionEventType = 0x0030
)

func (e SubscriptionEventType) GetFacility() SubscriptionEventType {
	return e & EventFacilityMask
}

func (e SubscriptionEventType) GetType() SubscriptionEventType {
	return e & EventTypeMask
}

func (e SubscriptionEventType) String() string {
	var res string
	switch e.GetType() {
	case EventNew:
		res += "new"
	case EventChange:
		res += "change"
	case EventRemove:
		res += "remove"
	default:
		return "<invalid type>"
	}
	res += " "
	switch e.GetFacility() {
	case EventSink:
		res += "sink"
	case EventSource:
		res += "source"
	case EventSinkSinkInput:
		res += "sink input"
	case EventSinkSourceOutput:
		res += "source output"
	case EventModule:
		res += "module"
	case EventClient:
		res += "client"
	case EventSampleCache:
		res += "sample cache"
	case EventServer:
		res += "server"
	case EventAutoload:
		res += "autoload"
	case EventCard:
		res += "card"
	default:
		return "<invalid facility>"
	}
	return res
}

type PropList map[string]PropListEntry

type PropListEntry []byte

func PropListString(s string) PropListEntry {
	e := make(PropListEntry, len(s)+1)
	copy(e, s)
	return e
}
func (e PropListEntry) String() string {
	if len(e) == 0 || e[len(e)-1] != '\x00' {
		return "<not a string>"
	}
	return string(e[:len(e)-1])
}
