# Consume Stock Exchange ䷧

O propósito deste repositório é fornecer uma solução para consumir dados em tempo real provenientes do streaming do NATS e armazená-los em um banco de dados. O projeto consiste em um conjunto de componentes e funcionalidades que permitem receber as mensagens do NATS e realizar o processo de persistência no banco de dados.
Com essa solução, é possível capturar e armazenar os dados em um formato adequado para análises futuras, processamento de dados ou qualquer outra finalidade necessária.

## Iniciando ▶️

### Pré-requesitos

É necessário efetuar a instalação do gerenciador de container:

- Docker: https://www.docker.com/

### Instalação

1. Clone the repositório: https://github.com/rafaelsanzio/go-consume-stock-exchange
2. Crie um arquivo `.env` com base no exemplo `.env.example` fornecido. Se necessário, ajuste os valores conforme sua preferência.
3. É necessário criar uma network no docker, para conexão com outros serviços:
   ```sh
   docker network create app_network
   ```
4. Por fim, para iniciar a aplicação basta rodar os comandos de build e up.
   ```sh
   docker-compose build && docker-compose up
   ```
