package minfraud

// Score executes Score API.
func (s *MinFraud) Score(req BaseRequest) (*ScoreResponse, error) {
	resp := ScoreResponse{}
	err := s.client.CallPOST("/v2.0/score", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ScoreByIP executes Score API with only ip address parameter.
func (s *MinFraud) ScoreByIP(ipaddr string) (*ScoreResponse, error) {
	return s.Score(BaseRequest{
		Device: &DeviceData{
			IPAddress: ipaddr,
		},
	})
}

// ScoreResponse has response from Score API.
type ScoreResponse struct {
	BaseResponse
	IPAddress IPAddress `json:"ip_address"`
}

func (r ScoreResponse) APIResponse() APIResponse {
	return APIResponse{
		BaseResponse: r.BaseResponse,
		IPAddress:    r.IPAddress,
	}
}
