/* eslint-disable react-hooks/exhaustive-deps */
// import { useEffect } from 'react';
// import { getData } from '../services/api';

import { useEffect, useState } from 'react';
import { register, setFormData } from '../services/api';
import { Link, useNavigate } from 'react-router-dom';

export default function Register() {
	let [data, setData] = useState({
		name: '',
		email: '',
		password: '',
		type: '',
	});
	let [message, setMessage] = useState('');
	const navigate = useNavigate();

	useEffect(() => {
		if (localStorage.getItem('auth-token')) navigate('/');
	}, []);

	return (
		<section className="flex flex-col space-y-5 justify-center items-center mt-10">
			<h2 className="text-center font-bold text-3xl text-dark-blue my-5">
				Register
			</h2>
			<form
				onSubmit={(e) => {
					register(e, data, setMessage, navigate);
				}}
				className="w-full flex flex-col space-y-10 px-10 max-w-3xl"
			>
				<div className="flex flex-col space-y-2">
					<label htmlFor="name" className="text-dark-blue font-semibold ml-2">
						Name
					</label>
					<input
						onChange={(e) => {
							setFormData(e, data, setData);
						}}
						type="text"
						name="name"
						id="name"
						placeholder="name"
						autoComplete="name"
						className="border-2 focus:outline-none px-3 py-4 rounded-lg placeholder-dark-white  border-slate-300 text-dark-blue font-semibold focus:border-dark-blue"
					/>
				</div>
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
						autoComplete="email"
						placeholder="email"
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

				<div className="flex justify-between items-center">
					<p className="text-dark-blue font-semibold ml-2">Type</p>
					<div className="flex justify- items-center mr-3 px-5 space-x-5">
						<div className="flex justify-center items-center">
							<label htmlFor="doctor" className="mx-3 cursor-pointer">
								Doctor
							</label>
							<input
								onChange={(e) => {
									setFormData(e, data, setData);
								}}
								type="radio"
								name="type"
								id="doctor"
								value={'doctor'}
								className="cursor-pointer"
							/>
						</div>
						<div className="flex justify-center items-center">
							<label htmlFor="patient" className="mx-3 cursor-pointer">
								Patient
							</label>
							<input
								onChange={(e) => {
									setFormData(e, data, setData);
								}}
								type="radio"
								name="type"
								id="patient"
								value={'patient'}
								className="cursor-pointer"
							/>
						</div>
					</div>
				</div>
				<div className="flex flex-col">
					<button
						type="submit"
						className="bg-dark-blue px-3 py-4 rounded-lg text-plain-white font-semibold outline-plain-white"
					>
						Register
					</button>
					<p className="mt-2 ms-2 text-lg">
						Already hava an account,{' '}
						<Link to={'/login'} className="text-dark-blue hover:underline">
							login
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
