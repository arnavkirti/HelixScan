"use client";
import { Dialog, Transition } from "@headlessui/react";
import Link from "next/link";
import { FormEvent, Fragment, useState } from "react";

export default function Home() {
  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [isSignupOpen, setIsSignupOpen] = useState(false);

  return (
    <div className="min-h-screen bg-gradient-to-b from-black via-purple-900 to-black text-white">
      <div className="container mx-auto px-4 py-16">
        {/* Hero Section */}
        <div className="text-center mb-20">
          <h1 className="text-7xl md:text-8xl font-bold mb-6 bg-clip-text text-transparent bg-gradient-to-r from-purple-400 to-pink-600">
            HelixScan
          </h1>
          <p className="text-xl text-gray-300 mb-8">
            Your Gateway to Seamless Blockchain Data Indexing
          </p>
          <div className="flex gap-6 justify-center mb-12">
            <button 
              onClick={() => setIsLoginOpen(true)}
              className="px-8 py-3 bg-purple-600 rounded-full hover:bg-purple-700 transition-all duration-300 font-semibold"
            >
              Login
            </button>
            <button 
              onClick={() => setIsSignupOpen(true)}
              className="px-8 py-3 bg-transparent border-2 border-purple-600 rounded-full hover:bg-purple-600/20 transition-all duration-300 font-semibold"
            >
              Sign Up
            </button>
          </div>
        </div>

        {/* Features Section */}
        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8 mb-20">
          <FeatureCard 
            title="Real-time NFT Data"
            description="Track NFT bids and prices across marketplaces instantly with our advanced indexing system"
            icon="🎭"
          />
          <FeatureCard 
            title="Token Analytics"
            description="Monitor token prices and borrowing rates across various platforms in real-time"
            icon="📊"
          />
          <FeatureCard 
            title="Custom Indexing"
            description="Define your own indexing parameters and get exactly the data you need"
            icon="⚙️"
          />
          <FeatureCard 
            title="Postgres Integration"
            description="Seamlessly store indexed data in your Postgres database with simple configuration"
            icon="🗄️"
          />
          <FeatureCard 
            title="Helius Powered"
            description="Leverage the power of Helius webhooks for reliable and fast blockchain data"
            icon="⚡"
          />
          <FeatureCard 
            title="Zero Infrastructure"
            description="No need to run your own RPC, Geyser, or Validator nodes"
            icon="☁️"
          />
        </div>

        {/* Bottom CTA */}
        <div className="text-center max-w-2xl mx-auto">
          <h2 className="text-2xl font-bold mb-4">Ready to Get Started?</h2>
          <p className="text-gray-400 mb-6 text-sm">
            Join the next generation of blockchain data indexing
          </p>
          <Link 
            href="/signup"
            className="px-8 py-3 bg-gradient-to-r from-purple-600 to-pink-600 rounded-full hover:opacity-90 transition-all duration-300 font-semibold inline-block text-sm"
          >
            Start Indexing Now
          </Link>
        </div>
      </div>
      
      {/* Add these two lines */}
      <LoginDialog isOpen={isLoginOpen} setIsOpen={setIsLoginOpen} />
      <SignupDialog isOpen={isSignupOpen} setIsOpen={setIsSignupOpen} />
    </div>
  );
}

const LoginDialog = ({ isOpen, setIsOpen }: { isOpen: boolean; setIsOpen: (isOpen: boolean) => void; }) => {
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    email: '',
    password: ''
  });

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
      });
      if (response.ok) {
        setIsOpen(false);
        // Handle successful login
      }
    } catch (error) {
      console.error('Login failed:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog as="div" className="relative z-10" onClose={() => setIsOpen(false)}>
        <div className="fixed inset-0 overflow-y-auto">
          <div className="flex min-h-full items-center justify-center p-4">
            <Dialog.Panel className="w-full max-w-md transform overflow-hidden rounded-2xl bg-gray-900/95 backdrop-blur-xl p-8 text-left align-middle shadow-xl transition-all border border-purple-500/20">
              <div className="flex justify-between items-center mb-6">
                <div>
                  <Dialog.Title as="h3" className="text-2xl font-bold text-white mb-2">
                    Welcome Back
                  </Dialog.Title>
                  <p className="text-gray-400 text-sm">Sign in to continue to HelixScan</p>
                </div>
                <button onClick={() => setIsOpen(false)} className="text-gray-400 hover:text-white">
                  <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
              <form className="space-y-5">
                <div>
                  <label className="block text-sm font-medium text-gray-400">Email</label>
                  <input 
                    type="email" 
                    value={formData.email}
                    onChange={(e) => setFormData({...formData, email: e.target.value})}
                    className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2" 
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-gray-400">Password</label>
                  <input 
                    type="password"
                    value={formData.password}
                    onChange={(e) => setFormData({...formData, password: e.target.value})}
                    className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2" 
                  />
                </div>
                <button 
                  type="submit"
                  disabled={isLoading}
                  className="w-full bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-lg p-3 font-medium hover:opacity-90 transition-all duration-300 disabled:opacity-50 flex items-center justify-center gap-2"
                >
                  {isLoading && (
                    <svg className="animate-spin h-5 w-5" viewBox="0 0 24 24">
                      <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none"/>
                      <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
                    </svg>
                  )}
                  {isLoading ? 'Signing in...' : 'Sign in'}
                </button>
              </form>
            </Dialog.Panel>
          </div>
        </div>
      </Dialog>
    </Transition>
  );
};

const SignupDialog = ({ isOpen, setIsOpen }: { isOpen: boolean; setIsOpen: (isOpen: boolean) => void; }) => {
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    database: {
      host: 'localhost',
      port: '5432',
      dbName: '',
      username: 'postgres',
      password: ''
    }
  });

  const handleSubmit = async (e:any) => {
    e.preventDefault();
    setIsLoading(true);
    
    try {
      const response = await fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          email: formData.email,
          password: formData.password,
          database: {
            host: formData.database.host,
            port: formData.database.port,
            db_name: formData.database.dbName,
            username: formData.database.username,
            password: formData.database.password
          }
        })
      });

      const data = await response.json();
      
      if (!response.ok) {
        throw new Error(data.message || 'Signup failed');
      }

      alert('Account created successfully!');
      setIsOpen(false);
      
    } catch (error:any) {
      alert(`Error: ${error.message}`);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Transition appear show={isOpen} as={Fragment}>
      <Dialog as="div" className="relative z-10" onClose={() => setIsOpen(false)}>
        {/* ... existing dialog structure ... */}
        <form onSubmit={handleSubmit} className="space-y-5">
          <div>
            <label className="block text-sm font-medium text-gray-400">Email</label>
            <input 
              type="email"
              value={formData.email}
              onChange={(e) => setFormData({...formData, email: e.target.value})}
              className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2" 
              required
            />
          </div>
          
          <div>
            <label className="block text-sm font-medium text-gray-400">Password</label>
            <input 
              type="password"
              value={formData.password}
              onChange={(e) => setFormData({...formData, password: e.target.value})}
              className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2" 
              required
            />
          </div>

          <h4 className="text-lg font-medium text-gray-300 mt-6">Database Configuration</h4>
          
          <div>
            <label className="block text-sm font-medium text-gray-400">Host</label>
            <input 
              type="text"
              value={formData.database.host}
              onChange={(e) => setFormData({
                ...formData,
                database: {...formData.database, host: e.target.value}
              })}
              className="mt-1 w-full rounded-md bg-gray-800 border border-gray-700 text-white px-3 py-2" 
              required
            />
          </div>
          <button
            type="submit"
            disabled={isLoading}
            className="w-full bg-gradient-to-r from-purple-600 to-pink-600 text-white rounded-lg p-3 font-medium hover:opacity-90 transition-all duration-300 disabled:opacity-50 flex items-center justify-center gap-2"
          >
            {isLoading ? 'Creating account...' : 'Create Account'}
          </button>
        </form>
      </Dialog>
    </Transition>
  );
};

const FeatureCard = ({ 
  title, 
  description, 
  icon 
}: {
  title: string;
  description: string;
  icon: string;
}) => (
  <div className="p-6 rounded-xl bg-purple-900/20 backdrop-blur-sm border border-purple-800/50 hover:border-purple-600/50 transition-all duration-300">
    <div className="text-4xl mb-4">{icon}</div>
    <h3 className="text-xl font-semibold mb-3">{title}</h3>
    <p className="text-gray-400">{description}</p>
  </div>
);
