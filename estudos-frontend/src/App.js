import React, {useEffect, useState} from 'react';

function App() {
  const [estudos, setEstudos] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/estudos')
      .then(response => response.json())
      .then(data => setEstudos(data))
      .catch(error => console.error("Erro ao buscar estudos:", error))
  }, []);

  return (
    <div style={{
      padding: '2rem', 
      fontFamily: 'Segoe UI, sans-serif',
      backgroundColor: 'lightblue',
      minHeight: '100vh',
      }}>

      <h1 style={{textAlign: 'center', marginBottom: '2rem'}}>Diário de Estudos</h1>

      <div style={{
        display: 'flex',
        flexDirection: 'column',
        gap: '1rem',
        maxWidth: '600px',
        margin: '0 auto'
      }}>
          {estudos.map((estudos, index) => (
            <div key={index} style={{
              backgroundColor: '#fff',
              padding: '1.5rem',
              borderRadius: '10px',
              boxShadow: '0 4px 12px rgba(0, 0, 0, 0.1)',
              lineHeight: '1.6'
            }}>

              <p><strong>Tema:</strong> {estudos.tema}</p>
              <p><strong>Horas:</strong> {estudos.horas}</p>
              {estudos.materia && <p><strong>Matéria:</strong> {estudos.materia}</p>}
            </div>
        ))}
      </div>
    </div>
  );
}

export default App;
