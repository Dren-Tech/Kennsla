import { test, expect } from '@playwright/test';

test('Task: Accordion show content on click', async ({ page }) => {
	// Go to http://localhost:3000/task/test
	await page.goto('http://localhost:3000/task/test');

	expect(await page.textContent('.block.hint')).not.toContain('Hinweis-Text 1.');
	expect(await page.textContent('.block.hint')).not.toContain('Hinweis-Text 2.');

	// Click button:has-text("Hinweis Nr. 1")
	await page.locator('button:has-text("Hinweis Nr. 1")').click();
	expect(await page.textContent('.block.hint')).toContain('Hinweis-Text 1.');

	// Click button:has-text("Hinweis Nr. 2")
	await page.locator('button:has-text("Hinweis Nr. 2")').click();
	expect(await page.textContent('.block.hint')).toContain('Hinweis-Text 2.');
});
