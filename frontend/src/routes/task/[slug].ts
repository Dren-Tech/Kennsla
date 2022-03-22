export async function get({ params }) {
	params.slug;

	const response = await fetch('http://localhost:3000/task.json');
	const task = await response.json();

	if (task) {
		return {
			body: { task }
		};
	}

	return {
		status: 404
	};
}
