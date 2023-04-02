#!/usr/bin/env bash

setContext() {
  goPath="${GOPATH}";
  cd $goPath/blaid;
}

addUnitFile() {
  cp ./blaid.service /etc/systemd/system;
  sudo chmod 664 /etc/systemd/system/blaid.service;
}

startService() {
  sudo /usr/bin/systemctl daemon-reload;
  sudo /usr/bin/systemctl start blaid.service;
  sudo /usr/bin/systemctl enable blaid.service;
}

main() {
    setContext;
    addUnitFile;
    startService;
}

main;