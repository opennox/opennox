//go:build !server

package opennox

import (
	"github.com/opennox/opennox/v1/client"
)

func (c *Client) Nox_xxx_client_4984B0_drawable(dr *client.Drawable) bool {
	if dr.DrawFuncPtr == nil {
		return false
	}
	if dr == c.ClientPlayerUnit() {
		return true
	}
	return c.Sight.Nox_xxx_client_4984B0_drawable_A(c.Viewport(), dr)
}

func (c *Client) nox_xxx_drawBlack_496150_B() {
	arr := c.Sight.Nox_xxx_drawBlack_496150_C()
	if len(arr) == 0 {
		return
	}
	c.sub_4C52E0(arr)
}
