---
AWSTemplateFormatVersion: '2010-09-09'
Resources:
  {{ range .queues }}
  MyQueue:
    Type: AWS::SQS::Queue
    Name: {{ .QueueName }}
    Properties:
      VisibilityTimeout: value
  {{ end }}
