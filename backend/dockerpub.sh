
if [ $# -eq 0 ]
then
	echo "Usage dockerpub :versionNumber"
	exit 1
fi

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/MARIE .
docker build -t marie-backend:$1 .
docker tag marie-backend:$1 zenika/marie-backend:$1
docker push zenika/marie-backend:$1
