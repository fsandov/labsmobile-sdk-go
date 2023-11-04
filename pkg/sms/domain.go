package sms

type SendSMSRequest struct {
	Message   string      `json:"message"`             // Required
	Recipient []Recipient `json:"recipient"`           // Required
	Tpoa      *string     `json:"tpoa,omitempty"`      // Optional
	SubID     *string     `json:"subid,omitempty"`     // Optional
	Label     *string     `json:"label,omitempty"`     // Optional
	Test      *int        `json:"test,omitempty"`      // Optional
	AckURL    *string     `json:"ackurl,omitempty"`    // Optional
	ShortLink *int        `json:"shortlink,omitempty"` // Optional
	ClickURL  *string     `json:"clickurl,omitempty"`  // Optional
	Scheduled *string     `json:"scheduled,omitempty"` // Optional
	Long      *int        `json:"long,omitempty"`      // Optional
	Crt       *string     `json:"crt,omitempty"`       // Optional
	CrtID     *string     `json:"crt_id,omitempty"`    // Optional
	CrtName   *string     `json:"crt_name,omitempty"`  // Optional
	UCS2      *int        `json:"ucs2,omitempty"`      // Optional
	NoFilter  *int        `json:"nofilter,omitempty"`  // Optional
}

type Recipient struct {
	Msisdn string `json:"msisdn"` // Required

}

type SendSMSResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	SubID   string `json:"subid,omitempty"`
}

type GetBalanceResponse struct {
	Code    interface{} `json:"code"`
	Credits string      `json:"credits"`
}
