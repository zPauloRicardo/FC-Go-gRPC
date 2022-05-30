# gRPC

É um framework de comunicação do google com suporte multiplataforma e multilinguagem.

Tem como vantagens o uso do HTTP/2 e trafego de dados binarios, tornando ele mais leve e rapido que o REST.

Este repositorio contem exemplos de implementação unidirecional e bidirecional nos modelos:
 
- API Unary (Modelo Request/Response Unica)
- API Server Streaming (Uma request e varias responses)
- API Client Streaming (Varias requests e uma response)
- API Bidirecional Streaming (Varias requests e varias responses)

Diferente do REST e do HTTP/1.1 utilizamos uma unica conexão para os casos que temos o envio de mais de
uma requisição, o que traz mais velocidade no trafego de dados.

Em um resumo, varias requests ou varias responses em uma unica conexão como servidor, menos carga e
mais desempenho.
