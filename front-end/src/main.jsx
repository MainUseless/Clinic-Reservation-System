import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App.jsx';
import './index.css';
import dotenv from 'dotenv';

if (typeof process !== 'undefined') {
	// This code will only run in a Node.js environment
	dotenv.config();
}
ReactDOM.createRoot(document.getElementById('root')).render(
	<React.StrictMode>
		<App />
	</React.StrictMode>
);
