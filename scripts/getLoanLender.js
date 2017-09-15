var rp = require('request-promise');

// In the URL, `mca` is the name of the chaincode, and `smbblockorderer` is the name of the channel
var url = "http://129.158.70.201:8889/api/v1/chaincodes/mca/channels/smbblockorderer/invoke";
var body = {
    "request": {
        chaincodeVersion: "v1",
        fcn: "getLoanSMB",
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
}
var lender = {
	"lender_federal_ein": "4865854-10",
	"lender_license_number": "4848-0912",
}
var loan = {
	"loan_id": "22222",
	"loan_type": "Big one",
}

body.request.args = [JSON.stringify(smb), JSON.stringify(lender), JSON.stringify(loan)];
var options = {
    method: 'POST',
    uri: url,
    body: body,
    json: true
};
rp.post(options).then(response => {
    var loan = JSON.parse(response.metadata);
    console.log(loan);
});
