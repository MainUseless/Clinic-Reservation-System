import { jwtDecode } from 'jwt-decode';

export const getDataFromToken = () => {
	return jwtDecode(localStorage.getItem('auth-token'));
};
