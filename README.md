# Url Shortener
Url shortener service built in Go with DynamoDB support.

Effectively a simple storage and retrieval service, where you can encode some text and are returned a hash to retrieve it later.

Has two endpoints:
- /e which encodes the full link passed in (i.e. shortlink.com/e?q=https://google.com)
- /d which decodes the shortlink passed in (i.e. shortlink.com/d?q=shortLink)

For setup, checks for the following environment variables:
- `DYNAMO_ENDPOINT`: The endpoint to your dynamo db table
- `DYNAMO_REGION`: The region of your dynamo db table
- `DYNAMO_TABLE_NAME`: The name of your dynamo db table

Or, if those don't exist, just checks for a `.config` file at the root, which expects the values for those variables in that order. For example, the file might look like:
```
https://my-dynamo-endpoint.com/
us-west-2
my-table-name
```
The port can be configured with the environment variable `PORT`, otherwise defaults to 5000.

The service just stores the mappings from short link to the full link in the table.
