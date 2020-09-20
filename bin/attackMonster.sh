#!/usr/bin/env bash
curl \
-X PUT \
-H 'Content-Type: application/json' \
-d '{"monsterId":"1"}' \
"http://localhost:8080/hunters/1/attack" | jq
