# typemisskey
Typologitics on SNS

## Usage
```
git clone https://github.com/WisterViolet/typemisskey.git
cd typemisskey
docker build .
docker run -v (pwd):/home/src/typemisskey --env-file ./.env YOUR_IMAGE_NAME 
```
`docker run`:Don't use `--rm` option
