#!/bin/bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
    echo "usage: $0 <file.sql>" >&2
    exit 1
fi

# use .env, if it has neareby
[[ -f .env ]] && { set -a; . ./.env; set +a; }

: "${MYSQL_PWD:?MYSQL_PWD not set - copy .env.example in .env}"

docker exec -e MYSQL_PWD="$MYSQL_PWD" -i ill-mysql8 mysql -u root < "$1"
