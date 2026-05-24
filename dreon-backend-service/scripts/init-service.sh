#!/usr/bin/env sh
set -eu

if [ "$#" -ne 2 ]; then
	echo "usage: scripts/init-service.sh <module-path> <service-name>"
	echo "example: scripts/init-service.sh github.com/hiamthach108/dreon-payment dreon-payment"
	exit 1
fi

module_path="$1"
service_name="$2"
database_name="$(printf '%s' "$service_name" | tr '-' '_')"

files="$(find . -type f \
	-not -path './.git/*' \
	-not -path './tmp/*' \
	-not -path './vendor/*')"

printf '%s\n' "$files" | xargs perl -pi -e "s#github.com/hiamthach108/dreon-backend-service#$module_path#g"
printf '%s\n' "$files" | xargs perl -pi -e "s#dreon-backend-service#$service_name#g"
printf '%s\n' "$files" | xargs perl -pi -e "s#dreon_backend_service#$database_name#g"

go mod tidy
gofmt -w .

echo "initialized $service_name ($module_path)"
