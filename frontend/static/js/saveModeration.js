function submitModerationUpdate({ enabled = null, modlog = null, triggerStatus = null, triggerInput = null, responseStatus = null, responseInput = null, roles = null, muteRole = null, managedMuteRole = null, updateRoles = null, update }) {
	const baseurl = document.body.getAttribute('data-url');
	const guildID = document.body.getAttribute('data-guild-id');
	const url = `${baseurl}/dashboard/${guildID}/manage/updateModeration`;

	const body = { update };
	if (enabled !== null) body.enabled = enabled;
	if (modlog !== null) body.modlog = modlog;
	if (triggerStatus !== null) body.triggerStatus = triggerStatus;
	if (triggerInput !== null) body.triggerInput = triggerInput;
	if (responseStatus !== null) body.responseStatus = responseStatus;
	if (responseInput !== null) body.responseInput = responseInput;
	if (roles !== null) body.roles = roles;
	if (muteRole !== null) body.muteRole = muteRole
	if (managedMuteRole !== null) body.managedMuteRole = managedMuteRole
	if (updateRoles !== null) body.updateRoles = updateRoles;

	fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(body)
	}).then(() => {
		if (enabled == null && responseStatus == null && triggerStatus == null) {
			console.log(body)
			const alert = document.getElementById('successAlert');
			if (alert) {
				alert.style.display = 'block';
				alert.style.opacity = 1;
				setTimeout(() => {
					location.reload();
				}, 1500);
			}
		}
	}).catch(console.error);
}

function debounce(func, delay = 300) {
	let timeout;
	return function (...args) {
		clearTimeout(timeout);
		timeout = setTimeout(() => func.apply(this, args), delay);
	};
}

// Save all input fields
document.getElementById('saveAllButton').addEventListener('click', function (e) {
	e.preventDefault();

	const modlogInput = document.getElementById('modlogInput').value.trim();
	const triggerInput = document.getElementById('triggerInput').value.trim();
	const responseInput = document.getElementById('responseInput').value.trim();
	const roleActions = ['Warn', 'Mute', 'Unmute', 'Kick', 'Ban', 'Unban'];
	const roles = {};
	const muteRoleInput = document.getElementById('muteRoleInput').value.trim();
	const updateRolesInput = document.getElementById('updateRolesInput').value.trim();

	roleActions.forEach(action => {
		const input = document.getElementById(`requiredRoles${action}Input`);
		roles[action] = JSON.parse(input.value || "[]");
	});

	submitModerationUpdate({
		modlog: modlogInput,
		roles: roles,
		triggerInput: triggerInput,
		responseInput: responseInput,
		muteRole: muteRoleInput,
		updateRoles: JSON.parse(updateRolesInput),
		update: "all"
	});
});

// Save modlog input
document.getElementById('saveModlogButton').addEventListener('click', function (e) {
	e.preventDefault();

	const modlogInput = document.getElementById('modlogInput').value.trim();
	submitModerationUpdate({
		modlog: modlogInput,
		update: "modlog"
	});
});

// Save trigger delay
document.getElementById('saveTriggerButton').addEventListener('click', function (e) {
	e.preventDefault();
	
	const triggerInput = parseInt(document.getElementById('triggerInput').value, 10) || 0;
	submitModerationUpdate({
		triggerInput: triggerInput,
		update: "triggerInput"
	});
});

// Save response delay
document.getElementById('saveResponseButton').addEventListener('click', function (e) {
	e.preventDefault();
	
	const responseInput = parseInt(document.getElementById('responseInput').value, 10) || 0;
	submitModerationUpdate({
		responseInput: responseInput,
		update: "responseInput"
	});
});

// Save individual roles input
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

// Save mute role
document.getElementById('saveMuteRoleButton').addEventListener('click', function (e) {
	e.preventDefault();

	const muteRoleInput = document.getElementById('muteRoleInput').value.trim();
	submitModerationUpdate({
		muteRole: muteRoleInput,
		update: "muteRole"
	});
});

// Save managed mute roles
document.getElementById('saveMuteRolesButton').addEventListener('click', function (e) {
	e.preventDefault();

	const updateRolesInput = document.getElementById('updateRolesInput').value.trim();
	submitModerationUpdate({
		updateRoles: updateRolesInput,
		update: "updateRoles"
	});
});

// Save toggle statuses
document.addEventListener('DOMContentLoaded', () => {
	// Moderation toggle
	const moderationToggle = document.getElementById('toggleModeration');
	if (moderationToggle) {
		moderationToggle.addEventListener('change', debounce(() => {
			submitModerationUpdate({
				enabled: moderationToggle.checked,
				update: "status"
			});
		}));
	}

	// Trigger toggle
	const triggerToggle = document.getElementById('toggleTriggerDeletion');
	if (triggerToggle) {
		triggerToggle.addEventListener('change', debounce(() => {
			submitModerationUpdate({
				triggerStatus: triggerToggle.checked,
				update: "triggerStatus"
			});
		}));
	}

	// Response toggle
	const responseToggle = document.getElementById('toggleResponseDeletion');
	if (responseToggle) {
		responseToggle.addEventListener('change', debounce(() => {
			submitModerationUpdate({
				responseStatus: responseToggle.checked,
				update: "responseStatus"
			});
		}));
	}

	// Managed mute role toggle
	const muteRoleToggle = document.getElementById('toggleManagedMuteRole');
	if (muteRoleToggle) {
		muteRoleToggle.addEventListener('change', debounce(() => {
			submitModerationUpdate({
				managedMuteRole: muteRoleToggle.checked,
				update: "manageMuteRole"
			});
		}));
	}
});