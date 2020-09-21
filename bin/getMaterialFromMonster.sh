#!/usr/bin/env bash
curl \
-X POST \
-H 'Content-Type: application/json' \
-d '{"monsterId": "1"}' \
"http://localhost:8080/hunters/1/get_material_from_monster" | jq