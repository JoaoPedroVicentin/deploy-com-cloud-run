
# Desafio GoExpert - Deploy com Cloud Run

O desafio inclui implementar um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin). Esse sistema deverá ser publicado no Google Cloud Run.

## 📦 Como rodar o projeto

### 2. Rodando localmente

### 2.1. Clone o repositório

```bash
git clone https://github.com/JoaoPedroVicentin/deploy-com-cloud-run.git
cd deploy-com-cloud-run
```

### 2.2. Crie o arquivo .env com a sua API_KEY da WeatherAPI com base no .env.example

### 2.3. Rode o programa
```bash
cd cmd/server
go run main.go
```

### 3. Acessando a API via Cloud Run acesse: https://deploy-com-cloud-run-44570442432.us-central1.run.app/{cep}

<div align="center">
<h3>👨‍💻</h3>
    <h3> Criado por João Pedro Vicentin!</h3>
    <div>
        <h3>
            <a href="https://www.linkedin.com/in/joaopedrovicentin/" target="_blank">Linkedin</a>
            <a href='https://github.com/JoaoPedroVicentin' target='_blank'>Github</a>
            <a href="https://contate.me/joao-pedro-lopes-vicentin" target="_blank">Whatsapp</a>
        </h3>
    </div>
</div>
