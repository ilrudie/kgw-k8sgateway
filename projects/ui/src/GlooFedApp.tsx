import React from 'react';
import { Global } from '@emotion/core';
import styled from '@emotion/styled';
import { ThemeProvider, useTheme } from '@emotion/react';
import { globalStyles } from './Styles/globalStyles';
import { Footer } from 'Components/SiteStructure/Footer';
import { Content } from 'Components/SiteStructure/Content';
import { MainMenu } from 'Components/SiteStructure/MainMenu';
import { BrowserRouter } from 'react-router-dom';
import './Styles/styles.css';
import { ConfirmModalProvider } from 'Components/Context/ConfirmModalContext';
import { Toaster } from 'react-hot-toast';
import { AppSettingsProvider } from 'Components/Context/AppSettingsContext';

const AppContainer = styled.div`
  display: grid;
  height: 100vh;
  grid-template-rows: 55px 1fr 62px;
`;

function GlooFedApp() {
  const theme = useTheme();
  return (
    <ThemeProvider theme={theme}>
      <AppSettingsProvider>
        <ConfirmModalProvider>
          <Toaster
            position='bottom-right'
            reverseOrder={false}
            toastOptions={{
              duration: 8000,
              style: { borderRadius: '2px' },
            }}
          />
          <BrowserRouter>
            <Global styles={globalStyles} />
            <AppContainer>
              <MainMenu />
              <Content />
              <Footer />
            </AppContainer>
          </BrowserRouter>
        </ConfirmModalProvider>
      </AppSettingsProvider>
    </ThemeProvider>
  );
}

export default GlooFedApp;