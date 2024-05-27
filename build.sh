#!/bin/bash

printf "Building core\n"
yarn core build

printf "\nBuilding dependencies\n"
yarn db build &
apid=$!
yarn auth build &
bpid=$!
yarn files build

wait $apid
wait $bpid

printf "\nBuilding api and ui\n"
yarn api build &
apid=$!
yarn ui build
wait $apid