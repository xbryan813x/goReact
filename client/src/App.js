import { useEffect } from 'react';
import { connect, sendMsg } from "./api";

function App() {

  useEffect(() => {
    connect();
  }, []);
  
  const send = () => {
    console.log("hello");
    sendMsg("hello");
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <p>
          <button onClick={send}>Hit</button>
        </p>
      </header>
    </div>
  );
}

export default App;
