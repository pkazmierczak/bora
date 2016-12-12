---
AWSTemplateFormatVersion: '2010-09-09'
Resources:
  {{ range .queues }}
  MyQueue:
    Type: AWS::SQS::Queue
    Properties:
      QueueName: {{ .QueueName }}
  {{ end }}
  S3Bucket:
    Type: AWS::S3::Bucket
    Properties:
      AccessControl: PublicRead
      WebsiteConfiguration:
        IndexDocument: {{ .indexdocument }}
        ErrorDocument: error.html
    DeletionPolicy: Retain

Outputs:
  WebsiteURL:
    Value: !GetAtt S3Bucket.WebsiteURL
    Description: URL for the website hosted on S3
  S3BucketSecureURL:
    Value: !Sub
        - https://${Domain}
        - Domain: !GetAtt S3Bucket.DomainName
    Description: Name of the S3 bucket to hold website content
