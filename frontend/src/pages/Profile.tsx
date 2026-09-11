import { useState, useEffect } from 'react';
import api from '../api';

interface User {
  id: number;
  name: string;
  email: string;
  profile_image_url: string;
}

export default function Profile() {
  const [user, setUser] = useState<User | null>(null);
  const [uploading, setUploading] = useState(false);

  useEffect(() => {
    api.get('/auth/me')
      .then(res => setUser(res.data))
      .catch(err => console.error(err));
  }, []);

  const handleImageUpload = async () => {
    // Phase 8 stub: simulate upload
    setUploading(true);
    try {
      const res = await api.post('/profile/image');
      setUser(prev => prev ? { ...prev, profile_image_url: res.data.url } : null);
      alert('Simulated image upload successful (Phase 8 placeholder)');
    } catch (err) {
      console.error(err);
    } finally {
      setUploading(false);
    }
  };

  if (!user) return <div>Loading profile...</div>;

  return (
    <div className="max-w-3xl mx-auto bg-white p-8 rounded-lg shadow-sm">
      <h1 className="text-2xl font-bold mb-6">Your Profile</h1>
      
      <div className="flex items-center space-x-6 mb-8">
        <div className="w-24 h-24 bg-gray-200 rounded-full flex items-center justify-center overflow-hidden border">
          {user.profile_image_url ? (
            <img src={user.profile_image_url} alt="Profile" className="w-full h-full object-cover" />
          ) : (
            <span className="text-gray-400">No Image</span>
          )}
        </div>
        <div>
          <button 
            onClick={handleImageUpload}
            disabled={uploading}
            className="bg-blue-50 border border-blue-200 text-blue-700 px-4 py-2 rounded text-sm hover:bg-blue-100 disabled:opacity-50"
          >
            {uploading ? 'Uploading...' : 'Upload New Picture (Stub)'}
          </button>
          <p className="text-xs text-gray-500 mt-2">Max size 2MB</p>
        </div>
      </div>

      <div className="space-y-4">
        <div>
          <label className="block text-sm font-medium text-gray-700">Name</label>
          <div className="mt-1 p-3 bg-gray-50 rounded border text-gray-900">{user.name}</div>
        </div>
        <div>
          <label className="block text-sm font-medium text-gray-700">Email Address</label>
          <div className="mt-1 p-3 bg-gray-50 rounded border text-gray-900">{user.email}</div>
        </div>
      </div>
    </div>
  );
}
