var rp = require('request-promise');

var url = "http://129.158.64.209:8901/bcsgw/rest/v1/transaction/invocation"
var body = {
	"channel": "test",
	"chaincode": "loan",
	"chaincodeVer": "v1",
	"method": "updateGeneralLedger",
	"args": []
}
var smb = {
    "smb_federal_ein":"12345-10",
    "smb_business_name":"Business A",
    "smb_mailing_address":"Address 1",
    "smb_contact_name":"Bob Ross",
    "smb_email":"email@example.com",
    "smb_phone":"2345353",
    "smb_proj_avg_mon_revenue":"1200.02",
    "smb_proj_avg_mon_cc_receipts":"40949.32",
    "smb_cash_flows_from_gl":"10.00",
    "smb_debt_to_equity_ratio":"10.00",
    "smb_working_capital":"10.00",
    "smb_currency":"10.00",
    "smb_gl_schedule":"2015-09-15T12:00:00.000Z",
    "smb_gl_schedule_begin_day":"2015-09-15T12:00:00.000Z",
    "smb_gl_schedule_end_day":"2015-09-16T12:00:00.000Z",
    "smb_approval_1_name":"Bob Ross",
    "smb_approval_1_role":"Boss",
    "smb_approval_1_email":"a@email.com",
    "smb_approval_2_name":"Bob Ross2",
    "smb_approval_2_role":"Boss",
    "smb_approval_2_email":"a@email.com",
    "smb_approval_3_name":"Bob Ross3",
    "smb_approval_3_role":"Boss",
    "smb_approval_3_email":"10.00"
}
var lender = {
    "lender_federal_ein":"4865854-10",
    "lender_license_number":"4848-0912"
}
var loan = {
    "loan_id":"11111",
    "loan_type":"Big one"
}

body.args = [JSON.stringify(smb), JSON.stringify(lender), JSON.stringify(loan)]
var options = {
    method: 'POST',
    uri: url,
    body: body,
    json: true
};
rp.post(options).then(console.log);
