import React, { useEffect, useState } from 'react';
import { BackButton, InputRow, DownloadButton, StatusMsg, LoadingOverlay } from './components.jsx';
import * as api from './wailsApi.js';

export default function AssemblyInstallerScreen({ onBack, onQRNeeded }) {
  const [assemblies, setAssemblies] = useState([]);
  const [selectedIdx, setSelectedIdx] = useState(null);
  const [steamFolder, setSteamFolder] = useState('');
  const [steamConfirmed, setSteamConfirmed] = useState(false);
  const [loading, setLoading] = useState(false);
  const [loadingText, setLoadingText] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  useEffect(() => {
    api.GetAssemblies().then(setAssemblies);
    api.GetDefaultSteamFolder().then(setSteamFolder);
  }, []);

  async function handleBrowse() {
    try {
      const folder = await api.BrowseForSteamFolder();
      if (folder) setSteamFolder(folder);
    } catch {}
  }

  async function handleDownload() {
    setError('');
    setSuccess('');
    if (selectedIdx === null) { setError('Please select an assembly to install.'); return; }
    if (!steamConfirmed) { setError('Please confirm your Steam folder first.'); return; }

    setLoading(true);
    setLoadingText('Downloading — Steam QR login will appear in terminal...');
    try {
      await api.InstallCompleteAssembly(selectedIdx, steamFolder);
      setSuccess('Assembly installed successfully! Folder opened.');
    } catch (e) {
      setError(String(e));
    } finally {
      setLoading(false);
      setLoadingText('');
    }
  }

  return (
    <div className="screen">
      {loading && <LoadingOverlay text={loadingText} />}

      <div className="header-row">
        <BackButton onClick={onBack} />
      </div>

      {/* Steam folder */}
      <div className="input-row">
        <input
          className="input-field"
          value={steamFolder}
          onChange={e => { setSteamFolder(e.target.value); setSteamConfirmed(false); }}
          placeholder="Steam folder path"
        />
        <div className="confirm-row" style={{ gap: 8, display: 'flex' }}>
          <button className="btn-confirm" style={{ background: '#7ec8e8', marginRight: 'auto' }}
            onClick={handleBrowse}>
            BROWSE
          </button>
          <button className="btn-confirm" onClick={() => setSteamConfirmed(true)}>
            CONFIRM
          </button>
        </div>
        {steamConfirmed && <div className="status-msg success">✓ Steam folder confirmed</div>}
      </div>

      {/* Assembly list */}
      <div className="list-container" style={{ scrollbarWidth: 'none' }}>
        {assemblies.map((asm, i) => (
          <button
            key={i}
            className={`list-item${selectedIdx === i ? ' selected' : ''}`}
            onClick={() => setSelectedIdx(i)}
          >
            {asm.name}
          </button>
        ))}
        {assemblies.length === 0 && (
          <div style={{ padding: '16px', color: '#aaa', fontSize: '0.85rem' }}>
            Loading assemblies…
          </div>
        )}
      </div>

      <StatusMsg type="error">{error}</StatusMsg>
      <StatusMsg type="success">{success}</StatusMsg>

      <div style={{ display: 'flex', justifyContent: 'flex-end' }}>
        <DownloadButton onClick={handleDownload} loading={loading} />
      </div>
    </div>
  );
}
