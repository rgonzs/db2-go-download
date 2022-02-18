$env:DB2_HOME = "C:\Users\Renato.Gonzales\go\pkg\mod\github.com\ibmdb\clidriver"
$env:CGO_CFLAGS = "-IC:\Users\Renato.Gonzales\go\pkg\mod\github.com\ibmdb\clidriver\include"
$env:CGO_LDFLAGS = "-LC:\Users\Renato.Gonzales\go\pkg\mod\github.com\ibmdb\clidriver\lib"
$env:LD_LIBRARY_PATH = ""
$env:GO111MODULE = "on"
echo $env:TCLLIBPATH
go build download.go