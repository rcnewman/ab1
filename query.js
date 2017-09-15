var rp = require('request-promise');

function json_escape(obj) {
	return JSON.stringify(obj);
}

var url = "http://129.158.70.201:8889/api/v1/chaincodes/mca/channels/smbblockorderer/invoke";
var body = {
    "request": {
        chaincodeVersion: "v1",
        fcn: "onboardLoan",
        args: []
    },
	orderer: {
		Addr: "orderer.smbblock.com",
		Port: "7050"
	},
	"peers": [
		{"Addr":"peer0.smbblock.com","Port":"7051"},
		{"Addr":"peer1.smbblock.com","Port":"7051"}
	]
}

var smb = {
	"smb_federal_ein": "12345-10",
	"smb_business_name": "Business A",
	"smb_mailing_address": "Address 1",
	"smb_contact_name": "Bob Ross",
	"smb_email": "email@example.com",
	"smb_phone": "2345353",
	"smb_proj_avg_mon_revenue": 1200.02,
	"smb_proj_avg_mon_cc_receipts": 40949.32
}
var lender = {
	"lender_federal_ein": "4865854-10",
	"lender_license_number": "4848-0912",
	"lender_business_name": "EGGGGZ",
	"lender_mailing_address": "This is a street",
	"lender_contact_name": "Jim",
	"lender_email": "jim@egggz.com",
	"lender_phone": "349494904"
}
var loan = {
	"loan_id": "11111",
	"loan_type": "Big one",
	"loan_total_loaned_amount": "10.30",
	"loan_term": "An amount of itme",
	"loan_repayment_freq": "Whenever",
	"loan_cc_split": "10.00",
	"loan_cc_split_surcharge_percentage": "10.00",
	"loan_funded_date": "2015-09-15T12:00:00.000Z",
	"loan_termination_threshold": "1"
}

body.request.args = [json_escape(smb), json_escape(lender), json_escape(loan)];
var options = {
    method: 'POST',
    uri: url,
    body: body,
    json: true
};
rp.post(options).then(console.log);
