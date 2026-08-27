// Wails auto-generates bindings at runtime under window.go.main.App
// This shim provides a clean API and a mock fallback for browser dev.

const isBrowser = typeof window !== 'undefined' && !window.go;

const mock = {
  GetAssemblies: async () => [{ name: "v73 Wesley's Basic" }],
  GetLCVersions: async () => [
    { name: "v45", lastDate: "2023-12-09" },
    { name: "v50", lastDate: "2024-06-28" },
    { name: "v56", lastDate: "2024-08-17" },
    { name: "v69", lastDate: "2024-12-13" },
    { name: "v73", lastDate: "2026-03-29" },
  ],
  GetDefaultSteamFolder: async () => "C:\\Program Files (x86)\\Steam",
  BrowseForSteamFolder: async () => "C:\\Program Files (x86)\\Steam",
  BrowseForLCInstallation: async () => "C:\\Program Files (x86)\\Steam\\steamapps\\content\\app_1966720\\v73 MyMods",
  InstallCompleteAssembly: async (idx, folder) => { await delay(2000); },
  InstallNewInstance: async (folder, vIdx, name) => { await delay(2000); return folder + "\\v73 " + name; },
  CheckHasBepInEx: async (path) => false,
  SetExistingInstallation: async (path, vIdx, erase) => {},
  IsValidThunderstoreLink: async (link) => link.startsWith("https://thunderstore.io/c/lethal-company/p/"),
  InstallMod: async (link, path, vIdx) => { await delay(1500); },
  OpenFolder: async (path) => {},
  ParseThunderstoreLink: async (link) => ["Author", "ModName"],
};

function delay(ms) { return new Promise(r => setTimeout(r, ms)); }

function wails(method, ...args) {
  if (isBrowser) return mock[method](...args);
  return window.go.main.App[method](...args);
}

export const GetAssemblies = () => wails('GetAssemblies');
export const GetLCVersions = () => wails('GetLCVersions');
export const GetDefaultSteamFolder = () => wails('GetDefaultSteamFolder');
export const BrowseForSteamFolder = () => wails('BrowseForSteamFolder');
export const BrowseForLCInstallation = () => wails('BrowseForLCInstallation');
export const InstallCompleteAssembly = (idx, folder) => wails('InstallCompleteAssembly', idx, folder);
export const InstallNewInstance = (folder, vIdx, name) => wails('InstallNewInstance', folder, vIdx, name);
export const CheckHasBepInEx = (path) => wails('CheckHasBepInEx', path);
export const SetExistingInstallation = (path, vIdx, erase) => wails('SetExistingInstallation', path, vIdx, erase);
export const IsValidThunderstoreLink = (link) => wails('IsValidThunderstoreLink', link);
export const InstallMod = (link, path, vIdx) => wails('InstallMod', link, path, vIdx);
export const OpenFolder = (path) => wails('OpenFolder', path);
export const ParseThunderstoreLink = (link) => wails('ParseThunderstoreLink', link);
