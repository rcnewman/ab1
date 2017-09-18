var rp = require('request-promise');

var url = "http://129.158.64.209:8901/bcsgw/rest/v1/transaction/query"
var body = {
	"channel": "test",
	"chaincode": "loan",
	"chaincodeVer": "v1",
	"method": "getLoanLender",
	"args": []
}

var smb = {
	"smb_federal_ein": "12345-10",
	"smb_business_name": "Business A",
}
var lender = {
	"lender_federal_ein": "4865854-10",
	"lender_license_number": "4848-0912",
}
var loan = {
	"loan_id": "22222",
	"loan_type": "Big one",
}

body.args = [JSON.stringify(smb), JSON.stringify(lender), JSON.stringify(loan)];
var options = {
    method: 'POST',
    uri: url,
    body: body,
    json: true
};
rp.post(options).then(response => {
    console.log(response.result);
});
