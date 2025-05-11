import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter, Route, Routes} from "react-router";
import {UserProvider} from "./shared/hooks/useUser";
import Navbar from "./features/navbar/Navbar";
import MainPage from "./features/main-page/MainPage";
import CalculatePage from "./features/calculate-page/CalculatePage";
import ExpressionPage from "./features/expression-page/ExpressionPage";
import ExpressionsPage from "./features/expressions-page/ExpressionsPage";
import LoginPage from "./features/login/LoginPage";
import './index.css';
import SignupPage from "./features/signup/SignupPage";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <UserProvider>
    <BrowserRouter>
      <Navbar/>
      <Routes>
        <Route index element={<MainPage/>}/>
        <Route path="/calculate" element={<CalculatePage/>}/>
        <Route path="/expressions/:id" element={<ExpressionPage/>}/>
        <Route path="/expressions" element={<ExpressionsPage/>}/>
        <Route path="/login" element={<LoginPage/>}/>
        <Route path="/signup" element={<SignupPage/>}/>
      </Routes>
    </BrowserRouter>
  </UserProvider>
);
