<h1>List of golang commands and usage</h1>

1. **go version**
    >Displays the version of go bin installed.

1. **go env**
    >Displays the go environmental variables set.

1. **go help [commands]**
    >Gives information about commands.

1. **go fmt <dir/file>**
    >Formats the code as per go standards.

1. **go run <filename>**
    > Builds and Executes the program in the file. 

1. **go build [filename]**
    >Builds the executable.  
     Provide the file name if file name isn't main.go .

1. **go get <location>**
    >Fetch the contents from the location, where location can be a link to repo
     or any url

1. **go mod init [locationOfModule/ModuleName]**
    >Create  a go module and Initializes the go *.mod* file.  
     give the location of module if outside go PATH.
     Egs: go mod init moduleExample/hello 

1. **go list -m all**
    >Prints all the modules dependencies.

1. **go mod tidy**
    >Cleans up unused dependencies.

1. **go list -m -versions <module/dependency>**
    >Lists all the versions of the module.                                                                                               
 
1. **go get <module/dependency>[@version]**
    >Fetches the module/dep on the specified/latest.                       