import { test, expect } from '@playwright/test';

test.beforeAll(async ({ request }) => {
	// Create a new repository
	const response = await request.post(`${process.env['VITE_API_URL']}/task`, {
		data: {
			slug: 'test',
			blocks: [
				{
					type: 'hint',
					payload: {
						title: 'Hinweis Nr. 1',
						text: 'Hinweis-Text 1.'
					}
				},
				{
					type: 'hint',
					payload: {
						title: 'Hinweis Nr. 2',
						text: 'Hinweis-Text 2.'
					}
				}
			]
		}
	});
	console.log(await response.body());
	expect(response.ok()).toBeTruthy();
});

test('Task: Accordion show content on click', async ({ page }) => {
	// Go to http://localhost:3000/task/test
	await page.goto('http://localhost:3000/task/test');

	expect(await page.textContent('.block.hint')).not.toContain('Hinweis-Text 1.');
	expect(await page.textContent('.block.hint')).not.toContain('Hinweis-Text 2.');

	// Click button:has-text("Hinweis Nr. 1")
	await page.locator('button:has-text("Hinweis Nr. 1")').click();
	await page.waitForTimeout(502);
	expect(await page.textContent('.block.hint')).toContain('Hinweis-Text 1.');

	// Click button:has-text("Hinweis Nr. 2")
	await page.locator('button:has-text("Hinweis Nr. 2")').click();
	await page.waitForTimeout(502);
	expect(await page.textContent(':nth-match(.block.hint, 2)')).toContain('Hinweis-Text 2.');
});
