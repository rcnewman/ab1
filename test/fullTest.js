var fs = require('fs');
var rp = require('request-promise');

var bcs_functions = { 
    API_CALL_Lender_On_Board_MCA_Smart_Contract: "onboardLoan",
    API_CALL_Lender_Get_MCA_Smart_Contract: "getLoanLender",
    API_CALL_SMB_Post_GL_MCA_Smart_Contract: "updateGeneralLedger",
    API_CALL_Lender_Post_Credit_Receipts_MCA_Smart_Contract: "updateCreditReceipts"
}

var csv = fs.readFileSync("test_seed.csv", "utf8")
var lines = csv.split('\n')
               .map(l => l.split(','));
lines.pop(); // get rid of empty line

var keys = lines.shift().map(key => key.trim());

var rows = []
for (i in lines) {
    var row = {}
    for (j in keys) {
        var value = lines[i][j];
        if (value && value.includes('/'))
            value = modifyDate(value);
        if (value && value != 'na')
            row[keys[j]] = value;
    }
    rows.push(row);
}

var t = 0;
rows
    .forEach(function (row) {
        setTimeout(function () {
            processRow(row);
        }, t)
        t += 5000;
    });

function modifyDate(input) {
    var parts = input.split('/');
    var month = parts[0];
    var date = parts[1];
    var year = parts[2];

    if (month.length == 1)
        month = "0" + month;
    if (date.length == 1)
        date = "0" + date;
    return `${year}-${month}-${date}T00:00:00Z`;
}

function processRow (row) {
    var smb = {}
    var lender = {}
    var loan = {}

    var keys = Object.keys(row).filter(key => key.includes('smb'));
    keys.forEach(key => smb[key] = row[key]);

    var keys = Object.keys(row).filter(key => key.includes('lender'));
    keys.forEach(key => lender[key] = row[key]);

    var keys = Object.keys(row).filter(key => key.includes('loan'));
    keys.forEach(key => loan[key] = row[key]);

    var body = {
        "channel": "test",
        "chaincode": "loan",
        "chaincodeVer": "v1",
        "method": bcs_functions[row.api],
        "args": [JSON.stringify(smb), JSON.stringify(lender), JSON.stringify(loan)]
    }

    if (bcs_functions[row.api] == 'getLoanLender')
        var url = "http://129.158.64.209:8901/bcsgw/rest/v1/transaction/query"
    else
        var url = "http://129.158.64.209:8901/bcsgw/rest/v1/transaction/invocation"
    var options = {
        method: 'POST',
        uri: url,
        body: body,
        json: true
    };
    rp.post(options).then(function (response) {
        logEvent(response);
        console.log(response);
    });
}

function logEvent(event) {
    var timestamp = new Date();
    var event = JSON.stringify(event);
    var line = `${timestamp}: ${event}\n`;
    fs.appendFileSync('test.log', line);
}
