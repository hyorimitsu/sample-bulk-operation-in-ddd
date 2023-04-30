FROM flyway/flyway:9.17.0-alpine

COPY ./local/conf /flyway/conf
COPY ./local/seeds /flyway/sql/local
COPY ./common/migrations /flyway/sql/common-migrations
COPY ./common/seeds /flyway/sql/common-seeds
