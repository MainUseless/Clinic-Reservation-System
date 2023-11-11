/* eslint-disable react/prop-types */
import { Link } from 'react-router-dom';
import { BellIcon } from '@heroicons/react/24/outline';
export default function Navbar({ logOut, userData }) {
	return (
		<nav className="bg-dark-blue">
			<div className="flex justify-between items-center text-plain-white px-5 md:px-20">
				<Link to={'/'}>
					<h1 className="font-bold text-xl md:text-2xl mx-4 my-4 md:my-6">
						Clinica
					</h1>
				</Link>
				<div className="not-authenticated">
					<ul className="flex justify-center items-center">
						{userData === null ? (
							<>
								<Link to={'login'}>
									<li className="px-3 hover:font-semibold hover:cursor-pointer md:text-lg">
										Login
									</li>
								</Link>
								<Link to={'register'}>
									<li className="px-3 hover:font-semibold hover:cursor-pointer md:text-lg">
										Register
									</li>
								</Link>
							</>
						) : (
							<>
								<Link to={'notification'}>
									<BellIcon className="w-8 h-8 hover:fill-white"></BellIcon>
								</Link>
								<li
									onClick={logOut}
									className="px-3 hover:font-semibold hover:cursor-pointer md:text-lg"
								>
									Logout
								</li>
							</>
						)}
					</ul>
				</div>
			</div>
		</nav>
	);
}
