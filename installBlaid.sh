#!/usr/bin/env bash

addUnitFile() {
    cp ./blaid.service /etc/systemd/system;
    sudo chmod 664 /etc/systemd/system/blaid.service;
}

startService() {
    sudo systemctl daemon-reload;
    sudo systemctl start blaid.service;
    sudo systemctl enable blaid.service;
}

main() {
    addUnitFile;
    startService;
}

main;