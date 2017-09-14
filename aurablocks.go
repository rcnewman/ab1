// aurablocks.go

package main

import (
"fmt"
"time"
"encoding/json"
"github.com/hyperledger/fabric/core/chaincode/shim"
"github.com/hyperledger/fabric/protos/peer"
)
// #########################
//    Type Definitions
// #########################

type AuraBlock struct {
}

type SMB struct {
	FederalEIN string `json:"smb_federal_ein"`
	BusinessName string `json:"smb_business_name"`
	MailingAddress string `json:"smb_mailing_address"`
	ContactName string `json:"smb_contact_name"`
	Email string `json:"smb_email"`
	Phone string `json:"smb_phone"`
	
	CashFlows float64 `json:"smb_cash_flows_from_gl"`
	DebtEquityRatio float64 `json:"smb_debt_to_equity_ratio"`
	WorkingCapital float64 `json:"smb_working_capital"`
	Currency string `json:"smb_currency"`
	GLSchedule string `json:"smb_gl_schedule"`
	GLScheduleBeginDay time.Time `json:"smb_gl_schedule_begin_day"`
	GLScheduleEndDay time.Time `json:"smb_gl_schedule_end_day"`

	SMBApiTriggerUrl1 string `json:"smb_api_trigger_api_url1"`
	SMBApiTriggerUrl2 string `json:"smb_api_trigger_api_url2"`

	Approval1Name string `json:"smb_approval_1_name"`
	Approval1Role string `json:"smb_approval_1_role"`
	Approval1Email string `json:"smb_approval_1_email"`
	Approval2Name string `json:"smb_approval_2_name"`
	Approval2Role string `json:"smb_approval_2_role"`
	Approval2Email string `json:"smb_approval_2_email"`
	Approval3Name string `json:"smb_approval_3_name"`
	Approval3Role string `json:"smb_approval_3_role"`
	Approval3Email string `json:"smb_approval_3_email"`

	NetCreditReceipts float64 `json:"smb_net_credit_receipts"`
	ReceiptsSchedule string `json:"smb_receipts_schedule"`
	ReceiptsBeginDay time.Time `json:"smb_receipts_begin_day"`
	ReceiptsEndDay time.Time `json:"smb_receipts_end_day"`

	ProjAvgMonRevenue float64 `json:"smb_proj_avg_mon_revenue"`
	ProjAvgMonCCReceipts float64 `json:"smb_proj_avg_mon_cc_receipts"`
	CumuCashFlows float64 `json:"smb_cumu_cash_flows_from_gl"`
	CumuNetCreditReceipts float64 `json:"smb_cumu_net_credit_receipts"`
}

type Lender struct {
	FederalEIN string `json:"lender_federal_ein"`
	LicenseNumber string `json:"lender_license_number"`
	BusinessName string `json:"lender_business_name"`
	MailingAddress string `json:"lender_mailing_address"`
	ContactName string `json:"lender_contact_name"`
	Email string `json:"lender_email"`
	Phone string `json:"lender_phone"`
	LenderApiTriggerUrl1 string `json:"lender_api_trigger_url1"`
	LenderApiTriggerUrl2 string `json:"lender_api_trigger_url2"`
	// LenderApiTriggerUrl3 string `json:"lender_api_trigger_url3"`
	// LenderApiTriggerUrl4 string `json:"lender_api_trigger_url4"`
	// LenderApiTriggerUrl5 string `json:"lender_api_trigger_url5"`
}


type Auditor struct {
	FederalEIN string `json:"auditor_federal_ein"`
	LicenseNumber string `json:"auditor_license_number"`
	BusinessName string `json:"auditor_business_name"`
	MailingAddress string `json:"auditor_mailing_address"`
	ContactName string `json:"auditor_contact_name"`
	Email string `json:"auditor_email"`
	Phone string `json:"auditor_phone"`
}

type Loan struct {
	LoanId string `json:"loan_id"`
	Type string `json:"loan_type"`
	TotalLoanedAmount float64 `json:"loan_total_loaned_amount"`
	Term string `json:"loan_term"`
	RepaymentFreq string `json:"loan_repayment_freq"`
	CCSplitPercentage float64 `json:"loan_cc_split_percentage"`
	CCSplitSurchargePercentage float64 `json:"loan_cc_split_surcharge_percentage"`
	AvgExpMonPayment float64 `json:"loan_avg_exp_mon_payment"`
	CumuAvgExpMonPayment float64 `json:"loan_cumu_avg_exp_mon_payment"`
	MonPayProjAvgMonRevenue float64 `json:"loan_mon_pay_proj_avg_mon_revenue"`
	EstimatedAPR float64 `json:"loan_est_apr_based_on_est_payments"`
	FundedDate time.Time `json:"loan_funded_date"`
	RealAPR float64 `json:"loan_real_apr_based_on_payments_made"`
	CCSplitPayment float64 `json:"loan_cc_split_payment"`
	CumuCCSplitpayment float64 `json:"loan_cumu_cc_split_payment"`

	ActualPayment float64 `json:"loan_actual_payment"`
	CumuActualPayment float64 `json:"loan_cumu_actual_payment"`
	CumuRepaymentPercentage float64 `json:"loan_cumu_repayment_percentage"`

	LoanPerformance string `json:"loan_performance"`
	SplitPercentageCurMonPayment float64 `json:"loan_split_percent_cur_mon_payment"`
	FedRateAtLoanOrigination float64 `json:"loan_fed_rate_at_loan_origination"`
	FedCurrentRate float64 `json:"loan_fed_current_rate"`
	TerminateThreshold float64 `json:"loan_termination_threshold`
	TerminationCount int64 `json:"loan_termination_count"`
	OnTrackPaymentCount int64 `json:"loan_on_track_payment_count"`
	TriggerInterestRateReview string `json:"loan_trigger_interest_rate_review"`
	Active bool `json:"loan_active"`
}

type Covenant struct {
}

// ##################
//       INIT
// ##################


func (t *AuraBlock) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// ########################
//     Invocations
// ########################


func (t *AuraBlock) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
        // Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println("DEBUG: invoke is running " + fn)
	fmt.Println("DEBUG: args %+v",args)

	if fn == "onboardLoan" {
		return t.onboardLoan(stub, args)
	} else if fn == "getLoanSMB" {
		return t.getLoanSMB(stub, args)
	} else if fn == "getLoanLender" {
		return t.getLoanLender(stub, args)
	} else if fn == "updateGeneralLedger" {
		return t.updateGeneralLedger(stub, args)
	} else if fn == "updateCreditReceipts" {
		return t.updateCreditReceipts(stub, args)
	}

	fmt.Println("invoke did not find function: " + fn)

	return shim.Error("Recieved unknown function invocation")
}

func (t *AuraBlock) onboardLoan(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting onboardLoan")
	fmt.Println("- starting onboardLoan")
	var smbJSON SMB
	var lenderJSON Lender
	var loanJSON Loan
	var covenantJSON Covenant
	var err error

	err = json.Unmarshal([]byte(args[0]), &smbJSON)
	if err != nil {
		return shim.Error("Failed to decode JSON of smb: " + args[0])
	}
	fmt.Println("DEBUG: generated SMB %+v", smbJSON)

	err = json.Unmarshal([]byte(args[1]), &lenderJSON)
	if err != nil {
		return shim.Error("Failed to decode JSON of lender: " + args[1])
	}
	fmt.Println("DEBUG: generated lender %+v", lenderJSON)

	err = json.Unmarshal([]byte(args[2]), &loanJSON)
	if err != nil {
		return shim.Error("Failed to decode JSON of loan: " + args[2])
	}
	fmt.Println("DEBUG: generated loan %+v", loanJSON)

	err = json.Unmarshal([]byte(args[3]), &covenantJSON)
	if err != nil {
		return shim.Error("Failed to decode JSON of covenant: " + args[3])
	}
	fmt.Println("DEBUG: generated covenant %+v", covenantJSON)


	fmt.Println("- end onboardLoan")
	return shim.Success(nil)
}


func (t *AuraBlock) getLoanSMB(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting getLoanSMB")

	fmt.Println("- end getLoanSMB")
	return shim.Success(nil)
}

func (t *AuraBlock) getLoanLender(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting getLoanLender")

	fmt.Println("- end getLoanLender")
	return shim.Success(nil)
}

func (t *AuraBlock) updateCreditReceipts(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting updateCreditReceipts")

	fmt.Println("- end updateCreditReceipts")
	return shim.Success(nil)
}

func (t *AuraBlock) updateGeneralLedger(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting updateGeneralLedger")

	fmt.Println("- end updateGeneralLedger")
	return shim.Success(nil)
}
// ================
// MAIN
// ================

func main() {
	if err := shim.Start(new(AuraBlock)); err != nil {
		fmt.Printf("Error starting AuraBlock chaincode: %s", err)
	}
}
