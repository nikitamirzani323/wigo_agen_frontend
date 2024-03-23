import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		table_header_font: "13px",
		table_body_font: "13px",
		path_websocket: "isbskor.com",
	}
});

export default app;