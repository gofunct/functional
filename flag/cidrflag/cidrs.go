package cidrflag

import (
	"net"
	"strings"
)

// CIDRs is a `flag.Value` for CIDR notation IP address and prefix length arguments.
type CIDRs struct {
	Values []struct {
		IPNet *net.IPNet
		IP    net.IP
	}
	Texts []string
}

// Help returns a string suitable for inclusion in a flag help message.
func (fv *CIDRs) Help() string {
	return "a CIDR notation IP address and prefix length"
}

// Set is flag.Value.Set
func (fv *CIDRs) Set(v string) error {
	ip, ipNet, err := net.ParseCIDR(v)
	if err != nil {
		return err
	}
	fv.Texts = append(fv.Texts, v)
	fv.Values = append(fv.Values, struct {
		IPNet *net.IPNet
		IP    net.IP
	}{IP: ip, IPNet: ipNet})
	return nil
}

func (fv *CIDRs) String() string {
	return strings.Join(fv.Texts, ",")
}
