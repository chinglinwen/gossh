#!/bin/sh
. ~/.bashrc
rm -f gossh*
build32
winbuild
#solabuild -o gossh.solaris

ver="$( grep 'version=' *.go | awk '{ print $2 }' FS='=' | \
              awk '{ print $1 }' FS=',' )"
tar -czf gossh.v$ver.tar.gz gossh gossh.exe
