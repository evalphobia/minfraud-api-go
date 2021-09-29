package minfraud

// BaseResponse is common response of minFraud API.
// ref: https://dev.maxmind.com/minfraud/api-documentation/responses
type BaseResponse struct {
	ErrData
	Disposition      Disposition `json:"disposition"`
	FundsRemaining   float64     `json:"funds_remaining"`
	ID               string      `json:"id"`
	QueriesRemaining int64       `json:"queries_remaining"`
	RiskScore        float64     `json:"risk_score"` // from 0.01 to 99
	Warnings         []Warning   `json:"warnings"`
}

func (r BaseResponse) HasError() bool {
	return r.ErrData.Code != ""
}

type ErrData struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type Disposition struct {
	Action    string `json:"action"`
	Reason    string `json:"reason"`
	RuleLabel string `json:"rule_label"`
}

type Warning struct {
	Code         string `json:"code"`
	InputPointer string `json:"input_pointer"`
	Warning      string `json:"warning"`
}

type IPAddress struct {
	Risk               float64            `json:"risk"` // from 0.01 to 99
	City               City               `json:"city"`
	Continent          Continent          `json:"continent"`
	Country            Country            `json:"country"`
	Location           Location           `json:"location"`
	Postal             Postal             `json:"postal"`
	RegisteredCountry  RegisteredCountry  `json:"registered_country"`
	RepresentedCountry RepresentedCountry `json:"represented_country"`
	RiskReasons        []RiskReason       `json:"risk_reasons"`
	Subdivisions       []Subdivision      `json:"subdivisions"`
	Traits             Traits             `json:"traits"`
}

type City struct {
	Confidence Confidence `json:"confidence"`
	GeoNameID  GeoNameID  `json:"geoname_id"`
	Names      Names      `json:"names"`
}

type Continent struct {
	Code      string    `json:"code"` // [AF, AN, AS, EU, NA, OC, SA]
	GeoNameID GeoNameID `json:"geoname_id"`
	Names     Names     `json:"names"`
}

// common data for country type object.
type CountryData struct {
	GeoNameID         GeoNameID `json:"geoname_id"`
	IsInEuropeanUnion bool      `json:"is_in_european_union"`
	ISOCode           string    `json:"iso_code"` // ISO 3166-1 alpha-2
	Names             Names     `json:"names"`
}

type Country struct {
	Confidence Confidence `json:"confidence"`
	CountryData
}

type Location struct {
	AccuracyRadius    int64   `json:"accuracy_radius"` // in Killometers
	AverageIncome     float64 `json:"average_income"`  // in USD
	Latitude          float64 `json:"latitude"`
	LocalTime         string  `json:"local_time"`
	Longitude         float64 `json:"longitude"`
	MetroCode         int64   `json:"metro_code"`
	PopulationDensity int64   `json:"population_density"` // only for US
	TimeZone          string  `json:"time_zone"`
}

type Postal struct {
	Code       string     `json:"code"`
	Confidence Confidence `json:"confidence"`
}

type RegisteredCountry struct {
	CountryData
}

type RepresentedCountry struct {
	CountryData
	Type string `json:"type"`
}

type RiskReason struct {
	Code   string `json:"code"`
	Reason string `json:"reason"`
}

type Subdivision struct {
	GeoNameID  GeoNameID  `json:"geoname_id"`
	ISOCode    string     `json:"iso_code"` // ISO 3166-1 alpha-2
	Names      Names      `json:"names"`
	Confidence Confidence `json:"confidence"`
}

type Traits struct {
	AutonomousSystemNumber       int64   `json:"autonomous_system_number"`
	AutonomousSystemOrganization string  `json:"autonomous_system_organization"`
	Domain                       string  `json:"domain"`
	IPAddress                    string  `json:"ip_address"`
	IsAnonymous                  bool    `json:"is_anonymous"`
	IsAnonymousVPN               bool    `json:"is_anonymous_vpn"`
	IsHostingProvider            bool    `json:"is_hosting_provider"`
	IsPublicProxy                bool    `json:"is_public_proxy"`
	IsResidentialProxy           bool    `json:"is_residential_proxy"`
	IsTorExitNode                bool    `json:"is_tor_exit_node"`
	ISP                          string  `json:"isp"`
	Network                      string  `json:"network"`
	Organization                 string  `json:"organization"`
	StaticIPScore                float64 `json:"static_ip_score"` // 0 ~ 99.99
	UserCount                    int64   `json:"user_count"`
	UserType                     string  `json:"user_type"`
}

type BillingAddress struct {
	DistanceToIPLocation int64   `json:"distance_to_ip_location"`
	IsInIPCountry        bool    `json:"is_in_ip_country"`
	IsPostalInCity       bool    `json:"is_postal_in_city"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
}

type CreditCard struct {
	Brand                           string `json:"brand"`
	Country                         string `json:"country"`
	IsBusiness                      bool   `json:"is_business"`
	IsIssuedInBillingAddressCountry bool   `json:"is_issued_in_billing_address_country"`
	IsPrepaid                       bool   `json:"is_prepaid"`
	IsVirtual                       bool   `json:"is_virtual"`
	Issuer                          Issuer `json:"issuer"`
	Type                            string `json:"type"`
}

type Issuer struct {
	MatchesProvidedName        bool   `json:"matches_provided_name"`
	MatchesProvidedPhoneNumber bool   `json:"matches_provided_phone_number"`
	Name                       string `json:"name"`
	PhoneNumber                string `json:"phone_number"`
}

type Device struct {
	Confidence Confidence `json:"confidence"`
	ID         string     `json:"id"`
	LastSeen   string     `json:"last_seen"`
	LocalTime  string     `json:"local_time"`
}

type Email struct {
	Domain       Domain `json:"domain"`
	FirstSeen    string `json:"first_seen"`
	IsDisposable bool   `json:"is_disposable"`
	IsFree       bool   `json:"is_free"`
	IsHighRisk   bool   `json:"is_high_risk"`
}

type Domain struct {
	FirstSeen string `json:"first_seen"`
}

type ShippingAddress struct {
	DistanceToBillingAddress int64   `json:"distance_to_billing_address"`
	DistanceToIPLocation     int64   `json:"distance_to_ip_location"`
	IsHighRisk               bool    `json:"is_high_risk"`
	IsInIPCountry            bool    `json:"is_in_ip_country"`
	IsPostalInCity           bool    `json:"is_postal_in_city"`
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
}

type Subscores struct {
	AVSResult                           float64 `json:"avs_result"`
	BillingAddress                      float64 `json:"billing_address"`
	BillingAddressDistanceToIPLocation  float64 `json:"billing_address_distance_to_ip_location"`
	Browser                             float64 `json:"browser"`
	Chargeback                          float64 `json:"chargeback"`
	Country                             float64 `json:"country"`
	CountryMismatch                     float64 `json:"country_mismatch"`
	CVVResult                           float64 `json:"cvv_result"`
	Device                              float64 `json:"device"`
	EmailAddress                        float64 `json:"email_address"`
	EmailDomain                         float64 `json:"email_domain"`
	EmailLocalPart                      float64 `json:"email_local_part"`
	IssuerIDNumber                      float64 `json:"issuer_id_number"`
	OrderAmount                         float64 `json:"order_amount"`
	PhoneNumber                         float64 `json:"phone_number"`
	ShippingAddress                     float64 `json:"shipping_address"`
	ShippingAddressDistanceToIPLocation float64 `json:"shipping_address_distance_to_ip_location"`
	TimeOfDay                           float64 `json:"time_of_day"`
}

type Names struct {
	DE   string `json:"de"`
	EN   string `json:"en"`
	ES   string `json:"es"`
	FR   string `json:"fr"`
	JA   string `json:"ja"`
	PtBR string `json:"pt-BR"`
	RU   string `json:"ru"`
	ZhCN string `json:"zh-CN"`
}

// ======= single value type ===========

// from 0 to 100
type Confidence int64

// ref: https://www.geonames.org/
type GeoNameID int64

// APIResponse is used to align the type from API request.
type APIResponse struct {
	BaseResponse
	IPAddress       IPAddress       `json:"ip_address"`
	BillingAddress  BillingAddress  `json:"billing_address"`
	CreditCard      CreditCard      `json:"credit_card"`
	Device          Device          `json:"device"`
	Email           Email           `json:"email"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
	Subscores       Subscores       `json:"subscores"`
}
