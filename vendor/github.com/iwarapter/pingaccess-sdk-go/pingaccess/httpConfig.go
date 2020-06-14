package pingaccess

import (
	"fmt"
	"net/http"
	"net/url"
)

type HttpConfigService service

type HttpConfigAPI interface {
	DeleteHttpMonitoringCommand() (resp *http.Response, err error)
	GetHttpMonitoringCommand() (result *HttpMonitoringView, resp *http.Response, err error)
	UpdateHttpMonitoringCommand(input *UpdateHttpMonitoringCommandInput) (result *HttpMonitoringView, resp *http.Response, err error)
	DeleteHostSourceCommand() (resp *http.Response, err error)
	GetHostSourceCommand() (result *HostMultiValueSourceView, resp *http.Response, err error)
	UpdateHostSourceCommand(input *UpdateHostSourceCommandInput) (result *HostMultiValueSourceView, resp *http.Response, err error)
	DeleteIpSourceCommand() (resp *http.Response, err error)
	GetIpSourceCommand() (result *IpMultiValueSourceView, resp *http.Response, err error)
	UpdateIpSourceCommand(input *UpdateIpSourceCommandInput) (result *IpMultiValueSourceView, resp *http.Response, err error)
	DeleteProtoSourceCommand() (resp *http.Response, err error)
	GetProtoSourceCommand() (result *ProtocolSourceView, resp *http.Response, err error)
	UpdateProtocolSourceCommand(input *UpdateProtocolSourceCommandInput) (result *ProtocolSourceView, resp *http.Response, err error)
}

//DeleteHttpMonitoringCommand - Resets the HTTP monitoring auditLevel to default value
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteHttpMonitoringCommand() (resp *http.Response, err error) {
	path := "/httpConfig/monitoring"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetHttpMonitoringCommand - Get the HTTP monitoring auditLevel
//RequestType: GET
//Input:
func (s *HttpConfigService) GetHttpMonitoringCommand() (result *HttpMonitoringView, resp *http.Response, err error) {
	path := "/httpConfig/monitoring"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateHttpMonitoringCommand - Update the HTTP monitoring auditLevel
//RequestType: PUT
//Input: input *UpdateHttpMonitoringCommandInput
func (s *HttpConfigService) UpdateHttpMonitoringCommand(input *UpdateHttpMonitoringCommandInput) (result *HttpMonitoringView, resp *http.Response, err error) {
	path := "/httpConfig/monitoring"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateHttpMonitoringCommandInput struct {
	Body HttpMonitoringView
}

//DeleteHostSourceCommand - Resets the HTTP request Host Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteHostSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetHostSourceCommand - Get the HTTP request Host Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetHostSourceCommand() (result *HostMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateHostSourceCommand - Update the HTTP request Host Source type
//RequestType: PUT
//Input: input *UpdateHostSourceCommandInput
func (s *HttpConfigService) UpdateHostSourceCommand(input *UpdateHostSourceCommandInput) (result *HostMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/hostSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateHostSourceCommandInput struct {
	Body HostMultiValueSourceView
}

//DeleteIpSourceCommand - Resets the HTTP request IP Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteIpSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetIpSourceCommand - Get the HTTP request IP Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetIpSourceCommand() (result *IpMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateIpSourceCommand - Update the HTTP request IP Source type
//RequestType: PUT
//Input: input *UpdateIpSourceCommandInput
func (s *HttpConfigService) UpdateIpSourceCommand(input *UpdateIpSourceCommandInput) (result *IpMultiValueSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/ipSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateIpSourceCommandInput struct {
	Body IpMultiValueSourceView
}

//DeleteProtoSourceCommand - Resets the HTTP request Protocol Source type to default values
//RequestType: DELETE
//Input:
func (s *HttpConfigService) DeleteProtoSourceCommand() (resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("DELETE", rel, nil)
	if err != nil {
		return nil, err
	}

	resp, err = s.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

//GetProtoSourceCommand - Get the HTTP request Protocol Source type
//RequestType: GET
//Input:
func (s *HttpConfigService) GetProtoSourceCommand() (result *ProtocolSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

//UpdateProtocolSourceCommand - Update the HTTP request Protocol Source type
//RequestType: PUT
//Input: input *UpdateProtocolSourceCommandInput
func (s *HttpConfigService) UpdateProtocolSourceCommand(input *UpdateProtocolSourceCommandInput) (result *ProtocolSourceView, resp *http.Response, err error) {
	path := "/httpConfig/request/protocolSource"
	rel := &url.URL{Path: fmt.Sprintf("%s%s", s.client.Context, path)}
	req, err := s.client.newRequest("PUT", rel, input.Body)
	if err != nil {
		return nil, nil, err
	}

	resp, err = s.client.do(req, &result)
	if err != nil {
		return result, resp, err
	}
	return result, resp, nil

}

type UpdateProtocolSourceCommandInput struct {
	Body ProtocolSourceView
}
