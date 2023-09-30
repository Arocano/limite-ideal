import React, { useEffect, useState } from 'react';

function App() {
  const [Words, setSentence] = useState('');
  const [limit, setLimit] = useState('');


  const handleSubmit = async (e) => {
    e.preventDefault();
    const  response = await fetch(import.meta.env.VITE_API +'/sentences',{
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({Words})
    })
    const data = await response.json()
    setLimit(data.Ideal_limit);
    console.log(data);
    console.log(limit);
  }
  return (
    <div>
<form onSubmit={handleSubmit}>
  <input type="sentence" placeholder="Ingresa la oración" onChange={(e)=>setSentence(e.target.value)}/>
  <button>Calcular</button>
</form>
     <div>

        <h1>El limite ideal es: {limit}</h1>
     </div>
    
    </div>
     
     )

}

export default App;