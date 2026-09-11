import { Link, useNavigate } from 'react-router-dom';
import api from '../api';

export default function Navbar({ onLogout }: { onLogout: () => void }) {
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      await api.post('/auth/logout');
      onLogout();
      navigate('/login');
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <nav className="bg-white shadow-sm border-b">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between h-16">
          <div className="flex items-center space-x-8">
            <Link to="/" className="text-xl font-bold text-blue-600">CloudBoard</Link>
            <div className="hidden sm:flex sm:space-x-4">
              <Link to="/" className="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">Dashboard</Link>
              <Link to="/profile" className="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium">Profile</Link>
            </div>
          </div>
          <div className="flex items-center">
            <button 
              onClick={handleLogout}
              className="text-gray-500 hover:text-gray-700 font-medium text-sm"
            >
              Logout
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
}
