### Building Images

This is a small attempt at generating bar charts as SVGs from a CSV given by a user.  
This makes use of svgo, a helpful module in go for the same.  

To run this:
`go build file.go` 
`./file -input=input.csv > test.svg`
