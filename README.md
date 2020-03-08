# Url Shortener
Url shortener service built in Go with DynamoDB support.

Has three endpoints:
- /e/ which encodes the full link passed in (i.e. shortlink.com/e/https://google.com)
- /d/ which decodes the shortlink passed in
- /r/ which decodes and redirects the user to the full link

For setup, checks for the following environment variables:
- DYNAMO_ENDPOINT: The endpoint to your dynamo db table
- DYNAMO_REGION: The region of your dynamo db table
- DYNAMO_TABLE_NAME: The name of your dynamo db table

Or, if those don't exist, just checks for a `.config` file at the root, which expects the values for those variables in that order. For example, the file might look like:
```
https://my-dynamo-endpoint.com/
us-west-2
my-table-name
```

The service just stores the mappings from short link to the full link in the table.
