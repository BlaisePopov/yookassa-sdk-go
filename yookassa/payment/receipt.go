package yoopayment

import yoocommon "github.com/BlaisePopov/yookassa-sdk-go/yookassa/common"

type Receipt struct {
	// Customer info
	Customer *yoocommon.Customer `json:"customer"`

	// Array with the data regarding products or services (maximum 6 positions)
	Items []*yoocommon.Item `json:"items"`
}
