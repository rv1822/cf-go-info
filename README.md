Go program for clound foundry

The source code must be in 'src' folder

* Building 

need godep

`
 go get github.com/kr/godep
`

and save your dependencies
`
godep save
`


* Launching
set PORT=9000
go run main.go

* Deploying

cf push cf-info-go -m 10M
