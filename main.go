package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/network"
)

const (
	pluginUDSocket = "plugin"
)

type TestDriver struct {
}

func (t *TestDriver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return &network.CapabilitiesResponse{Scope: "global"}, nil
}

func (t *TestDriver) AllocateNetwork(r *network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	logrus.Infof("Allocate Network %v", r.NetworkID)
	return nil, nil
}

func (t *TestDriver) FreeNetwork(r *network.FreeNetworkRequest) error {
	logrus.Infof("Free Network %v", r.NetworkID)
	return nil
}

func (t *TestDriver) CreateNetwork(r *network.CreateNetworkRequest) error {
	logrus.Infof("Create Network %v", r.NetworkID)
	return nil
}

func (t *TestDriver) DeleteNetwork(r *network.DeleteNetworkRequest) error {
	logrus.Infof("Delete Network %v", r.NetworkID)
	return nil
}

func (t *TestDriver) CreateEndpoint(r *network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	logrus.Infof("Create Endpoint %v", r.NetworkID)
	return &network.CreateEndpointResponse{}, nil
}

func (t *TestDriver) DeleteEndpoint(r *network.DeleteEndpointRequest) error {
	return nil
}

func (t *TestDriver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	return &network.InfoResponse{}, nil
}

func (t *TestDriver) DiscoverNew(*network.DiscoveryNotification) error {
	return nil
}

func (t *TestDriver) DiscoverDelete(*network.DiscoveryNotification) error {
	return nil
}

func (t *TestDriver) Join(r *network.JoinRequest) (*network.JoinResponse, error) {
	return &network.JoinResponse{}, nil
}

func (t *TestDriver) Leave(r *network.LeaveRequest) error {
	return nil
}

func (t *TestDriver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	return nil
}

func (t *TestDriver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	return nil
}

func main() {
	d := &TestDriver{}
	handler := network.NewHandler(d)
	// logrus.Error(handler.ServeTCP("test", ":32234"))
	logrus.Error(handler.ServeUnix("", pluginUDSocket))
}
