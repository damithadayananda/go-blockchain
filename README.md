# go-blockchain
blockchain written in go for educational purpose

# build single node blockchain 
`make -f Makefile.SingleNode build-single-node
`

# build three node blockchain
`make -f Makefile.MultipleNode build-multiple-node
`

# build ui
`make -f Makefile.ui all
`

# check service running
1. visit http://localhost:8080/chain
2. if build ui make target is already executed visit http://localhost:3000/

# unit test 
run the shell scrip inside scripts directory `./unit-test.sh`

# curl for triggering validate node
curl --location --request GET 'https://localhost:8080/node/validate' \
--header 'Content-Type: application/json' \
--data '{
"address":"d5efaa68cb559d469ee7f5b703b66f4900d60d4a"
}'

# certificate generation
./scripts/certificate-generation.sh <output dir>
./scripts/certificate-generation.sh ./secrets/node5

# running unit tests
go test -v -run '^(?!.*_IT$).*' ./...


