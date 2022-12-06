curl --header "Content-Type: application/json-patch+json" \
--request PATCH \
--data '[{"op": "remove", "path": "/status/capacity/cnvrg.io~1metagpu"}]' \
http://localhost:8001/api/v1/nodes/devopsdaystlv01/status