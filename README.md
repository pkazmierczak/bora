# bora—a simple AWS Cloudformation wrapper
[![Build Status](https://travis-ci.org/pkazmierczak/bora.svg)](https://travis-ci.org/pkazmierczak/bora)

## Installation
bora requires AWS SDK for Go as well as some other open source packages. To install all its dependencies, run `go get ./...` in the project directory. 

## Usage
bora requires:

- AWS credentials, you can read about how to set these up [here](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs)
- and a `config.yml` file, like this one:
```yml
region: "eu-central-1"
```

