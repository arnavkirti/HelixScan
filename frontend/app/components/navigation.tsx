"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Navigation() {
  const pathname = usePathname();

  return (
    <nav className="fixed top-0 left-0 right-0 bg-gray-900/80 backdrop-blur-xl border-b border-purple-500/20 z-50">
      <div className="max-w-6xl mx-auto px-4 py-4">
        <div className="flex items-center justify-between">
          <Link href="/" className="text-2xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-600">
            HelixScan
          </Link>
          <div className="flex items-center space-x-6">
            <Link
              href="/dashboard"
              className={`text-sm font-medium transition-colors duration-200 ${pathname === "/dashboard" ? "text-purple-400" : "text-gray-300 hover:text-white"}`}
            >
              Dashboard
            </Link>
            <button
              onClick={() => {
                // TODO: Implement logout
                console.log("Logout clicked");
              }}
              className="text-sm font-medium text-gray-300 hover:text-white transition-colors duration-200"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
}