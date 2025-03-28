"use client";

import { FormEvent, useState } from "react";

interface DatabaseConfig {
  host: string;
  port: string;
  dbName: string;
  username: string;
  password: string;
}

interface IndexingPreference {
  nftBids: boolean;
  nftPrices: boolean;
  borrowableTokens: boolean;
  tokenPrices: boolean;
  customFilters: Record<string, unknown>;
}

interface SyncStatus {
  lastSynced: string;
  syncedBlocks: number;
  errorLog: string;
}

export default function Dashboard() {
  const [dbConfig, setDbConfig] = useState<DatabaseConfig>({
    host: "",
    port: "",
    dbName: "",
    username: "",
    password: "",
  });

  const [indexingPrefs, setIndexingPrefs] = useState<IndexingPreference>({
    nftBids: true,
    nftPrices: true,
    borrowableTokens: false,
    tokenPrices: false,
    customFilters: {},
  });

  const [syncStatus, setSyncStatus] = useState<SyncStatus>({
    lastSynced: "",
    syncedBlocks: 0,
    errorLog: "",
  });

  const handleDbConfigSubmit = async (e: FormEvent) => {
    e.preventDefault();
    try {
      const response = await fetch("/api/config/database", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(dbConfig),
      });
      if (!response.ok) throw new Error("Failed to update database config");
    } catch (error) {
      console.error("Error updating database config:", error);
    }
  };

  const handlePrefsSubmit = async (e: FormEvent) => {
    e.preventDefault();
    try {
      const response = await fetch("/api/config/indexing", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(indexingPrefs),
      });
      if (!response.ok) throw new Error("Failed to update indexing preferences");
    } catch (error) {
      console.error("Error updating indexing preferences:", error);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-b from-black via-purple-900 to-black text-white p-8">
      <div className="max-w-6xl mx-auto space-y-8">
        <h1 className="text-4xl font-bold mb-8">Dashboard</h1>

        {/* Database Configuration */}
        <div className="bg-gray-900/80 backdrop-blur-xl p-6 rounded-xl border border-purple-500/20">
          <h2 className="text-2xl font-semibold mb-4">Database Configuration</h2>
          <form onSubmit={handleDbConfigSubmit} className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div>
                <label className="block text-sm font-medium text-gray-400">Host</label>
                <input
                  type="text"
                  value={dbConfig.host}
                  onChange={(e) => setDbConfig({ ...dbConfig, host: e.target.value })}
                  className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-400">Port</label>
                <input
                  type="text"
                  value={dbConfig.port}
                  onChange={(e) => setDbConfig({ ...dbConfig, port: e.target.value })}
                  className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-400">Database Name</label>
                <input
                  type="text"
                  value={dbConfig.dbName}
                  onChange={(e) => setDbConfig({ ...dbConfig, dbName: e.target.value })}
                  className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-gray-400">Username</label>
                <input
                  type="text"
                  value={dbConfig.username}
                  onChange={(e) => setDbConfig({ ...dbConfig, username: e.target.value })}
                  className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2"
                />
              </div>
              <div className="md:col-span-2">
                <label className="block text-sm font-medium text-gray-400">Password</label>
                <input
                  type="password"
                  value={dbConfig.password}
                  onChange={(e) => setDbConfig({ ...dbConfig, password: e.target.value })}
                  className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2"
                />
              </div>
            </div>
            <button
              type="submit"
              className="w-full bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-lg p-3 font-medium hover:opacity-90 transition-all duration-300"
            >
              Save Database Configuration
            </button>
          </form>
        </div>

        {/* Indexing Preferences */}
        <div className="bg-gray-900/80 backdrop-blur-xl p-6 rounded-xl border border-purple-500/20">
          <h2 className="text-2xl font-semibold mb-4">Indexing Preferences</h2>
          <form onSubmit={handlePrefsSubmit} className="space-y-4">
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <div className="flex items-center space-x-3">
                <input
                  type="checkbox"
                  id="nftBids"
                  checked={indexingPrefs.nftBids}
                  onChange={(e) => setIndexingPrefs({ ...indexingPrefs, nftBids: e.target.checked })}
                  className="w-4 h-4 rounded border-gray-700 text-purple-600 focus:ring-purple-500"
                />
                <label htmlFor="nftBids" className="text-sm font-medium text-gray-300">
                  NFT Bids
                </label>
              </div>
              <div className="flex items-center space-x-3">
                <input
                  type="checkbox"
                  id="nftPrices"
                  checked={indexingPrefs.nftPrices}
                  onChange={(e) => setIndexingPrefs({ ...indexingPrefs, nftPrices: e.target.checked })}
                  className="w-4 h-4 rounded border-gray-700 text-purple-600 focus:ring-purple-500"
                />
                <label htmlFor="nftPrices" className="text-sm font-medium text-gray-300">
                  NFT Prices
                </label>
              </div>
              <div className="flex items-center space-x-3">
                <input
                  type="checkbox"
                  id="borrowableTokens"
                  checked={indexingPrefs.borrowableTokens}
                  onChange={(e) => setIndexingPrefs({ ...indexingPrefs, borrowableTokens: e.target.checked })}
                  className="w-4 h-4 rounded border-gray-700 text-purple-600 focus:ring-purple-500"
                />
                <label htmlFor="borrowableTokens" className="text-sm font-medium text-gray-300">
                  Borrowable Tokens
                </label>
              </div>
              <div className="flex items-center space-x-3">
                <input
                  type="checkbox"
                  id="tokenPrices"
                  checked={indexingPrefs.tokenPrices}
                  onChange={(e) => setIndexingPrefs({ ...indexingPrefs, tokenPrices: e.target.checked })}
                  className="w-4 h-4 rounded border-gray-700 text-purple-600 focus:ring-purple-500"
                />
                <label htmlFor="tokenPrices" className="text-sm font-medium text-gray-300">
                  Token Prices
                </label>
              </div>
            </div>
            <button
              type="submit"
              className="w-full bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-lg p-3 font-medium hover:opacity-90 transition-all duration-300"
            >
              Save Indexing Preferences
            </button>
          </form>
        </div>

        {/* Sync Status */}
        <div className="bg-gray-900/80 backdrop-blur-xl p-6 rounded-xl border border-purple-500/20">
          <h2 className="text-2xl font-semibold mb-4">Sync Status</h2>
          <div className="space-y-4">
            <div>
              <p className="text-sm text-gray-400">Last Synced</p>
              <p className="text-lg font-mono">{syncStatus.lastSynced || "Never"}</p>
            </div>
            <div>
              <p className="text-sm text-gray-400">Synced Blocks</p>
              <p className="text-lg font-mono">{syncStatus.syncedBlocks}</p>
            </div>
            {syncStatus.errorLog && (
              <div>
                <p className="text-sm text-gray-400">Error Log</p>
                <p className="text-red-400 font-mono text-sm mt-1">{syncStatus.errorLog}</p>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}