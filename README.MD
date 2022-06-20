
```bash
docker exec -it my_postgres psql -U postgres
```

```psql
create database ks_aggregator_rates TEMPLATE template0;
drop database ks_aggregator_rates;
```

```bash
brew install golang-migrate
export DB_URL=postgres://postgres:my_password@localhost:54320/ks_aggregator_rates?sslmode=disable
migrate -path db/migrations -database "$DB_URL" -verbose up
```

```
psql postgres://postgres:my_password@localhost:54320/ks_aggregator_rates
```