import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import PageNotFound from './pages/PageNotFound';
import Home from './pages/Home';
import Register from './pages/Register';
import Login from './pages/Login';
import Navbar from './components/Navbar';
import { useEffect, useState } from 'react';
import { jwtDecode } from 'jwt-decode';
import Notification from './pages/Notification';

export const App = () => {
	let [userData, setUserData] = useState(null);

	const saveUserData = () => {
		const data = jwtDecode(localStorage.getItem('auth-token'));
		setUserData(data);
	};

	const logOut = () => {
		setUserData(null);
		localStorage.removeItem('auth-token');
		return <Navigate to={'/login'} />;
	};
	// protect routes if not authenticated users
	const Protect = (props) => {
		if (localStorage.getItem('auth-token') === null) {
			return <Navigate to={'/login'} />;
		}
		// eslint-disable-next-line react/prop-types
		return props.children;
	};
	// if user refresh will not clear the "userData" useState
	useEffect(() => {
		if (localStorage.getItem('auth-token')) saveUserData();
	}, []);

	return (
		<BrowserRouter>
			<Navbar logOut={logOut} userData={userData} />
			<Routes>
				{/* <Protect></Protect> */}
				<Route
					path="/"
					element={
						<Protect>
							<Home userData={userData} />
						</Protect>
					}
				/>
				<Route path="/register" element={<Register />} />
				<Route
					path="/login"
					element={
						<Login saveUserData={saveUserData} setUserData={setUserData} />
					}
				/>
				<Route
					path="/notification"
					element={
						<Protect>
							<Notification />
						</Protect>
					}
				/>
				<Route path="*" element={<PageNotFound />} />
			</Routes>
		</BrowserRouter>
	);
};

export default App;
