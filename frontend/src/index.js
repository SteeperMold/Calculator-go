import React from 'react';
import ReactDOM from 'react-dom/client';
import {BrowserRouter, Route, Routes} from "react-router";
import MainPage from "./features/main-page/MainPage";
import CalculatePage from "./features/calculate-page/CalculatePage";
import ExpressionPage from "./features/expression-page/ExpressionPage";
import ExpressionsPage from "./features/expressions-page/ExpressionsPage";
import './index.css';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <BrowserRouter>
    <Routes>
      <Route index element={<MainPage/>}/>
      <Route path="/calculate" element={<CalculatePage/>}/>
      <Route path="/expressions/:id" element={<ExpressionPage/>}/>
      <Route path="/expressions" element={<ExpressionsPage/>}/>
    </Routes>
  </BrowserRouter>
);
