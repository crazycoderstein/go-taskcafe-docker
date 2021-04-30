import React, { useState, useEffect } from 'react';
import { createBrowserHistory } from 'history';
import { Router } from 'react-router';
import { PopupProvider } from 'shared/components/PopupMenu';
import styled, { ThemeProvider } from 'styled-components';
import NormalizeStyles from './NormalizeStyles';
import BaseStyles from './BaseStyles';
import theme from './ThemeStyles';
import Routes from './Routes';
import ToastedContainer from './Toast';
import { UserContext } from './context';

import 'react-toastify/dist/ReactToastify.css';
import './fonts.css';

const history = createBrowserHistory();

const App = () => {
  const [user, setUser] = useState<string | null>(null);

  return (
    <>
      <UserContext.Provider value={{ user, setUser }}>
        <ThemeProvider theme={theme}>
          <NormalizeStyles />
          <BaseStyles />
          <Router history={history}>
            <PopupProvider>
              <Routes history={history} />
            </PopupProvider>
          </Router>
          <ToastedContainer
            position="bottom-right"
            autoClose={5000}
            hideProgressBar
            newestOnTop
            closeOnClick
            rtl={false}
            pauseOnFocusLoss
            draggable
            pauseOnHover
            limit={5}
          />
        </ThemeProvider>
      </UserContext.Provider>
    </>
  );
};

export default App;
