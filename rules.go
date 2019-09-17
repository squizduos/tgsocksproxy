package main

import (
	"fmt"
	"strings"

	"net"

	"github.com/armon/go-socks5"
	"golang.org/x/net/context"
)

type Rules struct {
	Adresses []string
	Networks []string
	Domains  []string

	ipAddresses []net.IP
	ipNetworks  []*net.IPNet
	ipDomains   []string
}

func (p *Rules) Load() error {
	p.ipAddresses = nil
	for _, address := range p.Adresses {
		ipAddr := net.ParseIP(address)
		if ipAddr == nil {
			return fmt.Errorf("Wrong IP: %s", address)
		}
		p.ipAddresses = append(p.ipAddresses, ipAddr)
	}

	p.ipNetworks = nil
	for _, network := range p.Networks {
		_, ipNet, err := net.ParseCIDR(network)
		if err != nil {
			return fmt.Errorf("Wrong network: %s", network)
		}
		p.ipNetworks = append(p.ipNetworks, ipNet)
	}

	p.ipDomains = nil
	for _, domain := range p.Domains {
		p.ipDomains = append(p.ipDomains, domain)
	}

	return nil
}

func (p *Rules) Allow(ctx context.Context, req *socks5.Request) (context.Context, bool) {
	for _, ip := range p.ipAddresses {
		if ip.Equal(req.DestAddr.IP) {
			return ctx, true
		}
	}

	for _, ipNet := range p.ipNetworks {
		if ipNet.Contains(req.DestAddr.IP) {
			return ctx, true
		}
	}

	for _, domain := range p.ipDomains {
		if strings.HasSuffix(req.DestAddr.FQDN, domain) {
			return ctx, true
		}
	}

	return ctx, false
}
