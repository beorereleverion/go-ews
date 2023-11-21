package goews

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/Azure/go-ntlmssp"
	"github.com/MihaylovNikitos/go-ews/elements"
	"github.com/MihaylovNikitos/go-ews/operations"
)

type Envelope interface {
	GetEnvelopeBytes() ([]byte, error)
}

type Client interface {
	SendAndReceive(e Envelope) ([]byte, error)
	GetServerAddr() string
	GetUsername() string
	DoRequest(e Envelope, oe operations.Element) error
	FindPeople(eItem *elements.FindPeople) (*elements.FindPeopleResponse, error)
	GetFolder(eItem *elements.GetFolder) (*elements.GetFolderResponse, error)
	FindItem(eItem *elements.FindItem) (*elements.FindItemResponse, error)
	CreateItem(eItem *elements.CreateItem) (*elements.CreateItemResponse, error)
	DeleteItem(eItem *elements.DeleteItem) (*elements.DeleteItemResponse, error)
}

type client struct {
	serverAddress string
	username      string
	password      string
	config        *Config
}

type Config struct {
	Dump    bool
	NTLM    bool
	SkipTLS bool
}

func (c *client) GetServerAddr() string {
	return c.serverAddress
}

func (c *client) GetUsername() string {
	return c.username
}

func NewClient(serverAddress, username, password string, config Config) Client {
	return &client{
		serverAddress: serverAddress,
		username:      username,
		password:      password,
		config:        &config,
	}
}

func (c *client) DoRequest(e Envelope, oe operations.Element) error {
	bArr, err := c.SendAndReceive(e)
	if err != nil {
		return err
	}
	return operations.Unmarshal(bArr, oe)
}

func (c *client) SendAndReceive(e Envelope) ([]byte, error) {
	bb, err := e.GetEnvelopeBytes()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.serverAddress, bytes.NewReader(bb))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	c.logRequest(req)

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "text/xml")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{},
	}
	applyConfig(c.config, client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	c.logResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, NewError(resp)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, err
}

func applyConfig(config *Config, client *http.Client) {
	if config.NTLM {
		client.Transport = ntlmssp.Negotiator{
			RoundTripper: &http.Transport{
				TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{},
			},
		}
	}

	if config.SkipTLS {
		client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
}

func (c *client) logRequest(req *http.Request) {
	if c.config != nil && c.config.Dump {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Request:\n%v\n----\n", string(dump))
	}
}

func (c *client) logResponse(resp *http.Response) {
	if c.config != nil && c.config.Dump {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Response:\n%v\n----\n", string(dump))
	}
}
