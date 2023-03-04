
# Handan Yayla

## Summary
This project has been written using Golang language. The project which is an e-commerce platform is about creating a new campaign. 
There are some commands about creating product, campaign and order. Also, the concept of time is also determined by commands.
İnitially,commands are read from the input file line-by-line. Each command that is splitted according to space is operated according to relevant bussiness logic. Finally, Result of commands is written to output file. So, to run the project, the names and paths of the input and output files must be entered as arguments from the command line.


### Prerequisites
```
Golang
```

### Run and Build
While in the directory where the main.go file is located, below command is run from the control line.
go run main.go [inputFileName] [outputFileName]

### Test

Test files is located under the model package. If you want to run all test files, you must use the below command.
- go test -v
P.S: I don't have enough time to write unit tests. Just a few unit tests related to the product
