---
AWSTemplateFormatVersion: '2010-09-09'
Resources:
  {{ range .queues }}
  MyQueue:
    Type: AWS::SQS::Queue
    Properties:
      QueueName: {{ .QueueName }}
  {{ end }}
