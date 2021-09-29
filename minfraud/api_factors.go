package minfraud

// Factors executes Factors API.
func (s *MinFraud) Factors(req BaseRequest) (*FactorsResponse, error) {
	resp := FactorsResponse{}
	err := s.client.CallPOST("/v2.0/factors", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// FactorsResponse has response from Factors API.
type FactorsResponse struct {
	BaseResponse
	IPAddress       IPAddress       `json:"ip_address"`
	BillingAddress  BillingAddress  `json:"billing_address"`
	CreditCard      CreditCard      `json:"credit_card"`
	Device          Device          `json:"device"`
	Email           Email           `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
	Subscores       Subscores       `json:"subscores"`
}

func (r FactorsResponse) APIResponse() APIResponse {
	return APIResponse(r)
}
