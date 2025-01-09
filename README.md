# Stats service

Service for gathering game statistics and serving them over GraphQL.

## Configuration

See `defaults.env` for default configuration values.
Environment variables will override the defaults.

Explanation of configuration values:
- `NATS_URL`: URL of the NATS server
- `LOG_JSON`: If set to `true`, logs will be in JSON format
- `REDIS_URL`: URL of the Redis server
- `GRAPHQL_PORT`: Port on which the GraphQL server will listen

## API

GraphQL schema is defined in [schema.graphqls](./graph/schema.graphqls).
GraphQL playground is available on the `/api/stats/public/playground` endpoint.