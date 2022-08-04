import React from 'react';
import ReactDOM from 'react-dom';
import {
  BrowserRouter,
  Routes,
  Route,
} from 'react-router-dom';
import './index.css';
import App from './App';
import { GlobalStateProvider } from './GlobalState';

ReactDOM.render(
  // Uncomment when Ant Design v5 is released
  // <React.StrictMode>
  <GlobalStateProvider>
    <BrowserRouter>
      <Routes>
        <Route
          path="/"
          element={<App />}
        />
        <Route
          path="/:project"
          element={<App />}
        />
      </Routes>
    </BrowserRouter>
    </GlobalStateProvider>,
  // </React.StrictMode>,
  document.getElementById('root')
);
