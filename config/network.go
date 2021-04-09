package config

import "net"

// HealthCheck is the health check route to use.
const HealthCheck = "/health"

// Network config
type Network struct {
	// Host is default adapter to listen on.
	// default: localhost
	Host string `json:"host"`
	// HealthAddress is the address for the health check server.
	// default: [Host]:7720
	HealthAddress string `json:"health_address" yaml:"health_address" mapstructure:"health_address"`
	// RPCAddress is the address for the gRPC server.
	// default: [Host]:7721
	RPCAddress string `json:"rpc_address" yaml:"rpc_address" mapstructure:"rpc_address"`
	// AdminAddress is the address for the admin server.
	// default: [Host]:7722
	AdminAddress string `json:"admin_address" yaml:"admin_address" mapstructure:"admin_address"`
	// HTTP is the address for the HTTP server.
	// default: [Host]:7727
	RESTAddress string `json:"rest_address" yaml:"rest_address" mapstructure:"rest_address"`
	// RPCWebAddress is the address for the gRPC-Web server.
	// default: [Host]:7729
	RPCWebAddress string `json:"rpcweb_address" yaml:"rpcweb_address" mapstructure:"rpcweb_address"`

	// TODO: use RequestID
	// RequestID is the request id to use
	RequestID string `json:"request_id" yaml:"request_id" mapstructure:"request_id"`
}

func (n *Network) normalize(Service) (err error) {
	dc := networkDefaults
	if n.Host == dc.Host {
		return
	}
	updateHost := func(addr, host string) (string, error) {
		var p string
		_, p, err = net.SplitHostPort(addr)
		if err != nil {
			return "", err
		}
		return net.JoinHostPort(host, p), nil
	}
	if n.RPCAddress == dc.RPCAddress {
		n.RPCAddress, err = updateHost(n.RPCAddress, n.Host)
		if err != nil {
			return
		}
	}
	if n.AdminAddress == dc.AdminAddress {
		n.AdminAddress, err = updateHost(n.AdminAddress, n.Host)
		if err != nil {
			return
		}
	}
	if n.RESTAddress == dc.RESTAddress {
		n.RESTAddress, err = updateHost(n.RESTAddress, n.Host)
		if err != nil {
			return
		}
	}
	if n.RPCWebAddress == dc.RPCWebAddress {
		n.RPCWebAddress, err = updateHost(n.RPCWebAddress, n.Host)
		if err != nil {
			return
		}
	}
	if n.HealthAddress == dc.HealthAddress {
		n.HealthAddress, err = updateHost(n.HealthAddress, n.Host)
		if err != nil {
			return
		}
	}
	return
}
