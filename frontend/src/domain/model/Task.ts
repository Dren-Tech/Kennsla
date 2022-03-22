import type Block from './Block';

export default interface Task {
	id: number;
	slug: string;
	blocks: Array<Block>;
}
