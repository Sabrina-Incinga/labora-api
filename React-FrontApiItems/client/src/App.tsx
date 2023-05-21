import { Routes, Route } from "react-router-dom";
import './App.css'
import ItemsComponent from './features/item/items.component';


function App() {

  return (
    <div className="App">
      <Routes>
        <Route path="/" element={<ItemsComponent />} />
      </Routes>
    </div>
  )
}

export default App
