import React, { useState } from 'react';
import HomeScreen from './HomeScreen.jsx';
import AssemblyInstallerScreen from './AssemblyInstallerScreen.jsx';
import IndividualModInstallerScreen from './IndividualModInstallerScreen.jsx';

export default function App() {
  const [screen, setScreen] = useState('home');

  return (
    <>
      {screen === 'home' && (
        <HomeScreen
          onInstallAssembly={() => setScreen('assembly')}
          onInstallMods={() => setScreen('individual')}
        />
      )}
      {screen === 'assembly' && (
        <AssemblyInstallerScreen
          onBack={() => setScreen('home')}
        />
      )}
      {screen === 'individual' && (
        <IndividualModInstallerScreen
          onBack={() => setScreen('home')}
        />
      )}
    </>
  );
}
