import React, { useEffect } from "react";
import "./App.css";
import { connect, sendMsg } from "./services/messages";

function App() {

  useEffect(() => {
    connect((msg) => {
      console.log("received " + msg.data)
    })
  })

  return (
    <div className="App">
      <button
        onClick={() => {
          sendMsg("smth");
        }}
      >
        Hits
      </button>
    </div>
  );
}

export default App;
