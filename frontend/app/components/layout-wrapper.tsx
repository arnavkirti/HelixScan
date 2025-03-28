"use client";

import Navigation from "./navigation";

export default function LayoutWrapper({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-gradient-to-b from-black via-purple-900 to-black text-white">
      <Navigation />
      <div className="pt-16">{children}</div>
    </div>
  );
}