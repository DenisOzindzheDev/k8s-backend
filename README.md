build 
```
docker build --tag k8s-backend .  <-- бидим 
```
run 
```
docker run --publish 8080:8080 k8s-backend
```

table 
```
create table queries (
id int NOT NULL GENERATED ALWAYS AS IDENTITY,
queries varchar,
CONSTRAINT requests_pkey PRIMARY KEY (id)
)

```