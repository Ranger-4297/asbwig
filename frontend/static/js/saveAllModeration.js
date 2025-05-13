function submitModerationUpdate({ modlog = null, roles = null, update }) {
	const baseurl = document.body.getAttribute('data-url');
	const guildID = document.body.getAttribute('data-guild-id');
	const url = `${baseurl}/dashboard/${guildID}/manage/updateModeration`;

	const body = { update };
	if (modlog !== null) body.modlog = modlog;
	if (roles !== null) body.roles = roles;

	fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(body)
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
}

// Save all button
document.getElementById('saveAllButton').addEventListener('click', function (e) {
	e.preventDefault();

	const modlogInput = document.getElementById('modlogInput').value.trim();
	const roleActions = ['Warn', 'Mute', 'Unmute', 'Kick', 'Ban', 'Unban'];
	const roles = {};

	roleActions.forEach(action => {
		const input = document.getElementById(`requiredRoles${action}Input`);
		roles[action] = JSON.parse(input.value || "[]");
	});

	submitModerationUpdate({
		modlog: modlogInput,
		roles: roles,
		update: "all"
	});
});

// Save modlog only
document.getElementById('saveModlogButton').addEventListener('click', function (e) {
	e.preventDefault();

	const modlogInput = document.getElementById('modlogInput').value.trim();
	submitModerationUpdate({
		modlog: modlogInput,
		update: "modlog"
	});
});

// Save individual roles
document.querySelectorAll('.saveRolesButton').forEach(button => {
	button.addEventListener('click', function (e) {
		e.preventDefault();

		const action = this.getAttribute('data-action');

		const input = document.getElementById(`requiredRoles${action}Input`);
		const roles = {};
		roles[action] = JSON.parse(input.value || "[]");

		submitModerationUpdate({
			roles: roles,
			update: action
		});
	});
});