<script lang="ts">
	import PocketBase from 'pocketbase';
	import { onMount } from 'svelte';

	const pb = new PocketBase('http://127.0.0.1:8090');

	const create = async () => {
		const data = {
			datetime: '2022-01-01 10:00:00.123Z',
			title: 'test',
			content: 'test'
		};

		const record = await pb.collection('reminders').create(data);
	};

	const get = async () => {
		const records = await pb.collection('reminders').getFullList(200 /* batch size */, {
			sort: '-created'
		});

		return records;
	};

	let data;

	get().then((res) => (data = res));

	// onMount(() => {
	// 	const interval = setInterval(() => {
	// 		get().then((res) => (data = res));
	// 	}, 1000);

	// 	return () => {
	// 		clearInterval(interval);
	// 	};
	// });
</script>

<div class="hero min-h-screen bg-base-200">
	<div class="hero-content text-center">
		<div class="max-w-md">
			<h1 class="text-5xl font-bold">Better Email Reminders</h1>
			<p class="py-6">It's easier than schedule sending to yourself!</p>
			<a href="#create" class="btn btn-primary"
				>Get Started</a
			>
		</div>
	</div>
</div>

<div class="divider"></div>

<br />
<div id="create">
	<button class="btn" on:click={create}>Test</button>
</div>

<!-- 

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<button class="btn" on:click={create}>Test</button>
<pre>{JSON.stringify(data, null, 2)}</pre> -->
