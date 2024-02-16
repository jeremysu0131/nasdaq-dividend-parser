package main

type Response struct {
	Data    Data        `json:"data"`
	Message interface{} `json:"message"`
	Status  Status      `json:"status"`
}

type Data struct {
	DividendHeaderValues []DividendHeaderValue `json:"dividendHeaderValues"`
	ExDividendDate       string                `json:"exDividendDate"`
	DividendPaymentDate  string                `json:"dividendPaymentDate"`
	Yield                string                `json:"yield"`
	AnnualizedDividend   string                `json:"annualizedDividend"`
	PayoutRatio          string                `json:"payoutRatio"`
	Dividends            Dividends             `json:"dividends"`
}

type DividendHeaderValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Dividends struct {
	AsOf    interface{} `json:"asOf"`
	Headers Headers     `json:"headers"`
	Rows    []Headers   `json:"rows"`
}

type Headers struct {
	ExOrEFFDate     string    `json:"exOrEffDate"`
	Type            TypeEnum  `json:"type"`
	Amount          string    `json:"amount"`
	DeclarationDate string    `json:"declarationDate"`
	RecordDate      string    `json:"recordDate"`
	PaymentDate     string    `json:"paymentDate"`
	Currency        *Currency `json:"currency,omitempty"`
}

type Status struct {
	RCode            int64       `json:"rCode"`
	BCodeMessage     interface{} `json:"bCodeMessage"`
	DeveloperMessage interface{} `json:"developerMessage"`
}

type Currency string

const (
	Usd Currency = "USD"
)

type TypeEnum string

const (
	Cash TypeEnum = "Cash"
	Type TypeEnum = "Type"
)
