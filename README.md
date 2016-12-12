# bora—a simple AWS Cloudformation wrapper
[![Build Status](https://travis-ci.org/pkazmierczak/bora.svg)](https://travis-ci.org/pkazmierczak/bora)

## Installation
bora requires AWS SDK for Go as well as some other open source packages
to compile. To install all its dependencies, run `go get ./...` in the
project directory. 

Pre-build binaries for Darwin and Linux are available. 

## Usage
bora requires:

- AWS credentials, you can read about how to set these up [here](http://blogs.aws.amazon.com/security/post/Tx3D6U6WSFGOK2H/A-New-and-Standardized-Way-to-Manage-Credentials-in-the-AWS-SDKs)
- A config file (YAML) and a matching meta-template (YAML with Go's
  `text/template` markup). Examples are available under
  `exampleTemplates/` directory. 

## Roadmap and status
bora is in very early stages of development. Currently it correctly
interprets meta-templates, converts them into CF files, deploys and
terminates stacks. It is not yet able to update stacks correctly.
Upcoming features:

- Stack updates;
- Role assumption/multi-account deployment;
- Automatic lookup of VPCs, AMIs, subnets and SGs by name; 
- Tests...
