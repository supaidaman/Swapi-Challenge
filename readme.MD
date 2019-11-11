FORMATOS DAS URLS

   /api/v1/planets -> Lista Planetas(GET)
	/api/v1/planets/name/{name} -> Pega planeta por nome(GET)
	/api/v1/planets/{id}-> Pega planeta por ID(GET)
	/api/v1/planets -> Insere novo planeta (POST)
	/api/v1/planets/{id} -> Atualiza planeta (POST)
	/api/v1/planets/{id} -> Deleta planeta("DELETE")


    Todas as dependencias estão no repositório já.

    Exemplo de json no postman para inserir novo planeta: 
    { "name":"Hoth",
     "climate":"test",
     "terrain":"test"  
    }
