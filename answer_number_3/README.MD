# ANSWER 

go.mod file which both defines projects requirements and locks dependencies to their correct version

A module is a collection of packages that are released, versioned, and distributed together. Modules may be downloaded directly from version control repositories or from module proxy servers.

A module is identified by a module path, which is declared in a go.mod file, together with information about the module’s dependencies. The module root directory is the directory that contains the go.mod file. The main module is the module containing the directory where the go command is invoked.

Each package within a module is a collection of source files in the same directory that are compiled together. A package path is the module path joined with the subdirectory containing the package (relative to the module root). For example, the module "golang.org/x/net" contains a package in the directory "html". That package’s path is "golang.org/x/net/html".


# How to import a third-party package in a Go project

1. go get <PACKAGE NAME><br/>


2. write third party package on go file :<br/>
    ![alt text](<Screenshot 2024-07-10 at 13.37.27.png>)
    <br/>

3. for update package we can use command : <br/>go mod tidy 