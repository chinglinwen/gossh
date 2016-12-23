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
  -t int
        timeout for a host in second (default 10)
  -u string
        user name
  -v    show version.

Examples: 
   ./gossh ip command
   echo date | ./gossh ip
   ./gossh -l ip.list command

Use as scp
  ./gossh -c srcfile host targetfile
  ./gossh -l ip.list -c srcfile targetfile

Environment variables:
  USER,PASS,PORT,HOST

Flag specify and Environment set the global user and pass

ip.list is a filename and can specify any file

Format exmaple:
ip user
ip
ip  user  pass

If omit user or pass, using global setting
Delimited by whitespace(or many continue whitespace)
It can specify different password for every ip(or domain name)

```

end.
