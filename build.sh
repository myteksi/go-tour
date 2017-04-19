set -e
docker build -t gotour .
docker tag gotour:latest zhchang/gotour:latest
docker push zhchang/gotour:latest
