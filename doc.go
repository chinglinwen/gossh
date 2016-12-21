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

`
