
if [ $# -eq 0 ]
then
	echo "Usage dockerpub :versionNumber"
	exit 1
fi

npm run build
docker build -t marie-backoffice:$1 .
docker tag marie-backoffice:$1 zenika/marie-backoffice:$1
docker push zenika/marie-backoffice:$1
