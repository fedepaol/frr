// SPDX-License-Identifier:Apache-2.0

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"net/http"

	v1beta1 "github.com/metallb/frr-k8s/core/v1beta1"
	"github.com/metallb/frr-k8s/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CoreV1beta1Interface interface {
	RESTClient() rest.Interface
	FRRConfigurationsGetter
}

// CoreV1beta1Client is used to interact with features provided by the  group.
type CoreV1beta1Client struct {
	restClient rest.Interface
}

func (c *CoreV1beta1Client) FRRConfigurations(namespace string) FRRConfigurationInterface {
	return newFRRConfigurations(c, namespace)
}

// NewForConfig creates a new CoreV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CoreV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new CoreV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CoreV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &CoreV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new CoreV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1beta1Client for the given RESTClient.
func New(c rest.Interface) *CoreV1beta1Client {
	return &CoreV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/api"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
