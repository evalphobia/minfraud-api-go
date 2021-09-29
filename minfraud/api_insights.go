package minfraud

// Insights executes Insights API.
func (s *MinFraud) Insights(req BaseRequest) (*InsightsResponse, error) {
	resp := InsightsResponse{}
	err := s.client.CallPOST("/v2.0/insights", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InsightsByIP executes Insights API with only ip address parameter.
func (s *MinFraud) InsightsByIP(ipaddr string) (*InsightsResponse, error) {
	return s.Insights(BaseRequest{
		Device: &DeviceData{
			IPAddress: ipaddr,
		},
	})
}

// InsightsResponse has response from Insights API.
type InsightsResponse struct {
	BaseResponse
	IPAddress       IPAddress       `json:"ip_address"`
	BillingAddress  BillingAddress  `json:"billing_address"`
	CreditCard      CreditCard      `json:"credit_card"`
	Device          Device          `json:"device"`
	Email           Email           `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

func (r InsightsResponse) APIResponse() APIResponse {
	return APIResponse{
		BaseResponse:    r.BaseResponse,
		IPAddress:       r.IPAddress,
		BillingAddress:  r.BillingAddress,
		CreditCard:      r.CreditCard,
		Device:          r.Device,
		Email:           r.Email,
		ShippingAddress: r.ShippingAddress,
	}
}
