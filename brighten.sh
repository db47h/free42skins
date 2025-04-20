#!/bin/bash
#
# This script generates brigher versions for dark theme users
# where free42 darkens the skin by 15%. This is a bit hackish
# but does the job well enough.
# 
go build -o brighten/brighten brighten/brighten.go

for f in *[lw].gif; do
  skin=$(dirname f)/$(basename -s .gif $f)
  gifsicle --transform-colormap ./brighten/brighten $f >${skin}_HC.gif
  cp ${skin}.layout ${skin}_HC.layout
  sed -i 's/a5b1a5 000022/c2d0c2 000028/' ${skin}_HC.layout
done
