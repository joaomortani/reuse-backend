#!/bin/sh

echo ">>> Iniciando aplicação com Air..."
echo "Conteúdo de air.toml:"
cat /app/air.toml

exec /usr/local/bin/air -c /app/air.toml