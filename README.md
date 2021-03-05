# mongo-documentdb-compat

This is a Go port of the [Amazon DocumentDB compatibility tool](https://github.com/awslabs/amazon-documentdb-tools/tree/master/compat-tool).

It exposes a single method: `CheckKeys`, which can be used to ensure that your operations don't use any of [DocumentDB's unsupported keys](https://docs.aws.amazon.com/documentdb/latest/developerguide/mongo-apis.html).

Like the Python tool, it does not currently check for any of the [functional differences](https://docs.aws.amazon.com/documentdb/latest/developerguide/functional-differences.html).
