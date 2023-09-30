import React, { useEffect, useState } from 'react';
import './App.css';
function App() {
  const [Words, setSentence] = useState('');
  const [limit, setLimit] = useState('');


  const handleSubmit = async (e) => {
    e.preventDefault();
    const response = await fetch(import.meta.env.VITE_API + '/sentences', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ Words })
    })
    const data = await response.json()
    setLimit(data.Ideal_limit);
    console.log(data);
    console.log(limit);
  }
  return (
    <div>
      <div className="centered-text">
        <h1 className='header-text' >Calculadora del límite ideal</h1>
      </div>

      <form  className ="form" onSubmit={handleSubmit}>
        <label className='label'>Ingrese el texto:</label>
        <textarea className='input' type="sentence" placeholder="hola mundo" onChange={(e) => setSentence(e.target.value)} />
        <button className='button'>Calcular</button>
      </form>
      <div className='limit'>
        <h2>El límite ideal es:</h2>
        <p className='pa'> {limit}</p>
      </div>

    </div>

  )

}

export default App;