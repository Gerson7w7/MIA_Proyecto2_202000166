import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Inicio from "./components/Inicio";
import Rep from "./components/Rep";
import User from "./components/User";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Inicio />} />
        <Route path="/login" element={<User />} />
        <Route path="/reportes" element={<Rep />} />
      </Routes>
    </Router>
  );
}

export default App;
