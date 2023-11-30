import axios from 'axios';
import { Constants } from './constants';

export const setFormData = (e, data, setData) => {
	let dataCopy = { ...data };
	dataCopy[e.target.name] = e.target.value;
	setData(dataCopy);
};

export const register = (e, registerData, setMessage, navigate) => {
	e.preventDefault();
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/account/signup`;
	axios
		.post(api, registerData)
		.then(() => {
			navigate('/login');
		})
		.catch((error) => {
			setMessage(error.response.data.error);
			// Handle error
		});
};

export const login = (e, loginData, setMessage, saveUserData, navigate) => {
	e.preventDefault();
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/account/signin`;
	axios
		.post(api, loginData)
		.then(({ data }) => {
			localStorage.setItem('auth-token', data.token);
			saveUserData();
			navigate('/');
		})
		.catch((error) => {
			setMessage(error.response.data.error);
			// Handle error
		});
};

export const createSlot = async (e, slotData, setMessage) => {
	e.preventDefault();
	let timestamp;
	if (slotData.date.length == 0 || slotData.time.length == 0) timestamp = '';
	else timestamp = slotData.date + ' ' + slotData.time;
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/doctor/appointment`;
	await axios
		.post(api, null, {
			params: { timestamp },
			headers: {
				Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
			},
		})
		.then((res) => {
			// if (!res.data)
			// else
			if (res.status === 204) setMessage('Invalid');
		})
		.catch(() => {
			setMessage('Already have slot in this time');
		});
};

export const deleteSlot = async (id) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/doctor/appointment`;
	await axios.delete(api, {
		headers: {
			Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
		},
		params: { appointment_id: id },
	});
};

export const getDoctorSlots = async (callback) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/doctor/appointment`;
	const { data } = await axios.get(api, {
		headers: {
			Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
		},
	});
	const appointments = data.appointments;
	if (appointments == null) callback([]);
	else callback(appointments);
};

export const getAllDoctors = async (callback) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/doctors`;
	const { data } = await axios.get(api, {
		headers: {
			Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
		},
	});
	const doctors = data.doctors;
	if (doctors === null) callback([]);
	else callback(doctors);
};

export const getPatientDoctors = async (callback) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/appointment`;

	const { data } = await axios.get(api, {
		headers: {
			Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
		},
	});
	const doctors = data.doctors;
	if (doctors !== null || doctors !== undefined || doctors.length !== 0)
		callback(doctors);
};

export const getDrSlotsById = async (id, setSlots) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/appointment`;
	const { data } = await axios.get(api, {
		headers: {
			Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
		},
		params: { doctor_id: id },
	});
	if (data.appointments === null) setSlots([]);
	else setSlots(data.appointments);
};

export const reserveSlot = async (e, id) => {
	e.preventDefault();
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/appointment`;
	await axios.post(api, null, {
		headers: {
			Authorization: 'Bearer ' + localStorage.getItem('auth-token'),
		},
		params: { appointment_id: id },
	});
};

export const getAppointments = async (callback) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/appointment`;
	const { data } = await axios.get(api, {
		headers: { Authorization: 'Bearer ' + localStorage.getItem('auth-token') },
	});
	if (data.appointments === null) callback([]);
	else callback(data.appointments);
};

export const deleteAppointment = async (id) => {
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/appointment`;
	await axios.delete(api, {
		headers: { Authorization: 'Bearer ' + localStorage.getItem('auth-token') },
		params: {
			appointment_id: id,
		},
	});
};
export const editAppointmentSlot = async (e, editedData) => {
	e.preventDefault();
	const api = `http://localhost:${Constants.BACKEND_PORT}/api/patient/appointment`;
	await axios.put(api, null, {
		headers: { Authorization: 'Bearer ' + localStorage.getItem('auth-token') },
		params: editedData,
	});
};
