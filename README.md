# LambdaC

Run code in response to events.

**Warning**: currently in alpha stage.

Inspired by [force12](https://github.com/force12io/force12), [lambda-docker](https://github.com/tobegit3hub/lambda-docker) and [AWS Lambda](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html).

Create **runtimes** by specifying an image and container engine (Docker is the first for now, but not limited to it), then create **functions** to be called when events arrive.

No scaling is needed. Function **instances** are created, started and stopped on demand.
