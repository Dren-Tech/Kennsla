import type Task from '../../domain/model/Task';

export async function get({ params }) {
	const response = await fetch(`${import.meta.env.VITE_API_URL}/task/${params.slug}`, {
		headers: {
			'Authorization': `Bearer ${import.meta.env.VITE_API_KEY}`
		}
	});
	const task: Task = await response.json();

	if (task) {
		return {
			body: { task }
		};
	}

	return {
		status: 404
	};
}
