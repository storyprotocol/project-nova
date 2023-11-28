Web3 gateway is the gateway layer for web3 services

## API Examples

## Upload content
grpcurl -plaintext -d '{"content":"eyJjb250ZW50IjoiR00gZnJvbSBwYWVsbGEifQ==","content_type":"MARKDOWN","storage":"ARWEAVE","tags":[{"name":"Content-Type","value":"application/json"}]}' localhost:8090 web3_gateway.StorageService/UploadContent

## fund irys account
grpcurl -plaintext -d '{"amountInWei":2113097404936000}' localhost:10002 web3_gateway.StorageService/FundIrysAccount