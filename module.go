package pulse

import "github.com/jfreymuth/pulse/proto"

// LoadModule loads module described by the definition.
func (c *Client) LoadModule(name string, args string) (uint32, error) {
	var res proto.LoadModuleReply
	err := c.c.Request(&proto.LoadModule{Name: name, Args: args}, &res)
	if err != nil {
		return 0, err
	}
	return res.ModuleIndex, nil
}

// UnloadModule unloads module by index.
func (c *Client) UnloadModule(idx uint32) error {
	err := c.c.Request(&proto.UnloadModule{ModuleIndex: idx}, nil)
	if err != nil {
		return err
	}
	return nil
}

// ListModules list all loaded modules.
func (c *Client) ListModules() ([]*proto.GetModuleInfoReply, error) {
	var res proto.GetModuleInfoListReply
	err := c.c.Request(&proto.GetModuleInfoList{}, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
