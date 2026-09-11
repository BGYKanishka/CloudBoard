import { useState, useEffect } from 'react';
import api from '../api';

interface Task {
  id: number;
  title: string;
  description: string;
  status: string;
  priority: string;
}

export default function Dashboard() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');

  const fetchTasks = async () => {
    try {
      const res = await api.get('/tasks');
      setTasks(res.data || []);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    // eslint-disable-next-line
    fetchTasks();
  }, []);

  const handleCreate = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;
    try {
      await api.post('/tasks', { title, description, status: 'pending', priority: 'medium' });
      setTitle('');
      setDescription('');
      fetchTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const handleStatusChange = async (id: number, status: string) => {
    try {
      await api.patch(`/tasks/${id}/status`, { status });
      fetchTasks();
    } catch (err) {
      console.error(err);
    }
  };

  const handleDelete = async (id: number) => {
    try {
      await api.delete(`/tasks/${id}`);
      fetchTasks();
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Dashboard</h1>
      
      <div className="bg-white p-6 rounded-lg shadow-sm mb-8">
        <h2 className="text-xl font-semibold mb-4">Create New Task</h2>
        <form onSubmit={handleCreate} className="space-y-4">
          <div>
            <input 
              type="text" placeholder="Task Title"
              className="w-full px-4 py-2 border rounded-md"
              value={title} onChange={e => setTitle(e.target.value)}
            />
          </div>
          <div>
            <textarea 
              placeholder="Description (optional)"
              className="w-full px-4 py-2 border rounded-md"
              value={description} onChange={e => setDescription(e.target.value)}
            />
          </div>
          <button type="submit" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Add Task</button>
        </form>
      </div>

      <div className="space-y-4">
        <h2 className="text-xl font-semibold mb-4">Your Tasks</h2>
        {tasks.length === 0 ? (
          <p className="text-gray-500">No tasks yet. Create one above!</p>
        ) : (
          tasks.map(task => (
            <div key={task.id} className="bg-white p-4 rounded-lg shadow-sm border border-gray-100 flex justify-between items-center">
              <div>
                <h3 className={`font-semibold ${task.status === 'completed' ? 'line-through text-gray-500' : 'text-gray-800'}`}>
                  {task.title}
                </h3>
                <p className="text-gray-600 text-sm mt-1">{task.description}</p>
                <div className="flex space-x-2 mt-2">
                  <span className={`text-xs px-2 py-1 rounded ${task.priority === 'high' ? 'bg-red-100 text-red-800' : 'bg-gray-100 text-gray-800'}`}>
                    {task.priority}
                  </span>
                </div>
              </div>
              <div className="flex space-x-3">
                <select 
                  className="border rounded px-2 py-1 text-sm bg-gray-50"
                  value={task.status}
                  onChange={(e) => handleStatusChange(task.id, e.target.value)}
                >
                  <option value="pending">Pending</option>
                  <option value="in_progress">In Progress</option>
                  <option value="completed">Completed</option>
                </select>
                <button 
                  onClick={() => handleDelete(task.id)}
                  className="text-red-500 hover:text-red-700"
                >
                  Delete
                </button>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}
