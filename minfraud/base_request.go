package minfraud

// BaseRequest is common request of minFraud API.
// ref: https://dev.maxmind.com/minfraud/api-documentation/requests
type BaseRequest struct {
	Account      *AccountData       `json:"account,omitempty"`
	Billing      *BillingData       `json:"billing,omitempty"`
	CreditCard   *CreditCardData    `json:"credit_card,omitempty"`
	Device       *DeviceData        `json:"device,omitempty"`
	Email        *EmailData         `json:"email,omitempty"`
	Event        *EventData         `json:"event,omitempty"`
	Order        *OrderData         `json:"order,omitempty"`
	Payment      *PaymentData       `json:"payment,omitempty"`
	Shipping     *ShippingData      `json:"shipping,omitempty"`
	ShoppingCart []ShoppingCartData `json:"shopping_cart,omitempty"`

	// ref: https://support.maxmind.com/custom-inputs-guide/
	CustomInputs interface{} `json:"custom_inputs,omitempty"`
}

type AccountData struct {
	UserID      string `json:"user_id"`
	UserNameMD5 string `json:"username_md5"`
}

type BillingData struct {
	Address          string `json:"address"`
	Address2         string `json:"address_2"`
	City             string `json:"city"`
	Company          string `json:"company"`
	Country          string `json:"country"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneCountryCode string `json:"phone_country_code"`
	PhoneNumber      string `json:"phone_number"`
	Postal           string `json:"postal"`
	Region           string `json:"region"`
}

type CreditCardData struct {
	AVSResult             string `json:"avs_result"`
	BankName              string `json:"bank_name"`
	BankPhoneCountryCode  string `json:"bank_phone_country_code"`
	BankPhoneNumber       string `json:"bank_phone_number"`
	CVVResult             string `json:"cvv_result"`
	IssuerIDNumber        string `json:"issuer_id_number"`
	Last4Digits           string `json:"last_4_digits"`
	Token                 string `json:"token"`
	Was3DSecureSuccessful bool   `json:"was_3d_secure_successful"`
}

type DeviceData struct {
	AcceptLanguage string  `json:"accept_language"`
	IPAddress      string  `json:"ip_address"`
	SessionAge     float64 `json:"session_age"`
	SessionID      string  `json:"session_id"`
	UserAgent      string  `json:"user_agent"`
}

type EmailData struct {
	Address string `json:"address,omitempty"` // plain text or MD5 hased email address
	Domain  string `json:"domain,omitempty"`
}

type EventData struct {
	ShopID        string `json:"shop_id"`
	Time          string `json:"time"`
	TransactionID string `json:"transaction_id"`
	Type          string `json:"type"`
}

type OrderData struct {
	AffiliateID    string  `json:"affiliateID"`
	Amount         float64 `json:"amount"`
	Currency       string  `json:"currency"`
	DiscountCode   string  `json:"discount_code"`
	HasGiftMessage bool    `json:"has_gift_message"`
	IsGift         bool    `json:"is_gift"`
	ReferrerURI    string  `json:"referrer_uri"`
	SubaffiliateID string  `json:"subaffiliate_id"`
}

type PaymentData struct {
	DeclineCode   string `json:"decline_code"`
	Processor     string `json:"processor"`
	WasAuthorized bool   `json:"was_authorized"`
}

type ShippingData struct {
	Address          string `json:"address"`
	Address2         string `json:"address_2"`
	City             string `json:"city"`
	Company          string `json:"company"`
	Country          string `json:"country"`
	DeliverySpeed    string `json:"delivery_speed"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	PhoneCountryCode string `json:"phone_country_code"`
	PhoneNumber      string `json:"phone_number"`
	Postal           string `json:"postal"`
	Region           string `json:"region"`
}

type ShoppingCartData struct {
	Category string  `json:"category"`
	ItemID   string  `json:"item_id"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}
