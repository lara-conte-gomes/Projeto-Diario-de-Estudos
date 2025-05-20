## Explicação código App.js

import React, {useEffect, useState} from 'react'; -> 

useEffect -> busca dados na API criada
useState -> cria estados no código

## Definindo o componente principal
function App(){

   //Armazena os dados vindos da API da Go, criado como um array vazio 
  const [estudos, setEstudos] = useState([]);

  useEffect(() => { -> executa essa função somente uma vez
    fetch('http://localhost:8080/estudos') -> faz uma requisição para a API
      .then(response => response.json()) -> transforma a API em JSON
      .then(data => setEstudos(data)) -> atualiza o estado com os dados
      .catch(error => console.error("Erro ao buscar estudos:", error)) -> captura erros
  }, []);

padding -> espaçamento interno
minHeight -> altura mínima da janela seja 100% (100vh)

usar map para percorrer os dados do array

{estudos.materia && <p><strong>Matéria:</strong> {estudos.materia}</p>} 
condição && conteúdo

Só vai ser mostrado estudos.materia se ele existir, se a condição for falsa não será mostrado nada



