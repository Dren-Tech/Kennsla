// @ts-check
const { devices } = require('@playwright/test');

/** @type {import('@playwright/test').PlaywrightTestConfig} */
const config = {
	use: {
		baseURL: 'http://localhost:3000',
		screenshot: 'only-on-failure'
	},
	retries: process.env.CI ? 2 : 0,
	webServer: {
		command: 'npm run build && npm run preview',
		port: 3000,
		timeout: 120 * 1000,
		reuseExistingServer: !process.env.CI
	},
	projects: [
		{
			name: 'chromium',
			use: { ...devices['Desktop Chrome'] }
		},
		{
			name: 'firefox',
			use: { ...devices['Desktop Firefox'] }
		},
		{
			name: 'webkit',
			use: { ...devices['Desktop Safari'] }
		}
	]
};
module.exports = config;
