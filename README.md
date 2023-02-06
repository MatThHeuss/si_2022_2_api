## Api sistemas de informação.

### Aplicativo de doações

Para executar o projeto execute os comandos abaixo:

```bash
    docker-compose up -d
```

### Endpoints disponiveis

````
- Criação de usuário
POST "http://localhost:8080/users"
curl --location --request POST 'http://localhost:8080/users' \
--form 'name=""' \
--form 'email=""' \
--form 'password=""' \
--form 'profile_image=@"path/to/image"'


- Buscar um usuário pelo email
GET "http://localhost:8080/users"

curl --location --request GET 'http://localhost:8080/users' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":""
}'



- Criar um anuncio
POST "http://localhost:8080/announcements
curl --location --request POST 'http://localhost:8080/announcements' \
--form 'name=" "' \
--form 'description=" "' \
--form 'category=" "' \
--form 'address=" "' \
--form 'postal_code=" "' \
--form 'user_id=""' \
--form 'image_1=@"/path/to/image"' \
--form 'image_2=@"path/to/image"'



- Listar todos os anuncios
GET "http://localhost:8080/announcements
curl --location --request GET 'http://localhost:8080/announcements'


- Listar anuncio por um ID 
GET "http://localhost:8080/announcements/{id}"
curl --location --request GET 'http://localhost:8080/announcements/id'




