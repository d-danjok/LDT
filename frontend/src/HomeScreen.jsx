import React from 'react';
import { LDTLogo, BackButton } from './components.jsx';

export default function HomeScreen({ onInstallAssembly, onInstallMods }) {
  return (
    <div className="screen">
      <div className="header-row">
        <div className="logo-area">
          <LDTLogo />
        </div>
      </div>

      <button className="menu-card" onClick={onInstallAssembly}>
        <span className="menu-card-title">INSTALL COMPLETE ASSEMBLY</span>
      </button>

      <button className="menu-card" onClick={onInstallMods}>
        <span className="menu-card-title">INSTALL INDIVIDUAL MODS</span>
      </button>
    </div>
  );
}
