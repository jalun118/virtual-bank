import { Route, Routes } from "react-router-dom";
import Navbar from "./components/Navbar";
import HomePage from "./pages/home/page";
import LoginPage from "./pages/login/page";
import RegisterPage from "./pages/register/page";
import TopUpPage from "./pages/top-up/page";

function App() {
  return (
    <div>
      <Navbar />
      <main className="mt-3">
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/top-up" element={<TopUpPage />} />
        </Routes>
      </main>
    </div>
  );
}

export default App;
