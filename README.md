# Autocross Timer

## About

This project is to allow autocross events to have a live timing leaderboard. It is designed to be easily set up, with automated infrastructure deployment to AWS, or a Dockerfile for anywhere you feel like running the containers.

## To Setup

There are several variables that need to be set before you can run this in your own environment. As well an AWS CLI account with sufficient access to build the terraform resources, and SSM access to allow SSH must be provisioned if using the automated build and deploy scripts.

### Github actions secrets

- `AWS_ACCESS_KEY_ID`: AWS CLI Access Key
- `AWS_SECRET_ACCESS_KEY`: AWS CLI Secret Access Key
- `DB_PASSWORD`: The password that will be used for the API account to access the sql database
- `SSH_PRIVATE_KEY`: The SSH private key that will be used to authenticate to the docker host

### Github actions variables

- `BASE_URL`: The base URL of the website. Used by the Go backend for CORS, and the `init-letsencrypt.sh` script for the SSL certificate.
- `VITE_API_URL`: The API url. By default it will be `BASE_URL/api`, but if you've changed the API routing it may be different.

### Terraform config

- `main.tf`
  - `region`: The region you will be deploying into
  - `subnet`: The primary subnet you will be deploying into
  - `backend`: The terraform backend you will be using

## To Deploy

It is very simple to deploy this project. Configure all of the variables above, and then run the three Github Actions scripts in order.

- `terraform.yml`: This script deploys all of the infrastructure
- `setup-host.yml`: This script configures the host to run Docker
- `deploy-app.yml`: This script deploys the app onto the Docker host and builds/runs it

## TODO

- Add Docker setup into Terraform config
- Implement updateRun in the API backend
- Add validation to API
- Add "get new runs" functionality to front end
  - Store runs in local storage
- Add Docker health checks to web and API
- Make driver view card better
  - Make best times only check valid runs
- Make programatic run simulation
