# bookstore-outh-api

## Creacion Cassandra Table

- Crear Keyspace
````
CREATE KEYSPACE IF NOT EXISTS oauth WITH replication = {'class':'SimpleStrategy', 'replication_factor':1};
````

- Usar Keyspace
````
use oauth;
````

- Crear Tabla
````
create table access_token
(
    access_token varchar primary key,
    user_id bigint,
    client_id bigint,
    expires bigint
);
````