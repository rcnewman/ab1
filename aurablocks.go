// aurablocks.go

package main

import (
"fmt"
"time"
"encoding/json"
"github.com/hyperledger/fabric/core/chaincode/shim"
"github.com/hyperledger/fabric/protos/peer"
"errors"
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

	CashFlows float64 `json:"smb_cash_flows_from_gl,string"`
	DebtEquityRatio float64 `json:"smb_debt_to_equity_ratio,string"`
	WorkingCapital float64 `json:"smb_working_capital,string"`
	Currency string `json:"smb_currency"`
	GLSchedule time.Time `json:"smb_gl_schedule"`
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

	NetCreditReceipts float64 `json:"smb_net_credit_receipts,string"`
	ReceiptsSchedule time.Time `json:"smb_receipts_schedule"`
	ReceiptsBeginDay time.Time `json:"smb_receipts_begin_day"`
	ReceiptsEndDay time.Time `json:"smb_receipts_end_day"`

	ProjAvgMonRevenue float64 `json:"smb_proj_avg_mon_revenue,string"`
	ProjAvgMonCCReceipts float64 `json:"smb_proj_avg_mon_cc_receipts,string"`
	CumuCashFlows float64 `json:"smb_cumu_cash_flows_from_gl,string"`
	CumuNetCreditReceipts float64 `json:"smb_cumu_net_credit_receipts,string"`
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
	TotalLoanedAmount float64 `json:"loan_total_loaned_amount,string"`
	Term string `json:"loan_term"`
	RepaymentFreq string `json:"loan_repayment_freq"`
	CCSplitPercentage float64 `json:"loan_cc_split_percentage,string"`
	CCSplitSurchargePercentage float64 `json:"loan_cc_split_surcharge_percentage,string"`
	AvgExpMonPayment float64 `json:"loan_avg_exp_mon_payment,string"`
	CumuAvgExpMonPayment float64 `json:"loan_cumu_avg_exp_mon_payment"`
	MonPayProjAvgMonRevenue float64 `json:"loan_mon_pay_proj_avg_mon_revenue,string"`
	EstimatedAPR float64 `json:"loan_est_apr_based_on_est_payments,string"`
	FundedDate time.Time `json:"loan_funded_date"`
	RealAPR float64 `json:"loan_real_apr_based_on_payments_made,string"`
	CCSplitPayment float64 `json:"loan_cc_split_payment,string"`
	CumuCCSplitPayment float64 `json:"loan_cumu_cc_split_payment,string"`

	ActualPayment float64 `json:"loan_actual_payment,string"`
	CumuActualPayment float64 `json:"loan_cumu_actual_payment,string"`
	CumuRepaymentPercentage float64 `json:"loan_cumu_repayment_percentage,string"`

	LoanPerformance string `json:"loan_performance"`
	SplitPercentageCurMonPayment float64 `json:"loan_split_percent_cur_mon_payment,string"`
	FedRateAtLoanOrigination float64 `json:"loan_fed_rate_at_loan_origination,string"`
	FedCurrentRate float64 `json:"loan_fed_current_rate,string"`
	TerminateThreshold float64 `json:"loan_termination_threshold,string"`
	TerminationCount int64 `json:"loan_termination_count,string"`
	OnTrackPaymentCount int64 `json:"loan_on_track_payment_count,string"`
	TriggerInterestRateReview string `json:"loan_trigger_interest_rate_review"`
	Active bool `json:"loan_active"`
}

type Transaction struct {
	TxLoan Loan
	TxSMB SMB
	TxLender Lender
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
	var smbJSON SMB
	var lenderJSON Lender
	var loanJSON Loan
	var err error

	err = marshallRequest(args, &smbJSON, &lenderJSON, &loanJSON)
	if err != nil { return shim.Error("Failed to marshall request" + err.Error())}


        loanJSON.AvgExpMonPayment = smbJSON.ProjAvgMonCCReceipts * loanJSON.CCSplitPercentage
        loanJSON.MonPayProjAvgMonRevenue = loanJSON.AvgExpMonPayment / smbJSON.ProjAvgMonRevenue


	var tx Transaction
	tx.TxSMB = smbJSON
	tx.TxLender = lenderJSON
	tx.TxLoan = loanJSON

	key, err := stub.CreateCompositeKey("txKey", []string{smbJSON.FederalEIN, smbJSON.BusinessName, lenderJSON.FederalEIN, lenderJSON.LicenseNumber, loanJSON.LoanId, loanJSON.Type})
	if err != nil { return shim.Error(err.Error())}

	txAsBytes, err := json.Marshal(tx)
	if err != nil { return shim.Error(err.Error())}

	stub.PutState(key, txAsBytes)
	fmt.Println("- end onboardLoan")
	return shim.Success(nil)
}


func (t *AuraBlock) getLoanSMB(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting getLoanSMB")
	var smbQuery SMB
	var lenderQuery Lender
	var loanQuery Loan
	var err error

	err = marshallRequest(args, &smbQuery, &lenderQuery, &loanQuery)
        if err != nil { return shim.Error("Failed to marshall request: " + err.Error())}

	key, err := stub.CreateCompositeKey("txKey", []string{smbQuery.FederalEIN, smbQuery.BusinessName, lenderQuery.FederalEIN, lenderQuery.LicenseNumber, loanQuery.LoanId, loanQuery.Type})
	if err != nil { return shim.Error(err.Error())}

	txBytes, err  := stub.GetState(key)
	if err != nil {
		return shim.Error("Failed to get tx: " + err.Error())
	} else if txBytes == nil {
		return shim.Error("Tx does not exist. ")
	}

	var tx Transaction
	err = json.Unmarshal(txBytes, &tx)
	if err != nil { return shim.Error(err.Error()) }

	// Remove unwantedf fields for SMB
	tx.TxLender.LenderApiTriggerUrl1 = ""
	tx.TxLender.LenderApiTriggerUrl2 = "" 
	tx.TxLoan.AvgExpMonPayment = 0
	tx.TxLoan.CumuAvgExpMonPayment = 0
	tx.TxLoan.MonPayProjAvgMonRevenue = 0
	tx.TxLoan.EstimatedAPR = 0
	tx.TxLoan.RealAPR = 0
	tx.TxLoan.CCSplitPayment = 0
	tx.TxLoan.CumuCCSplitPayment = 0
	tx.TxLoan.CumuRepaymentPercentage = 0
	tx.TxLoan.LoanPerformance = ""
	tx.TxLoan.SplitPercentageCurMonPayment = 0
	tx.TxLoan.FedRateAtLoanOrigination = 0
	tx.TxLoan.FedCurrentRate = 0
	tx.TxLoan.TerminateThreshold = 0
	tx.TxLoan.TerminationCount = 0
	tx.TxLoan.OnTrackPaymentCount = 0
	tx.TxLoan.TriggerInterestRateReview = ""
	tx.TxLoan.Active = false


	txBytesOut, err := json.Marshal(tx)
	if err != nil { return shim.Error(err.Error())}

	fmt.Println("- end getLoanSMB")
	return shim.Success(txBytesOut)
}

func (t *AuraBlock) getLoanLender(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting getLoanLender")
        var smbQuery SMB
        var lenderQuery Lender
        var loanQuery Loan
        var err error

	err = marshallRequest(args, &smbQuery, &lenderQuery, &loanQuery)
        if err != nil { return shim.Error("Failed to marshall request: " + err.Error())}

        key, err := stub.CreateCompositeKey("txKey", []string{smbQuery.FederalEIN, smbQuery.BusinessName, lenderQuery.FederalEIN, lenderQuery.LicenseNumber, loanQuery.LoanId, loanQuery.Type})
        if err != nil { return shim.Error(err.Error())}

        txBytes, err  := stub.GetState(key)
        if err != nil {
                return shim.Error("Failed to get tx: " + err.Error())
        } else if txBytes == nil {
                return shim.Error("Tx does not exist. ")
        }

	fmt.Println("- end getLoanLender")
	return shim.Success(txBytes)
}


func (t *AuraBlock) updateCreditReceipts(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting updateCreditReceipts")

	var smbQuery SMB
	var lenderQuery Lender
	var loanQuery Loan
	var err error

	err = marshallRequest(args, &smbQuery, &lenderQuery, &loanQuery)
        if err != nil { return shim.Error("Failed to marshall request: " + err.Error())}

	key, err := stub.CreateCompositeKey("txKey", []string{smbQuery.FederalEIN, smbQuery.BusinessName, lenderQuery.FederalEIN, lenderQuery.LicenseNumber, loanQuery.LoanId, loanQuery.Type})
        if err != nil { return shim.Error(err.Error())}

	txBytes, err  := stub.GetState(key)
	if err != nil { return shim.Error(err.Error())}

	tx := Transaction{}
	err = json.Unmarshal(txBytes, &tx)
	if err != nil { return shim.Error(err.Error()) }


	tx.TxSMB.NetCreditReceipts = smbQuery.NetCreditReceipts
	tx.TxSMB.CumuNetCreditReceipts = tx.TxSMB.CumuNetCreditReceipts + smbQuery.NetCreditReceipts
	tx.TxSMB.ReceiptsSchedule = smbQuery.ReceiptsSchedule
	tx.TxSMB.ReceiptsBeginDay = smbQuery.ReceiptsBeginDay
	tx.TxSMB.ReceiptsEndDay = smbQuery.ReceiptsEndDay


	core(&smbQuery, &lenderQuery, &loanQuery, &tx)

	txAsBytes, err := json.Marshal(tx)
	if err != nil { return shim.Error(err.Error())}

	stub.PutState(key, txAsBytes)
	
	fmt.Println("- end updateCreditReceipts")
	return shim.Success(nil)
}

func (t *AuraBlock) updateGeneralLedger(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	fmt.Println("- starting updateGeneralLedger")

	var smbQuery SMB
    var lenderQuery Lender
    var loanQuery Loan
    var err error

	err = marshallRequest(args, &smbQuery, &lenderQuery, &loanQuery)
        if err != nil { return shim.Error("Failed to marshall request: " + err.Error())}

	key, err := stub.CreateCompositeKey("txKey", []string{smbQuery.FederalEIN, smbQuery.BusinessName, lenderQuery.FederalEIN, lenderQuery.LicenseNumber, loanQuery.LoanId, loanQuery.Type})
        if err != nil { return shim.Error(err.Error())}

    txBytes, err  := stub.GetState(key)
    if err != nil { return shim.Error(err.Error())}

	tx := Transaction{}
	err = json.Unmarshal(txBytes, &tx)
	if err != nil { return shim.Error(err.Error()) }

	tx.TxSMB.CashFlows = smbQuery.CashFlows
	tx.TxSMB.DebtEquityRatio = smbQuery.DebtEquityRatio
	tx.TxSMB.WorkingCapital = smbQuery.WorkingCapital
	tx.TxSMB.Currency = smbQuery.Currency
	tx.TxSMB.GLSchedule = smbQuery.GLSchedule
	tx.TxSMB.GLScheduleBeginDay = smbQuery.GLScheduleBeginDay
	tx.TxSMB.GLScheduleEndDay = smbQuery.GLScheduleEndDay 

	tx.TxSMB.Approval1Name = smbQuery.Approval1Name
	tx.TxSMB.Approval1Role = smbQuery.Approval1Role
	tx.TxSMB.Approval1Email = smbQuery.Approval1Email
	tx.TxSMB.Approval2Name = smbQuery.Approval2Name
	tx.TxSMB.Approval2Role = smbQuery.Approval2Role
	tx.TxSMB.Approval2Email = smbQuery.Approval2Email
	tx.TxSMB.Approval3Name = smbQuery.Approval3Name
	tx.TxSMB.Approval3Role = smbQuery.Approval3Role
	tx.TxSMB.Approval3Email = smbQuery.Approval3Email

	core(&smbQuery, &lenderQuery, &loanQuery, &tx)

	txAsBytes, err := json.Marshal(tx)
	if err != nil { return shim.Error(err.Error())}

	stub.PutState(key, txAsBytes)
	fmt.Println("- end updateGeneralLedger")
	return shim.Success(nil)
}

// ================
// UTILS
// ===============

func core(smbQuery *SMB, lenderQuery *Lender, loanQuery *Loan, tx *Transaction) {
	if (smbQuery.GLScheduleBeginDay.Equal(tx.TxSMB.ReceiptsSchedule) && tx.TxSMB.ReceiptsBeginDay.Equal(smbQuery.GLScheduleBeginDay) && tx.TxSMB.ReceiptsEndDay.Equal(smbQuery.GLScheduleEndDay)) {
		fmt.Println("executing function core")	

		tx.TxLoan.CumuAvgExpMonPayment = tx.TxLoan.CumuAvgExpMonPayment + tx.TxLoan.AvgExpMonPayment
		tx.TxLoan.CCSplitPayment = tx.TxSMB.NetCreditReceipts * tx.TxLoan.CCSplitPercentage
		tx.TxSMB.CumuCashFlows = tx.TxSMB.CumuCashFlows + smbQuery.CashFlows

		tx.TxLoan.CumuCCSplitPayment = tx.TxLoan.CumuCCSplitPayment + tx.TxLoan.CCSplitPayment
		tx.TxLoan.ActualPayment = tx.TxLoan.CCSplitPayment
		tx.TxLoan.CumuActualPayment = tx.TxLoan.CumuActualPayment + tx.TxLoan.CCSplitPayment

		tx.TxLoan.LoanPerformance = "Good"
		tx.TxLoan.OnTrackPaymentCount = tx.TxLoan.OnTrackPaymentCount + 1

		if (((tx.TxLoan.CumuActualPayment / tx.TxSMB.CumuCashFlows ) * 100) < 12){
			tx.TxLoan.ActualPayment = tx.TxLoan.ActualPayment + (tx.TxSMB.NetCreditReceipts * tx.TxLoan.CCSplitSurchargePercentage)
			tx.TxLoan.CumuActualPayment = tx.TxLoan.CumuActualPayment + tx.TxLoan.ActualPayment
			tx.TxLoan.LoanPerformance = "Surcharge"
			tx.TxLoan.SplitPercentageCurMonPayment = tx.TxLoan.CCSplitPercentage + tx.TxLoan.CCSplitSurchargePercentage
			tx.TxLoan.OnTrackPaymentCount = 0
		}
	}
}

func marshallRequest(args []string, smb *SMB, lender *Lender, loan *Loan) error {
	var err error
	if len(args) != 3 {
                return errors.New("Incorrect number of arguments, expecting 3")
        }

	err = json.Unmarshal([]byte(args[0]), &smb)
	if err != nil {	return err }
	fmt.Println("DEBUG: generated SMB %+v", smb)


	err = json.Unmarshal([]byte(args[1]), &lender)
	if err != nil {	return err }
	fmt.Println("DEBUG: generated lender %+v", lender)

	err = json.Unmarshal([]byte(args[2]), &loan)
	if err != nil {	return err }
	fmt.Println("DEBUG: generated loan %+v", loan)

	return nil
}

// ================
// MAIN
// ================

func main() {
	if err := shim.Start(new(AuraBlock)); err != nil {
		fmt.Printf("Error starting AuraBlock chaincode: %s", err)
	}
}
