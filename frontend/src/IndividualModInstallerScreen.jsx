import React, { useEffect, useState } from 'react';
import { BackButton, InputRow, DownloadButton, StatusMsg, LoadingOverlay } from './components.jsx';
import * as api from './wailsApi.js';

// Sub-step: select LC version (used for both new instance and existing)
function LCVersionSelector({ versions, selectedIdx, onSelect }) {
    return (
        <div className="list-container" style={{ scrollbarWidth: 'none' }}>
            {versions.map((v, i) => (
                <button
                  key={i}
                  className={`list-item${selectedIdx === i ? ' selected' : ''}`}
                  onClick={() => onSelect(i)}
                  style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}
                >
                <span>{v.name.toUpperCase()}</span>
                <span style={{ fontSize: '0.75rem', color: '#666', fontWeight: 'normal' }}>{v.lastDate}</span>
                </button>
            ))}
        </div>
    );
}

export default function IndividualModInstallerScreen({ onBack }) {
  const [versions, setVersions] = useState([]);

  // Step: 'choose-type' | 'new-instance' | 'existing-instance' | 'install-mod'
  const [step, setStep] = useState('choose-type');

  // Shared state
  const [steamFolder, setSteamFolder] = useState('');
  const [steamConfirmed, setSteamConfirmed] = useState(false);
  const [selectedVersionIdx, setSelectedVersionIdx] = useState(null);
  const [versionConfirmed, setVersionConfirmed] = useState(false);
  const [installPath, setInstallPath] = useState('');

  // New instance specific
  const [instanceName, setInstanceName] = useState('');

  // Existing installation specific
  const [existingPath, setExistingPath] = useState('');
  const [hasBepInEx, setHasBepInEx] = useState(false);
  const [bepInExChoice, setBepInExChoice] = useState(null); // 'alongside' | 'erase'

  // Mod install
  const [modLink, setModLink] = useState('');
  const [modLinkConfirmed, setModLinkConfirmed] = useState(false);

  const [loading, setLoading] = useState(false);
  const [loadingText, setLoadingText] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  useEffect(() => {
    api.GetLCVersions().then(setVersions);
    api.GetDefaultSteamFolder().then(setSteamFolder);
  }, []);

  function clearError() { setError(''); setSuccess(''); }

  // ── NEW INSTANCE ──────────────────────────────────────────────

  async function handleNewInstanceSetup() {
    clearError();
    if (!steamConfirmed) { setError('Confirm your Steam folder first.'); return; }
    if (selectedVersionIdx === null) { setError('Select an LC version.'); return; }
    if (!instanceName.trim()) { setError('Enter an instance name.'); return; }

    setLoading(true);
    setLoadingText('Downloading LC version — Steam QR login will appear in terminal…');
    try {
      const path = await api.InstallNewInstance(steamFolder, selectedVersionIdx, instanceName.trim());
      setInstallPath(path);
      setStep('install-mod');
    } catch (e) {
      setError(String(e));
    } finally {
      setLoading(false);
      setLoadingText('');
    }
  }

  // ── EXISTING INSTALLATION ─────────────────────────────────────

  async function handleBrowseExisting() {
    try {
      const p = await api.BrowseForLCInstallation();
      if (p) {
        setExistingPath(p);
        const has = await api.CheckHasBepInEx(p);
        setHasBepInEx(has);
        setBepInExChoice(null);
      }
    } catch (e) { setError(String(e)); }
  }

  async function handleConfirmExisting() {
    clearError();
    if (!existingPath) { setError('Browse to select an existing installation.'); return; }
    if (hasBepInEx && bepInExChoice === null) { setError('Choose how to handle existing mods.'); return; }
    if (selectedVersionIdx === null) { setError('Select the LC version you have.'); return; }

    setLoading(true);
    setLoadingText('Preparing installation…');
    try {
      await api.SetExistingInstallation(existingPath, selectedVersionIdx, bepInExChoice === 'erase');
      setInstallPath(existingPath);
      setStep('install-mod');
    } catch (e) {
      setError(String(e));
    } finally {
      setLoading(false);
      setLoadingText('');
    }
  }

  // ── MOD INSTALL ────────────────────────────────────────────────

  async function handleInstallMod() {
    clearError();
    if (!modLinkConfirmed || !modLink) { setError('Enter and confirm a Thunderstore link.'); return; }

    const valid = await api.IsValidThunderstoreLink(modLink);
    if (!valid) { setError('Invalid Thunderstore link. Expected format:\nhttps://thunderstore.io/c/lethal-company/p/[author]/[mod]/'); return; }

    setLoading(true);
    setLoadingText('Installing mod and dependencies…');
    try {
      await api.InstallMod(modLink, installPath, selectedVersionIdx ?? 0);
      setSuccess('Mod installed! You can install another or close.');
      setModLink('');
      setModLinkConfirmed(false);
    } catch (e) {
      setError(String(e));
    } finally {
      setLoading(false);
      setLoadingText('');
    }
  }

  function handleOpenFolder() {
    api.OpenFolder(installPath);
  }

  // ── RENDER ────────────────────────────────────────────────────

  function handleBack() {
    clearError();
    if (step === 'choose-type') onBack();
    else if (step === 'new-instance' || step === 'existing-instance') setStep('choose-type');
    else if (step === 'install-mod') {
      // Go back to whichever setup we came from
      setStep('choose-type');
      setInstallPath('');
    }
  }

  return (
    <div className="screen">
      {loading && <LoadingOverlay text={loadingText} />}

      <div className="header-row">
        <BackButton onClick={handleBack} />
        {/* SELECT LC VERSION pill shown on version-select sub-steps */}
        {(step === 'new-instance' || step === 'existing-instance') && (
          <button className="pill-btn select-version">SELECT LC VERSION</button>
        )}
      </div>

      {/* ── CHOOSE TYPE ── */}
      {step === 'choose-type' && (
        <>
          <div className="input-row">
            <input
              className="input-field"
              value={steamFolder}
              onChange={e => { setSteamFolder(e.target.value); setSteamConfirmed(false); }}
              placeholder="Steam folder path"
            />
            <div className="confirm-row" style={{ gap: 8, display: 'flex' }}>
              <button className="btn-confirm" style={{ background: '#7ec8e8', marginRight: 'auto' }}
                onClick={async () => {
                  try { const p = await api.BrowseForSteamFolder(); if (p) setSteamFolder(p); } catch {}
                }}>
                BROWSE
              </button>
              <button className="btn-confirm" onClick={() => setSteamConfirmed(true)}>CONFIRM</button>
            </div>
            {steamConfirmed && <div className="status-msg success">✓ Steam folder confirmed</div>}
          </div>

          <button className="menu-card" onClick={() => { clearError(); setStep('new-instance'); }}>
            <span className="menu-card-title">NEW INSTANCE</span>
          </button>

          <div className="menu-card" onClick={() => { clearError(); setStep('existing-instance'); }}>
            <span className="menu-card-title">EXISTING INSTALLATION</span>
          </div>

          <StatusMsg type="error">{error}</StatusMsg>
        </>
      )}

      {/* ── NEW INSTANCE ── */}
      {step === 'new-instance' && (
        <>
          <LCVersionSelector versions={versions} selectedIdx={selectedVersionIdx} onSelect={setSelectedVersionIdx} />

          <div className="input-row">
            <input
              className="input-field"
              value={instanceName}
              onChange={e => setInstanceName(e.target.value)}
              placeholder="Instance name (e.g. MyMods)"
              onKeyDown={e => e.key === 'Enter' && handleNewInstanceSetup()}
            />
            <div className="confirm-row">
              <button className="btn-confirm" onClick={handleNewInstanceSetup}>CONFIRM</button>
            </div>
          </div>

          <StatusMsg type="error">{error}</StatusMsg>
        </>
      )}

      {/* ── EXISTING INSTALLATION ── */}
      {step === 'existing-instance' && (
        <>
          <LCVersionSelector versions={versions} selectedIdx={selectedVersionIdx} onSelect={setSelectedVersionIdx} />

          <div className="existing-card">
            <span className="existing-title">EXISTING INSTALLATION</span>
            <div className="existing-path">{existingPath || 'No folder selected — click Browse'}</div>
            <div className="confirm-row" style={{ gap: 8, display: 'flex' }}>
              <button className="btn-confirm" style={{ background: '#7ec8e8', marginRight: 'auto' }}
                onClick={handleBrowseExisting}>
                BROWSE
              </button>
              <button className="btn-confirm" onClick={handleConfirmExisting}>CONFIRM</button>
            </div>
          </div>

          {hasBepInEx && (
            <div className="section-card">
              <span className="section-label">Existing mods detected — how to proceed?</span>
              <button
                className={`list-item${bepInExChoice === 'alongside' ? ' selected' : ''}`}
                onClick={() => setBepInExChoice('alongside')}
              >
                INSTALL ALONGSIDE
              </button>
              <button
                className={`list-item${bepInExChoice === 'erase' ? ' selected' : ''}`}
                onClick={() => setBepInExChoice('erase')}
                style={{ color: '#c04040' }}
              >
                ERASE EXISTING
              </button>
            </div>
          )}

          <StatusMsg type="error">{error}</StatusMsg>
        </>
      )}

      {/* ── MOD INSTALL ── */}
      {step === 'install-mod' && (
        <>
          {installPath && (
            <div className="existing-card">
              <span className="existing-title">INSTALL PATH</span>
              <div className="existing-path">{installPath}</div>
              <div className="confirm-row">
                <button className="btn-confirm" style={{ background: '#7ec8e8' }} onClick={handleOpenFolder}>
                  OPEN FOLDER
                </button>
              </div>
            </div>
          )}

          <div className="input-row">
            <input
              className="input-field"
              value={modLink}
              onChange={e => { setModLink(e.target.value); setModLinkConfirmed(false); }}
              placeholder="https://thunderstore.io/c/lethal-company/p/..."
              onKeyDown={e => e.key === 'Enter' && setModLinkConfirmed(true)}
            />
            <div className="confirm-row">
              <button className="btn-confirm" onClick={() => setModLinkConfirmed(true)}>CONFIRM</button>
            </div>
            {modLinkConfirmed && modLink && (
              <div className="status-msg info">✓ Link confirmed</div>
            )}
          </div>

          <StatusMsg type="error">{error}</StatusMsg>
          <StatusMsg type="success">{success}</StatusMsg>

          <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
            <DownloadButton onClick={handleInstallMod} loading={loading} />
          </div>
        </>
      )}
    </div>
  );
}
