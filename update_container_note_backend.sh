#!/bin/bash

echo "📥 Pulling latest images..."
docker-compose pull note-backend

echo "🚀 Restarting containers with new images..."
docker-compose up -d note-backend

echo "🧹 Cleaning up unused images..."
docker image prune -a -f

echo "✅ Done!"
