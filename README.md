# Modify Google Microservice to use DynamoDB

_Work in Progress_

See [`dynamodb-json-demo`](https://github.com/garystafford/dynamodb-json-demo) GitHub project for DynamoDB-related files.

## Commands
```bash
# set credentials
export AWS_REGION="us-east-1"
export AWS_ACCESS_KEY_ID=""
export AWS_SECRET_ACCESS_KEY=""
export AWS_SESSION_TOKEN=""

# config go
GOROOT=/usr/local/go
GOPATH=/Users/garystaf/Documents/projects/go

# build and run
go build
./productcatalogservice
```

## Output
```text
{"message":"Contents of the Product Catalog: :{[id:\"OLJCESPC7Z\" name:\"Vintage Typewriter\" description:\"This 
typewriter looks good in your living room.\" picture:\"/static/img/products/typewriter.jpg\" price_usd:
\u003ccurrency_code:\"USD\" units:67 nanos:990000000 \u003e categories:\"vintage\"  id:\"66VCHSJNUP\" name:
\"Vintage Camera Lens\" description:\"You won't have a camera to use it and it probably doesn't work anyway.
\" picture:\"/static/img/products/camera-lens.jpg\" price_usd:\u003ccurrency_code:\"USD\" units:12 nanos:490000000 
\u003e categories:\"photography\" categories:\"vintage\"  id:\"1YMWWN1N4O\" name:\"Home Barista Kit\" description:
\"Always wanted to brew coffee with Chemex and Aeropress at home?\" picture:\"/static/img/products/barista-kit.jpg
\" price_usd:\u003ccurrency_code:\"USD\" units:124 \u003e categories:\"cookware\"  id:\"9SIQT8TOJO\" name:\"City Bike
\" description:\"This single gear bike probably cannot climb the hills of San Francisco.\" picture:\"/static/img/
products/city-bike.jpg\" price_usd:\u003ccurrency_code:\"USD\" units:789 nanos:500000000 \u003e categories:
\"cycling\"  id:\"0PUK6V6EV0\" name:\"Vintage Record Player\" description:\"It still works.\" picture:
\"/static/img/products/record-player.jpg\" price_usd:\u003ccurrency_code:\"USD\" units:65 nanos:500000000 
\u003e categories:\"music\" categories:\"vintage\"  id:\"L9ECAV7KIM\" name:\"Terrarium\" description:\"This terrarium 
will looks great in your white painted living room.\" picture:\"/static/img/products/terrarium.jpg\" price_usd:
\u003ccurrency_code:\"USD\" units:36 nanos:450000000 \u003e categories:\"gardening\"  id:\"2ZYFJ3GM2N\" name:
\"Film Camera\" description:\"This camera looks like it's a film camera, but it's actually digital.\" picture:
\"/static/img/products/film-camera.jpg\" price_usd:\u003ccurrency_code:\"USD\" units:2245 \u003e categories:
\"photography\" categories:\"vintage\"  id:\"6E92ZMYYFZ\" name:\"Air Plant\" description:\"Have you ever wondered 
whether air plants need water? Buy one and figure out.\" picture:\"/static/img/products/air-plant.jpg\" price_usd:
\u003ccurrency_code:\"USD\" units:12 nanos:300000000 \u003e categories:\"gardening\"  id:\"LS4PSXUNUM\" name:
\"Metal Camping Mug\" description:\"You probably don't go camping that often but this is better than plastic cups.
\" picture:\"/static/img/products/camp-mug.jpg\" price_usd:\u003ccurrency_code:\"USD\" units:24 nanos:330000000 
\u003e categories:\"cookware\" ] {}  %!s(int32=0)}","severity":"info","timestamp":"2020-02-01T22:46:26.138322-08:00"}
{"message":"starting grpc server at :3550","severity":"info","timestamp":"2020-02-01T22:46:26.138667-08:00"}
```

## AWS Linux2 EC2 Test
```bash
# ssh to EC2 host
yes | sudo yum update
yes | sudo yum install git
wget https://dl.google.com/go/go1.13.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.13.7.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
# go version go1.13.7 linux/amd64

GOROOT=/usr/local/go
GOPATH=/Users/ec2-user/go
GOBIN=$GOPATH/bin

wget https://github.com/garystafford/productcatalogservice/archive/master.zip
unzip master.zip
mkdir -p /home/ec2-user/go/src/github.com/garystafford/productcatalogservice
mv /home/ec2-user/productcatalogservice-master/* /home/ec2-user/go/src/github.com/garystafford/productcatalogservice
cd /home/ec2-user/go/src/github.com/garystafford/productcatalogservice

go get ./...
go build

# export AWS_REGION="us-east-1"
./productcatalogservice
```

## References
- <https://github.com/GoogleCloudPlatform/microservices-demo>
- <https://github.com/GoogleCloudPlatform/microservices-demo/blob/master/src/productcatalogservice/products.json>
- <https://github.com/blueCycle>
- <https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/dynamodb>
