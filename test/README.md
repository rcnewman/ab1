# Biz2Credit Testing

## Full Test

To run the test script:
   
1. A default `test_seed.csv` file is provided. If you want to use your own 
custom test seeds, replace this file with a file of the same name with your
seed data.

2. Install dependencies: Run `npm install`

3. Execute the test script: `node fullTest.js`

## Function Testing

If you would like to test individual functions, use the scripts provided for
each function. For example, to test the onboardLoan function, `node onboardLoan.js`

## Logs

Full testing logs can be found in the `test.log` file. 

If you are looking for just the returned query data from the blockchain to compare
against your seed data, the query info is logged to `loaninfo.csv` in CSV format. 
Compare it side by side against `test_seed.csv` to quickly validate the tests. 

By default, `fullTest.js` will append to the `loaninfo.csv` file and preserve old 
tests. To clear the data, delete the existing `loaninfo.csv`  file before running 
the test.
