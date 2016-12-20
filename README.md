# gossh
A ssh client tool written in Go

Provide ssh and scp function without interactive input

## Usage

```
Usage of gossh:
  -c string
        scp file to copy
  -l string
        list file of hosts
  -p string
        password
  -port string
        ssh port (default "22")
  -u string
        user name
  -v    show version.

Examples: 
   ./gossh ip command
   echo date | ./gossh ip
   ./gossh -l ip.list command

Use as scp
  ./gossh -c srcfile host targetfile

Environment variables:
  USER,PASS,PORT,HOST

```

end.