document.addEventListener('DOMContentLoaded', () => {
	initModlogSelector();
	initRoleCheckboxes();
	populateInitialRoleConfig();
	initMuteRoleSelector();
	initMuteRolesCheckboxes();
	populateInitialMuteRoleConfig();
});

// Helper functions

function initModlogSelector() {
	const channelItems = document.querySelectorAll('.channelListItem');
	const modLogLabel = document.getElementById('modlogDropdownLabel');
	const modlogHiddenInput = document.getElementById('modlogInput');
	const modlogID = document.body.getAttribute('data-modlog-id');

	channelItems.forEach(item => {
		item.addEventListener('click', e => {
			e.preventDefault();
			modLogLabel.textContent = item.textContent;
			modlogHiddenInput.value = item.getAttribute('data-value');
		});
	});

	if (modlogID) {
		const selectedItem = document.querySelector(`a[data-value="${modlogID}"]`);
		if (selectedItem) {
			modLogLabel.textContent = selectedItem.textContent;
			selectedItem.classList.add('selected');
		}
	}
}

function initMuteRoleSelector() {
	const channelItems = document.querySelectorAll('.roleListItem');
	const muteRoleLabel = document.getElementById('muteRoleLabel');
	const muteRoleHiddenInput = document.getElementById('muteRoleInput');
	const muteRoleID = document.body.getAttribute('data-muterole-id');
	
	channelItems.forEach(item => {
		item.addEventListener('click', e => {
			e.preventDefault();
			muteRoleLabel.textContent = item.textContent;
			muteRoleHiddenInput.value = item.getAttribute('data-value');
		});
	});
	
	if (muteRoleID) {
		const selectedItem = document.querySelector(`a[data-value="${muteRoleID}"]`);
		if (selectedItem) {
			muteRoleLabel.textContent = selectedItem.textContent;
			selectedItem.classList.add('selected');
		}
	}
}

function initRoleCheckboxes() {
	document.querySelectorAll('.roleCheckbox').forEach(checkbox => {
		checkbox.addEventListener('change', () => {
			const container = checkbox.closest('.input-group');
			const checked = container.querySelectorAll('.roleCheckbox:checked');

			const names = Array.from(checked).map(cb => cb.parentElement.textContent.trim());
			const ids = Array.from(checked).map(cb => cb.value);

			container.querySelector('#requiredRolesLabel').textContent = names.length ? names.join(', ') : 'Select roles';
			container.querySelector('input[type="hidden"]').value = JSON.stringify(ids);
		});
	});

	document.querySelectorAll('.requiredRolesDropdown + .dropdown-menu').forEach(menu => {
		menu.addEventListener('click', e => e.stopPropagation());
	});
}

function populateInitialRoleConfig() {
	const commandRestrictions = JSON.parse(document.getElementById('restrictions').getAttribute('data-restrictions'));

	for (const [action, roleIDs] of Object.entries(commandRestrictions)) {
		if (!Array.isArray(roleIDs)) continue;

		const inputGroup = document.querySelector(`.saveRolesButton[data-action="${action}"]`)?.closest('.input-group');
		if (!inputGroup) continue;

		const label = inputGroup.querySelector('#requiredRolesLabel');
		const hiddenInput = inputGroup.querySelector('input[type="hidden"]');
		const checkboxes = inputGroup.querySelectorAll('.roleCheckbox');

		const selectedNames = [];

		checkboxes.forEach(cb => {
			if (roleIDs.includes(cb.value)) {
				cb.checked = true;
				selectedNames.push(cb.parentElement.textContent.trim());
			}
		});
	
		label.textContent = selectedNames.length ? selectedNames.join(', ') : 'Select roles';
		hiddenInput.value = JSON.stringify(roleIDs);
	}
}

function initMuteRolesCheckboxes() {
	document.querySelectorAll('.muteRoleCheckbox').forEach(checkbox => {
		checkbox.addEventListener('change', () => {
			const container = checkbox.closest('.input-group');
			const checked = container.querySelectorAll('.muteRoleCheckbox:checked');

			const names = Array.from(checked).map(cb => cb.parentElement.textContent.trim());
			const ids = Array.from(checked).map(cb => cb.value);

			container.querySelector('#updateRolesLabel').textContent = names.length ? names.join(', ') : 'Select roles';
			container.querySelector('input[type="hidden"]').value = JSON.stringify(ids);
		});
	});

	document.querySelectorAll('.requiredRolesDropdown + .dropdown-menu').forEach(menu => {
		menu.addEventListener('click', e => e.stopPropagation());
	});
}

function populateInitialMuteRoleConfig() {
	const muteRoles = JSON.parse(document.getElementById('restrictions').getAttribute('data-updatemuteroles'));

	if (!Array.isArray(muteRoles)) return;

	const inputGroup = document.querySelector(`.saveMuteRolesButton[data-action="updateRoles"]`)?.closest('.input-group');
	if (!inputGroup) return;

	const label = inputGroup.querySelector('#updateRolesLabel');
	const hiddenInput = inputGroup.querySelector('input[type="hidden"]');
	const checkboxes = inputGroup.querySelectorAll('.muteRoleCheckbox');

	const selectedNames = [];

	checkboxes.forEach(cb => {
		if (muteRoles.includes(cb.value)) {
			cb.checked = true;
			selectedNames.push(cb.parentElement.textContent.trim());
		}
	});

	label.textContent = selectedNames.length ? selectedNames.join(', ') : 'Select roles';
	hiddenInput.value = JSON.stringify(muteRoles);
}