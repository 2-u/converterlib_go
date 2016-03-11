# converterlib_go
Common type conversions written in Go

# Make sure you set your GOPATH correctly
cd $GOPATH/src

# If you make changes, please run gofmt
gofmt -w github.com/mraym/converterlib_go/*.go

# Build the library
go build github.com/mraym/converterlib_go

# Run the test to make sure that there is output.
# You should see everything FAIL.
# That's ok, I just want to see the outputs.
go test github.com/mraym/converterlib_go
