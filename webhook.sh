#!/bin/bash

curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
     "url": "https://us-central1-uc-internal-sandbox.cloudfunctions.net/bot"
   }' \
   https://api.telegram.org/bot389481972:AAGdjqFkdE-RndG_xTfyzfNNiojFJ4zr2JA/setWebhook