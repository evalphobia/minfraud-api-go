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
	UserID      string `json:"user_id,omitempty"`
	UserNameMD5 string `json:"username_md5,omitempty"`
}

type BillingData struct {
	Address          string `json:"address,omitempty"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city,omitempty"`
	Company          string `json:"company,omitempty"`
	Country          string `json:"country,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	Postal           string `json:"postal,omitempty"`
	Region           string `json:"region,omitempty"`
}

type CreditCardData struct {
	AVSResult             string `json:"avs_result,omitempty"`
	BankName              string `json:"bank_name,omitempty"`
	BankPhoneCountryCode  string `json:"bank_phone_country_code,omitempty"`
	BankPhoneNumber       string `json:"bank_phone_number,omitempty"`
	CVVResult             string `json:"cvv_result,omitempty"`
	IssuerIDNumber        string `json:"issuer_id_number,omitempty"`
	Last4Digits           string `json:"last_4_digits,omitempty"`
	Token                 string `json:"token,omitempty"`
	Was3DSecureSuccessful *bool   `json:"was_3d_secure_successful,omitempty"`
}

type DeviceData struct {
	AcceptLanguage string  `json:"accept_language,omitempty"`
	IPAddress      string  `json:"ip_address,omitempty"`
	SessionAge     float64 `json:"session_age,omitempty"`
	SessionID      string  `json:"session_id,omitempty"`
	UserAgent      string  `json:"user_agent,omitempty"`
}

type EmailData struct {
	Address string `json:"address,omitempty"` // plain text or MD5 hased email address
	Domain  string `json:"domain,omitempty"`
}

type EventData struct {
	ShopID        string `json:"shop_id,omitempty"`
	Time          string `json:"time,omitempty"`
	TransactionID string `json:"transaction_id,omitempty"`
	Type          string `json:"type,omitempty"`
}

type OrderData struct {
	AffiliateID    string  `json:"affiliateID,omitempty"`
	Amount         *float64 `json:"amount,omitempty"`
	Currency       string  `json:"currency,omitempty"`
	DiscountCode   string  `json:"discount_code,omitempty"`
	HasGiftMessage *bool    `json:"has_gift_message,omitempty"`
	IsGift         *bool    `json:"is_gift,omitempty"`
	ReferrerURI    string  `json:"referrer_uri,omitempty"`
	SubaffiliateID string  `json:"subaffiliate_id,omitempty"`
}

type PaymentData struct {
	DeclineCode   string `json:"decline_code,omitempty"`
	Processor     string `json:"processor,omitempty"`
	WasAuthorized *bool   `json:"was_authorized,omitempty"`
}

type ShippingData struct {
	Address          string `json:"address,omitempty"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city,omitempty"`
	Company          string `json:"company,omitempty"`
	Country          string `json:"country,omitempty"`
	DeliverySpeed    string `json:"delivery_speed,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	Postal           string `json:"postal,omitempty"`
	Region           string `json:"region,omitempty"`
}

type ShoppingCartData struct {
	Category string  `json:"category,omitempty"`
	ItemID   string  `json:"item_id,omitempty"`
	Price    *float64 `json:"price,omitempty"`
	Quantity *int64   `json:"quantity,omitempty"`
}
