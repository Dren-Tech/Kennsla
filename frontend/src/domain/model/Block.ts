export default interface Block {
	id: number;
	type: string;
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	payload: any;
}
