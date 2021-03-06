package client

import (
	"github.com/rancher/norman/types"
)

const (
	GatewayType           = "gateway"
	GatewayFieldCreated   = "created"
	GatewayFieldLabels    = "labels"
	GatewayFieldName      = "name"
	GatewayFieldNamespace = "namespace"
	GatewayFieldRemoved   = "removed"
	GatewayFieldSelector  = "selector"
	GatewayFieldServers   = "servers"
	GatewayFieldUUID      = "uuid"
)

type Gateway struct {
	types.Resource
	Created   string            `json:"created,omitempty" yaml:"created,omitempty"`
	Labels    map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name      string            `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Removed   string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Selector  map[string]string `json:"selector,omitempty" yaml:"selector,omitempty"`
	Servers   []Server          `json:"servers,omitempty" yaml:"servers,omitempty"`
	UUID      string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type GatewayCollection struct {
	types.Collection
	Data   []Gateway `json:"data,omitempty"`
	client *GatewayClient
}

type GatewayClient struct {
	apiClient *Client
}

type GatewayOperations interface {
	List(opts *types.ListOpts) (*GatewayCollection, error)
	Create(opts *Gateway) (*Gateway, error)
	Update(existing *Gateway, updates interface{}) (*Gateway, error)
	Replace(existing *Gateway) (*Gateway, error)
	ByID(id string) (*Gateway, error)
	Delete(container *Gateway) error
}

func newGatewayClient(apiClient *Client) *GatewayClient {
	return &GatewayClient{
		apiClient: apiClient,
	}
}

func (c *GatewayClient) Create(container *Gateway) (*Gateway, error) {
	resp := &Gateway{}
	err := c.apiClient.Ops.DoCreate(GatewayType, container, resp)
	return resp, err
}

func (c *GatewayClient) Update(existing *Gateway, updates interface{}) (*Gateway, error) {
	resp := &Gateway{}
	err := c.apiClient.Ops.DoUpdate(GatewayType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *GatewayClient) Replace(obj *Gateway) (*Gateway, error) {
	resp := &Gateway{}
	err := c.apiClient.Ops.DoReplace(GatewayType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *GatewayClient) List(opts *types.ListOpts) (*GatewayCollection, error) {
	resp := &GatewayCollection{}
	err := c.apiClient.Ops.DoList(GatewayType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *GatewayCollection) Next() (*GatewayCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &GatewayCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *GatewayClient) ByID(id string) (*Gateway, error) {
	resp := &Gateway{}
	err := c.apiClient.Ops.DoByID(GatewayType, id, resp)
	return resp, err
}

func (c *GatewayClient) Delete(container *Gateway) error {
	return c.apiClient.Ops.DoResourceDelete(GatewayType, &container.Resource)
}
