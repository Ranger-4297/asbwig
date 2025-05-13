document.getElementById('saveModlogButton').addEventListener('click', function() {
	const input = document.getElementById('modlogInput').value.trim();
	document.getElementById('successAlert').style.display = 'block';
	document.getElementById('successAlert').style.opacity = 1; // Fade in the alert
	const baseurl = document.body.getAttribute('data-url');
	const guildID = document.body.getAttribute('data-guild-id');
	const url = `${baseurl}/dashboard/${guildID}/manage/update-channel`;
	fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify({ channel: input })
	}).then(() => {
		const alert = document.getElementById('successAlert');
		if (alert) {
			alert.style.display = 'block';
			alert.style.opacity = 1;
			setTimeout(() => {
				location.reload();
			}, 1500);
		}
	}).catch(console.error);
});