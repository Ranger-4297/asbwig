function submitModerationUpdate({ enabled = null, modlog = null, roles = null, triggerStatus = null, triggerInput = null, responseStatus = null, responseInput = null, update }) {
	const baseurl = document.body.getAttribute('data-url');
	const guildID = document.body.getAttribute('data-guild-id');
	const url = `${baseurl}/dashboard/${guildID}/manage/updateModeration`;
	
	const body = { update };
	if (enabled !== null) body.enabled = enabled;
	if (modlog !== null) body.modlog = modlog;
	if (roles !== null) body.roles = roles;
	if (triggerStatus !== null) body.triggerStatus = triggerStatus;
	if (triggerInput !== null) body.triggerInput = triggerInput;
	if (responseStatus !== null) body.responseStatus = responseStatus;
	if (responseInput !== null) body.triggerInput = responseInput;
	
	fetch(url, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(body)
	}).then(() => {
		if (enabled == null && responseStatus == null && triggerStatus == null) {
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

// Save all button
document.getElementById('saveAllButton').addEventListener('click', function (e) {
	e.preventDefault();
	
	const modlogInput = document.getElementById('modlogInput').value.trim();
	const triggerInput = document.getElementById('triggerInput').value.trim();
	const responseInput = document.getElementById('responseInput').value.trim();
	const roleActions = ['Warn', 'Mute', 'Unmute', 'Kick', 'Ban', 'Unban'];
	const roles = {};
	
	roleActions.forEach(action => {
		const input = document.getElementById(`requiredRoles${action}Input`);
		roles[action] = JSON.parse(input.value || "[]");
	});
	
	submitModerationUpdate({
		modlog: modlogInput,
		roles: roles,
		triggerInput: triggerInput,
		responseInput: responseInput,
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

// Save moderation status
document.addEventListener('DOMContentLoaded', () => {
	const moderationToggle = document.getElementById('toggleModeration');

	if (moderationToggle) {
		moderationToggle.addEventListener('change', () => {
			const isEnabled = moderationToggle.checked;

			submitModerationUpdate({
				enabled: isEnabled,
				update: "status"
			});
		});
	}
});

// Save trigger deletion status
document.addEventListener('DOMContentLoaded', () => {
	const triggerToggle = document.getElementById('toggleTriggerDeletion');

	if (triggerToggle) {
		triggerToggle.addEventListener('change', () => {
			const isEnabled = triggerToggle.checked;

			submitModerationUpdate({
				triggerStatus: isEnabled,
				update: "triggerStatus"
			});
		});
	}
});

// Save trigger time
document.getElementById('saveTriggerButton').addEventListener('click', function (e) {
	e.preventDefault();
	
	const triggerInput = parseInt(document.getElementById('triggerInput').value, 10) || 0;
	submitModerationUpdate({
		triggerInput: triggerInput,
		update: "triggerInput"
	});
});

// Save response deletion status
document.addEventListener('DOMContentLoaded', () => {
	const responseToggle = document.getElementById('toggleResponseDeletion');

	if (responseToggle) {
		responseToggle.addEventListener('change', () => {
			const isEnabled = responseToggle.checked;

			submitModerationUpdate({
				responseStatus: isEnabled,
				update: "responseStatus"
			});
		});
	}
});

// Save response time
document.getElementById('saveResponseButton').addEventListener('click', function (e) {
	e.preventDefault();
	
	const responseInput = parseInt(document.getElementById('responseInput').value, 10) || 0;
	submitModerationUpdate({
		responseInput: responseInput,
		update: "responseInput"
	});
});

// Helper functions

// Display current modlog config
document.addEventListener('DOMContentLoaded', function () {
	const channelItems = document.querySelectorAll('.channelListItem');
	const modLogLabel = document.getElementById('modlogDropdownLabel');
	const modlogHiddenInput = document.getElementById('modlogInput');
	
	const modlogID = document.querySelector('body').getAttribute('data-modlog-id');
	
	channelItems.forEach(item => {
		item.addEventListener('click', function (e) {
			e.preventDefault();
			const value = this.getAttribute('data-value');
			const text = this.textContent;
			modLogLabel.textContent = text;
			modlogHiddenInput.value = value;
		});
	});
	
	if (modlogID) {
		const selectedItem = document.querySelector(`a[data-value="${modlogID}"]`);
		if (selectedItem) {
			modLogLabel.textContent = selectedItem.textContent;
			selectedItem.classList.add('selected');
		}
	}
});

// Keep roles modal open
document.addEventListener('DOMContentLoaded', () => {
	document.querySelectorAll('.roleCheckbox').forEach(roleCheckbox => {
		roleCheckbox.addEventListener('change', function () {
			const container = roleCheckbox.closest('.input-group');
			const selected = Array.from(container.querySelectorAll('.roleCheckbox:checked'))
			.map(cb => cb.parentElement.textContent.trim());
			const selectedIDs = Array.from(container.querySelectorAll('.roleCheckbox:checked'))
			.map(cb => cb.value);
			
			const label = container.querySelector('#requiredRolesLabel');
			const hiddenInput = container.querySelector('input[type="hidden"]');
			
			label.textContent = selected.length > 0 ? selected.join(', ') : 'Select roles';
			hiddenInput.value = JSON.stringify(selectedIDs);
		});
	});
	document.querySelectorAll('.requiredRolesDropdown + .dropdown-menu').forEach(menu => {
		menu.addEventListener('click', function (e) {
			e.stopPropagation();
		});
	});
});

// Display current role config
document.addEventListener('DOMContentLoaded', function () {
	const commandRestrictions = JSON.parse(document.getElementById('settings').getAttribute('data-restrictions'));
	for (const [action, roleIDs] of Object.entries(commandRestrictions)) {
		if (!Array.isArray(roleIDs)) continue;
	
		const inputGroup = document.querySelector(`.saveRolesButton[data-action="${action}"]`)?.closest('.input-group');
		if (!inputGroup) continue;
		
		const label = inputGroup.querySelector('#requiredRolesLabel');
		const hiddenInput = inputGroup.querySelector('input[type="hidden"]');
		const checkboxes = inputGroup.querySelectorAll('.roleCheckbox');
		
		let selectedNames = [];
		
		checkboxes.forEach(cb => {
			if (roleIDs.includes(cb.value)) {
				cb.checked = true;
				selectedNames.push(cb.parentElement.textContent.trim());
			}
		});
		
		if (label) label.textContent = selectedNames.length > 0 ? selectedNames.join(', ') : 'Select roles';
		if (hiddenInput) hiddenInput.value = JSON.stringify(roleIDs);
	}
});