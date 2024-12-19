import { Route, Routes } from "react-router-dom";
import Layout from "./layout";
import HomePage from "./pages/home/page";
import LoginPage from "./pages/login/page";
import MyAccount from "./pages/my-account/page";
import ChangePasswordPage from "./pages/my-info/change-password/page";
import EditInformationPage from "./pages/my-info/edit/page";
import MyInfo from "./pages/my-info/page";
import NotFoundPage from "./pages/not-found";
import RegisterPage from "./pages/register/page";
import DetailTopUpPage from "./pages/top-up/detail/page";
import TopUpPage from "./pages/top-up/page";
import TransactionDetailPage from "./pages/transaction/detail/page";
import TransactionPage from "./pages/transaction/page";
import TransferPage from "./pages/transfer/page";

export default function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/my-info" element={<MyInfo />} />
        <Route path="/my-info/edit" element={<EditInformationPage />} />
        <Route
          path="/my-info/change-password"
          element={<ChangePasswordPage />}
        />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/my-account" element={<MyAccount />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/transfer" element={<TransferPage />} />
        <Route path="/top-up" element={<TopUpPage />} />
        <Route path="/top-up/detail" element={<DetailTopUpPage />} />
        <Route path="/transaction" element={<TransactionPage />} />
        <Route path="/transaction/:id" element={<TransactionDetailPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </Layout>
  );
}
