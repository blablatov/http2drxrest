## sap2gRPCdirx
## On Russian

Интеграция с Directum RX, через REST, OData. Представлены тестовые модули для обмена данными.

Схема обмена данными:

> REST(imitation from web browser) <---> web-server_URL-parser/http2drxrest <---> REST_OData/restpost

Для проверки, запустить сервер http2drxrest, из строки браузера создать запрос типа:

    http://localhost:8080/Doc:Простой документ



## On English

Integration with Directum RX, via REST, OData. 

Sheme exchange of data:

> REST(imitation from web browser) <---> web-server_URL-parser/http2drxrest <---> REST_OData/restpost

To check, start the http2drxrest server and create a test request from the browser line: 

    http://localhost:8080/Doc:Simple file
