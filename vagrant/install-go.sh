#!/bin/bash

# install golang

bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)

[[ -s "$HOME/.gvm/scripts/gvm" ]] && source "$HOME/.gvm/scripts/gvm"

gvm install go1.4

echo "gvm use go1.4" >> ~/.bashrc



