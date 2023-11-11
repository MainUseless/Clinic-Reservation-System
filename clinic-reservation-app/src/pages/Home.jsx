import { useEffect, useState } from 'react';
import { getDataFromToken } from '../services/helpers';
import {
	createSlot,
	setFormData,
	getAllDoctors,
	getDoctorSlots,
	deleteSlot,
	getDrSlotsById,
	reserveSlot,
	getAppointments,
	deleteAppointment,
	editAppointmentSlot,
} from '../services/api';
import {
	TrashIcon,
	PencilSquareIcon,
	XCircleIcon,
} from '@heroicons/react/24/outline';
/* eslint-disable react/prop-types */
export default function Home() {
	const [user, setUser] = useState(getDataFromToken());
	const [isOpen, setIsOpen] = useState(false);
	const [data, setData] = useState({ date: '', time: '' });
	const [message, setMessage] = useState('');
	const [doctors, setDoctors] = useState([]);
	const [slots, setSlots] = useState([]);
	const [selectedDoctor, setSelectedDoctor] = useState(null);
	const [selectedSlots, setSelectedSlots] = useState([]);
	const [appointements, setAppointments] = useState([]);
	const [targetedSlot, setTargetedSlot] = useState('');
	const [editedApp, setEditedApp] = useState('');
	const [editedSlots, setEditedSlots] = useState('');
	const [editAppApi, setEditApi] = useState({
		appointment_id: '',
		new_appointment_id: '',
	});
	console.log('docId', editedApp);
	console.log('Api', editAppApi);
	useEffect(() => {
		setUser(getDataFromToken());
		// for patient layout
		if (user.type == 'patient') {
			getAllDoctors(setDoctors);
			getAppointments(setAppointments);
		} else {
			// for doctor layout
			getDoctorSlots(setSlots);
		}
	}, []);

	useEffect(() => {
		return () => {
			getAppointments(setAppointments);
			getDoctorSlots(setSlots);
			getDrSlotsById(selectedDoctor, setSelectedSlots);
			getDrSlotsById(editedApp, setEditedSlots);
		};
	}, [appointements, slots, selectedDoctor, editedApp]);

	return (
		<section className="md:w-5/6 m-auto p-5 md:p-7 space-y-10 md:space-y-20 relative">
			<h2 className="text-xl md:text-3xl">
				Hello <span className="text-dark-blue">{user.name.split(' ')[0]}</span>,
				<span className="text-base ms-3">({user.type})</span>
			</h2>
			{user.type === 'doctor' && (
				<>
					<div className="flex flex-col">
						<h3 className="text-2xl md:text-4xl font-semibold text-dark-blue my-5 md:my-9 text-center">
							My Slots
						</h3>
						<div className="flex justify-around text-lg md:text-xl font-semibold">
							<span>Date</span>
							<span>Time</span>
							<span>Is Reserved</span>
							<span className="text-red-500">Cancel</span>
						</div>
						<span className="h-[0.2px] w-full bg-slate-500 my-3"></span>
						<div className="space-y-7">
							{slots.map((slot, i) => (
								<div key={i} className="flex justify-around items-center">
									<p className="md:text-base text-sm">
										{slot.appointment_time.split(' ')[0]}
									</p>
									<p className="md:text-base text-sm mr-5">
										{slot.appointment_time.split(' ')[1]}
									</p>
									{slot.patient_id !== null ? (
										<span className=" bg-red-500 font-semibold text-plain-white px-3 py-2 rounded-md text-xs md:text-base mr-9">
											Reserved
										</span>
									) : (
										<span className="w-36"></span>
									)}
									<TrashIcon
										onClick={() => {
											deleteSlot(slot.id);
										}}
										className="w-5 md:w-8 cursor-pointer text-red-500 mr-7"
									></TrashIcon>
								</div>
							))}
						</div>
					</div>
					<div className="bg-slate-500 w-2/3 m-auto rounded-lg p-3 md:p-5 flex flex-col items-center">
						<h3 className="text-xl font-semibold text-plain-white">
							Add new Slot
						</h3>
						<form
							onSubmit={(e) => {
								setMessage('');
								createSlot(e, data, setMessage);
							}}
							className="flex flex-col w-full md:w-2/3"
						>
							<input
								onChange={(e) => {
									setFormData(e, data, setData);
								}}
								type="date"
								name="date"
								id="date"
								className="rounded-lg px-2 py-3 text-lg my-3 cursor-pointer"
							/>
							<input
								onChange={(e) => {
									setFormData(e, data, setData);
								}}
								type="time"
								name="time"
								id="time"
								className="rounded-lg px-2 py-3 text-lg cursor-pointer"
							/>
							<button
								type="submit"
								className="bg-dark-blue px-3 py-4 rounded-lg text-plain-white font-semibold outline-plain-white mt-7"
							>
								Add slot
							</button>
							{message.length > 0 ? (
								<span className="bg-red-300 rounded-lg text-lg text-gray-700 font-medium px-5 py-4 mt-2">
									{message}
								</span>
							) : (
								''
							)}
						</form>
					</div>
				</>
			)}
			{user.type === 'patient' && (
				<>
					{isOpen && (
						<div className="absolute top-10 left-0 right-0 bottom-0 bg-slate-500 flex flex-col items-center justify-center mx-3 p-5 rounded-lg">
							{' '}
							<XCircleIcon
								onClick={() => {
									setIsOpen(false);
								}}
								className="absolute top-5 right-5 w-10 h-10 cursor-pointer text-plain-white"
							></XCircleIcon>
							<h3 className="text-2xl md:text-4xl font-semibold text-plain-white my-5 md:my-9 text-center">
								Edit Appointment
							</h3>
							<form
								onSubmit={(e) => {
									setIsOpen(false);
									editAppointmentSlot(e, editAppApi);
								}}
								className="flex flex-col w-full md:w-2/3"
							>
								<label
									htmlFor="app"
									className="px-2 py-1 font-semibold text-plain-white"
								>
									Slot
								</label>
								<select
									onChange={(e) => {
										setEditApi({
											appointment_id: editedApp,
											new_appointment_id: parseInt(e.target.value),
										});
									}}
									id="app"
									className="rounded-lg px-2 py-3 text-lg cursor-pointer"
								>
									<>
										<option value=""></option>
										{editedSlots.map(
											(slot, i) =>
												slot.patient_id === null && (
													<option key={i} value={slot.id}>
														{slot.appointment_time}
													</option>
												)
										)}
									</>
								</select>
								<button
									type="submit"
									className="bg-dark-blue px-3 py-4 rounded-lg text-plain-white font-semibold outline-plain-white mt-7"
								>
									Edit
								</button>
							</form>
						</div>
					)}
					<div className="flex flex-col">
						<h3 className="text-2xl md:text-4xl font-semibold text-dark-blue my-5 md:my-9 text-center">
							My Appointments
						</h3>
						<div className="flex justify-around text-lg md:text-xl font-semibold">
							<span>Date</span>
							<span>Time</span>
							<span>Doctor</span>
							<span className="text-yellow-600">Edit</span>
							<span className="text-red-500">Cancel</span>
						</div>
						<span className="h-[0.2px] w-full bg-slate-500 my-3"></span>
						<div className="space-y-7">
							{appointements.map((appointement, i) => (
								<div key={i} className="flex justify-around items-center ">
									<p className="md:text-base text-sm">
										{appointement.appointment_time.split(' ')[0]}
									</p>
									<p className="md:text-base text-sm mr-5">
										{appointement.appointment_time.split(' ')[1]}
									</p>
									<p className="md:text-base text-sm mr-5">
										Dr.{appointement.name}
									</p>
									<PencilSquareIcon
										onClick={() => {
											setEditedApp(appointement.id);
											setIsOpen(true);
										}}
										className="w-5 md:w-8 cursor-pointer text-yellow-600 mr-12"
									></PencilSquareIcon>
									<TrashIcon
										onClick={() => {
											deleteAppointment(appointement.id);
										}}
										className="w-5 md:w-8 cursor-pointer text-red-500 mr-7"
									></TrashIcon>
								</div>
							))}
						</div>
					</div>
					<div className="bg-slate-500 w-2/3 m-auto rounded-lg p-3 md:p-5 flex flex-col items-center">
						<h3 className="text-xl font-semibold text-plain-white">
							Add new Appointment
						</h3>
						<form
							onSubmit={(e) => {
								reserveSlot(e, targetedSlot);
							}}
							className={`flex flex-col w-full md:w-2/3 space-y-3`}
						>
							<div className="flex flex-col">
								<label
									htmlFor="appointment_id"
									className="px-2 py-1 font-semibold text-plain-white"
								>
									Doctor
								</label>
								<select
									onChange={(e) => {
										setSelectedDoctor(e.target.value);
									}}
									id="appointment_id"
									className="rounded-lg px-2 py-3 text-lg cursor-pointer"
								>
									<option value=""></option>
									{doctors.map((doc, i) => (
										<option key={i} value={doc.id}>
											Dr.{doc.name}
										</option>
									))}
								</select>
							</div>
							<div className="flex flex-col">
								<label
									htmlFor="slot"
									className="px-2 py-1 font-semibold text-plain-white"
								>
									Slot
								</label>
								<select
									onChange={(e) => {
										setTargetedSlot(e.target.value);
									}}
									id="slot"
									className="rounded-lg px-2 py-3 text-lg cursor-pointer"
								>
									<option value=""></option>
									{selectedSlots.map(
										(slot, i) =>
											slot.patient_id === null && (
												<option key={i} value={slot.id}>
													{slot.appointment_time}
												</option>
											)
									)}
								</select>
							</div>
							<button
								type="submit"
								className="bg-dark-blue px-3 py-4 rounded-lg text-plain-white font-semibold outline-plain-white mt-7"
							>
								Reserve
							</button>
						</form>
					</div>
				</>
			)}
		</section>
	);
}
