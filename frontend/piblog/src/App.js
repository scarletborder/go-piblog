
// src/App.js
import logo from './logo.svg';
import './App.css';
import Markdown from 'react-markdown'
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import PostPage from './pages/PostPage';
import ArchivesPage from "./pages/ArchivesPage";

const App = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path='archives' element={<ArchivesPage />} />
        <Route path="/post/:id" element={<PostPage />} />
      </Routes>
    </Router>
  );
};

export default App;
