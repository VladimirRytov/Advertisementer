package httpsender

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/VladimirRytov/advertisementer/internal/datatransferobjects"
	"github.com/VladimirRytov/advertisementer/internal/encodedecoder"
)

type Sender struct {
	client   *http.Client
	addr     string
	apiPath  string
	connInfo datatransferobjects.ServerDSN
	token    string
}

func NewSender() *Sender {
	s := &Sender{
		client: &http.Client{},
	}
	return s
}

func (s *Sender) Initialize(ctx context.Context, addr datatransferobjects.ServerDSN) error {
	s.connInfo = addr
	s.addr = "http://" + addr.Source + ":" + strconv.Itoa(int(addr.Port))
	s.apiPath = "api/v1"
	err := s.authenticate(ctx, addr.UserName, addr.Password)
	if err != nil {
		s.addr = "https://" + addr.Source + ":" + strconv.Itoa(int(addr.Port))
		err := s.authenticate(ctx, addr.UserName, addr.Password)
		if err != nil {
			return err
		}
	}
	return err
}

func (s *Sender) SetAuthToken(token string) {
	s.token = token
}

func (s *Sender) SetAdress(addr string) {
	s.addr = addr
}

func (s *Sender) SetApiPath(path string) {
	s.apiPath = path
}

func (s *Sender) ConnectionInfo() map[string]string {
	m := make(map[string]string)
	m["adress"] = s.addr
	m["apiPath"] = s.apiPath
	m["token"] = s.token
	return m
}

func (s *Sender) authenticate(ctx context.Context, user, password string) error {
	path, err := url.JoinPath(s.addr, s.apiPath, "auth")
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(user, password)
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	switch {
	case resp.StatusCode == 200:
		var tok datatransferobjects.Token
		err = encodedecoder.FromJSON(&tok, resp.Body)
		if err != nil {
			return err
		}
		s.token = string(tok.AccessToken)
		return nil
	default:
		var respMsg datatransferobjects.ResponceMessage
		err = encodedecoder.FromJSON(&respMsg, resp.Body)
		if err != nil {
			return err
		}
		return errors.New(respMsg.Message)
	}
}

func (s *Sender) GetRequest(ctx context.Context, params *datatransferobjects.RequestDTO, body io.Reader) ([]byte, error) {
	var b bytes.Buffer
	path, err := url.JoinPath(s.addr, s.apiPath, params.Kind, params.Name)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, body)
	if err != nil {
		return nil, err
	}
	queryVals := req.URL.Query()
	for k := range params.Queries {
		queryVals.Add(k, params.Queries[k])
	}
	req.URL.RawQuery = queryVals.Encode()
	req.Header.Add("Authorization", "Bearer "+s.token)
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(&b, resp.Body)
	switch {
	case resp.StatusCode == 200:
		return b.Bytes(), err
	default:
		var rawErr error = errors.New(b.String())
		var respMsg datatransferobjects.ResponceMessage
		err = encodedecoder.FromJSON(&respMsg, resp.Body)
		if err != nil {
			return nil, rawErr
		}
		return b.Bytes(), errors.New(respMsg.Message)
	}
}

func (s *Sender) CreateRequest(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader) ([]byte, error) {
	path, err := url.JoinPath(s.addr, s.apiPath, params.Kind, params.Name)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+s.token)
	queryVals := req.URL.Query()
	for k := range params.Queries {
		queryVals.Add(k, params.Queries[k])
	}
	req.URL.RawQuery = queryVals.Encode()
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	_, err = io.Copy(&b, resp.Body)

	switch {
	case resp.StatusCode == 201:
		return b.Bytes(), err
	default:
		var rawErr error = errors.New(b.String())
		var respMsg datatransferobjects.ResponceMessage
		err = encodedecoder.FromJSON(&respMsg, &b)
		if err != nil {
			return nil, rawErr
		}
		return b.Bytes(), errors.New(respMsg.Message)
	}
}

func (s *Sender) UploadMultipart(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader, contentType string) ([]byte, error) {
	path, err := url.JoinPath(s.addr, s.apiPath, params.Kind, params.Name)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+s.token)
	queryVals := req.URL.Query()
	for k := range params.Queries {
		queryVals.Add(k, params.Queries[k])
	}
	req.URL.RawQuery = queryVals.Encode()
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	_, err = io.Copy(&b, resp.Body)

	switch {
	case resp.StatusCode == 201:
		return b.Bytes(), err
	default:
		var rawErr error = errors.New(b.String())
		var respMsg datatransferobjects.ResponceMessage
		err = encodedecoder.FromJSON(&respMsg, &b)
		if err != nil {
			return nil, rawErr
		}
		return b.Bytes(), errors.New(respMsg.Message)
	}
}

func (s *Sender) UpdateRequest(ctx context.Context, params *datatransferobjects.RequestDTO, data io.Reader) ([]byte, error) {
	path, err := url.JoinPath(s.addr, s.apiPath, params.Kind, params.Name)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, path, data)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+s.token)
	queryVals := req.URL.Query()

	for k := range params.Queries {
		queryVals.Add(k, params.Queries[k])
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	_, err = io.Copy(&b, resp.Body)
	switch {
	case resp.StatusCode == 200 || resp.StatusCode == 201:
		return b.Bytes(), err
	default:
		var rawErr error = errors.New(b.String())
		var respMsg datatransferobjects.ResponceMessage
		err = encodedecoder.FromJSON(&respMsg, resp.Body)
		if err != nil {
			return nil, rawErr
		}
		return b.Bytes(), errors.New(respMsg.Message)
	}
}

func (s *Sender) DeleteRequest(ctx context.Context, params *datatransferobjects.RequestDTO) ([]byte, error) {
	path, err := url.JoinPath(s.addr, s.apiPath, params.Kind, params.Name)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+s.token)
	queryVals := req.URL.Query()

	for k := range params.Queries {
		queryVals.Add(k, params.Queries[k])
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var b bytes.Buffer
	_, err = io.Copy(&b, resp.Body)

	switch {
	case resp.StatusCode == 204:
		return b.Bytes(), err
	default:
		var rawErr error = errors.New(b.String())
		var respMsg datatransferobjects.ResponceMessage
		err = encodedecoder.FromJSON(&respMsg, resp.Body)
		if err != nil {
			return nil, rawErr
		}
		return b.Bytes(), errors.New(respMsg.Message)
	}
}
