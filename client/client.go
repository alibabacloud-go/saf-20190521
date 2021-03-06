// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ExecuteExtendServiceRequest struct {
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ExecuteExtendServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExtendServiceRequest) GoString() string {
	return s.String()
}

func (s *ExecuteExtendServiceRequest) SetRegion(v string) *ExecuteExtendServiceRequest {
	s.Region = &v
	return s
}

func (s *ExecuteExtendServiceRequest) SetService(v string) *ExecuteExtendServiceRequest {
	s.Service = &v
	return s
}

func (s *ExecuteExtendServiceRequest) SetServiceParameters(v string) *ExecuteExtendServiceRequest {
	s.ServiceParameters = &v
	return s
}

type ExecuteExtendServiceResponseBody struct {
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	HttpStatusCode *string                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Data           *ExecuteExtendServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Success        *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ExecuteExtendServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExtendServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteExtendServiceResponseBody) SetMessage(v string) *ExecuteExtendServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteExtendServiceResponseBody) SetRequestId(v string) *ExecuteExtendServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteExtendServiceResponseBody) SetHttpStatusCode(v string) *ExecuteExtendServiceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ExecuteExtendServiceResponseBody) SetData(v *ExecuteExtendServiceResponseBodyData) *ExecuteExtendServiceResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteExtendServiceResponseBody) SetCode(v string) *ExecuteExtendServiceResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteExtendServiceResponseBody) SetSuccess(v bool) *ExecuteExtendServiceResponseBody {
	s.Success = &v
	return s
}

type ExecuteExtendServiceResponseBodyData struct {
	InvokeResult *string `json:"InvokeResult,omitempty" xml:"InvokeResult,omitempty"`
}

func (s ExecuteExtendServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExtendServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ExecuteExtendServiceResponseBodyData) SetInvokeResult(v string) *ExecuteExtendServiceResponseBodyData {
	s.InvokeResult = &v
	return s
}

type ExecuteExtendServiceResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteExtendServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteExtendServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteExtendServiceResponse) GoString() string {
	return s.String()
}

func (s *ExecuteExtendServiceResponse) SetHeaders(v map[string]*string) *ExecuteExtendServiceResponse {
	s.Headers = v
	return s
}

func (s *ExecuteExtendServiceResponse) SetBody(v *ExecuteExtendServiceResponseBody) *ExecuteExtendServiceResponse {
	s.Body = v
	return s
}

type ExecuteRequestRequest struct {
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ExecuteRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestRequest) GoString() string {
	return s.String()
}

func (s *ExecuteRequestRequest) SetServiceParameters(v string) *ExecuteRequestRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ExecuteRequestRequest) SetService(v string) *ExecuteRequestRequest {
	s.Service = &v
	return s
}

type ExecuteRequestResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ExecuteRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteRequestResponseBody) SetMessage(v string) *ExecuteRequestResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteRequestResponseBody) SetRequestId(v string) *ExecuteRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteRequestResponseBody) SetData(v map[string]interface{}) *ExecuteRequestResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteRequestResponseBody) SetCode(v int32) *ExecuteRequestResponseBody {
	s.Code = &v
	return s
}

type ExecuteRequestResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestResponse) GoString() string {
	return s.String()
}

func (s *ExecuteRequestResponse) SetHeaders(v map[string]*string) *ExecuteRequestResponse {
	s.Headers = v
	return s
}

func (s *ExecuteRequestResponse) SetBody(v *ExecuteRequestResponseBody) *ExecuteRequestResponse {
	s.Body = v
	return s
}

type ExecuteRequestMLRequest struct {
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ExecuteRequestMLRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestMLRequest) GoString() string {
	return s.String()
}

func (s *ExecuteRequestMLRequest) SetServiceParameters(v string) *ExecuteRequestMLRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ExecuteRequestMLRequest) SetService(v string) *ExecuteRequestMLRequest {
	s.Service = &v
	return s
}

type ExecuteRequestMLResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ExecuteRequestMLResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestMLResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteRequestMLResponseBody) SetMessage(v string) *ExecuteRequestMLResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteRequestMLResponseBody) SetRequestId(v string) *ExecuteRequestMLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteRequestMLResponseBody) SetData(v map[string]interface{}) *ExecuteRequestMLResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteRequestMLResponseBody) SetCode(v int32) *ExecuteRequestMLResponseBody {
	s.Code = &v
	return s
}

type ExecuteRequestMLResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteRequestMLResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteRequestMLResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestMLResponse) GoString() string {
	return s.String()
}

func (s *ExecuteRequestMLResponse) SetHeaders(v map[string]*string) *ExecuteRequestMLResponse {
	s.Headers = v
	return s
}

func (s *ExecuteRequestMLResponse) SetBody(v *ExecuteRequestMLResponseBody) *ExecuteRequestMLResponse {
	s.Body = v
	return s
}

type ExecuteRequestSGRequest struct {
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ExecuteRequestSGRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestSGRequest) GoString() string {
	return s.String()
}

func (s *ExecuteRequestSGRequest) SetServiceParameters(v string) *ExecuteRequestSGRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ExecuteRequestSGRequest) SetService(v string) *ExecuteRequestSGRequest {
	s.Service = &v
	return s
}

type ExecuteRequestSGResponseBody struct {
	Message   *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ExecuteRequestSGResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestSGResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteRequestSGResponseBody) SetMessage(v string) *ExecuteRequestSGResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteRequestSGResponseBody) SetRequestId(v string) *ExecuteRequestSGResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteRequestSGResponseBody) SetData(v map[string]interface{}) *ExecuteRequestSGResponseBody {
	s.Data = v
	return s
}

func (s *ExecuteRequestSGResponseBody) SetCode(v int32) *ExecuteRequestSGResponseBody {
	s.Code = &v
	return s
}

type ExecuteRequestSGResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteRequestSGResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteRequestSGResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestSGResponse) GoString() string {
	return s.String()
}

func (s *ExecuteRequestSGResponse) SetHeaders(v map[string]*string) *ExecuteRequestSGResponse {
	s.Headers = v
	return s
}

func (s *ExecuteRequestSGResponse) SetBody(v *ExecuteRequestSGResponseBody) *ExecuteRequestSGResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": tea.String("saf.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("saf"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteExtendServiceWithOptions(request *ExecuteExtendServiceRequest, runtime *util.RuntimeOptions) (_result *ExecuteExtendServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteExtendServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteExtendService"), tea.String("2019-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteExtendService(request *ExecuteExtendServiceRequest) (_result *ExecuteExtendServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteExtendServiceResponse{}
	_body, _err := client.ExecuteExtendServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteRequestWithOptions(request *ExecuteRequestRequest, runtime *util.RuntimeOptions) (_result *ExecuteRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteRequestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteRequest"), tea.String("2019-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteRequest(request *ExecuteRequestRequest) (_result *ExecuteRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteRequestResponse{}
	_body, _err := client.ExecuteRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteRequestMLWithOptions(request *ExecuteRequestMLRequest, runtime *util.RuntimeOptions) (_result *ExecuteRequestMLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteRequestMLResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteRequestML"), tea.String("2019-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteRequestML(request *ExecuteRequestMLRequest) (_result *ExecuteRequestMLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteRequestMLResponse{}
	_body, _err := client.ExecuteRequestMLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteRequestSGWithOptions(request *ExecuteRequestSGRequest, runtime *util.RuntimeOptions) (_result *ExecuteRequestSGResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteRequestSGResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteRequestSG"), tea.String("2019-05-21"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteRequestSG(request *ExecuteRequestSGRequest) (_result *ExecuteRequestSGResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteRequestSGResponse{}
	_body, _err := client.ExecuteRequestSGWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
