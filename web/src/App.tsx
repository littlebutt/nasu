import React, {useState} from 'react';
import './App.css';
import SideBar from "./components/sidebar";


function App() {
    const [active, setActive] = useState('overview')
  return (
    <div>
        <SideBar setActive={setActive}/>
    </div>
  );
}

export default App;
