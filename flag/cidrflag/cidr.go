package cidrflag

import (
	"net"
)

// CIDR is a `flag.Value` for CIDR notation IP address and prefix length arguments.
type CIDR struct {
	Value struct {
		IPNet *net.IPNet
		IP    net.IP
	}
	Text string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *CIDR) Help() string {
	return "a CIDR notation IP address and prefix length"
}

// Set is flag.Value.Set
func (fv *CIDR) Set(v string) error {
	ip, ipNet, err := net.ParseCIDR(v)
	if err != nil {
		return err
	}
	fv.Text = v
	fv.Value = struct {
		IPNet *net.IPNet
		IP    net.IP
	}{IP: ip, IPNet: ipNet}
	return nil
}

func (fv *CIDR) String() string {
	return fv.Text
}
