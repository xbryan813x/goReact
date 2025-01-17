import { useEffect, useState } from 'react';
import { connect, sendMsg } from "./api";

function App() {
  const [name, setName] = useState("")
  const [age, setAge] = useState("")
  const [note, setNote] = useState("")
  const [beResponse, setBeResponse] = useState({})

  useEffect(() => {
    connect();
  }, []);

  const saveNote = (config) => {
    fetch('http://localhost:8080/', {
      method: 'POST', // or 'PUT'
      body: JSON.stringify(config),
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      setBeResponse(data)
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  
  const send = () => {
    const configData = {
      name,
      age,
      note
    }
    sendMsg(`hello`);
    saveNote(configData);
  };

  const buttonNames = ['add', 'remove', 'update', 'delete'];

  const logInput = (element, keyName) => {
    console.log(keyName);
    switch(keyName) {
      case 'name':
        setName(element.target.value)
        break;
      case 'age':
        setAge(element.target.value)
        break;
      case 'note':
        setNote(element.target.value)
        break;
        default:
        setNote(element.target.value)
    }
  };

  return (
    <div className="App">
      <header className="App-header">
      </header>
      <h2>
        Enter two numbers in the input field and then select type of computation.
      </h2>
      <input type="text" placeholder="name" onInput={(s) => logInput(s, 'name')}/>
      <input type="tel" placeholder="age" onInput={(s) => logInput(s, 'age')} />
      <input type="text" placeholder="Note" onInput={(s) => logInput(s, 'note')} />

      <div style={{ 
          display: 'flex',
          justifyContent: 'space-around'}}>
        {buttonNames.map((title) => 
          (
            <p key={title}>
              <button onClick={send}>{title}</button>
            </p>
          )
        )}
        <p>
          {JSON.stringify(beResponse)}
        </p>
      </div>
    </div>
  );
}

export default App;
