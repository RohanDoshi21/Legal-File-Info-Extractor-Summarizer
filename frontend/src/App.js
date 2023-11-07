import React, { useState } from 'react';
import Login from './login.js';
import AdminPage from './AdminDashboard.js';
import UserPage from './UserDashboard.js';
import Navbar from './Navbar.js';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import AdminViewInfo from './AdminViewInfo';

function App() {
  const [jwtToken, setJwtToken] = useState(localStorage.getItem('jwtToken'));

  // Function to update the jwtToken when the user logs in
  const updateJwtToken = (token) => {
    setJwtToken(token);
    localStorage.setItem('jwtToken', token);
  };

  return (
    <div className="App">
      <BrowserRouter>
        {/* Pass the jwtToken as a prop */}
        <Routes>
          <Route path="/" element={<Login  />} /> {/* Pass the updateJwtToken function to Login */}
          <Route path="/AdminViewInfo" element={<AdminViewInfo />} />
          <Route path="/AdminPage" element={<AdminPage />} />
          <Route path="/UserPage" element={<UserPage />} />
          <Route path="*" element={<Navigate to="/" />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
