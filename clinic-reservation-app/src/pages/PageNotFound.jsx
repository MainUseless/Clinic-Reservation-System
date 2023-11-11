import errorImage from '../assets/404.webp';
export default function PageNotFound() {
	return (
		<div className="w-screen h-screen flex justify-center items-center">
			<div className="p-5 md:p-0">
				<img src={errorImage} alt="error page 404" className="w-full" />
			</div>
		</div>
	);
}
