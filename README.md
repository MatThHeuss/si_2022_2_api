## Api sistemas de informação.

### Aplicativo de doações

Para executar o projeto execute os comandos abaixo:

```bash
    docker-compose up -d
```

### Endpoints disponiveis


- Criação de usuário
POST "http://localhost:8080/users"

- Buscar um usuário pelo email
GET "http://localhost:8080/users"


- Criar um anuncio
POST "http://localhost:8080/announcements

- Listar todos os anuncios
GET "http://localhost:8080/announcements

- Listar anuncio por um ID 
GET "http://localhost:8080/announcements/{id}"



