#!/bin/bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
    echo "usage: $0 <file.sql>" >&2
    exit 1
fi

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
env_file="$script_dir/../.env"

[[ -f "$env_file" ]] && { set -a; . "$env_file"; set +a; }

: "${MYSQL_PWD:?MYSQL_PWD not set - copy .env.example to .env}"

docker exec -e MYSQL_PWD="$MYSQL_PWD" -i ill-mysql8 mysql -u root < "$1"
