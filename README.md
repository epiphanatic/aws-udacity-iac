# CD12352 - Infrastructure as Code Project Solution

# Jason Simpson

## Spin up instructions

1. Ensure you have the aws cli installed on your local machine and configured with your credentials

2. From the root of the project cd into the network directory and run the following command to create the network stack:

`./run.sh deploy us-east-1 UdagramNetwork network.yml network-parameters.json`

3. After the network stack has been created, a bucket will need to be created first, then the local index.html file will need to be copied over to the bucket, then the application stack needs to be updated. To do this, cd into the application directory in the root of the project and run the following command to create the application stack:

```
./run.sh deploy us-east-1 UdagramApplication udagram-bucket.yml udagram-bucket-parameters.json &&
aws s3 cp index.html s3://udagram-s3-bucket-2635189-8570238947520 &&
./run.sh deploy us-east-1 UdagramApplication udagram.yml udagram-parameters.json
```

## Tear down instructions

1. First delete the index.html file in the S3 bucket created by the application stack.

2. Then, from the root of the project cd into the application directory and run the following command to delete the application stack:

`./run.sh delete us-east-1 UdagramApplication`

3. After the application stack has been deleted, cd into the network directory back in the root of the project and run the following command to delete the network stack:

`./run.sh delete us-east-1 UdagramNetwork`
