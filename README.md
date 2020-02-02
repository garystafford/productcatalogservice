# Modify Google Microservice productcatalogservice to use DynamoDB

See [`dynamodb-json-demo`](https://github.com/garystafford/dynamodb-json-demo) GitHub project for DynamoDB-related files. Build the CloudFormation stack, then run the Python script to write products to DynamoDB from JSON file.

## Commands
### Build and Test
```bash
# set credentials to run locally
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

# test server.go
# set creds
go test -v
```

### Build and Run Docker Container
```bash
# build Dockerfile
# use old Gopkg.toml
dep ensure
dep ensure -update

docker build -t garystafford/productcatalogservice:1.0.0 . --no-cache

# run the image
docker run -d \
    --publish 3550:3550 \
    --env AWS_ACCESS_KEY_ID \
    --env AWS_SECRET_ACCESS_KEY \
    --env AWS_SESSION_TOKEN \
    --name productcatalogservice garystafford/productcatalogservice:1.0.0
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

# productcatalogservice

Run the following command to restore dependencies to `vendor/` directory:

    dep ensure --vendor-only

## Dynamic catalog reloading / artificial delay

This service has a "dynamic catalog reloading" feature that is purposefully
not well implemented. The goal of this feature is to allow you to modify the
`products.json` file and have the changes be picked up without having to
restart the service.

However, this feature is bugged: the catalog is actually reloaded on each
request, introducing a noticeable delay in the frontend. This delay will also
show up in profiling tools: the `parseCatalog` function will take more than 80%
of the CPU time.

You can trigger this feature (and the delay) by sending a `USR1` signal and
remove it (if needed) by sending a `USR2` signal:

```
# Trigger bug
kubectl exec \
    $(kubectl get pods -l app=productcatalogservice -o jsonpath='{.items[0].metadata.name}') \
    -c server -- kill -USR1 1
# Remove bug
kubectl exec \
    $(kubectl get pods -l app=productcatalogservice -o jsonpath='{.items[0].metadata.name}') \
    -c server -- kill -USR2 1
```

## Latency injection

This service has an `EXTRA_LATENCY` environment variable. This will inject a sleep for the specified [time.Duration](https://golang.org/pkg/time/#ParseDuration) on every call to
to the server.

For example, use `EXTRA_LATENCY="5.5s"` to sleep for 5.5 seconds on every request.
