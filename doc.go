package main

var example = `
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

If omit user or pass, using global setting (for the first entry)
For others entries ( using previous line's user and password )

Delimited by whitespace(or many continue whitespace)
It can specify different password for every ip(or domain name)

`
