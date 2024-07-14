package models

type QRData struct {
	RawData                         string
	PayloadFormatIndicator          string
	PointOfInitiationMethod         string
	MerchantAccountInfo             map[string]MerchantAccountInfo
	MerchantCategoryCode            string
	TransactionCurrency             string
	TransactionAmount               string
	TipOrConvenienceIndicator       string
	ValueOfConvenienceFeeFixed      string
	ValueOfConvenienceFeePercentage string
	CountryCode                     string
	MerchantName                    string
	MerchantCity                    string
	PostalCode                      string
	AdditionalDataFieldTemplate     string
	CRC                             string
}

type MerchantAccountInfo struct {
	GlobalUniqueIdentifier string
	MerchantPAN            string
	MerchantID             string
	MerchantCriteria       string
}

type Response struct {
	Success bool
	Message string
	Data    QRData
}
