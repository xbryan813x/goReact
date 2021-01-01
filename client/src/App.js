import { useEffect, useState } from 'react';
import { connect, sendMsg } from "./api";

function App() {
  const [name, setName] = useState("")
  const [age, setAge] = useState("")
  const [note, setNote] = useState("")

  useEffect(() => {
    connect();
  }, []);

  const saveNote = (config) => {
    fetch('http://localhost:8080/', {
      method: 'POST', // or 'PUT'
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(config),
    })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
    })
    .catch((error) => {
      console.error('Error:', error);
    });
  };
  
  const send = () => {
    const configData = {
      name: name,
      age: age,
      note: note
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
        default:
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
      </div>
    </div>
  );
}

export default App;
