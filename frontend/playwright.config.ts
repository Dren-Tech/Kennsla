import { devices } from '@playwright/test';

/** @type {import('@playwright/test').PlaywrightTestConfig} */
const config = {
	forbidOnly: !!process.env.CI,
	retries: process.env.CI ? 2 : 0,
	testMatch: 'tests/**/*.ts',
	use: {
		baseURL: 'http://localhost:3000',
		screenshot: 'only-on-failure',
		extraHTTPHeaders: {
			'Authorization': 'token TEST'
		}
	},
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
export default config;
