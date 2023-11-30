import { Constants } from '../services/constants';

export default function Notification() {
	// for rabbitmq
	const socket = new WebSocket(`ws://localhost:${Constants.BACKEND_PORT}/ws?JWT=${localStorage.getItem('auth-token')}`);

	socket.onopen = () => {
		console.log('Connected to server');
		socket.send(JSON.stringify({ type : "authenticate" , token : localStorage.getItem('auth-token') }));
	}

	socket.onmessage = (event) => {
		// Handle received messages from the server (RabbitMQ)
		alert(event.data);
		// const message = JSON.parse(event.data);
	};

	return (
		<section className="flex flex-col space-y-5 justify-center items-center mt-10">
			<div className="w-5/6">
				<h2 className="text-center font-bold text-3xl text-dark-blue my-5">
					Notifications
				</h2>
				<div className="notifications space-y-5">
					<div className="bg-slate-500 flex justify-around items-center py-4 md:py-6 rounded-lg">
						<div className="flex flex-col items-center text-plain-white font-semibold">
							<p className="font-semibold text-lg md:text-2xl">Doctor ID</p>
							<p className="font-medium">1</p>
						</div>
						<div className="flex flex-col items-center text-plain-white font-semibold">
							<p className="font-semibold text-lg md:text-2xl">Patient ID</p>
							<p className="font-medium">1</p>
						</div>
						<div className="flex flex-col items-center text-plain-white space-y-1">
							<p className="font-semibold text-lg md:text-2xl">Operation</p>
							<p className="font-medium">Created</p>
						</div>
					</div>
				</div>
			</div>
		</section>
	);
}
