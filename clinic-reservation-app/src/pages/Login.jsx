/* eslint-disable react/no-unescaped-entities */
/* eslint-disable react/prop-types */
import { useEffect, useState } from 'react';
import { login, setFormData } from '../services/api';
import { Link, useNavigate } from 'react-router-dom';

export default function Login({ saveUserData }) {
	let [data, setData] = useState({
		email: '',
		password: '',
	});
	let [message, setMessage] = useState('');
	const navigate = useNavigate();
	useEffect(() => {
		if (localStorage.getItem('auth-token')) navigate('/');
	}, []);
	return (
		<section className="flex flex-col space-y-5 justify-center items-center mt-10">
			<h2 className="text-center font-bold text-3xl text-dark-blue my-5">
				Login
			</h2>
			<form
				onSubmit={(e) => {
					login(e, data, setMessage, saveUserData, navigate);
				}}
				className="w-full flex flex-col space-y-10 px-10 max-w-3xl"
			>
				<div className="flex flex-col space-y-2 ">
					<label htmlFor="email" className="text-dark-blue font-semibold ml-2">
						Email
					</label>
					<input
						onChange={(e) => {
							setFormData(e, data, setData);
						}}
						type="text"
						name="email"
						id="email"
						placeholder="email"
						autoComplete="email"
						className="border-2 focus:outline-none px-3 py-4 rounded-lg placeholder-dark-white  border-slate-300 text-dark-blue font-semibold focus:border-dark-blue"
					/>
				</div>
				<div className="flex flex-col space-y-2 ">
					<label
						htmlFor="password"
						className="text-dark-blue font-semibold ml-2"
					>
						Password
					</label>
					<input
						onChange={(e) => {
							setFormData(e, data, setData);
						}}
						type="password"
						name="password"
						id="password"
						autoComplete="current-pass"
						placeholder="password"
						className="border-2 focus:outline-none px-3 py-4 rounded-lg placeholder-dark-white  border-slate-300 text-dark-blue font-semibold focus:border-dark-blue"
					/>
				</div>
				<div className="flex flex-col">
					<button
						type="submit"
						className="bg-dark-blue px-3 py-4 rounded-lg text-plain-white font-semibold outline-plain-white"
					>
						Login
					</button>
					<p className="mt-2 ms-2 text-lg">
						Don't hava an account,{' '}
						<Link to={'/register'} className="text-dark-blue hover:underline">
							register
						</Link>
					</p>
				</div>
				{message.length > 0 ? (
					<span className="bg-red-300 rounded-lg text-lg text-gray-700 font-medium px-5 py-4">
						{message}
					</span>
				) : (
					''
				)}
			</form>
		</section>
	);
}
