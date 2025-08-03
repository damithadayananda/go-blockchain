#!/bin/bash
set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <output-directory>"
  exit 1
fi

TARGET_DIR="$1"
mkdir -p "$TARGET_DIR"

# Generate with more explicit configuration
openssl req -x509 -nodes -days 365 \
  -newkey rsa:2048 \
  -keyout "$TARGET_DIR/key.pem" \
  -out "$TARGET_DIR/cert.pem" \
  -subj "/C=US/ST=State/L=City/O=Organization/CN=localhost" \
  -extensions v3_req \
  -config <(echo "[req]"; echo "distinguished_name=req"; echo "[v3_req]"; echo "subjectAltName=DNS:localhost,IP:127.0.0.1")

# Verify generation
if openssl x509 -in "$TARGET_DIR/cert.pem" -noout 2>/dev/null; then
  echo "✓ Certificate generated successfully at: $TARGET_DIR"
else
  echo "✗ Certificate generation failed"
  exit 1
fi

#Printing address
PUBLIC_KEY=$(openssl rsa -in "$TARGET_DIR/key.pem" -pubout 2>/dev/null)
ADDRESS=$(echo -n "$PUBLIC_KEY" | openssl dgst -sha256 -binary | xxd -p -c 20 | head -c 40)
echo "✓ Blockchain address: $ADDRESS"
