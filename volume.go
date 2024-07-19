package pulse

import "github.com/jfreymuth/pulse/proto"

// SetSinkVolume sets volume on a sink with the given name.
func (c *Client) SetSinkVolume(name string, volume proto.ChannelVolumes) error {
	err := c.c.Request(&proto.SetSinkVolume{SinkIndex: proto.Undefined, SinkName: name, ChannelVolumes: volume}, nil)
	if err != nil {
		return err
	}
	return nil
}

// SetSinkMute sets mute flag on a sink with the given name.
func (c *Client) SetSinkMute(name string, mute bool) error {
	err := c.c.Request(&proto.SetSinkMute{SinkIndex: proto.Undefined, SinkName: name, Mute: mute}, nil)
	if err != nil {
		return err
	}
	return nil
}
